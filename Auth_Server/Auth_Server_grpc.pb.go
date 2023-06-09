// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: Auth_Server.proto

package Auth_Server

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

// Req_DHClient is the client API for Req_DH service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type Req_DHClient interface {
	GetPQResponse(ctx context.Context, in *PQ_Request, opts ...grpc.CallOption) (*PQ_Response, error)
	GetDHResponse(ctx context.Context, in *DH_Request, opts ...grpc.CallOption) (*DH_Response, error)
}

type req_DHClient struct {
	cc grpc.ClientConnInterface
}

func NewReq_DHClient(cc grpc.ClientConnInterface) Req_DHClient {
	return &req_DHClient{cc}
}

func (c *req_DHClient) GetPQResponse(ctx context.Context, in *PQ_Request, opts ...grpc.CallOption) (*PQ_Response, error) {
	out := new(PQ_Response)
	err := c.cc.Invoke(ctx, "/Req_DH/GetPQResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *req_DHClient) GetDHResponse(ctx context.Context, in *DH_Request, opts ...grpc.CallOption) (*DH_Response, error) {
	out := new(DH_Response)
	err := c.cc.Invoke(ctx, "/Req_DH/GetDHResponse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Req_DHServer is the server API for Req_DH service.
// All implementations must embed UnimplementedReq_DHServer
// for forward compatibility
type Req_DHServer interface {
	GetPQResponse(context.Context, *PQ_Request) (*PQ_Response, error)
	GetDHResponse(context.Context, *DH_Request) (*DH_Response, error)
	mustEmbedUnimplementedReq_DHServer()
}

// UnimplementedReq_DHServer must be embedded to have forward compatible implementations.
type UnimplementedReq_DHServer struct {
}

func (UnimplementedReq_DHServer) GetPQResponse(context.Context, *PQ_Request) (*PQ_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPQResponse not implemented")
}
func (UnimplementedReq_DHServer) GetDHResponse(context.Context, *DH_Request) (*DH_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDHResponse not implemented")
}
func (UnimplementedReq_DHServer) mustEmbedUnimplementedReq_DHServer() {}

// UnsafeReq_DHServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to Req_DHServer will
// result in compilation errors.
type UnsafeReq_DHServer interface {
	mustEmbedUnimplementedReq_DHServer()
}

func RegisterReq_DHServer(s grpc.ServiceRegistrar, srv Req_DHServer) {
	s.RegisterService(&Req_DH_ServiceDesc, srv)
}

func _Req_DH_GetPQResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PQ_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Req_DHServer).GetPQResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Req_DH/GetPQResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Req_DHServer).GetPQResponse(ctx, req.(*PQ_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Req_DH_GetDHResponse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DH_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Req_DHServer).GetDHResponse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Req_DH/GetDHResponse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Req_DHServer).GetDHResponse(ctx, req.(*DH_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Req_DH_ServiceDesc is the grpc.ServiceDesc for Req_DH service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Req_DH_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Req_DH",
	HandlerType: (*Req_DHServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPQResponse",
			Handler:    _Req_DH_GetPQResponse_Handler,
		},
		{
			MethodName: "GetDHResponse",
			Handler:    _Req_DH_GetDHResponse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Auth_Server.proto",
}
