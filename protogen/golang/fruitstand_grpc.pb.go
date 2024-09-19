// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.1
// source: fruitstand.proto

package the_magical_fruit_stand

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

// FruitSageClient is the client API for FruitSage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FruitSageClient interface {
	AskAboutFruit(ctx context.Context, in *FruitQuery, opts ...grpc.CallOption) (*FruitWisdom, error)
}

type fruitSageClient struct {
	cc grpc.ClientConnInterface
}

func NewFruitSageClient(cc grpc.ClientConnInterface) FruitSageClient {
	return &fruitSageClient{cc}
}

func (c *fruitSageClient) AskAboutFruit(ctx context.Context, in *FruitQuery, opts ...grpc.CallOption) (*FruitWisdom, error) {
	out := new(FruitWisdom)
	err := c.cc.Invoke(ctx, "/FruitSage/AskAboutFruit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FruitSageServer is the server API for FruitSage service.
// All implementations must embed UnimplementedFruitSageServer
// for forward compatibility
type FruitSageServer interface {
	AskAboutFruit(context.Context, *FruitQuery) (*FruitWisdom, error)
	mustEmbedUnimplementedFruitSageServer()
}

// UnimplementedFruitSageServer must be embedded to have forward compatible implementations.
type UnimplementedFruitSageServer struct {
}

func (UnimplementedFruitSageServer) AskAboutFruit(context.Context, *FruitQuery) (*FruitWisdom, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AskAboutFruit not implemented")
}
func (UnimplementedFruitSageServer) mustEmbedUnimplementedFruitSageServer() {}

// UnsafeFruitSageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FruitSageServer will
// result in compilation errors.
type UnsafeFruitSageServer interface {
	mustEmbedUnimplementedFruitSageServer()
}

func RegisterFruitSageServer(s grpc.ServiceRegistrar, srv FruitSageServer) {
	s.RegisterService(&FruitSage_ServiceDesc, srv)
}

func _FruitSage_AskAboutFruit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FruitQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FruitSageServer).AskAboutFruit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FruitSage/AskAboutFruit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FruitSageServer).AskAboutFruit(ctx, req.(*FruitQuery))
	}
	return interceptor(ctx, in, info, handler)
}

// FruitSage_ServiceDesc is the grpc.ServiceDesc for FruitSage service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FruitSage_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "FruitSage",
	HandlerType: (*FruitSageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AskAboutFruit",
			Handler:    _FruitSage_AskAboutFruit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fruitstand.proto",
}
