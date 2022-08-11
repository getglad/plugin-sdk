// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: internal/pb/destination.proto

package pb

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

// DestinationClient is the client API for Destination service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DestinationClient interface {
	Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error)
	// Get an example configuration for the source plugin
	GetExampleConfig(ctx context.Context, in *GetExampleConfig_Request, opts ...grpc.CallOption) (*GetExampleConfig_Response, error)
	// Save resources
	Save(ctx context.Context, opts ...grpc.CallOption) (Destination_SaveClient, error)
	// Create tables
	CreateTables(ctx context.Context, in *CreateTables_Request, opts ...grpc.CallOption) (*CreateTables_Response, error)
}

type destinationClient struct {
	cc grpc.ClientConnInterface
}

func NewDestinationClient(cc grpc.ClientConnInterface) DestinationClient {
	return &destinationClient{cc}
}

func (c *destinationClient) Configure(ctx context.Context, in *Configure_Request, opts ...grpc.CallOption) (*Configure_Response, error) {
	out := new(Configure_Response)
	err := c.cc.Invoke(ctx, "/proto.Destination/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationClient) GetExampleConfig(ctx context.Context, in *GetExampleConfig_Request, opts ...grpc.CallOption) (*GetExampleConfig_Response, error) {
	out := new(GetExampleConfig_Response)
	err := c.cc.Invoke(ctx, "/proto.Destination/GetExampleConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *destinationClient) Save(ctx context.Context, opts ...grpc.CallOption) (Destination_SaveClient, error) {
	stream, err := c.cc.NewStream(ctx, &Destination_ServiceDesc.Streams[0], "/proto.Destination/Save", opts...)
	if err != nil {
		return nil, err
	}
	x := &destinationSaveClient{stream}
	return x, nil
}

type Destination_SaveClient interface {
	Send(*Save_Request) error
	CloseAndRecv() (*Save_Response, error)
	grpc.ClientStream
}

type destinationSaveClient struct {
	grpc.ClientStream
}

func (x *destinationSaveClient) Send(m *Save_Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *destinationSaveClient) CloseAndRecv() (*Save_Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Save_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *destinationClient) CreateTables(ctx context.Context, in *CreateTables_Request, opts ...grpc.CallOption) (*CreateTables_Response, error) {
	out := new(CreateTables_Response)
	err := c.cc.Invoke(ctx, "/proto.Destination/CreateTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DestinationServer is the server API for Destination service.
// All implementations must embed UnimplementedDestinationServer
// for forward compatibility
type DestinationServer interface {
	Configure(context.Context, *Configure_Request) (*Configure_Response, error)
	// Get an example configuration for the source plugin
	GetExampleConfig(context.Context, *GetExampleConfig_Request) (*GetExampleConfig_Response, error)
	// Save resources
	Save(Destination_SaveServer) error
	// Create tables
	CreateTables(context.Context, *CreateTables_Request) (*CreateTables_Response, error)
	mustEmbedUnimplementedDestinationServer()
}

// UnimplementedDestinationServer must be embedded to have forward compatible implementations.
type UnimplementedDestinationServer struct {
}

func (UnimplementedDestinationServer) Configure(context.Context, *Configure_Request) (*Configure_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (UnimplementedDestinationServer) GetExampleConfig(context.Context, *GetExampleConfig_Request) (*GetExampleConfig_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExampleConfig not implemented")
}
func (UnimplementedDestinationServer) Save(Destination_SaveServer) error {
	return status.Errorf(codes.Unimplemented, "method Save not implemented")
}
func (UnimplementedDestinationServer) CreateTables(context.Context, *CreateTables_Request) (*CreateTables_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTables not implemented")
}
func (UnimplementedDestinationServer) mustEmbedUnimplementedDestinationServer() {}

// UnsafeDestinationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DestinationServer will
// result in compilation errors.
type UnsafeDestinationServer interface {
	mustEmbedUnimplementedDestinationServer()
}

func RegisterDestinationServer(s grpc.ServiceRegistrar, srv DestinationServer) {
	s.RegisterService(&Destination_ServiceDesc, srv)
}

func _Destination_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Configure_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Destination/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).Configure(ctx, req.(*Configure_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destination_GetExampleConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExampleConfig_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).GetExampleConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Destination/GetExampleConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).GetExampleConfig(ctx, req.(*GetExampleConfig_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Destination_Save_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DestinationServer).Save(&destinationSaveServer{stream})
}

type Destination_SaveServer interface {
	SendAndClose(*Save_Response) error
	Recv() (*Save_Request, error)
	grpc.ServerStream
}

type destinationSaveServer struct {
	grpc.ServerStream
}

func (x *destinationSaveServer) SendAndClose(m *Save_Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *destinationSaveServer) Recv() (*Save_Request, error) {
	m := new(Save_Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Destination_CreateTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTables_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DestinationServer).CreateTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Destination/CreateTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DestinationServer).CreateTables(ctx, req.(*CreateTables_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Destination_ServiceDesc is the grpc.ServiceDesc for Destination service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Destination_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Destination",
	HandlerType: (*DestinationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _Destination_Configure_Handler,
		},
		{
			MethodName: "GetExampleConfig",
			Handler:    _Destination_GetExampleConfig_Handler,
		},
		{
			MethodName: "CreateTables",
			Handler:    _Destination_CreateTables_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Save",
			Handler:       _Destination_Save_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/pb/destination.proto",
}