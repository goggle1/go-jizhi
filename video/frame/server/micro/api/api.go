package api

import (
	"io/ioutil"
	"mime"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"git.tvblack.com/hardware/smart_lamp_server/api/module/micro/helper"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"
)

func APIHttp(w http.ResponseWriter, r *http.Request) {

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

	req := &api.Request{
		Path:   r.URL.Path,
		Method: r.Method,
		Header: make(map[string]*api.Pair),
		Get:    make(map[string]*api.Pair),
		Post:   make(map[string]*api.Pair),
		Url:    r.URL.String(),
	}
	// Get data
	for key, vals := range r.URL.Query() {
		header, ok := req.Get[key]
		if !ok {
			header = &api.Pair{
				Key: key,
			}
			req.Get[key] = header
		}
		header.Values = vals
	}
	request = req
	service = "iptv.server.srv"
	method = "Message.OnRequest"
	switch ct {
	case "application/json":
		data, _ := ioutil.ReadAll(r.Body)
		req.Body = string(data)

		ct, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil {
			ct = "application/x-www-form-urlencoded"
			r.Header.Set("Content-Type", ct)
		}
		// Set X-Forwarded-For if it does not exist
		if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
			if prior, ok := r.Header["X-Forwarded-For"]; ok {
				ip = strings.Join(prior, ", ") + ", " + ip
			}

			// Set the header
			req.Header["X-Forwarded-For"] = &api.Pair{
				Key:    "X-Forwarded-For",
				Values: []string{ip},
			}
		}

		// Get data
	default:
		data, _ := ioutil.ReadAll(r.Body)
		req.Body = string(data)
	}
	request = req

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
	reqOp := (*cmd.DefaultOptions().Client).NewRequest(service, method, request, client.WithContentType("application/json"))

	// create context
	ctx := helper.RequestToContext(r)

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
	err = (*cmd.DefaultOptions().Client).Call(ctx, reqOp, &response, opts...)
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
		w.Header().Set("Content-Type", ct)
	}

	// b, _ := response.MarshalJSON()
	// w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	// w.Write(response)
	w.WriteHeader(int(response.StatusCode))
	w.Write([]byte(response.Body))
}
