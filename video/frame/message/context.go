package message

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"

	"git.tvblack.com/video/frame/proto/p_common"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	micro "github.com/micro/go-micro"
	"github.com/sirupsen/logrus"
)

type ResponseType int

const (
	RspJson ResponseType = 0
	RspXml  ResponseType = 1
	RspText ResponseType = 2
)

type IContext interface {
	GetID() int32
	GetName() string
	GetLogger() *logrus.Entry
	GetObject() interface{}
	GetService() interface{}
	GetDeviceID() int32
	GetContext() context.Context
}

type UserContext struct {
	ID      int32
	Name    string
	Logger  *logrus.Entry
	Item    interface{}
	Service interface{}
	Device  int32
	Context context.Context
}

func (u *UserContext) GetID() int32 {
	return u.ID
}

func (u *UserContext) GetContext() context.Context {
	return u.Context
}

func (u *UserContext) GetName() string {
	return u.Name
}

func (u *UserContext) GetObject() interface{} {
	return u.Item
}

func (u *UserContext) GetService() interface{} {
	return u.Service
}

func (u *UserContext) GetLogger() *logrus.Entry {
	return u.Logger
}

func (u *UserContext) GetDeviceID() int32 {
	return u.Device
}

type ReqItem struct {
	items map[string]*any.Any
}

func (r *ReqItem) Get(key string) *any.Any {
	o, ok := r.items[key]
	if !ok {
		return nil
	}
	return o
}

func (r *ReqItem) GetMap() map[string]*any.Any {
	return r.items
}

func (r *ReqItem) GetInt32(key string) int32 {
	o := r.Get(key)
	if o == nil {
		return 0
	}
	pb := p_common.IntMsg{}
	if err := ptypes.UnmarshalAny(o, &pb); err != nil {
		strpb := p_common.StringMsg{}
		if err = ptypes.UnmarshalAny(o, &strpb); err != nil {
			logrus.Error("UnmarshalAny error :", err)
			return 0
		}
		if strpb.Value == "" {
			return 0
		}
		v, err := strconv.Atoi(strpb.Value)
		if err != nil {
			logrus.Error("strconv.Atoi error :", err)
			return 0
		}
		return int32(v)
	}

	return pb.Value
}

func (r *ReqItem) GetInts(key string) []int32 {
	v := r.GetString(key)
	if v == "" {
		return []int32{}
	}
	s := strings.Split(v, ",")
	var is []int32
	for _, sv := range s {
		intv, err := strconv.Atoi(sv)
		if err != nil {
			continue
		}
		is = append(is, int32(intv))
	}
	return is
}

func (r *ReqItem) GetInt(key string) int {
	return int(r.GetInt32(key))
}

func (r *ReqItem) GetString(key string) string {
	o := r.Get(key)
	if o == nil {
		return ""
	}
	pb := p_common.StringMsg{}
	if err := ptypes.UnmarshalAny(o, &pb); err != nil {
		logrus.Error("UnmarshalAny error :", err)
		return ""
	}

	return pb.Value
}

func (r *ReqItem) SetString(key, value string) {
	proto := &p_common.StringMsg{
		Value: value,
	}
	any, err := ptypes.MarshalAny(proto)
	if err != nil {
		logrus.Error("UnmarshalAny error :", err)
		return
	}
	r.items[key] = any
}

func (r *ReqItem) SetInt(key string, value int32) {
	proto := &p_common.IntMsg{
		Value: value,
	}
	any, err := ptypes.MarshalAny(proto)
	if err != nil {
		logrus.Error("UnmarshalAny error :", err)
		return
	}
	r.items[key] = any
}

func (r *ReqItem) GetStrings(key string) []string {
	o := r.GetString(key)
	if o == "" {
		return []string{}
	}

	s := strings.Split(o, ",")
	var is []string
	for _, sv := range s {
		// if sv == "" {
		// 	continue
		// }
		is = append(is, sv)
	}

	return is
}

type RspItem struct {
	Cmd   string
	Items map[string]interface{}
}

func (r *RspItem) Add(key string, obj interface{}) {
	_, ok := r.Items[key]
	if ok {
		return
	}
	r.Items[key] = obj
}

//ReqContext request
type ReqContext struct {
	Cmd           string
	Context       context.Context
	RequestItems  *ReqItem
	ResponseItems []*RspItem
	Publisher     micro.Publisher
	RspType       ResponseType
}

func (r *ReqContext) ToAny(cmd string) []*p_common.MessageItem {
	var rsp []*p_common.MessageItem
	for _, item := range r.ResponseItems {
		var reqItems []*p_common.MessageContent
		for key, v := range item.Items {
			msg, ok := v.(proto.Message)
			if !ok {
				continue
			}
			anyItems, err := ptypes.MarshalAny(msg)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"cmd": cmd,
					"key": key,
					"v":   v,
				}).Error("ToAny error :", err)
				continue
			}
			reqItems = append(reqItems, &p_common.MessageContent{
				Key: key,
				Obj: anyItems,
			})
		}
		rsp = append(rsp, &p_common.MessageItem{
			Cmd:       item.Cmd,
			Parameter: reqItems,
		})
	}

	return rsp
}

func (r *ReqContext) GetResponse() string {
	if r.RspType == RspJson {
		return r.toJson()
	} else if r.RspType == RspXml {
		return r.toXml()
	} else {
		return r.toText()
	}
}

func (r *ReqContext) toJson() string {
	b, err := json.Marshal(r.ResponseItems)
	if err != nil {
		logrus.Error("ToJson error :", err)
		return ""
	}
	return string(b)
}

func (r *ReqContext) toXml() string {
	if len(r.ResponseItems) <= 0 {
		return ""
	}
	item := r.ResponseItems[0].Items[""]
	// b, err := xml.Marshal(item)
	// if err != nil {
	// 	logrus.Error("ToXml error :", err)
	// 	fmt.Println("toXml err:", err)
	// 	return ""
	// }
	str, ok := item.(string)
	if !ok {
		return ""
	}
	return str
}

func (r *ReqContext) toText() string {
	if len(r.ResponseItems) <= 0 {
		return ""
	}
	item := r.ResponseItems[0].Items[""]
	return item.(string)
}

func (r *ReqContext) AddResponse(key string, obj interface{}) {
	response := r.ResponseItems[0]
	response.Add(key, obj)
}

func (r *ReqContext) AddNotify(cmd string, obj interface{}) {
	rspItem := &RspItem{
		Cmd:   cmd,
		Items: make(map[string]interface{}),
	}
	rspItem.Items["msg"] = obj
	r.ResponseItems = append(r.ResponseItems, rspItem)
}

func (r *ReqContext) SetText(content string) {
	r.RspType = RspText
	response := r.ResponseItems[0]
	response.Add("", content)
}

//NewReqContext request
func NewReqContext(cmd string, cxt context.Context, items []*p_common.MessageContent) *ReqContext {
	req := &ReqContext{
		Context: cxt,
		Cmd:     cmd,
	}
	reqItem := &ReqItem{
		items: make(map[string]*any.Any, len(items)),
	}
	for _, val := range items {
		reqItem.items[val.Key] = val.Obj
	}

	rspItem := &RspItem{
		Cmd:   cmd,
		Items: make(map[string]interface{}, len(items)),
	}

	req.RequestItems = reqItem
	req.ResponseItems = append(req.ResponseItems, rspItem)
	return req
}
