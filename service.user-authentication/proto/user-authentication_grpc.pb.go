// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// UserAuthenticationClient is the client API for UserAuthentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAuthenticationClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	IsAuthend(ctx context.Context, in *IsAuthedRequest, opts ...grpc.CallOption) (*IsAuthedResponse, error)
}

type userAuthenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAuthenticationClient(cc grpc.ClientConnInterface) UserAuthenticationClient {
	return &userAuthenticationClient{cc}
}

func (c *userAuthenticationClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/auth_proto.UserAuthentication/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/auth_proto.UserAuthentication/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthenticationClient) IsAuthend(ctx context.Context, in *IsAuthedRequest, opts ...grpc.CallOption) (*IsAuthedResponse, error) {
	out := new(IsAuthedResponse)
	err := c.cc.Invoke(ctx, "/auth_proto.UserAuthentication/IsAuthend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAuthenticationServer is the server API for UserAuthentication service.
// All implementations must embed UnimplementedUserAuthenticationServer
// for forward compatibility
type UserAuthenticationServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	IsAuthend(context.Context, *IsAuthedRequest) (*IsAuthedResponse, error)
	mustEmbedUnimplementedUserAuthenticationServer()
}

// UnimplementedUserAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedUserAuthenticationServer struct {
}

func (UnimplementedUserAuthenticationServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedUserAuthenticationServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserAuthenticationServer) IsAuthend(context.Context, *IsAuthedRequest) (*IsAuthedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAuthend not implemented")
}
func (UnimplementedUserAuthenticationServer) mustEmbedUnimplementedUserAuthenticationServer() {}

// UnsafeUserAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAuthenticationServer will
// result in compilation errors.
type UnsafeUserAuthenticationServer interface {
	mustEmbedUnimplementedUserAuthenticationServer()
}

func RegisterUserAuthenticationServer(s grpc.ServiceRegistrar, srv UserAuthenticationServer) {
	s.RegisterService(&UserAuthentication_ServiceDesc, srv)
}

func _UserAuthentication_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.UserAuthentication/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuthentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.UserAuthentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuthentication_IsAuthend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).IsAuthend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth_proto.UserAuthentication/IsAuthend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).IsAuthend(ctx, req.(*IsAuthedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAuthentication_ServiceDesc is the grpc.ServiceDesc for UserAuthentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAuthentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth_proto.UserAuthentication",
	HandlerType: (*UserAuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _UserAuthentication_Create_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserAuthentication_Login_Handler,
		},
		{
			MethodName: "IsAuthend",
			Handler:    _UserAuthentication_IsAuthend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.user-authentication/proto/user-authentication.proto",
}