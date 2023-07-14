// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: message.peer.proto

package peerpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	//MessageInsert 插入消息
	//二次开发时，需要实现该接口
	MessageInsert(ctx context.Context, in *MessageInsertReq, opts ...grpc.CallOption) (*MessageInsertResp, error)
	//MessageSend 发送消息
	MessageSend(ctx context.Context, in *MessageSendReq, opts ...grpc.CallOption) (*MessageSendResp, error)
	//MessagePush 推送消息
	MessagePush(ctx context.Context, in *MessagePushReq, opts ...grpc.CallOption) (*MessagePushResp, error)
	//GetConvMessageSeq 获取会话消息序列号
	GetConvMessageSeq(ctx context.Context, in *GetConvMessageSeqReq, opts ...grpc.CallOption) (*GetConvMessageSeqResp, error)
	//SyncMessage 同步消息
	SyncMessage(ctx context.Context, in *SyncMessageReq, opts ...grpc.CallOption) (*SyncMessageResp, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) MessageInsert(ctx context.Context, in *MessageInsertReq, opts ...grpc.CallOption) (*MessageInsertResp, error) {
	out := new(MessageInsertResp)
	err := c.cc.Invoke(ctx, "/pb.messageService/MessageInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) MessageSend(ctx context.Context, in *MessageSendReq, opts ...grpc.CallOption) (*MessageSendResp, error) {
	out := new(MessageSendResp)
	err := c.cc.Invoke(ctx, "/pb.messageService/MessageSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) MessagePush(ctx context.Context, in *MessagePushReq, opts ...grpc.CallOption) (*MessagePushResp, error) {
	out := new(MessagePushResp)
	err := c.cc.Invoke(ctx, "/pb.messageService/MessagePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) GetConvMessageSeq(ctx context.Context, in *GetConvMessageSeqReq, opts ...grpc.CallOption) (*GetConvMessageSeqResp, error) {
	out := new(GetConvMessageSeqResp)
	err := c.cc.Invoke(ctx, "/pb.messageService/GetConvMessageSeq", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SyncMessage(ctx context.Context, in *SyncMessageReq, opts ...grpc.CallOption) (*SyncMessageResp, error) {
	out := new(SyncMessageResp)
	err := c.cc.Invoke(ctx, "/pb.messageService/SyncMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility
type MessageServiceServer interface {
	//MessageInsert 插入消息
	//二次开发时，需要实现该接口
	MessageInsert(context.Context, *MessageInsertReq) (*MessageInsertResp, error)
	//MessageSend 发送消息
	MessageSend(context.Context, *MessageSendReq) (*MessageSendResp, error)
	//MessagePush 推送消息
	MessagePush(context.Context, *MessagePushReq) (*MessagePushResp, error)
	//GetConvMessageSeq 获取会话消息序列号
	GetConvMessageSeq(context.Context, *GetConvMessageSeqReq) (*GetConvMessageSeqResp, error)
	//SyncMessage 同步消息
	SyncMessage(context.Context, *SyncMessageReq) (*SyncMessageResp, error)
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (UnimplementedMessageServiceServer) MessageInsert(context.Context, *MessageInsertReq) (*MessageInsertResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageInsert not implemented")
}
func (UnimplementedMessageServiceServer) MessageSend(context.Context, *MessageSendReq) (*MessageSendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessageSend not implemented")
}
func (UnimplementedMessageServiceServer) MessagePush(context.Context, *MessagePushReq) (*MessagePushResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MessagePush not implemented")
}
func (UnimplementedMessageServiceServer) GetConvMessageSeq(context.Context, *GetConvMessageSeqReq) (*GetConvMessageSeqResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConvMessageSeq not implemented")
}
func (UnimplementedMessageServiceServer) SyncMessage(context.Context, *SyncMessageReq) (*SyncMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncMessage not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_MessageInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageInsertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.messageService/MessageInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageInsert(ctx, req.(*MessageInsertReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_MessageSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessageSendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessageSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.messageService/MessageSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessageSend(ctx, req.(*MessageSendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_MessagePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MessagePushReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).MessagePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.messageService/MessagePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).MessagePush(ctx, req.(*MessagePushReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_GetConvMessageSeq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConvMessageSeqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).GetConvMessageSeq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.messageService/GetConvMessageSeq",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).GetConvMessageSeq(ctx, req.(*GetConvMessageSeqReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SyncMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SyncMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.messageService/SyncMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SyncMessage(ctx, req.(*SyncMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.messageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MessageInsert",
			Handler:    _MessageService_MessageInsert_Handler,
		},
		{
			MethodName: "MessageSend",
			Handler:    _MessageService_MessageSend_Handler,
		},
		{
			MethodName: "MessagePush",
			Handler:    _MessageService_MessagePush_Handler,
		},
		{
			MethodName: "GetConvMessageSeq",
			Handler:    _MessageService_GetConvMessageSeq_Handler,
		},
		{
			MethodName: "SyncMessage",
			Handler:    _MessageService_SyncMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.peer.proto",
}

// NoticeServiceClient is the client API for NoticeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoticeServiceClient interface {
	//NoticeSend 通知发送
	NoticeSend(ctx context.Context, in *NoticeSendReq, opts ...grpc.CallOption) (*NoticeSendResp, error)
}

type noticeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoticeServiceClient(cc grpc.ClientConnInterface) NoticeServiceClient {
	return &noticeServiceClient{cc}
}

func (c *noticeServiceClient) NoticeSend(ctx context.Context, in *NoticeSendReq, opts ...grpc.CallOption) (*NoticeSendResp, error) {
	out := new(NoticeSendResp)
	err := c.cc.Invoke(ctx, "/pb.noticeService/NoticeSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoticeServiceServer is the server API for NoticeService service.
// All implementations must embed UnimplementedNoticeServiceServer
// for forward compatibility
type NoticeServiceServer interface {
	//NoticeSend 通知发送
	NoticeSend(context.Context, *NoticeSendReq) (*NoticeSendResp, error)
	mustEmbedUnimplementedNoticeServiceServer()
}

// UnimplementedNoticeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNoticeServiceServer struct {
}

func (UnimplementedNoticeServiceServer) NoticeSend(context.Context, *NoticeSendReq) (*NoticeSendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NoticeSend not implemented")
}
func (UnimplementedNoticeServiceServer) mustEmbedUnimplementedNoticeServiceServer() {}

// UnsafeNoticeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoticeServiceServer will
// result in compilation errors.
type UnsafeNoticeServiceServer interface {
	mustEmbedUnimplementedNoticeServiceServer()
}

func RegisterNoticeServiceServer(s grpc.ServiceRegistrar, srv NoticeServiceServer) {
	s.RegisterService(&NoticeService_ServiceDesc, srv)
}

func _NoticeService_NoticeSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoticeSendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoticeServiceServer).NoticeSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.noticeService/NoticeSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoticeServiceServer).NoticeSend(ctx, req.(*NoticeSendReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NoticeService_ServiceDesc is the grpc.ServiceDesc for NoticeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoticeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.noticeService",
	HandlerType: (*NoticeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NoticeSend",
			Handler:    _NoticeService_NoticeSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.peer.proto",
}
