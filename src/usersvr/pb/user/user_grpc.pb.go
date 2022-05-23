// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: proto/user.proto

package user

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

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	RegisterCore(ctx context.Context, in *RegisterCoreReq, opts ...grpc.CallOption) (*RegisterCoreResp, error)
	Register2(ctx context.Context, in *Register2Req, opts ...grpc.CallOption) (*Register2Resp, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	GetUserCore(ctx context.Context, in *GetUserCoreReq, opts ...grpc.CallOption) (*GetUserCoreResp, error)
	CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error)
	ModifyUserInfo(ctx context.Context, in *ModifyUserInfoReq, opts ...grpc.CallOption) (*NilResp, error)
	GetUserCoreList(ctx context.Context, in *GetUserCoreListReq, opts ...grpc.CallOption) (*GetUserCoreListResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/user.User/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RegisterCore(ctx context.Context, in *RegisterCoreReq, opts ...grpc.CallOption) (*RegisterCoreResp, error) {
	out := new(RegisterCoreResp)
	err := c.cc.Invoke(ctx, "/user.User/registerCore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Register2(ctx context.Context, in *Register2Req, opts ...grpc.CallOption) (*Register2Resp, error) {
	out := new(Register2Resp)
	err := c.cc.Invoke(ctx, "/user.User/register2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, "/user.User/getUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserCore(ctx context.Context, in *GetUserCoreReq, opts ...grpc.CallOption) (*GetUserCoreResp, error) {
	out := new(GetUserCoreResp)
	err := c.cc.Invoke(ctx, "/user.User/getUserCore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CheckToken(ctx context.Context, in *CheckTokenReq, opts ...grpc.CallOption) (*CheckTokenResp, error) {
	out := new(CheckTokenResp)
	err := c.cc.Invoke(ctx, "/user.User/checkToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ModifyUserInfo(ctx context.Context, in *ModifyUserInfoReq, opts ...grpc.CallOption) (*NilResp, error) {
	out := new(NilResp)
	err := c.cc.Invoke(ctx, "/user.User/modifyUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserCoreList(ctx context.Context, in *GetUserCoreListReq, opts ...grpc.CallOption) (*GetUserCoreListResp, error) {
	out := new(GetUserCoreListResp)
	err := c.cc.Invoke(ctx, "/user.User/getUserCoreList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	RegisterCore(context.Context, *RegisterCoreReq) (*RegisterCoreResp, error)
	Register2(context.Context, *Register2Req) (*Register2Resp, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	GetUserCore(context.Context, *GetUserCoreReq) (*GetUserCoreResp, error)
	CheckToken(context.Context, *CheckTokenReq) (*CheckTokenResp, error)
	ModifyUserInfo(context.Context, *ModifyUserInfoReq) (*NilResp, error)
	GetUserCoreList(context.Context, *GetUserCoreListReq) (*GetUserCoreListResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) RegisterCore(context.Context, *RegisterCoreReq) (*RegisterCoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCore not implemented")
}
func (UnimplementedUserServer) Register2(context.Context, *Register2Req) (*Register2Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register2 not implemented")
}
func (UnimplementedUserServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServer) GetUserCore(context.Context, *GetUserCoreReq) (*GetUserCoreResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCore not implemented")
}
func (UnimplementedUserServer) CheckToken(context.Context, *CheckTokenReq) (*CheckTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckToken not implemented")
}
func (UnimplementedUserServer) ModifyUserInfo(context.Context, *ModifyUserInfoReq) (*NilResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUserInfo not implemented")
}
func (UnimplementedUserServer) GetUserCoreList(context.Context, *GetUserCoreListReq) (*GetUserCoreListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserCoreList not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RegisterCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RegisterCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/registerCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RegisterCore(ctx, req.(*RegisterCoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Register2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Register2Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/register2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register2(ctx, req.(*Register2Req))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserCore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserCore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUserCore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserCore(ctx, req.(*GetUserCoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/checkToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckToken(ctx, req.(*CheckTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ModifyUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ModifyUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/modifyUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ModifyUserInfo(ctx, req.(*ModifyUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserCoreList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserCoreListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserCoreList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/getUserCoreList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserCoreList(ctx, req.(*GetUserCoreListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "registerCore",
			Handler:    _User_RegisterCore_Handler,
		},
		{
			MethodName: "register2",
			Handler:    _User_Register2_Handler,
		},
		{
			MethodName: "getUserInfo",
			Handler:    _User_GetUserInfo_Handler,
		},
		{
			MethodName: "getUserCore",
			Handler:    _User_GetUserCore_Handler,
		},
		{
			MethodName: "checkToken",
			Handler:    _User_CheckToken_Handler,
		},
		{
			MethodName: "modifyUserInfo",
			Handler:    _User_ModifyUserInfo_Handler,
		},
		{
			MethodName: "getUserCoreList",
			Handler:    _User_GetUserCoreList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}