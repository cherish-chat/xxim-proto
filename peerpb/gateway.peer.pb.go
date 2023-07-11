// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: gateway.peer.proto

package peerpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//GatewayWriteDataType 写入数据类型
type GatewayWriteDataType int32

const (
	//返回响应
	GatewayWriteDataType_Response GatewayWriteDataType = 0
	//主动推送message
	GatewayWriteDataType_PushMessage GatewayWriteDataType = 1
	//主动推送notice
	GatewayWriteDataType_PushNotice GatewayWriteDataType = 2
)

// Enum value maps for GatewayWriteDataType.
var (
	GatewayWriteDataType_name = map[int32]string{
		0: "Response",
		1: "PushMessage",
		2: "PushNotice",
	}
	GatewayWriteDataType_value = map[string]int32{
		"Response":    0,
		"PushMessage": 1,
		"PushNotice":  2,
	}
)

func (x GatewayWriteDataType) Enum() *GatewayWriteDataType {
	p := new(GatewayWriteDataType)
	*p = x
	return p
}

func (x GatewayWriteDataType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GatewayWriteDataType) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_peer_proto_enumTypes[0].Descriptor()
}

func (GatewayWriteDataType) Type() protoreflect.EnumType {
	return &file_gateway_peer_proto_enumTypes[0]
}

func (x GatewayWriteDataType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GatewayWriteDataType.Descriptor instead.
func (GatewayWriteDataType) EnumDescriptor() ([]byte, []int) {
	return file_gateway_peer_proto_rawDescGZIP(), []int{0}
}

//ListLongConnectionReq 获取长连接列表请求
type ListLongConnectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *RequestHeader                `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Filter *ListLongConnectionReq_Filter `protobuf:"bytes,2,opt,name=filter,proto3,oneof" json:"filter,omitempty"`
}

func (x *ListLongConnectionReq) Reset() {
	*x = ListLongConnectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_peer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLongConnectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLongConnectionReq) ProtoMessage() {}

func (x *ListLongConnectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_peer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLongConnectionReq.ProtoReflect.Descriptor instead.
func (*ListLongConnectionReq) Descriptor() ([]byte, []int) {
	return file_gateway_peer_proto_rawDescGZIP(), []int{0}
}

func (x *ListLongConnectionReq) GetHeader() *RequestHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListLongConnectionReq) GetFilter() *ListLongConnectionReq_Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

//ListLongConnectionResp 获取长连接列表响应
type ListLongConnectionResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header          *ResponseHeader   `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	LongConnections []*LongConnection `protobuf:"bytes,2,rep,name=longConnections,proto3" json:"longConnections,omitempty"`
}

func (x *ListLongConnectionResp) Reset() {
	*x = ListLongConnectionResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_peer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLongConnectionResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLongConnectionResp) ProtoMessage() {}

func (x *ListLongConnectionResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_peer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLongConnectionResp.ProtoReflect.Descriptor instead.
func (*ListLongConnectionResp) Descriptor() ([]byte, []int) {
	return file_gateway_peer_proto_rawDescGZIP(), []int{1}
}

func (x *ListLongConnectionResp) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ListLongConnectionResp) GetLongConnections() []*LongConnection {
	if x != nil {
		return x.LongConnections
	}
	return nil
}

//GatewayWriteDataContent 写入数据内容
type GatewayWriteDataContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//写入数据类型
	DataType GatewayWriteDataType `protobuf:"varint,1,opt,name=dataType,proto3,enum=pb.GatewayWriteDataType" json:"dataType,omitempty"`
	//响应
	Response *GatewayApiResponse `protobuf:"bytes,2,opt,name=response,proto3,oneof" json:"response,omitempty"`
	//主动推送message
	Message *Message `protobuf:"bytes,3,opt,name=message,proto3,oneof" json:"message,omitempty"`
	//主动推送notice
	Notice *Notice `protobuf:"bytes,4,opt,name=notice,proto3,oneof" json:"notice,omitempty"`
}

func (x *GatewayWriteDataContent) Reset() {
	*x = GatewayWriteDataContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_peer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GatewayWriteDataContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GatewayWriteDataContent) ProtoMessage() {}

func (x *GatewayWriteDataContent) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_peer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GatewayWriteDataContent.ProtoReflect.Descriptor instead.
func (*GatewayWriteDataContent) Descriptor() ([]byte, []int) {
	return file_gateway_peer_proto_rawDescGZIP(), []int{2}
}

func (x *GatewayWriteDataContent) GetDataType() GatewayWriteDataType {
	if x != nil {
		return x.DataType
	}
	return GatewayWriteDataType_Response
}

func (x *GatewayWriteDataContent) GetResponse() *GatewayApiResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *GatewayWriteDataContent) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *GatewayWriteDataContent) GetNotice() *Notice {
	if x != nil {
		return x.Notice
	}
	return nil
}

type ListLongConnectionReq_Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds *string `protobuf:"bytes,1,opt,name=userIds,proto3,oneof" json:"userIds,omitempty"` // length > 0 时有效
}

func (x *ListLongConnectionReq_Filter) Reset() {
	*x = ListLongConnectionReq_Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_peer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLongConnectionReq_Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLongConnectionReq_Filter) ProtoMessage() {}

func (x *ListLongConnectionReq_Filter) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_peer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLongConnectionReq_Filter.ProtoReflect.Descriptor instead.
func (*ListLongConnectionReq_Filter) Descriptor() ([]byte, []int) {
	return file_gateway_peer_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ListLongConnectionReq_Filter) GetUserIds() string {
	if x != nil && x.UserIds != nil {
		return *x.UserIds
	}
	return ""
}

var File_gateway_peer_proto protoreflect.FileDescriptor

var file_gateway_peer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x10, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x01, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12,
	0x3d, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x48, 0x00, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x1a, 0x33,
	0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x82,
	0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x81, 0x02, 0x0a, 0x17, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x34, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x41, 0x70, 0x69, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48,
	0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x01, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x06, 0x6e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x48, 0x02, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65,
	0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x2a, 0x45, 0x0a, 0x14, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a,
	0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x50, 0x75, 0x73, 0x68, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x10, 0x02, 0x32, 0x5c,
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x09, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x00, 0x32, 0x60, 0x0a, 0x0f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0a,
	0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gateway_peer_proto_rawDescOnce sync.Once
	file_gateway_peer_proto_rawDescData = file_gateway_peer_proto_rawDesc
)

func file_gateway_peer_proto_rawDescGZIP() []byte {
	file_gateway_peer_proto_rawDescOnce.Do(func() {
		file_gateway_peer_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_peer_proto_rawDescData)
	})
	return file_gateway_peer_proto_rawDescData
}

var file_gateway_peer_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gateway_peer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_gateway_peer_proto_goTypes = []interface{}{
	(GatewayWriteDataType)(0),            // 0: pb.GatewayWriteDataType
	(*ListLongConnectionReq)(nil),        // 1: pb.ListLongConnectionReq
	(*ListLongConnectionResp)(nil),       // 2: pb.ListLongConnectionResp
	(*GatewayWriteDataContent)(nil),      // 3: pb.GatewayWriteDataContent
	(*ListLongConnectionReq_Filter)(nil), // 4: pb.ListLongConnectionReq.Filter
	(*RequestHeader)(nil),                // 5: pb.RequestHeader
	(*ResponseHeader)(nil),               // 6: pb.ResponseHeader
	(*LongConnection)(nil),               // 7: pb.LongConnection
	(*GatewayApiResponse)(nil),           // 8: pb.GatewayApiResponse
	(*Message)(nil),                      // 9: pb.Message
	(*Notice)(nil),                       // 10: pb.Notice
}
var file_gateway_peer_proto_depIdxs = []int32{
	5,  // 0: pb.ListLongConnectionReq.header:type_name -> pb.RequestHeader
	4,  // 1: pb.ListLongConnectionReq.filter:type_name -> pb.ListLongConnectionReq.Filter
	6,  // 2: pb.ListLongConnectionResp.header:type_name -> pb.ResponseHeader
	7,  // 3: pb.ListLongConnectionResp.longConnections:type_name -> pb.LongConnection
	0,  // 4: pb.GatewayWriteDataContent.dataType:type_name -> pb.GatewayWriteDataType
	8,  // 5: pb.GatewayWriteDataContent.response:type_name -> pb.GatewayApiResponse
	9,  // 6: pb.GatewayWriteDataContent.message:type_name -> pb.Message
	10, // 7: pb.GatewayWriteDataContent.notice:type_name -> pb.Notice
	3,  // 8: pb.connectionService.WriteData:input_type -> pb.GatewayWriteDataContent
	1,  // 9: pb.internalService.ListLongConnection:input_type -> pb.ListLongConnectionReq
	3,  // 10: pb.connectionService.WriteData:output_type -> pb.GatewayWriteDataContent
	2,  // 11: pb.internalService.ListLongConnection:output_type -> pb.ListLongConnectionResp
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_gateway_peer_proto_init() }
func file_gateway_peer_proto_init() {
	if File_gateway_peer_proto != nil {
		return
	}
	file_types_peer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_gateway_peer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLongConnectionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_peer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLongConnectionResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_peer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GatewayWriteDataContent); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gateway_peer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLongConnectionReq_Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_gateway_peer_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_gateway_peer_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_gateway_peer_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gateway_peer_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_gateway_peer_proto_goTypes,
		DependencyIndexes: file_gateway_peer_proto_depIdxs,
		EnumInfos:         file_gateway_peer_proto_enumTypes,
		MessageInfos:      file_gateway_peer_proto_msgTypes,
	}.Build()
	File_gateway_peer_proto = out.File
	file_gateway_peer_proto_rawDesc = nil
	file_gateway_peer_proto_goTypes = nil
	file_gateway_peer_proto_depIdxs = nil
}