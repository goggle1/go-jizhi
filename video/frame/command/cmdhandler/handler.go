package cmdhandler

import (
	"context"
	"encoding/json"
	"fmt"

	"time"

	"git.tvblack.com/video/frame/core"
	"git.tvblack.com/video/frame/message"
	"git.tvblack.com/video/frame/proto/p_common"
	"github.com/golang/protobuf/ptypes"
	api "github.com/micro/go-api/proto"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/metadata"
	"github.com/sirupsen/logrus"
)

type RequestHandler struct {
	Publisher micro.Publisher
}

func (r *RequestHandler) getResult(c context.Context, reqHead *p_common.RequestHead) (*message.Result, *message.ReqContext, error) {
	cmd := core.App.GetCommand().Get(reqHead.Request.Cmd)
	if cmd == nil {
		logrus.Error("not find cmd:", reqHead.Request.Cmd)
		return nil, nil, errors.BadRequest("OnRequest", "not find cmd:"+reqHead.Request.Cmd)
	}

	context := message.NewReqContext(reqHead.Request.Cmd, c, reqHead.Request.Parameter)
	context.Publisher = r.Publisher
	result := cmd.DoRequest(cmd, context)
	return result, context, nil
}

func (r *RequestHandler) OnRequestRpc(c context.Context, reqHead *p_common.RequestHead, rspHead *p_common.ResponseHead) error {
	startTime := time.Now()
	if reqHead.Request == nil {
		logrus.Error("Request = nil")
		return errors.BadRequest("OnRequest", "not find request")
	}
	if reqHead.Request.Cmd == "" {
		logrus.Error("cmd = nil")
		return errors.BadRequest("OnRequest", "cmd is null")
	}

	result, context, err := r.getResult(c, reqHead)
	if err != nil {
		return err
	}
	rspHead.Code = result.Code
	rspHead.Msg = result.Msg
	rspHead.ErrCode = result.ErrCode

	if result.Success() {
		rspHead.Response = context.ToAny(reqHead.Request.Cmd)
	}

	endTime := time.Now().Sub(startTime)
	t := endTime / 1e6
	ServiceStatics.Statics(reqHead.Request.Cmd, float64(t))
	return nil
}

func (r *RequestHandler) OnRequestJson(c context.Context, reqHead *p_common.RequestHead, rspHead *p_common.RspJsonHead) error {
	startTime := time.Now()
	if reqHead.Request == nil {
		logrus.Error("Request = nil")
		return errors.BadRequest("OnRequest", "not find request")
	}
	if reqHead.Request.Cmd == "" {
		logrus.Error("cmd = nil")
		return errors.BadRequest("OnRequest", "cmd is null")
	}

	result, context, err := r.getResult(c, reqHead)
	if err != nil {
		return err
	}
	rspHead.Code = result.Code
	rspHead.Msg = result.Msg
	if result.Success() {
		rspHead.Response = context.GetResponse()
	}

	endTime := time.Now().Sub(startTime)
	t := endTime / 1e6
	ServiceStatics.Statics(reqHead.Request.Cmd, float64(t))

	return nil
}

func (r *RequestHandler) OnRequest(c context.Context, req *api.Request, rsp *api.Response) error {
	md, ok := metadata.FromContext(c)
	if !ok {
		md = metadata.Metadata{}
	}
	cmdPair, ok := req.Get["cmd"]
	if !ok {
		return errors.BadRequest("OnRequest", "cmd is null")
	}
	if len(cmdPair.Values) <= 0 {
		return errors.BadRequest("OnRequest", "cmd < 0")
	}
	cmd := cmdPair.Values[0]
	if cmd == "serviceStat" {
		data := ServiceStatics.GetStatics()
		rsp.StatusCode = 200
		rsp.Body = data
		return nil
	}
	rsp.Header = make(map[string]*api.Pair)
	reqItem := r.getHttpRequest(req)

	reqHead := &p_common.RequestHead{
		Request: &p_common.MessageItem{
			Cmd:       cmd,
			Parameter: reqItem,
		},
	}

	msgType, _ := md["MsgType"]

	if msgType == message.RpcMessage {
		var rspHead p_common.ResponseHead
		err := r.OnRequestRpc(c, reqHead, &rspHead)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"cmd": cmd,
			}).Error("system error:", err)
			fmt.Println(err.Error())
			return errors.BadRequest("OnRequest", "system error")
		}
		rsp.StatusCode = 200
		b, _ := json.Marshal(rspHead)
		rsp.Body = string(b)
	} else {
		var rspHead p_common.RspJsonHead
		result, context, err := r.getResult(c, reqHead)
		if err != nil {
			return err
		}
		rspHead.Code = result.Code
		rspHead.Msg = result.Msg
		rsp.StatusCode = 200
		if !result.Success() {
			b, _ := json.Marshal(rspHead)
			rsp.Body = string(b)
			return nil
		}
		rspHead.Response = context.GetResponse()
		if context.RspType == message.RspJson {
			b, _ := json.Marshal(rspHead)
			rsp.Body = string(b)
		} else if context.RspType == message.RspXml {

			rsp.Header["Content-Type"] = &api.Pair{
				Key:    "Content-Type",
				Values: []string{"text/xml"},
			}
			rsp.Body = rspHead.Response
		} else {
			rsp.Header["Content-Type"] = &api.Pair{
				Key:    "Content-Type",
				Values: []string{"application/json"},
			}
			rsp.Body = string(rspHead.Response)
		}
		// err := r.OnRequestJson(c, reqHead, &rspHead)
		// if err != nil {
		// 	logrus.WithFields(logrus.Fields{
		// 		"cmd": cmd,
		// 	}).Error("system error:", err)
		// 	fmt.Println(err.Error())
		// 	return errors.BadRequest("OnRequest", "system error")
		// }

		// rsp.StatusCode = 200
		// b, _ := json.Marshal(rspHead)
		// rsp.Body = string(b)

	}

	return nil
}

type rspItem struct {
	items []*p_common.MessageContent
}

func (r *RequestHandler) getHttpRequest(apiReq *api.Request) []*p_common.MessageContent {
	req := &rspItem{}
	r.pairToValue(req, apiReq.Get)
	r.pairToValue(req, apiReq.Post)

	proto := &p_common.StringMsg{
		Value: apiReq.Body,
	}
	any, err := ptypes.MarshalAny(proto)
	if err != nil {
		logrus.Error("getHttpRequest error:", err)
	}
	req.items = append(req.items, &p_common.MessageContent{
		Key: "body",
		Obj: any,
	})
	return req.items
}
func (r *RequestHandler) pairToValue(req *rspItem, ps map[string]*api.Pair) {
	for _, p := range ps {
		item := &p_common.MessageContent{}
		item.Key = p.Key
		if len(p.Values) > 1 {
			proto := &p_common.StringsMsg{}
			for _, v := range p.Values {
				proto.Value = append(proto.Value, v)
			}
			any, err := ptypes.MarshalAny(proto)
			if err != nil {
				logrus.Error("getHttpRequest error:", err)
			}
			item.Obj = any
		} else {
			proto := &p_common.StringMsg{
				Value: p.Values[0],
			}
			any, err := ptypes.MarshalAny(proto)
			if err != nil {
				logrus.Error("getHttpRequest error:", err)
			}
			item.Obj = any

		}
		req.items = append(req.items, item)
	}
}
