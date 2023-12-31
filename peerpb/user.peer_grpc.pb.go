// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user.peer.proto

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	//UserRegister 用户注册
	//二次开发人员可以传递自定义参数 如果不满足你的需求 你可以修改该proto文件
	UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
	//UserDestroy 用户注销
	//注销逻辑可以从这里修改
	UserDestroy(ctx context.Context, in *UserDestroyReq, opts ...grpc.CallOption) (*UserDestroyResp, error)
	//UserToken 用户登录
	//登录逻辑可以从这里修改 默认是密码/手机号登录
	UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error)
	//RefreshUserToken 刷新用户token
	//刷新逻辑可以从这里修改
	RefreshUserToken(ctx context.Context, in *RefreshUserTokenReq, opts ...grpc.CallOption) (*RefreshUserTokenResp, error)
	//RevokeUserToken 注销用户token
	//注销逻辑可以从这里修改
	RevokeUserToken(ctx context.Context, in *RevokeUserTokenReq, opts ...grpc.CallOption) (*RevokeUserTokenResp, error)
	//UpdateUserAccountMap 更新用户账号信息
	//更新账号信息逻辑可以从这里修改 默认：如果是修改密码，需要旧密码或手机验证码。如果修改绑定手机号，需要验证新手机号的验证码
	UpdateUserAccountMap(ctx context.Context, in *UpdateUserAccountMapReq, opts ...grpc.CallOption) (*UpdateUserAccountMapResp, error)
	//ResetUserAccountMap 重置用户账号信息
	//重置账号信息逻辑可以从这里修改
	ResetUserAccountMap(ctx context.Context, in *ResetUserAccountMapReq, opts ...grpc.CallOption) (*ResetUserAccountMapResp, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	out := new(UserRegisterResp)
	err := c.cc.Invoke(ctx, "/pb.accountService/UserRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UserDestroy(ctx context.Context, in *UserDestroyReq, opts ...grpc.CallOption) (*UserDestroyResp, error) {
	out := new(UserDestroyResp)
	err := c.cc.Invoke(ctx, "/pb.accountService/UserDestroy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UserToken(ctx context.Context, in *UserTokenReq, opts ...grpc.CallOption) (*UserTokenResp, error) {
	out := new(UserTokenResp)
	err := c.cc.Invoke(ctx, "/pb.accountService/UserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) RefreshUserToken(ctx context.Context, in *RefreshUserTokenReq, opts ...grpc.CallOption) (*RefreshUserTokenResp, error) {
	out := new(RefreshUserTokenResp)
	err := c.cc.Invoke(ctx, "/pb.accountService/RefreshUserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) RevokeUserToken(ctx context.Context, in *RevokeUserTokenReq, opts ...grpc.CallOption) (*RevokeUserTokenResp, error) {
	out := new(RevokeUserTokenResp)
	err := c.cc.Invoke(ctx, "/pb.accountService/RevokeUserToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateUserAccountMap(ctx context.Context, in *UpdateUserAccountMapReq, opts ...grpc.CallOption) (*UpdateUserAccountMapResp, error) {
	out := new(UpdateUserAccountMapResp)
	err := c.cc.Invoke(ctx, "/pb.accountService/UpdateUserAccountMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) ResetUserAccountMap(ctx context.Context, in *ResetUserAccountMapReq, opts ...grpc.CallOption) (*ResetUserAccountMapResp, error) {
	out := new(ResetUserAccountMapResp)
	err := c.cc.Invoke(ctx, "/pb.accountService/ResetUserAccountMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations must embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	//UserRegister 用户注册
	//二次开发人员可以传递自定义参数 如果不满足你的需求 你可以修改该proto文件
	UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error)
	//UserDestroy 用户注销
	//注销逻辑可以从这里修改
	UserDestroy(context.Context, *UserDestroyReq) (*UserDestroyResp, error)
	//UserToken 用户登录
	//登录逻辑可以从这里修改 默认是密码/手机号登录
	UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error)
	//RefreshUserToken 刷新用户token
	//刷新逻辑可以从这里修改
	RefreshUserToken(context.Context, *RefreshUserTokenReq) (*RefreshUserTokenResp, error)
	//RevokeUserToken 注销用户token
	//注销逻辑可以从这里修改
	RevokeUserToken(context.Context, *RevokeUserTokenReq) (*RevokeUserTokenResp, error)
	//UpdateUserAccountMap 更新用户账号信息
	//更新账号信息逻辑可以从这里修改 默认：如果是修改密码，需要旧密码或手机验证码。如果修改绑定手机号，需要验证新手机号的验证码
	UpdateUserAccountMap(context.Context, *UpdateUserAccountMapReq) (*UpdateUserAccountMapResp, error)
	//ResetUserAccountMap 重置用户账号信息
	//重置账号信息逻辑可以从这里修改
	ResetUserAccountMap(context.Context, *ResetUserAccountMapReq) (*ResetUserAccountMapResp, error)
	mustEmbedUnimplementedAccountServiceServer()
}

// UnimplementedAccountServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) UserRegister(context.Context, *UserRegisterReq) (*UserRegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedAccountServiceServer) UserDestroy(context.Context, *UserDestroyReq) (*UserDestroyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDestroy not implemented")
}
func (UnimplementedAccountServiceServer) UserToken(context.Context, *UserTokenReq) (*UserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserToken not implemented")
}
func (UnimplementedAccountServiceServer) RefreshUserToken(context.Context, *RefreshUserTokenReq) (*RefreshUserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshUserToken not implemented")
}
func (UnimplementedAccountServiceServer) RevokeUserToken(context.Context, *RevokeUserTokenReq) (*RevokeUserTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeUserToken not implemented")
}
func (UnimplementedAccountServiceServer) UpdateUserAccountMap(context.Context, *UpdateUserAccountMapReq) (*UpdateUserAccountMapResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAccountMap not implemented")
}
func (UnimplementedAccountServiceServer) ResetUserAccountMap(context.Context, *ResetUserAccountMapReq) (*ResetUserAccountMapResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetUserAccountMap not implemented")
}
func (UnimplementedAccountServiceServer) mustEmbedUnimplementedAccountServiceServer() {}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.accountService/UserRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UserRegister(ctx, req.(*UserRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UserDestroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDestroyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UserDestroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.accountService/UserDestroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UserDestroy(ctx, req.(*UserDestroyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.accountService/UserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UserToken(ctx, req.(*UserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_RefreshUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshUserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).RefreshUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.accountService/RefreshUserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).RefreshUserToken(ctx, req.(*RefreshUserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_RevokeUserToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeUserTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).RevokeUserToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.accountService/RevokeUserToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).RevokeUserToken(ctx, req.(*RevokeUserTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateUserAccountMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAccountMapReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateUserAccountMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.accountService/UpdateUserAccountMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateUserAccountMap(ctx, req.(*UpdateUserAccountMapReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_ResetUserAccountMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetUserAccountMapReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).ResetUserAccountMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.accountService/ResetUserAccountMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).ResetUserAccountMap(ctx, req.(*ResetUserAccountMapReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.accountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _AccountService_UserRegister_Handler,
		},
		{
			MethodName: "UserDestroy",
			Handler:    _AccountService_UserDestroy_Handler,
		},
		{
			MethodName: "UserToken",
			Handler:    _AccountService_UserToken_Handler,
		},
		{
			MethodName: "RefreshUserToken",
			Handler:    _AccountService_RefreshUserToken_Handler,
		},
		{
			MethodName: "RevokeUserToken",
			Handler:    _AccountService_RevokeUserToken_Handler,
		},
		{
			MethodName: "UpdateUserAccountMap",
			Handler:    _AccountService_UpdateUserAccountMap_Handler,
		},
		{
			MethodName: "ResetUserAccountMap",
			Handler:    _AccountService_ResetUserAccountMap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.peer.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	//UpdateUserProfileMap 更新用户个人信息
	//更新个人信息逻辑可以从这里修改
	UpdateUserProfileMap(ctx context.Context, in *UpdateUserProfileMapReq, opts ...grpc.CallOption) (*UpdateUserProfileMapResp, error)
	//UpdateUserExtraMap 更新用户扩展信息
	//更新扩展信息逻辑可以从这里修改
	UpdateUserExtraMap(ctx context.Context, in *UpdateUserExtraMapReq, opts ...grpc.CallOption) (*UpdateUserExtraMapResp, error)
	//UpdateUserCountMap 更新用户计数信息
	//更新计数信息逻辑可以从这里修改
	UpdateUserCountMap(ctx context.Context, in *UpdateUserCountMapReq, opts ...grpc.CallOption) (*UpdateUserCountMapResp, error)
	//GetSelfUserInfo 获取自己的用户信息
	//获取自己的用户信息逻辑可以从这里修改
	GetSelfUserInfo(ctx context.Context, in *GetSelfUserInfoReq, opts ...grpc.CallOption) (*GetSelfUserInfoResp, error)
	//GetUserInfo 获取用户信息
	//获取用户信息逻辑可以从这里修改
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	//GetUserModelById 获取用户模型
	//获取用户模型逻辑可以从这里修改
	GetUserModelById(ctx context.Context, in *GetUserModelByIdReq, opts ...grpc.CallOption) (*GetUserModelByIdResp, error)
	//GetUserModelByIds 批量获取用户模型
	//批量获取用户模型逻辑可以从这里修改
	GetUserModelByIds(ctx context.Context, in *GetUserModelByIdsReq, opts ...grpc.CallOption) (*GetUserModelByIdsResp, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UpdateUserProfileMap(ctx context.Context, in *UpdateUserProfileMapReq, opts ...grpc.CallOption) (*UpdateUserProfileMapResp, error) {
	out := new(UpdateUserProfileMapResp)
	err := c.cc.Invoke(ctx, "/pb.userService/UpdateUserProfileMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserExtraMap(ctx context.Context, in *UpdateUserExtraMapReq, opts ...grpc.CallOption) (*UpdateUserExtraMapResp, error) {
	out := new(UpdateUserExtraMapResp)
	err := c.cc.Invoke(ctx, "/pb.userService/UpdateUserExtraMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserCountMap(ctx context.Context, in *UpdateUserCountMapReq, opts ...grpc.CallOption) (*UpdateUserCountMapResp, error) {
	out := new(UpdateUserCountMapResp)
	err := c.cc.Invoke(ctx, "/pb.userService/UpdateUserCountMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetSelfUserInfo(ctx context.Context, in *GetSelfUserInfoReq, opts ...grpc.CallOption) (*GetSelfUserInfoResp, error) {
	out := new(GetSelfUserInfoResp)
	err := c.cc.Invoke(ctx, "/pb.userService/GetSelfUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, "/pb.userService/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserModelById(ctx context.Context, in *GetUserModelByIdReq, opts ...grpc.CallOption) (*GetUserModelByIdResp, error) {
	out := new(GetUserModelByIdResp)
	err := c.cc.Invoke(ctx, "/pb.userService/GetUserModelById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserModelByIds(ctx context.Context, in *GetUserModelByIdsReq, opts ...grpc.CallOption) (*GetUserModelByIdsResp, error) {
	out := new(GetUserModelByIdsResp)
	err := c.cc.Invoke(ctx, "/pb.userService/GetUserModelByIds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	//UpdateUserProfileMap 更新用户个人信息
	//更新个人信息逻辑可以从这里修改
	UpdateUserProfileMap(context.Context, *UpdateUserProfileMapReq) (*UpdateUserProfileMapResp, error)
	//UpdateUserExtraMap 更新用户扩展信息
	//更新扩展信息逻辑可以从这里修改
	UpdateUserExtraMap(context.Context, *UpdateUserExtraMapReq) (*UpdateUserExtraMapResp, error)
	//UpdateUserCountMap 更新用户计数信息
	//更新计数信息逻辑可以从这里修改
	UpdateUserCountMap(context.Context, *UpdateUserCountMapReq) (*UpdateUserCountMapResp, error)
	//GetSelfUserInfo 获取自己的用户信息
	//获取自己的用户信息逻辑可以从这里修改
	GetSelfUserInfo(context.Context, *GetSelfUserInfoReq) (*GetSelfUserInfoResp, error)
	//GetUserInfo 获取用户信息
	//获取用户信息逻辑可以从这里修改
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	//GetUserModelById 获取用户模型
	//获取用户模型逻辑可以从这里修改
	GetUserModelById(context.Context, *GetUserModelByIdReq) (*GetUserModelByIdResp, error)
	//GetUserModelByIds 批量获取用户模型
	//批量获取用户模型逻辑可以从这里修改
	GetUserModelByIds(context.Context, *GetUserModelByIdsReq) (*GetUserModelByIdsResp, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UpdateUserProfileMap(context.Context, *UpdateUserProfileMapReq) (*UpdateUserProfileMapResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfileMap not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserExtraMap(context.Context, *UpdateUserExtraMapReq) (*UpdateUserExtraMapResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserExtraMap not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserCountMap(context.Context, *UpdateUserCountMapReq) (*UpdateUserCountMapResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserCountMap not implemented")
}
func (UnimplementedUserServiceServer) GetSelfUserInfo(context.Context, *GetSelfUserInfoReq) (*GetSelfUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelfUserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserModelById(context.Context, *GetUserModelByIdReq) (*GetUserModelByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserModelById not implemented")
}
func (UnimplementedUserServiceServer) GetUserModelByIds(context.Context, *GetUserModelByIdsReq) (*GetUserModelByIdsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserModelByIds not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UpdateUserProfileMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileMapReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserProfileMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.userService/UpdateUserProfileMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserProfileMap(ctx, req.(*UpdateUserProfileMapReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserExtraMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserExtraMapReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserExtraMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.userService/UpdateUserExtraMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserExtraMap(ctx, req.(*UpdateUserExtraMapReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserCountMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserCountMapReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserCountMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.userService/UpdateUserCountMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserCountMap(ctx, req.(*UpdateUserCountMapReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetSelfUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSelfUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSelfUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.userService/GetSelfUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSelfUserInfo(ctx, req.(*GetSelfUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.userService/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserModelById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserModelByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserModelById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.userService/GetUserModelById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserModelById(ctx, req.(*GetUserModelByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserModelByIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserModelByIdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserModelByIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.userService/GetUserModelByIds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserModelByIds(ctx, req.(*GetUserModelByIdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.userService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateUserProfileMap",
			Handler:    _UserService_UpdateUserProfileMap_Handler,
		},
		{
			MethodName: "UpdateUserExtraMap",
			Handler:    _UserService_UpdateUserExtraMap_Handler,
		},
		{
			MethodName: "UpdateUserCountMap",
			Handler:    _UserService_UpdateUserCountMap_Handler,
		},
		{
			MethodName: "GetSelfUserInfo",
			Handler:    _UserService_GetSelfUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserModelById",
			Handler:    _UserService_GetUserModelById_Handler,
		},
		{
			MethodName: "GetUserModelByIds",
			Handler:    _UserService_GetUserModelByIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.peer.proto",
}

// CallbackServiceClient is the client API for CallbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallbackServiceClient interface {
	UserAfterOnline(ctx context.Context, in *UserAfterOnlineReq, opts ...grpc.CallOption) (*UserAfterOnlineResp, error)
	UserAfterOffline(ctx context.Context, in *UserAfterOfflineReq, opts ...grpc.CallOption) (*UserAfterOfflineResp, error)
	UserBeforeConnect(ctx context.Context, in *UserBeforeConnectReq, opts ...grpc.CallOption) (*UserBeforeConnectResp, error)
	UserBeforeRequest(ctx context.Context, in *UserBeforeRequestReq, opts ...grpc.CallOption) (*UserBeforeRequestResp, error)
	UserAfterKeepAlive(ctx context.Context, in *UserAfterKeepAliveReq, opts ...grpc.CallOption) (*UserAfterKeepAliveResp, error)
}

type callbackServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCallbackServiceClient(cc grpc.ClientConnInterface) CallbackServiceClient {
	return &callbackServiceClient{cc}
}

func (c *callbackServiceClient) UserAfterOnline(ctx context.Context, in *UserAfterOnlineReq, opts ...grpc.CallOption) (*UserAfterOnlineResp, error) {
	out := new(UserAfterOnlineResp)
	err := c.cc.Invoke(ctx, "/pb.callbackService/UserAfterOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbackServiceClient) UserAfterOffline(ctx context.Context, in *UserAfterOfflineReq, opts ...grpc.CallOption) (*UserAfterOfflineResp, error) {
	out := new(UserAfterOfflineResp)
	err := c.cc.Invoke(ctx, "/pb.callbackService/UserAfterOffline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbackServiceClient) UserBeforeConnect(ctx context.Context, in *UserBeforeConnectReq, opts ...grpc.CallOption) (*UserBeforeConnectResp, error) {
	out := new(UserBeforeConnectResp)
	err := c.cc.Invoke(ctx, "/pb.callbackService/UserBeforeConnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbackServiceClient) UserBeforeRequest(ctx context.Context, in *UserBeforeRequestReq, opts ...grpc.CallOption) (*UserBeforeRequestResp, error) {
	out := new(UserBeforeRequestResp)
	err := c.cc.Invoke(ctx, "/pb.callbackService/UserBeforeRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *callbackServiceClient) UserAfterKeepAlive(ctx context.Context, in *UserAfterKeepAliveReq, opts ...grpc.CallOption) (*UserAfterKeepAliveResp, error) {
	out := new(UserAfterKeepAliveResp)
	err := c.cc.Invoke(ctx, "/pb.callbackService/UserAfterKeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallbackServiceServer is the server API for CallbackService service.
// All implementations must embed UnimplementedCallbackServiceServer
// for forward compatibility
type CallbackServiceServer interface {
	UserAfterOnline(context.Context, *UserAfterOnlineReq) (*UserAfterOnlineResp, error)
	UserAfterOffline(context.Context, *UserAfterOfflineReq) (*UserAfterOfflineResp, error)
	UserBeforeConnect(context.Context, *UserBeforeConnectReq) (*UserBeforeConnectResp, error)
	UserBeforeRequest(context.Context, *UserBeforeRequestReq) (*UserBeforeRequestResp, error)
	UserAfterKeepAlive(context.Context, *UserAfterKeepAliveReq) (*UserAfterKeepAliveResp, error)
	mustEmbedUnimplementedCallbackServiceServer()
}

// UnimplementedCallbackServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCallbackServiceServer struct {
}

func (UnimplementedCallbackServiceServer) UserAfterOnline(context.Context, *UserAfterOnlineReq) (*UserAfterOnlineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAfterOnline not implemented")
}
func (UnimplementedCallbackServiceServer) UserAfterOffline(context.Context, *UserAfterOfflineReq) (*UserAfterOfflineResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAfterOffline not implemented")
}
func (UnimplementedCallbackServiceServer) UserBeforeConnect(context.Context, *UserBeforeConnectReq) (*UserBeforeConnectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserBeforeConnect not implemented")
}
func (UnimplementedCallbackServiceServer) UserBeforeRequest(context.Context, *UserBeforeRequestReq) (*UserBeforeRequestResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserBeforeRequest not implemented")
}
func (UnimplementedCallbackServiceServer) UserAfterKeepAlive(context.Context, *UserAfterKeepAliveReq) (*UserAfterKeepAliveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAfterKeepAlive not implemented")
}
func (UnimplementedCallbackServiceServer) mustEmbedUnimplementedCallbackServiceServer() {}

// UnsafeCallbackServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallbackServiceServer will
// result in compilation errors.
type UnsafeCallbackServiceServer interface {
	mustEmbedUnimplementedCallbackServiceServer()
}

func RegisterCallbackServiceServer(s grpc.ServiceRegistrar, srv CallbackServiceServer) {
	s.RegisterService(&CallbackService_ServiceDesc, srv)
}

func _CallbackService_UserAfterOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAfterOnlineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackServiceServer).UserAfterOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.callbackService/UserAfterOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackServiceServer).UserAfterOnline(ctx, req.(*UserAfterOnlineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallbackService_UserAfterOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAfterOfflineReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackServiceServer).UserAfterOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.callbackService/UserAfterOffline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackServiceServer).UserAfterOffline(ctx, req.(*UserAfterOfflineReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallbackService_UserBeforeConnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBeforeConnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackServiceServer).UserBeforeConnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.callbackService/UserBeforeConnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackServiceServer).UserBeforeConnect(ctx, req.(*UserBeforeConnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallbackService_UserBeforeRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserBeforeRequestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackServiceServer).UserBeforeRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.callbackService/UserBeforeRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackServiceServer).UserBeforeRequest(ctx, req.(*UserBeforeRequestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CallbackService_UserAfterKeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAfterKeepAliveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallbackServiceServer).UserAfterKeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.callbackService/UserAfterKeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallbackServiceServer).UserAfterKeepAlive(ctx, req.(*UserAfterKeepAliveReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CallbackService_ServiceDesc is the grpc.ServiceDesc for CallbackService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CallbackService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.callbackService",
	HandlerType: (*CallbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserAfterOnline",
			Handler:    _CallbackService_UserAfterOnline_Handler,
		},
		{
			MethodName: "UserAfterOffline",
			Handler:    _CallbackService_UserAfterOffline_Handler,
		},
		{
			MethodName: "UserBeforeConnect",
			Handler:    _CallbackService_UserBeforeConnect_Handler,
		},
		{
			MethodName: "UserBeforeRequest",
			Handler:    _CallbackService_UserBeforeRequest_Handler,
		},
		{
			MethodName: "UserAfterKeepAlive",
			Handler:    _CallbackService_UserAfterKeepAlive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.peer.proto",
}
