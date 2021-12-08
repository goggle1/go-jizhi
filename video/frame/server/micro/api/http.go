package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.tvblack.com/video/frame/common"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
)

type rpcRequest struct {
	Service string
	Method  string
	Address string
	Request interface{}
}

// RPC Handler passes on a JSON or form encoded RPC request to
// a service.
func RPC(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	badRequest := func(description string) {
		e := errors.BadRequest("go.micro.rpc", description)
		w.WriteHeader(400)
		w.Write([]byte(e.Error()))
	}

	var service, method, address string
	var request interface{}

	// response content type
	w.Header().Set("Content-Type", "application/json")

	ct := r.Header.Get("Content-Type")

	// Strip charset from Content-Type (like `application/json; charset=UTF-8`)
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}

	switch ct {
	case "application/json":
		var rpcReq rpcRequest

		d := json.NewDecoder(r.Body)
		d.UseNumber()

		if err := d.Decode(&rpcReq); err != nil {
			badRequest(err.Error())
			return
		}

		service = rpcReq.Service
		method = rpcReq.Method
		address = rpcReq.Address
		request = rpcReq.Request

		// JSON as string
		if req, ok := rpcReq.Request.(string); ok {
			d := json.NewDecoder(strings.NewReader(req))
			d.UseNumber()

			if err := d.Decode(&request); err != nil {
				badRequest("error decoding request string: " + err.Error())
				return
			}
		}
	default:
		apiReq, err := requestToProto(r)
		if err != nil {
			er := errors.InternalServerError("go.micro.api", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			w.Write([]byte(er.Error()))
			return
		}
		req := getRequest(apiReq)

		service = req.Service
		method = req.Method
		address = req.Address
		request = apiReq

	}

	if len(service) == 0 {
		badRequest("invalid service")
		return
	}

	if len(method) == 0 {
		badRequest("invalid method")
		return
	}

	// create request/response
	//var response json.RawMessage
	response := &api.Response{}
	var err error
	req := (*cmd.DefaultOptions().Client).NewRequest(service, method, request, client.WithContentType("application/json"))

	// create context
	ctx := common.RequestToContext(r)

	var opts []client.CallOption

	timeout, _ := strconv.Atoi(r.Header.Get("Timeout"))
	// set timeout
	if timeout > 0 {
		opts = append(opts, client.WithRequestTimeout(time.Duration(timeout)*time.Second))
	}

	// remote call
	if len(address) > 0 {
		opts = append(opts, client.WithAddress(address))
	}

	// remote call
	err = (*cmd.DefaultOptions().Client).Call(ctx, req, &response, opts...)
	if err != nil {
		ce := errors.Parse(err.Error())
		switch ce.Code {
		case 0:
			// assuming it's totally screwed
			ce.Code = 500
			ce.Id = "go.micro.rpc"
			ce.Status = http.StatusText(500)
			ce.Detail = "error during request: " + ce.Detail
			w.WriteHeader(500)
		default:
			w.WriteHeader(int(ce.Code))
		}
		w.Write([]byte(ce.Error()))
		return
	}
	response.StatusCode = http.StatusOK
	for _, header := range response.GetHeader() {
		for _, val := range header.Values {
			w.Header().Add(header.Key, val)
		}
	}

	if len(w.Header().Get("Content-Type")) == 0 {
		w.Header().Set("Content-Type", "application/json")
	}

	// b, _ := response.MarshalJSON()
	// w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	// w.Write(response)
	w.WriteHeader(int(response.StatusCode))
	w.Write([]byte(response.Body))
}
func getRequest(req *api.Request) rpcRequest {
	var request rpcRequest

	service, ok := req.Get["s"]
	if !ok {
		service, ok = req.Post["s"]
	}
	if ok && len(service.Values) > 0 {
		request.Service = service.Values[0]
	}

	method, ok := req.Get["m"]
	if !ok {
		method, ok = req.Get["m"]
	}
	if ok && len(method.Values) > 0 {
		request.Method = method.Values[0]
	}

	request.Request = req
	return request
}
