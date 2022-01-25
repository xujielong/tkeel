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

// Oauth2Client is the client API for Oauth2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Oauth2Client interface {
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	IssuePluginToken(ctx context.Context, in *IssuePluginTokenRequest, opts ...grpc.CallOption) (*IssueTokenResponse, error)
	AddPluginWhiteList(ctx context.Context, in *AddPluginWhiteListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	IssueAdminToken(ctx context.Context, in *IssueAdminTokenRequest, opts ...grpc.CallOption) (*IssueTokenResponse, error)
	VerifyToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type oauth2Client struct {
	cc grpc.ClientConnInterface
}

func NewOauth2Client(cc grpc.ClientConnInterface) Oauth2Client {
	return &oauth2Client{cc}
}

func (c *oauth2Client) IssuePluginToken(ctx context.Context, in *IssuePluginTokenRequest, opts ...grpc.CallOption) (*IssueTokenResponse, error) {
	out := new(IssueTokenResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.rudder.api.oauth2.v1.Oauth2/IssuePluginToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) AddPluginWhiteList(ctx context.Context, in *AddPluginWhiteListRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/io.tkeel.rudder.api.oauth2.v1.Oauth2/AddPluginWhiteList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) IssueAdminToken(ctx context.Context, in *IssueAdminTokenRequest, opts ...grpc.CallOption) (*IssueTokenResponse, error) {
	out := new(IssueTokenResponse)
	err := c.cc.Invoke(ctx, "/io.tkeel.rudder.api.oauth2.v1.Oauth2/IssueAdminToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oauth2Client) VerifyToken(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/io.tkeel.rudder.api.oauth2.v1.Oauth2/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Oauth2Server is the server API for Oauth2 service.
// All implementations must embed UnimplementedOauth2Server
// for forward compatibility
type Oauth2Server interface {
	// TKEEL_COMMENT
	// {"response":{"raw_data":true}}
	IssuePluginToken(context.Context, *IssuePluginTokenRequest) (*IssueTokenResponse, error)
	AddPluginWhiteList(context.Context, *AddPluginWhiteListRequest) (*emptypb.Empty, error)
	IssueAdminToken(context.Context, *IssueAdminTokenRequest) (*IssueTokenResponse, error)
	VerifyToken(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedOauth2Server()
}

// UnimplementedOauth2Server must be embedded to have forward compatible implementations.
type UnimplementedOauth2Server struct {
}

func (UnimplementedOauth2Server) IssuePluginToken(context.Context, *IssuePluginTokenRequest) (*IssueTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssuePluginToken not implemented")
}
func (UnimplementedOauth2Server) AddPluginWhiteList(context.Context, *AddPluginWhiteListRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPluginWhiteList not implemented")
}
func (UnimplementedOauth2Server) IssueAdminToken(context.Context, *IssueAdminTokenRequest) (*IssueTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueAdminToken not implemented")
}
func (UnimplementedOauth2Server) VerifyToken(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (UnimplementedOauth2Server) mustEmbedUnimplementedOauth2Server() {}

// UnsafeOauth2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Oauth2Server will
// result in compilation errors.
type UnsafeOauth2Server interface {
	mustEmbedUnimplementedOauth2Server()
}

func RegisterOauth2Server(s grpc.ServiceRegistrar, srv Oauth2Server) {
	s.RegisterService(&Oauth2_ServiceDesc, srv)
}

func _Oauth2_IssuePluginToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssuePluginTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).IssuePluginToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.rudder.api.oauth2.v1.Oauth2/IssuePluginToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).IssuePluginToken(ctx, req.(*IssuePluginTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_AddPluginWhiteList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPluginWhiteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).AddPluginWhiteList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.rudder.api.oauth2.v1.Oauth2/AddPluginWhiteList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).AddPluginWhiteList(ctx, req.(*AddPluginWhiteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_IssueAdminToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueAdminTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).IssueAdminToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.rudder.api.oauth2.v1.Oauth2/IssueAdminToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).IssueAdminToken(ctx, req.(*IssueAdminTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Oauth2_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Oauth2Server).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/io.tkeel.rudder.api.oauth2.v1.Oauth2/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Oauth2Server).VerifyToken(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Oauth2_ServiceDesc is the grpc.ServiceDesc for Oauth2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Oauth2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.tkeel.rudder.api.oauth2.v1.Oauth2",
	HandlerType: (*Oauth2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssuePluginToken",
			Handler:    _Oauth2_IssuePluginToken_Handler,
		},
		{
			MethodName: "AddPluginWhiteList",
			Handler:    _Oauth2_AddPluginWhiteList_Handler,
		},
		{
			MethodName: "IssueAdminToken",
			Handler:    _Oauth2_IssueAdminToken_Handler,
		},
		{
			MethodName: "VerifyToken",
			Handler:    _Oauth2_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/oauth2/v1/oauth2.proto",
}
