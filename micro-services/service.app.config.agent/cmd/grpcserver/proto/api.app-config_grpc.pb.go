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

// AppConfigurationClient is the client API for AppConfiguration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppConfigurationClient interface {
	Initialize(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	AppendPermission(ctx context.Context, in *AppendPermissionRequest, opts ...grpc.CallOption) (*AppendPermissionResponse, error)
	RollbackAppendPermission(ctx context.Context, in *RollbackAppendPermissionRequest, opts ...grpc.CallOption) (*RollbackAppendPermissionResponse, error)
	LockConfig(ctx context.Context, in *LockConfigRequest, opts ...grpc.CallOption) (*LockConfigResponse, error)
	UnlockConfig(ctx context.Context, in *UnlockConfigRequest, opts ...grpc.CallOption) (*UnlockConfigResponse, error)
	GetForClient(ctx context.Context, in *GetForClientRequest, opts ...grpc.CallOption) (*GetForClientResponse, error)
}

type appConfigurationClient struct {
	cc grpc.ClientConnInterface
}

func NewAppConfigurationClient(cc grpc.ClientConnInterface) AppConfigurationClient {
	return &appConfigurationClient{cc}
}

func (c *appConfigurationClient) Initialize(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/config_proto.AppConfiguration/Initialize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appConfigurationClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/config_proto.AppConfiguration/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appConfigurationClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/config_proto.AppConfiguration/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appConfigurationClient) AppendPermission(ctx context.Context, in *AppendPermissionRequest, opts ...grpc.CallOption) (*AppendPermissionResponse, error) {
	out := new(AppendPermissionResponse)
	err := c.cc.Invoke(ctx, "/config_proto.AppConfiguration/AppendPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appConfigurationClient) RollbackAppendPermission(ctx context.Context, in *RollbackAppendPermissionRequest, opts ...grpc.CallOption) (*RollbackAppendPermissionResponse, error) {
	out := new(RollbackAppendPermissionResponse)
	err := c.cc.Invoke(ctx, "/config_proto.AppConfiguration/RollbackAppendPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appConfigurationClient) LockConfig(ctx context.Context, in *LockConfigRequest, opts ...grpc.CallOption) (*LockConfigResponse, error) {
	out := new(LockConfigResponse)
	err := c.cc.Invoke(ctx, "/config_proto.AppConfiguration/LockConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appConfigurationClient) UnlockConfig(ctx context.Context, in *UnlockConfigRequest, opts ...grpc.CallOption) (*UnlockConfigResponse, error) {
	out := new(UnlockConfigResponse)
	err := c.cc.Invoke(ctx, "/config_proto.AppConfiguration/UnlockConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appConfigurationClient) GetForClient(ctx context.Context, in *GetForClientRequest, opts ...grpc.CallOption) (*GetForClientResponse, error) {
	out := new(GetForClientResponse)
	err := c.cc.Invoke(ctx, "/config_proto.AppConfiguration/GetForClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppConfigurationServer is the server API for AppConfiguration service.
// All implementations must embed UnimplementedAppConfigurationServer
// for forward compatibility
type AppConfigurationServer interface {
	Initialize(context.Context, *InitRequest) (*InitResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	AppendPermission(context.Context, *AppendPermissionRequest) (*AppendPermissionResponse, error)
	RollbackAppendPermission(context.Context, *RollbackAppendPermissionRequest) (*RollbackAppendPermissionResponse, error)
	LockConfig(context.Context, *LockConfigRequest) (*LockConfigResponse, error)
	UnlockConfig(context.Context, *UnlockConfigRequest) (*UnlockConfigResponse, error)
	GetForClient(context.Context, *GetForClientRequest) (*GetForClientResponse, error)
	mustEmbedUnimplementedAppConfigurationServer()
}

// UnimplementedAppConfigurationServer must be embedded to have forward compatible implementations.
type UnimplementedAppConfigurationServer struct {
}

func (UnimplementedAppConfigurationServer) Initialize(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initialize not implemented")
}
func (UnimplementedAppConfigurationServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAppConfigurationServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAppConfigurationServer) AppendPermission(context.Context, *AppendPermissionRequest) (*AppendPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendPermission not implemented")
}
func (UnimplementedAppConfigurationServer) RollbackAppendPermission(context.Context, *RollbackAppendPermissionRequest) (*RollbackAppendPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RollbackAppendPermission not implemented")
}
func (UnimplementedAppConfigurationServer) LockConfig(context.Context, *LockConfigRequest) (*LockConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockConfig not implemented")
}
func (UnimplementedAppConfigurationServer) UnlockConfig(context.Context, *UnlockConfigRequest) (*UnlockConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlockConfig not implemented")
}
func (UnimplementedAppConfigurationServer) GetForClient(context.Context, *GetForClientRequest) (*GetForClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForClient not implemented")
}
func (UnimplementedAppConfigurationServer) mustEmbedUnimplementedAppConfigurationServer() {}

// UnsafeAppConfigurationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppConfigurationServer will
// result in compilation errors.
type UnsafeAppConfigurationServer interface {
	mustEmbedUnimplementedAppConfigurationServer()
}

func RegisterAppConfigurationServer(s grpc.ServiceRegistrar, srv AppConfigurationServer) {
	s.RegisterService(&AppConfiguration_ServiceDesc, srv)
}

func _AppConfiguration_Initialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigurationServer).Initialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_proto.AppConfiguration/Initialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigurationServer).Initialize(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppConfiguration_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigurationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_proto.AppConfiguration/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigurationServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppConfiguration_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigurationServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_proto.AppConfiguration/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigurationServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppConfiguration_AppendPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigurationServer).AppendPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_proto.AppConfiguration/AppendPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigurationServer).AppendPermission(ctx, req.(*AppendPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppConfiguration_RollbackAppendPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RollbackAppendPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigurationServer).RollbackAppendPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_proto.AppConfiguration/RollbackAppendPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigurationServer).RollbackAppendPermission(ctx, req.(*RollbackAppendPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppConfiguration_LockConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigurationServer).LockConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_proto.AppConfiguration/LockConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigurationServer).LockConfig(ctx, req.(*LockConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppConfiguration_UnlockConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigurationServer).UnlockConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_proto.AppConfiguration/UnlockConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigurationServer).UnlockConfig(ctx, req.(*UnlockConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppConfiguration_GetForClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetForClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppConfigurationServer).GetForClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/config_proto.AppConfiguration/GetForClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppConfigurationServer).GetForClient(ctx, req.(*GetForClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppConfiguration_ServiceDesc is the grpc.ServiceDesc for AppConfiguration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppConfiguration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "config_proto.AppConfiguration",
	HandlerType: (*AppConfigurationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Initialize",
			Handler:    _AppConfiguration_Initialize_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AppConfiguration_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AppConfiguration_Update_Handler,
		},
		{
			MethodName: "AppendPermission",
			Handler:    _AppConfiguration_AppendPermission_Handler,
		},
		{
			MethodName: "RollbackAppendPermission",
			Handler:    _AppConfiguration_RollbackAppendPermission_Handler,
		},
		{
			MethodName: "LockConfig",
			Handler:    _AppConfiguration_LockConfig_Handler,
		},
		{
			MethodName: "UnlockConfig",
			Handler:    _AppConfiguration_UnlockConfig_Handler,
		},
		{
			MethodName: "GetForClient",
			Handler:    _AppConfiguration_GetForClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.app.config.agent/cmd/grpcserver/proto/api.app-config.proto",
}
