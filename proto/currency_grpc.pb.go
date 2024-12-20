// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: proto/currency.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CurrencyConversion_Convert_FullMethodName = "/currency.CurrencyConversion/Convert"
)

// CurrencyConversionClient is the client API for CurrencyConversion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrencyConversionClient interface {
	Convert(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error)
}

type currencyConversionClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrencyConversionClient(cc grpc.ClientConnInterface) CurrencyConversionClient {
	return &currencyConversionClient{cc}
}

func (c *currencyConversionClient) Convert(ctx context.Context, in *ConvertRequest, opts ...grpc.CallOption) (*ConvertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ConvertResponse)
	err := c.cc.Invoke(ctx, CurrencyConversion_Convert_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrencyConversionServer is the server API for CurrencyConversion service.
// All implementations must embed UnimplementedCurrencyConversionServer
// for forward compatibility.
type CurrencyConversionServer interface {
	Convert(context.Context, *ConvertRequest) (*ConvertResponse, error)
	mustEmbedUnimplementedCurrencyConversionServer()
}

// UnimplementedCurrencyConversionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCurrencyConversionServer struct{}

func (UnimplementedCurrencyConversionServer) Convert(context.Context, *ConvertRequest) (*ConvertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Convert not implemented")
}
func (UnimplementedCurrencyConversionServer) mustEmbedUnimplementedCurrencyConversionServer() {}
func (UnimplementedCurrencyConversionServer) testEmbeddedByValue()                            {}

// UnsafeCurrencyConversionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrencyConversionServer will
// result in compilation errors.
type UnsafeCurrencyConversionServer interface {
	mustEmbedUnimplementedCurrencyConversionServer()
}

func RegisterCurrencyConversionServer(s grpc.ServiceRegistrar, srv CurrencyConversionServer) {
	// If the following call pancis, it indicates UnimplementedCurrencyConversionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CurrencyConversion_ServiceDesc, srv)
}

func _CurrencyConversion_Convert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrencyConversionServer).Convert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CurrencyConversion_Convert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrencyConversionServer).Convert(ctx, req.(*ConvertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrencyConversion_ServiceDesc is the grpc.ServiceDesc for CurrencyConversion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrencyConversion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "currency.CurrencyConversion",
	HandlerType: (*CurrencyConversionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Convert",
			Handler:    _CurrencyConversion_Convert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/currency.proto",
}
