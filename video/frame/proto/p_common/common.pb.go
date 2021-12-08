// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package p_common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestType int32

const (
	RequestType_ReqType_Default RequestType = 0
	RequestType_ReqType_Json    RequestType = 1
)

var RequestType_name = map[int32]string{
	0: "ReqType_Default",
	1: "ReqType_Json",
}

var RequestType_value = map[string]int32{
	"ReqType_Default": 0,
	"ReqType_Json":    1,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}

func (RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type StatusType int32

const (
	StatusType_status_normal   StatusType = 0
	StatusType_status_disabled StatusType = 1
)

var StatusType_name = map[int32]string{
	0: "status_normal",
	1: "status_disabled",
}

var StatusType_value = map[string]int32{
	"status_normal":   0,
	"status_disabled": 1,
}

func (x StatusType) String() string {
	return proto.EnumName(StatusType_name, int32(x))
}

func (StatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type NoticeType int32

const (
	NoticeType_notice_reload NoticeType = 0
	NoticeType_notice_event  NoticeType = 1
)

var NoticeType_name = map[int32]string{
	0: "notice_reload",
	1: "notice_event",
}

var NoticeType_value = map[string]int32{
	"notice_reload": 0,
	"notice_event":  1,
}

func (x NoticeType) String() string {
	return proto.EnumName(NoticeType_name, int32(x))
}

func (NoticeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type RequestCode int32

const (
	//成功
	RequestCode_RequestCodeSuccess          RequestCode = 0
	RequestCode_RequestCodeErrorVerifi      RequestCode = 1
	RequestCode_RequestCodeErrorParameter   RequestCode = 2
	RequestCode_RequestCodeErrorPermissions RequestCode = 3
	RequestCode_RequestCodeErrorSystem      RequestCode = 4
	RequestCode_RequestCodeErrorAction      RequestCode = 5
	RequestCode_RequestCodeErrorSession     RequestCode = 6
)

var RequestCode_name = map[int32]string{
	0: "RequestCodeSuccess",
	1: "RequestCodeErrorVerifi",
	2: "RequestCodeErrorParameter",
	3: "RequestCodeErrorPermissions",
	4: "RequestCodeErrorSystem",
	5: "RequestCodeErrorAction",
	6: "RequestCodeErrorSession",
}

var RequestCode_value = map[string]int32{
	"RequestCodeSuccess":          0,
	"RequestCodeErrorVerifi":      1,
	"RequestCodeErrorParameter":   2,
	"RequestCodeErrorPermissions": 3,
	"RequestCodeErrorSystem":      4,
	"RequestCodeErrorAction":      5,
	"RequestCodeErrorSession":     6,
}

func (x RequestCode) String() string {
	return proto.EnumName(RequestCode_name, int32(x))
}

func (RequestCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

//运营渠道
type Operator int32

const (
	Operator_Operator_default Operator = 0
)

var Operator_name = map[int32]string{
	0: "Operator_default",
}

var Operator_value = map[string]int32{
	"Operator_default": 0,
}

func (x Operator) String() string {
	return proto.EnumName(Operator_name, int32(x))
}

func (Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

//设备
type OsType int32

const (
	OsType_os_none    OsType = 0
	OsType_os_android OsType = 1
	OsType_os_ios     OsType = 2
)

var OsType_name = map[int32]string{
	0: "os_none",
	1: "os_android",
	2: "os_ios",
}

var OsType_value = map[string]int32{
	"os_none":    0,
	"os_android": 1,
	"os_ios":     2,
}

func (x OsType) String() string {
	return proto.EnumName(OsType_name, int32(x))
}

func (OsType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

type RequestHead struct {
	Type                 RequestType  `protobuf:"varint,1,opt,name=type,proto3,enum=p_common.RequestType" json:"type,omitempty"`
	Request              *MessageItem `protobuf:"bytes,2,opt,name=Request,proto3" json:"Request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RequestHead) Reset()         { *m = RequestHead{} }
func (m *RequestHead) String() string { return proto.CompactTextString(m) }
func (*RequestHead) ProtoMessage()    {}
func (*RequestHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *RequestHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestHead.Unmarshal(m, b)
}
func (m *RequestHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestHead.Marshal(b, m, deterministic)
}
func (m *RequestHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestHead.Merge(m, src)
}
func (m *RequestHead) XXX_Size() int {
	return xxx_messageInfo_RequestHead.Size(m)
}
func (m *RequestHead) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestHead.DiscardUnknown(m)
}

var xxx_messageInfo_RequestHead proto.InternalMessageInfo

func (m *RequestHead) GetType() RequestType {
	if m != nil {
		return m.Type
	}
	return RequestType_ReqType_Default
}

func (m *RequestHead) GetRequest() *MessageItem {
	if m != nil {
		return m.Request
	}
	return nil
}

type MessageItem struct {
	Cmd                  string            `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Parameter            []*MessageContent `protobuf:"bytes,2,rep,name=parameter,proto3" json:"parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MessageItem) Reset()         { *m = MessageItem{} }
func (m *MessageItem) String() string { return proto.CompactTextString(m) }
func (*MessageItem) ProtoMessage()    {}
func (*MessageItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *MessageItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageItem.Unmarshal(m, b)
}
func (m *MessageItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageItem.Marshal(b, m, deterministic)
}
func (m *MessageItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageItem.Merge(m, src)
}
func (m *MessageItem) XXX_Size() int {
	return xxx_messageInfo_MessageItem.Size(m)
}
func (m *MessageItem) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageItem.DiscardUnknown(m)
}

var xxx_messageInfo_MessageItem proto.InternalMessageInfo

func (m *MessageItem) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

func (m *MessageItem) GetParameter() []*MessageContent {
	if m != nil {
		return m.Parameter
	}
	return nil
}

type ResponseHead struct {
	Code                 RequestCode    `protobuf:"varint,1,opt,name=code,proto3,enum=p_common.RequestCode" json:"code,omitempty"`
	ErrCode              int32          `protobuf:"varint,2,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	Msg                  string         `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Response             []*MessageItem `protobuf:"bytes,4,rep,name=Response,proto3" json:"Response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResponseHead) Reset()         { *m = ResponseHead{} }
func (m *ResponseHead) String() string { return proto.CompactTextString(m) }
func (*ResponseHead) ProtoMessage()    {}
func (*ResponseHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *ResponseHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseHead.Unmarshal(m, b)
}
func (m *ResponseHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseHead.Marshal(b, m, deterministic)
}
func (m *ResponseHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseHead.Merge(m, src)
}
func (m *ResponseHead) XXX_Size() int {
	return xxx_messageInfo_ResponseHead.Size(m)
}
func (m *ResponseHead) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseHead.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseHead proto.InternalMessageInfo

func (m *ResponseHead) GetCode() RequestCode {
	if m != nil {
		return m.Code
	}
	return RequestCode_RequestCodeSuccess
}

func (m *ResponseHead) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResponseHead) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *ResponseHead) GetResponse() []*MessageItem {
	if m != nil {
		return m.Response
	}
	return nil
}

type MessageContent struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Obj                  *any.Any `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageContent) Reset()         { *m = MessageContent{} }
func (m *MessageContent) String() string { return proto.CompactTextString(m) }
func (*MessageContent) ProtoMessage()    {}
func (*MessageContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *MessageContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageContent.Unmarshal(m, b)
}
func (m *MessageContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageContent.Marshal(b, m, deterministic)
}
func (m *MessageContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageContent.Merge(m, src)
}
func (m *MessageContent) XXX_Size() int {
	return xxx_messageInfo_MessageContent.Size(m)
}
func (m *MessageContent) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageContent.DiscardUnknown(m)
}

var xxx_messageInfo_MessageContent proto.InternalMessageInfo

func (m *MessageContent) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *MessageContent) GetObj() *any.Any {
	if m != nil {
		return m.Obj
	}
	return nil
}

type IntMsg struct {
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntMsg) Reset()         { *m = IntMsg{} }
func (m *IntMsg) String() string { return proto.CompactTextString(m) }
func (*IntMsg) ProtoMessage()    {}
func (*IntMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *IntMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMsg.Unmarshal(m, b)
}
func (m *IntMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMsg.Marshal(b, m, deterministic)
}
func (m *IntMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMsg.Merge(m, src)
}
func (m *IntMsg) XXX_Size() int {
	return xxx_messageInfo_IntMsg.Size(m)
}
func (m *IntMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMsg.DiscardUnknown(m)
}

var xxx_messageInfo_IntMsg proto.InternalMessageInfo

func (m *IntMsg) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type IntsMsg struct {
	Value                []int32  `protobuf:"varint,2,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntsMsg) Reset()         { *m = IntsMsg{} }
func (m *IntsMsg) String() string { return proto.CompactTextString(m) }
func (*IntsMsg) ProtoMessage()    {}
func (*IntsMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *IntsMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntsMsg.Unmarshal(m, b)
}
func (m *IntsMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntsMsg.Marshal(b, m, deterministic)
}
func (m *IntsMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntsMsg.Merge(m, src)
}
func (m *IntsMsg) XXX_Size() int {
	return xxx_messageInfo_IntsMsg.Size(m)
}
func (m *IntsMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_IntsMsg.DiscardUnknown(m)
}

var xxx_messageInfo_IntsMsg proto.InternalMessageInfo

func (m *IntsMsg) GetValue() []int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type StringMsg struct {
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringMsg) Reset()         { *m = StringMsg{} }
func (m *StringMsg) String() string { return proto.CompactTextString(m) }
func (*StringMsg) ProtoMessage()    {}
func (*StringMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *StringMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringMsg.Unmarshal(m, b)
}
func (m *StringMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringMsg.Marshal(b, m, deterministic)
}
func (m *StringMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringMsg.Merge(m, src)
}
func (m *StringMsg) XXX_Size() int {
	return xxx_messageInfo_StringMsg.Size(m)
}
func (m *StringMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StringMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StringMsg proto.InternalMessageInfo

func (m *StringMsg) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StringsMsg struct {
	Value                []string `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringsMsg) Reset()         { *m = StringsMsg{} }
func (m *StringsMsg) String() string { return proto.CompactTextString(m) }
func (*StringsMsg) ProtoMessage()    {}
func (*StringsMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

func (m *StringsMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringsMsg.Unmarshal(m, b)
}
func (m *StringsMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringsMsg.Marshal(b, m, deterministic)
}
func (m *StringsMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringsMsg.Merge(m, src)
}
func (m *StringsMsg) XXX_Size() int {
	return xxx_messageInfo_StringsMsg.Size(m)
}
func (m *StringsMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StringsMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StringsMsg proto.InternalMessageInfo

func (m *StringsMsg) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type RspJsonHead struct {
	Code                 RequestCode `protobuf:"varint,1,opt,name=code,proto3,enum=p_common.RequestCode" json:"code,omitempty"`
	ErrCode              int32       `protobuf:"varint,2,opt,name=err_code,json=errCode,proto3" json:"err_code,omitempty"`
	Msg                  string      `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	Response             string      `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RspJsonHead) Reset()         { *m = RspJsonHead{} }
func (m *RspJsonHead) String() string { return proto.CompactTextString(m) }
func (*RspJsonHead) ProtoMessage()    {}
func (*RspJsonHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{8}
}

func (m *RspJsonHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspJsonHead.Unmarshal(m, b)
}
func (m *RspJsonHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspJsonHead.Marshal(b, m, deterministic)
}
func (m *RspJsonHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspJsonHead.Merge(m, src)
}
func (m *RspJsonHead) XXX_Size() int {
	return xxx_messageInfo_RspJsonHead.Size(m)
}
func (m *RspJsonHead) XXX_DiscardUnknown() {
	xxx_messageInfo_RspJsonHead.DiscardUnknown(m)
}

var xxx_messageInfo_RspJsonHead proto.InternalMessageInfo

func (m *RspJsonHead) GetCode() RequestCode {
	if m != nil {
		return m.Code
	}
	return RequestCode_RequestCodeSuccess
}

func (m *RspJsonHead) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *RspJsonHead) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RspJsonHead) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type DeviceMsg struct {
	OsType               OsType   `protobuf:"varint,1,opt,name=os_type,json=osType,proto3,enum=p_common.OsType" json:"os_type,omitempty"`
	Manufacturer         string   `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	Model                string   `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	MacAddress           string   `protobuf:"bytes,4,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	Imei                 string   `protobuf:"bytes,5,opt,name=imei,proto3" json:"imei,omitempty"`
	Imsi                 string   `protobuf:"bytes,6,opt,name=imsi,proto3" json:"imsi,omitempty"`
	Sn                   string   `protobuf:"bytes,7,opt,name=sn,proto3" json:"sn,omitempty"`
	Brand                string   `protobuf:"bytes,8,opt,name=brand,proto3" json:"brand,omitempty"`
	AndroidId            string   `protobuf:"bytes,9,opt,name=android_id,json=androidId,proto3" json:"android_id,omitempty"`
	Version              string   `protobuf:"bytes,10,opt,name=version,proto3" json:"version,omitempty"`
	SdkVersion           string   `protobuf:"bytes,11,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`
	PackageName          string   `protobuf:"bytes,12,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	PackageVersion       string   `protobuf:"bytes,13,opt,name=package_version,json=packageVersion,proto3" json:"package_version,omitempty"`
	Identity             string   `protobuf:"bytes,14,opt,name=identity,proto3" json:"identity,omitempty"`
	OperatorId           Operator `protobuf:"varint,15,opt,name=operator_id,json=operatorId,proto3,enum=p_common.Operator" json:"operator_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceMsg) Reset()         { *m = DeviceMsg{} }
func (m *DeviceMsg) String() string { return proto.CompactTextString(m) }
func (*DeviceMsg) ProtoMessage()    {}
func (*DeviceMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{9}
}

func (m *DeviceMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceMsg.Unmarshal(m, b)
}
func (m *DeviceMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceMsg.Marshal(b, m, deterministic)
}
func (m *DeviceMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceMsg.Merge(m, src)
}
func (m *DeviceMsg) XXX_Size() int {
	return xxx_messageInfo_DeviceMsg.Size(m)
}
func (m *DeviceMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceMsg.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceMsg proto.InternalMessageInfo

func (m *DeviceMsg) GetOsType() OsType {
	if m != nil {
		return m.OsType
	}
	return OsType_os_none
}

func (m *DeviceMsg) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *DeviceMsg) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *DeviceMsg) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *DeviceMsg) GetImei() string {
	if m != nil {
		return m.Imei
	}
	return ""
}

func (m *DeviceMsg) GetImsi() string {
	if m != nil {
		return m.Imsi
	}
	return ""
}

func (m *DeviceMsg) GetSn() string {
	if m != nil {
		return m.Sn
	}
	return ""
}

func (m *DeviceMsg) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *DeviceMsg) GetAndroidId() string {
	if m != nil {
		return m.AndroidId
	}
	return ""
}

func (m *DeviceMsg) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DeviceMsg) GetSdkVersion() string {
	if m != nil {
		return m.SdkVersion
	}
	return ""
}

func (m *DeviceMsg) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *DeviceMsg) GetPackageVersion() string {
	if m != nil {
		return m.PackageVersion
	}
	return ""
}

func (m *DeviceMsg) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *DeviceMsg) GetOperatorId() Operator {
	if m != nil {
		return m.OperatorId
	}
	return Operator_Operator_default
}

type PageInfo struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Total                int32    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Type                 int32    `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageInfo) Reset()         { *m = PageInfo{} }
func (m *PageInfo) String() string { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()    {}
func (*PageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{10}
}

func (m *PageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageInfo.Unmarshal(m, b)
}
func (m *PageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageInfo.Marshal(b, m, deterministic)
}
func (m *PageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageInfo.Merge(m, src)
}
func (m *PageInfo) XXX_Size() int {
	return xxx_messageInfo_PageInfo.Size(m)
}
func (m *PageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PageInfo proto.InternalMessageInfo

func (m *PageInfo) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *PageInfo) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *PageInfo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PageInfo) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type VersionItem struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OsType               OsType     `protobuf:"varint,2,opt,name=os_type,json=osType,proto3,enum=p_common.OsType" json:"os_type,omitempty"`
	AppId                int32      `protobuf:"varint,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	PackageVersion       string     `protobuf:"bytes,4,opt,name=package_version,json=packageVersion,proto3" json:"package_version,omitempty"`
	ChannelId            int32      `protobuf:"varint,5,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	VersionSize          string     `protobuf:"bytes,6,opt,name=version_size,json=versionSize,proto3" json:"version_size,omitempty"`
	Status               StatusType `protobuf:"varint,7,opt,name=status,proto3,enum=p_common.StatusType" json:"status,omitempty"`
	DownloadUrl          string     `protobuf:"bytes,8,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
	Type                 int32      `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	UpdateId             int32      `protobuf:"varint,10,opt,name=update_id,json=updateId,proto3" json:"update_id,omitempty"`
	UpdateTitle          string     `protobuf:"bytes,11,opt,name=update_title,json=updateTitle,proto3" json:"update_title,omitempty"`
	UpdateMessage        string     `protobuf:"bytes,12,opt,name=update_message,json=updateMessage,proto3" json:"update_message,omitempty"`
	CreateTime           string     `protobuf:"bytes,13,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime           string     `protobuf:"bytes,14,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	FileDomain           string     `protobuf:"bytes,15,opt,name=file_domain,json=fileDomain,proto3" json:"file_domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *VersionItem) Reset()         { *m = VersionItem{} }
func (m *VersionItem) String() string { return proto.CompactTextString(m) }
func (*VersionItem) ProtoMessage()    {}
func (*VersionItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{11}
}

func (m *VersionItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionItem.Unmarshal(m, b)
}
func (m *VersionItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionItem.Marshal(b, m, deterministic)
}
func (m *VersionItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionItem.Merge(m, src)
}
func (m *VersionItem) XXX_Size() int {
	return xxx_messageInfo_VersionItem.Size(m)
}
func (m *VersionItem) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionItem.DiscardUnknown(m)
}

var xxx_messageInfo_VersionItem proto.InternalMessageInfo

func (m *VersionItem) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VersionItem) GetOsType() OsType {
	if m != nil {
		return m.OsType
	}
	return OsType_os_none
}

func (m *VersionItem) GetAppId() int32 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *VersionItem) GetPackageVersion() string {
	if m != nil {
		return m.PackageVersion
	}
	return ""
}

func (m *VersionItem) GetChannelId() int32 {
	if m != nil {
		return m.ChannelId
	}
	return 0
}

func (m *VersionItem) GetVersionSize() string {
	if m != nil {
		return m.VersionSize
	}
	return ""
}

func (m *VersionItem) GetStatus() StatusType {
	if m != nil {
		return m.Status
	}
	return StatusType_status_normal
}

func (m *VersionItem) GetDownloadUrl() string {
	if m != nil {
		return m.DownloadUrl
	}
	return ""
}

func (m *VersionItem) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *VersionItem) GetUpdateId() int32 {
	if m != nil {
		return m.UpdateId
	}
	return 0
}

func (m *VersionItem) GetUpdateTitle() string {
	if m != nil {
		return m.UpdateTitle
	}
	return ""
}

func (m *VersionItem) GetUpdateMessage() string {
	if m != nil {
		return m.UpdateMessage
	}
	return ""
}

func (m *VersionItem) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *VersionItem) GetUpdateTime() string {
	if m != nil {
		return m.UpdateTime
	}
	return ""
}

func (m *VersionItem) GetFileDomain() string {
	if m != nil {
		return m.FileDomain
	}
	return ""
}

type VersionList struct {
	Items                []*VersionItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *VersionList) Reset()         { *m = VersionList{} }
func (m *VersionList) String() string { return proto.CompactTextString(m) }
func (*VersionList) ProtoMessage()    {}
func (*VersionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{12}
}

func (m *VersionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionList.Unmarshal(m, b)
}
func (m *VersionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionList.Marshal(b, m, deterministic)
}
func (m *VersionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionList.Merge(m, src)
}
func (m *VersionList) XXX_Size() int {
	return xxx_messageInfo_VersionList.Size(m)
}
func (m *VersionList) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionList.DiscardUnknown(m)
}

var xxx_messageInfo_VersionList proto.InternalMessageInfo

func (m *VersionList) GetItems() []*VersionItem {
	if m != nil {
		return m.Items
	}
	return nil
}

//通知消息
type NoticeMsg struct {
	NoticeType           NoticeType `protobuf:"varint,1,opt,name=notice_type,json=noticeType,proto3,enum=p_common.NoticeType" json:"notice_type,omitempty"`
	Obj                  *any.Any   `protobuf:"bytes,2,opt,name=obj,proto3" json:"obj,omitempty"`
	Event                string     `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	ServiceType          int32      `protobuf:"varint,4,opt,name=service_type,json=serviceType,proto3" json:"service_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NoticeMsg) Reset()         { *m = NoticeMsg{} }
func (m *NoticeMsg) String() string { return proto.CompactTextString(m) }
func (*NoticeMsg) ProtoMessage()    {}
func (*NoticeMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{13}
}

func (m *NoticeMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NoticeMsg.Unmarshal(m, b)
}
func (m *NoticeMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NoticeMsg.Marshal(b, m, deterministic)
}
func (m *NoticeMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NoticeMsg.Merge(m, src)
}
func (m *NoticeMsg) XXX_Size() int {
	return xxx_messageInfo_NoticeMsg.Size(m)
}
func (m *NoticeMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_NoticeMsg.DiscardUnknown(m)
}

var xxx_messageInfo_NoticeMsg proto.InternalMessageInfo

func (m *NoticeMsg) GetNoticeType() NoticeType {
	if m != nil {
		return m.NoticeType
	}
	return NoticeType_notice_reload
}

func (m *NoticeMsg) GetObj() *any.Any {
	if m != nil {
		return m.Obj
	}
	return nil
}

func (m *NoticeMsg) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *NoticeMsg) GetServiceType() int32 {
	if m != nil {
		return m.ServiceType
	}
	return 0
}

func init() {
	proto.RegisterEnum("p_common.RequestType", RequestType_name, RequestType_value)
	proto.RegisterEnum("p_common.StatusType", StatusType_name, StatusType_value)
	proto.RegisterEnum("p_common.NoticeType", NoticeType_name, NoticeType_value)
	proto.RegisterEnum("p_common.RequestCode", RequestCode_name, RequestCode_value)
	proto.RegisterEnum("p_common.Operator", Operator_name, Operator_value)
	proto.RegisterEnum("p_common.OsType", OsType_name, OsType_value)
	proto.RegisterType((*RequestHead)(nil), "p_common.RequestHead")
	proto.RegisterType((*MessageItem)(nil), "p_common.MessageItem")
	proto.RegisterType((*ResponseHead)(nil), "p_common.ResponseHead")
	proto.RegisterType((*MessageContent)(nil), "p_common.MessageContent")
	proto.RegisterType((*IntMsg)(nil), "p_common.IntMsg")
	proto.RegisterType((*IntsMsg)(nil), "p_common.IntsMsg")
	proto.RegisterType((*StringMsg)(nil), "p_common.StringMsg")
	proto.RegisterType((*StringsMsg)(nil), "p_common.StringsMsg")
	proto.RegisterType((*RspJsonHead)(nil), "p_common.RspJsonHead")
	proto.RegisterType((*DeviceMsg)(nil), "p_common.DeviceMsg")
	proto.RegisterType((*PageInfo)(nil), "p_common.PageInfo")
	proto.RegisterType((*VersionItem)(nil), "p_common.VersionItem")
	proto.RegisterType((*VersionList)(nil), "p_common.VersionList")
	proto.RegisterType((*NoticeMsg)(nil), "p_common.NoticeMsg")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcd, 0x72, 0x1b, 0x45,
	0x10, 0xf6, 0x4a, 0x96, 0x2c, 0xf5, 0x3a, 0xca, 0x32, 0x38, 0x61, 0xe3, 0x54, 0x88, 0xb3, 0x55,
	0x40, 0x62, 0x28, 0xa7, 0xf2, 0x03, 0x07, 0x6e, 0xa9, 0x84, 0x2a, 0x94, 0x22, 0x21, 0xb5, 0x0e,
	0xe1, 0xc6, 0xd6, 0x78, 0xa7, 0x25, 0x06, 0xef, 0xce, 0x6c, 0x66, 0x46, 0xa6, 0x94, 0x1b, 0x8f,
	0xc1, 0x99, 0x77, 0xca, 0xbb, 0x70, 0xa3, 0xe6, 0x4f, 0x52, 0x62, 0xa7, 0xe0, 0xc2, 0xc9, 0xd3,
	0x5f, 0x7f, 0xd3, 0xdd, 0xd3, 0xfd, 0xb9, 0x57, 0xb0, 0x5b, 0xcb, 0xb6, 0x95, 0xe2, 0xa8, 0x53,
	0xd2, 0x48, 0x32, 0xea, 0x2a, 0x6f, 0xef, 0x5f, 0x9b, 0x4b, 0x39, 0x6f, 0xf0, 0xae, 0xc3, 0x4f,
	0x16, 0xb3, 0xbb, 0x54, 0x2c, 0x3d, 0xa9, 0xe0, 0x90, 0x96, 0xf8, 0x7a, 0x81, 0xda, 0x7c, 0x8f,
	0x94, 0x91, 0x3b, 0xb0, 0x6d, 0x96, 0x1d, 0xe6, 0xc9, 0x41, 0x72, 0x7b, 0x72, 0xff, 0xca, 0x51,
	0x0c, 0x71, 0x14, 0x48, 0x2f, 0x97, 0x1d, 0x96, 0x8e, 0x42, 0xee, 0xc2, 0x4e, 0x00, 0xf3, 0xde,
	0x41, 0x72, 0x3b, 0xdd, 0x64, 0x3f, 0x43, 0xad, 0xe9, 0x1c, 0xa7, 0x06, 0xdb, 0x32, 0xb2, 0x8a,
	0x9f, 0x21, 0xdd, 0xc0, 0x49, 0x06, 0xfd, 0xba, 0x65, 0x2e, 0xd3, 0xb8, 0xb4, 0x47, 0xf2, 0x0d,
	0x8c, 0x3b, 0xaa, 0x68, 0x8b, 0x06, 0x55, 0xde, 0x3b, 0xe8, 0xdf, 0x4e, 0xef, 0xe7, 0xe7, 0x62,
	0x3e, 0x96, 0xc2, 0xa0, 0x30, 0xe5, 0x9a, 0x5a, 0xfc, 0x99, 0xc0, 0x6e, 0x89, 0xba, 0x93, 0x42,
	0x63, 0x7c, 0x45, 0x2d, 0xd9, 0x87, 0x5f, 0xf1, 0x58, 0x32, 0x2c, 0x1d, 0x85, 0x5c, 0x83, 0x11,
	0x2a, 0x55, 0x39, 0xba, 0x7d, 0xc6, 0xa0, 0xdc, 0x41, 0xa5, 0x2c, 0xc1, 0x16, 0xd8, 0xea, 0x79,
	0xde, 0xf7, 0x05, 0xb6, 0x7a, 0x4e, 0xee, 0xc1, 0x28, 0xe6, 0xc9, 0xb7, 0x5d, 0x7d, 0x1f, 0x78,
	0xf3, 0x8a, 0x56, 0x3c, 0x85, 0xc9, 0xbb, 0x85, 0xdb, 0xb0, 0xa7, 0xb8, 0x8c, 0xef, 0x3e, 0xc5,
	0x25, 0xf9, 0x1c, 0xfa, 0xf2, 0xe4, 0xb7, 0xd0, 0xc5, 0xbd, 0x23, 0x3f, 0xac, 0xa3, 0x38, 0xac,
	0xa3, 0x47, 0x62, 0x59, 0x5a, 0x42, 0xf1, 0x29, 0x0c, 0xa7, 0xc2, 0x3c, 0xd3, 0x73, 0xb2, 0x07,
	0x83, 0x33, 0xda, 0x2c, 0x62, 0xc9, 0xde, 0x28, 0x6e, 0xc2, 0xce, 0x54, 0x18, 0xfd, 0x1e, 0xa1,
	0xbf, 0x26, 0xdc, 0x82, 0xf1, 0xb1, 0x51, 0x5c, 0xcc, 0xcf, 0xc5, 0x18, 0x47, 0x4a, 0x01, 0xe0,
	0x29, 0xe7, 0xc3, 0xac, 0x38, 0x7f, 0x24, 0x90, 0x96, 0xba, 0x7b, 0xaa, 0xa5, 0xf8, 0x5f, 0xdb,
	0xbd, 0x0f, 0x23, 0xb5, 0x6e, 0xb7, 0x85, 0x57, 0x76, 0xf1, 0xb6, 0x0f, 0xe3, 0x27, 0x78, 0xc6,
	0x6b, 0xb4, 0x75, 0xde, 0x81, 0x1d, 0xa9, 0xab, 0x0d, 0xe5, 0x66, 0xeb, 0x22, 0x7e, 0xd4, 0x4e,
	0xb4, 0x43, 0xe9, 0xfe, 0x92, 0x02, 0x76, 0x5b, 0x2a, 0x16, 0x33, 0x5a, 0x9b, 0x85, 0x72, 0x3a,
	0xb3, 0x81, 0xdf, 0xc1, 0xec, 0xb3, 0x5b, 0xc9, 0xb0, 0x09, 0xc5, 0x78, 0x83, 0xdc, 0x84, 0xb4,
	0xa5, 0x75, 0x45, 0x19, 0x53, 0xa8, 0x75, 0xa8, 0x08, 0x5a, 0x5a, 0x3f, 0xf2, 0x08, 0x21, 0xb0,
	0xcd, 0x5b, 0xe4, 0xf9, 0xc0, 0x79, 0xdc, 0xd9, 0x63, 0x9a, 0xe7, 0xc3, 0x88, 0x69, 0x4e, 0x26,
	0xd0, 0xd3, 0x22, 0xdf, 0x71, 0x48, 0x4f, 0x0b, 0x9b, 0xee, 0x44, 0x51, 0xc1, 0xf2, 0x91, 0x4f,
	0xe7, 0x0c, 0x72, 0x03, 0x80, 0x0a, 0xa6, 0x24, 0x67, 0x15, 0x67, 0xf9, 0xd8, 0xb9, 0xc6, 0x01,
	0x99, 0x32, 0x92, 0xc3, 0xce, 0x19, 0x2a, 0xcd, 0xa5, 0xc8, 0xc1, 0xf9, 0xa2, 0x69, 0xeb, 0xd4,
	0xec, 0xb4, 0x8a, 0xde, 0xd4, 0xd7, 0xa9, 0xd9, 0xe9, 0xab, 0x40, 0xb8, 0x05, 0xbb, 0x1d, 0xad,
	0x4f, 0xe9, 0x1c, 0x2b, 0x41, 0x5b, 0xcc, 0x77, 0x1d, 0x23, 0x0d, 0xd8, 0x73, 0xda, 0x22, 0xf9,
	0x02, 0x2e, 0x47, 0x4a, 0x8c, 0x73, 0xc9, 0xb1, 0x26, 0x01, 0x8e, 0xb1, 0xf6, 0x61, 0xc4, 0x19,
	0x0a, 0xc3, 0xcd, 0x32, 0x9f, 0xf8, 0x19, 0x45, 0x9b, 0x3c, 0x80, 0x54, 0x76, 0xa8, 0xa8, 0x91,
	0xca, 0x3e, 0xe1, 0xb2, 0x9b, 0x0c, 0xd9, 0x98, 0x4c, 0x70, 0x96, 0x10, 0x69, 0x53, 0x56, 0xfc,
	0x02, 0xa3, 0x17, 0xf6, 0xdf, 0x48, 0xcc, 0xa4, 0x6d, 0x5e, 0x47, 0xe7, 0x7e, 0xa6, 0x83, 0xd2,
	0x9d, 0x6d, 0xb3, 0x8c, 0x34, 0xb4, 0x89, 0xd2, 0x77, 0x86, 0x43, 0xb9, 0x69, 0x30, 0x4e, 0xcc,
	0x19, 0xf6, 0xbe, 0xd3, 0xc4, 0xb6, 0xbf, 0x6f, 0xcf, 0xc5, 0xdf, 0x7d, 0x48, 0x43, 0xf1, 0x6e,
	0x0d, 0x4d, 0xa0, 0xc7, 0x59, 0xc8, 0xd0, 0xe3, 0x6c, 0x53, 0x4a, 0xbd, 0x7f, 0x91, 0xd2, 0x15,
	0x18, 0xd2, 0xae, 0xb3, 0x4f, 0xeb, 0xfb, 0x5a, 0x68, 0xd7, 0x4d, 0xd9, 0x45, 0xbd, 0xdb, 0xbe,
	0xb0, 0x77, 0x37, 0x00, 0xea, 0x5f, 0xa9, 0x10, 0xd8, 0xd8, 0x18, 0x03, 0x17, 0x63, 0x1c, 0x90,
	0x29, 0xb3, 0x63, 0x0a, 0xf7, 0x2b, 0xcd, 0xdf, 0x60, 0x90, 0x50, 0x1a, 0xb0, 0x63, 0xfe, 0x06,
	0xc9, 0x57, 0x30, 0xd4, 0x86, 0x9a, 0x85, 0x76, 0x6a, 0x9a, 0xdc, 0xdf, 0x5b, 0xd7, 0x7a, 0xec,
	0x70, 0x5f, 0xaf, 0xe7, 0xd8, 0x80, 0x4c, 0xfe, 0x2e, 0x1a, 0x49, 0x59, 0xb5, 0x50, 0x4d, 0x90,
	0x5b, 0x1a, 0xb1, 0x9f, 0x54, 0xb3, 0xea, 0xd8, 0x78, 0xdd, 0x31, 0x72, 0x1d, 0xc6, 0x8b, 0x8e,
	0x51, 0x83, 0xb6, 0x4a, 0x70, 0x8e, 0x91, 0x07, 0x7c, 0x91, 0xc1, 0xe9, 0xfb, 0xef, 0xd5, 0x96,
	0x7a, 0xec, 0xa5, 0x9b, 0xc2, 0x67, 0x30, 0x09, 0x94, 0xd6, 0x6f, 0xc2, 0x20, 0xb8, 0x4b, 0x1e,
	0x0d, 0xeb, 0xd1, 0xca, 0xb6, 0x56, 0xe8, 0x23, 0xb5, 0x18, 0xe4, 0x06, 0x1e, 0x7a, 0xc9, 0x5b,
	0x47, 0x58, 0xa5, 0x6a, 0x31, 0xa8, 0x0d, 0x62, 0x26, 0x4f, 0x98, 0xf1, 0x06, 0x2b, 0x26, 0x5b,
	0xca, 0x85, 0xd3, 0xdb, 0xb8, 0x04, 0x0b, 0x3d, 0x71, 0x48, 0xf1, 0xed, 0x6a, 0xf4, 0x3f, 0x70,
	0x6d, 0xc8, 0x97, 0x30, 0xe0, 0x06, 0x5b, 0x9d, 0x27, 0xef, 0xef, 0xf2, 0x0d, 0x81, 0x94, 0x9e,
	0x53, 0xfc, 0x95, 0xc0, 0xf8, 0xb9, 0x34, 0x61, 0xe1, 0x7c, 0x0d, 0xa9, 0x70, 0xc6, 0xe6, 0xd2,
	0xd9, 0xe8, 0xbe, 0x67, 0xba, 0xee, 0x83, 0x58, 0x9d, 0xff, 0xeb, 0xa6, 0xb7, 0x72, 0xc6, 0x33,
	0x14, 0x26, 0xca, 0xd9, 0x19, 0xb6, 0xd7, 0x1a, 0xd5, 0xd9, 0x2a, 0xab, 0x97, 0x75, 0x1a, 0x30,
	0x9b, 0xe0, 0xf0, 0xe1, 0xea, 0x73, 0xee, 0xf2, 0x7d, 0x0c, 0x97, 0x4b, 0x7c, 0x6d, 0x8f, 0xd5,
	0x13, 0x9c, 0xd1, 0x45, 0x63, 0xb2, 0x2d, 0x92, 0xd9, 0xaf, 0xa5, 0x07, 0xed, 0x0a, 0xcf, 0x92,
	0xc3, 0x87, 0x76, 0xe9, 0x47, 0xb9, 0x90, 0x8f, 0xe0, 0x92, 0x17, 0x4c, 0x25, 0xa4, 0x6a, 0x69,
	0x93, 0x6d, 0xd9, 0x38, 0x01, 0x62, 0x5c, 0xd3, 0x93, 0x06, 0x59, 0x96, 0x1c, 0xde, 0x03, 0x58,
	0x3f, 0xd3, 0xde, 0x0a, 0x1d, 0x51, 0x68, 0xd5, 0xe4, 0x13, 0x05, 0xc8, 0xd5, 0x9f, 0x25, 0x87,
	0x6f, 0x93, 0x55, 0x7d, 0x6e, 0xe7, 0x5f, 0x05, 0xb2, 0x61, 0x1e, 0x2f, 0xea, 0x1a, 0xb5, 0xce,
	0xb6, 0xc8, 0x3e, 0x5c, 0xdd, 0xc0, 0xbf, 0x53, 0x4a, 0xaa, 0x57, 0xa8, 0xf8, 0x8c, 0x67, 0x09,
	0xb9, 0x01, 0xd7, 0xde, 0xf7, 0xbd, 0x88, 0x3f, 0x05, 0xb2, 0x1e, 0xb9, 0x09, 0xd7, 0xcf, 0xb9,
	0x51, 0xb5, 0x5c, 0xdb, 0x81, 0xea, 0xac, 0x7f, 0x51, 0xec, 0xe3, 0xa5, 0x36, 0xd8, 0x66, 0xdb,
	0x17, 0xf9, 0x1e, 0xd5, 0x86, 0x4b, 0x91, 0x0d, 0xc8, 0x75, 0xf8, 0xe4, 0xdc, 0x3d, 0x74, 0x51,
	0xb3, 0xe1, 0xe1, 0x01, 0x8c, 0xe2, 0x36, 0x23, 0x7b, 0x90, 0xc5, 0x73, 0xc5, 0x62, 0xd7, 0x0f,
	0xef, 0xc1, 0xd0, 0xaf, 0x0f, 0x92, 0xba, 0x0d, 0x23, 0xa4, 0xc0, 0x6c, 0x8b, 0x4c, 0x00, 0xa4,
	0xae, 0xc2, 0x5a, 0xcf, 0x12, 0x02, 0x30, 0x94, 0xba, 0xe2, 0x52, 0x67, 0xbd, 0x93, 0xa1, 0x13,
	0xc6, 0x83, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x35, 0xbf, 0x72, 0xd7, 0x09, 0x00, 0x00,
}