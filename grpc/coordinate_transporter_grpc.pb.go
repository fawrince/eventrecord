// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: grpc/coordinate_transporter.proto

package grpc

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

// CoordinateTransporterClient is the client API for CoordinateTransporter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoordinateTransporterClient interface {
	PostCoordinates(ctx context.Context, opts ...grpc.CallOption) (CoordinateTransporter_PostCoordinatesClient, error)
}

type coordinateTransporterClient struct {
	cc grpc.ClientConnInterface
}

func NewCoordinateTransporterClient(cc grpc.ClientConnInterface) CoordinateTransporterClient {
	return &coordinateTransporterClient{cc}
}

func (c *coordinateTransporterClient) PostCoordinates(ctx context.Context, opts ...grpc.CallOption) (CoordinateTransporter_PostCoordinatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoordinateTransporter_ServiceDesc.Streams[0], "/CoordinateTransporter/PostCoordinates", opts...)
	if err != nil {
		return nil, err
	}
	x := &coordinateTransporterPostCoordinatesClient{stream}
	return x, nil
}

type CoordinateTransporter_PostCoordinatesClient interface {
	Send(*PostCoordinateRequest) error
	Recv() (*PostCoordinateResponse, error)
	grpc.ClientStream
}

type coordinateTransporterPostCoordinatesClient struct {
	grpc.ClientStream
}

func (x *coordinateTransporterPostCoordinatesClient) Send(m *PostCoordinateRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *coordinateTransporterPostCoordinatesClient) Recv() (*PostCoordinateResponse, error) {
	m := new(PostCoordinateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CoordinateTransporterServer is the server API for CoordinateTransporter service.
// All implementations must embed UnimplementedCoordinateTransporterServer
// for forward compatibility
type CoordinateTransporterServer interface {
	PostCoordinates(CoordinateTransporter_PostCoordinatesServer) error
	mustEmbedUnimplementedCoordinateTransporterServer()
}

// UnimplementedCoordinateTransporterServer must be embedded to have forward compatible implementations.
type UnimplementedCoordinateTransporterServer struct {
}

func (UnimplementedCoordinateTransporterServer) PostCoordinates(CoordinateTransporter_PostCoordinatesServer) error {
	return status.Errorf(codes.Unimplemented, "method PostCoordinates not implemented")
}
func (UnimplementedCoordinateTransporterServer) mustEmbedUnimplementedCoordinateTransporterServer() {}

// UnsafeCoordinateTransporterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoordinateTransporterServer will
// result in compilation errors.
type UnsafeCoordinateTransporterServer interface {
	mustEmbedUnimplementedCoordinateTransporterServer()
}

func RegisterCoordinateTransporterServer(s grpc.ServiceRegistrar, srv CoordinateTransporterServer) {
	s.RegisterService(&CoordinateTransporter_ServiceDesc, srv)
}

func _CoordinateTransporter_PostCoordinates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CoordinateTransporterServer).PostCoordinates(&coordinateTransporterPostCoordinatesServer{stream})
}

type CoordinateTransporter_PostCoordinatesServer interface {
	Send(*PostCoordinateResponse) error
	Recv() (*PostCoordinateRequest, error)
	grpc.ServerStream
}

type coordinateTransporterPostCoordinatesServer struct {
	grpc.ServerStream
}

func (x *coordinateTransporterPostCoordinatesServer) Send(m *PostCoordinateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *coordinateTransporterPostCoordinatesServer) Recv() (*PostCoordinateRequest, error) {
	m := new(PostCoordinateRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CoordinateTransporter_ServiceDesc is the grpc.ServiceDesc for CoordinateTransporter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoordinateTransporter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CoordinateTransporter",
	HandlerType: (*CoordinateTransporterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostCoordinates",
			Handler:       _CoordinateTransporter_PostCoordinates_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/coordinate_transporter.proto",
}