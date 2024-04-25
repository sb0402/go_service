// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: liquid.proto

package liquidparsingpb

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

const (
	LiquidParsing_ParseAndRenderString_FullMethodName = "/liquid.LiquidParsing/ParseAndRenderString"
)

// LiquidParsingClient is the client API for LiquidParsing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiquidParsingClient interface {
	ParseAndRenderString(ctx context.Context, in *Template, opts ...grpc.CallOption) (*ParsedResult, error)
}

type liquidParsingClient struct {
	cc grpc.ClientConnInterface
}

func NewLiquidParsingClient(cc grpc.ClientConnInterface) LiquidParsingClient {
	return &liquidParsingClient{cc}
}

func (c *liquidParsingClient) ParseAndRenderString(ctx context.Context, in *Template, opts ...grpc.CallOption) (*ParsedResult, error) {
	out := new(ParsedResult)
	err := c.cc.Invoke(ctx, LiquidParsing_ParseAndRenderString_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiquidParsingServer is the server API for LiquidParsing service.
// All implementations must embed UnimplementedLiquidParsingServer
// for forward compatibility
type LiquidParsingServer interface {
	ParseAndRenderString(context.Context, *Template) (*ParsedResult, error)
	mustEmbedUnimplementedLiquidParsingServer()
}

// UnimplementedLiquidParsingServer must be embedded to have forward compatible implementations.
type UnimplementedLiquidParsingServer struct {
}

func (UnimplementedLiquidParsingServer) ParseAndRenderString(context.Context, *Template) (*ParsedResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseAndRenderString not implemented")
}
func (UnimplementedLiquidParsingServer) mustEmbedUnimplementedLiquidParsingServer() {}

// UnsafeLiquidParsingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiquidParsingServer will
// result in compilation errors.
type UnsafeLiquidParsingServer interface {
	mustEmbedUnimplementedLiquidParsingServer()
}

func RegisterLiquidParsingServer(s grpc.ServiceRegistrar, srv LiquidParsingServer) {
	s.RegisterService(&LiquidParsing_ServiceDesc, srv)
}

func _LiquidParsing_ParseAndRenderString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Template)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiquidParsingServer).ParseAndRenderString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiquidParsing_ParseAndRenderString_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiquidParsingServer).ParseAndRenderString(ctx, req.(*Template))
	}
	return interceptor(ctx, in, info, handler)
}

// LiquidParsing_ServiceDesc is the grpc.ServiceDesc for LiquidParsing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiquidParsing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "liquid.LiquidParsing",
	HandlerType: (*LiquidParsingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseAndRenderString",
			Handler:    _LiquidParsing_ParseAndRenderString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "liquid.proto",
}