// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PluginClient is the client API for Plugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginClient interface {
	RegisterPlugin(ctx context.Context, in *RegisterPluginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeletePlugin(ctx context.Context, in *DeletePluginRequest, opts ...grpc.CallOption) (*DeletePluginResponse, error)
	GetPlugin(ctx context.Context, in *GetPluginRequest, opts ...grpc.CallOption) (*GetPluginResponse, error)
	ListPlugin(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPluginResponse, error)
}

type pluginClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginClient(cc grpc.ClientConnInterface) PluginClient {
	return &pluginClient{cc}
}

func (c *pluginClient) RegisterPlugin(ctx context.Context, in *RegisterPluginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/api.plugin.v1.Plugin/RegisterPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) DeletePlugin(ctx context.Context, in *DeletePluginRequest, opts ...grpc.CallOption) (*DeletePluginResponse, error) {
	out := new(DeletePluginResponse)
	err := c.cc.Invoke(ctx, "/api.plugin.v1.Plugin/DeletePlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) GetPlugin(ctx context.Context, in *GetPluginRequest, opts ...grpc.CallOption) (*GetPluginResponse, error) {
	out := new(GetPluginResponse)
	err := c.cc.Invoke(ctx, "/api.plugin.v1.Plugin/GetPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginClient) ListPlugin(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListPluginResponse, error) {
	out := new(ListPluginResponse)
	err := c.cc.Invoke(ctx, "/api.plugin.v1.Plugin/ListPlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServer is the server API for Plugin service.
// All implementations must embed UnimplementedPluginServer
// for forward compatibility
type PluginServer interface {
	RegisterPlugin(context.Context, *RegisterPluginRequest) (*emptypb.Empty, error)
	DeletePlugin(context.Context, *DeletePluginRequest) (*DeletePluginResponse, error)
	GetPlugin(context.Context, *GetPluginRequest) (*GetPluginResponse, error)
	ListPlugin(context.Context, *emptypb.Empty) (*ListPluginResponse, error)
	mustEmbedUnimplementedPluginServer()
}

// UnimplementedPluginServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServer struct {
}

func (UnimplementedPluginServer) RegisterPlugin(context.Context, *RegisterPluginRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPlugin not implemented")
}
func (UnimplementedPluginServer) DeletePlugin(context.Context, *DeletePluginRequest) (*DeletePluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlugin not implemented")
}
func (UnimplementedPluginServer) GetPlugin(context.Context, *GetPluginRequest) (*GetPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlugin not implemented")
}
func (UnimplementedPluginServer) ListPlugin(context.Context, *emptypb.Empty) (*ListPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlugin not implemented")
}
func (UnimplementedPluginServer) mustEmbedUnimplementedPluginServer() {}

// UnsafePluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServer will
// result in compilation errors.
type UnsafePluginServer interface {
	mustEmbedUnimplementedPluginServer()
}

func RegisterPluginServer(s grpc.ServiceRegistrar, srv PluginServer) {
	s.RegisterService(&Plugin_ServiceDesc, srv)
}

func _Plugin_RegisterPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).RegisterPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.plugin.v1.Plugin/RegisterPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).RegisterPlugin(ctx, req.(*RegisterPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_DeletePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).DeletePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.plugin.v1.Plugin/DeletePlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).DeletePlugin(ctx, req.(*DeletePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_GetPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).GetPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.plugin.v1.Plugin/GetPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).GetPlugin(ctx, req.(*GetPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Plugin_ListPlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServer).ListPlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.plugin.v1.Plugin/ListPlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServer).ListPlugin(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Plugin_ServiceDesc is the grpc.ServiceDesc for Plugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.plugin.v1.Plugin",
	HandlerType: (*PluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterPlugin",
			Handler:    _Plugin_RegisterPlugin_Handler,
		},
		{
			MethodName: "DeletePlugin",
			Handler:    _Plugin_DeletePlugin_Handler,
		},
		{
			MethodName: "GetPlugin",
			Handler:    _Plugin_GetPlugin_Handler,
		},
		{
			MethodName: "ListPlugin",
			Handler:    _Plugin_ListPlugin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/plugin/v1/plugin.proto",
}
