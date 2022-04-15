// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: edpc.proto

package edpcrpc

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

// EDPCerClient is the client API for EDPCer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EDPCerClient interface {
	Commander(ctx context.Context, in *CommanderRequest, opts ...grpc.CallOption) (*ReplyHeader, error)
	Docked(ctx context.Context, in *DockedRequest, opts ...grpc.CallOption) (*ReplyHeader, error)
}

type eDPCerClient struct {
	cc grpc.ClientConnInterface
}

func NewEDPCerClient(cc grpc.ClientConnInterface) EDPCerClient {
	return &eDPCerClient{cc}
}

func (c *eDPCerClient) Commander(ctx context.Context, in *CommanderRequest, opts ...grpc.CallOption) (*ReplyHeader, error) {
	out := new(ReplyHeader)
	err := c.cc.Invoke(ctx, "/EDPCer/Commander", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eDPCerClient) Docked(ctx context.Context, in *DockedRequest, opts ...grpc.CallOption) (*ReplyHeader, error) {
	out := new(ReplyHeader)
	err := c.cc.Invoke(ctx, "/EDPCer/Docked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EDPCerServer is the server API for EDPCer service.
// All implementations must embed UnimplementedEDPCerServer
// for forward compatibility
type EDPCerServer interface {
	Commander(context.Context, *CommanderRequest) (*ReplyHeader, error)
	Docked(context.Context, *DockedRequest) (*ReplyHeader, error)
	mustEmbedUnimplementedEDPCerServer()
}

// UnimplementedEDPCerServer must be embedded to have forward compatible implementations.
type UnimplementedEDPCerServer struct {
}

func (UnimplementedEDPCerServer) Commander(context.Context, *CommanderRequest) (*ReplyHeader, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commander not implemented")
}
func (UnimplementedEDPCerServer) Docked(context.Context, *DockedRequest) (*ReplyHeader, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Docked not implemented")
}
func (UnimplementedEDPCerServer) mustEmbedUnimplementedEDPCerServer() {}

// UnsafeEDPCerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EDPCerServer will
// result in compilation errors.
type UnsafeEDPCerServer interface {
	mustEmbedUnimplementedEDPCerServer()
}

func RegisterEDPCerServer(s grpc.ServiceRegistrar, srv EDPCerServer) {
	s.RegisterService(&EDPCer_ServiceDesc, srv)
}

func _EDPCer_Commander_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommanderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EDPCerServer).Commander(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EDPCer/Commander",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EDPCerServer).Commander(ctx, req.(*CommanderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EDPCer_Docked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EDPCerServer).Docked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EDPCer/Docked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EDPCerServer).Docked(ctx, req.(*DockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EDPCer_ServiceDesc is the grpc.ServiceDesc for EDPCer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EDPCer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EDPCer",
	HandlerType: (*EDPCerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Commander",
			Handler:    _EDPCer_Commander_Handler,
		},
		{
			MethodName: "Docked",
			Handler:    _EDPCer_Docked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "edpc.proto",
}
