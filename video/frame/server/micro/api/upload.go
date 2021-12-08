package api

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-log/log"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	api "github.com/micro/micro/api/proto"
)

type ApiUpload func(response api.Response, r *http.Request) []byte

//UploadMessage UploadMessage
func UploadMessage(w http.ResponseWriter, r *http.Request, upCall ApiUpload) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
	request, err := requestToProto(r)
	if err != nil {
		log.Log("Registry == nil")
		http.Error(w, "Registry == nil", 500)
		return
	}

	req := getRequest(request)

	if req.Service == "" {
		log.Log("Service == nil")
		http.Error(w, "Service == nil", 500)
		return
	}
	if req.Method == "" {
		log.Log("Method == nil")
		http.Error(w, "Method == nil", 500)
		return
	}

	cli := (*cmd.DefaultOptions().Client).NewRequest(req.Service, req.Method, req.Request, client.WithContentType("application/json"))
	// create context
	ctx := requestToContext(r)

	var opts []client.CallOption

	timeout, _ := strconv.Atoi(r.Header.Get("Timeout"))
	// set timeout
	if timeout > 0 {
		opts = append(opts, client.WithRequestTimeout(time.Duration(timeout)*time.Second))
	}

	// remote call
	if len(req.Address) > 0 {
		opts = append(opts, client.WithAddress(req.Address))
	}

	var response api.Response
	// remote call
	err = (*cmd.DefaultOptions().Client).Call(ctx, cli, &response, opts...)
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
	if response.StatusCode != 200 {
		http.Error(w, response.Body, 500)
		return
	}

	bc := upCall(response, r)

	w.Header().Set("Content-Length", strconv.Itoa(len(bc)))
	w.Write(bc)
}
func requestToContext(r *http.Request) context.Context {
	ctx := context.Background()
	md := make(metadata.Metadata)
	for k, v := range r.Header {
		md[k] = strings.Join(v, ",")
	}
	return metadata.NewContext(ctx, md)
}
