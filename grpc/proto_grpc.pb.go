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

// ChittyChatServiceClient is the client API for ChittyChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatServiceClient interface {
	JoinServer(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	Connected(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_ConnectedClient, error)
}

type chittyChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatServiceClient(cc grpc.ClientConnInterface) ChittyChatServiceClient {
	return &chittyChatServiceClient{cc}
}

func (c *chittyChatServiceClient) JoinServer(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/ChittyChatService/joinServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatServiceClient) Connected(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_ConnectedClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[0], "/ChittyChatService/connected", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatServiceConnectedClient{stream}
	return x, nil
}

type ChittyChatService_ConnectedClient interface {
	Send(*PostMessage) error
	Recv() (*PostMessage, error)
	grpc.ClientStream
}

type chittyChatServiceConnectedClient struct {
	grpc.ClientStream
}

func (x *chittyChatServiceConnectedClient) Send(m *PostMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chittyChatServiceConnectedClient) Recv() (*PostMessage, error) {
	m := new(PostMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittyChatServiceServer is the server API for ChittyChatService service.
// All implementations must embed UnimplementedChittyChatServiceServer
// for forward compatibility
type ChittyChatServiceServer interface {
	JoinServer(context.Context, *JoinRequest) (*JoinResponse, error)
	Connected(ChittyChatService_ConnectedServer) error
	mustEmbedUnimplementedChittyChatServiceServer()
}

// UnimplementedChittyChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServiceServer struct {
}

func (UnimplementedChittyChatServiceServer) JoinServer(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinServer not implemented")
}
func (UnimplementedChittyChatServiceServer) Connected(ChittyChatService_ConnectedServer) error {
	return status.Errorf(codes.Unimplemented, "method Connected not implemented")
}
func (UnimplementedChittyChatServiceServer) mustEmbedUnimplementedChittyChatServiceServer() {}

// UnsafeChittyChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServiceServer will
// result in compilation errors.
type UnsafeChittyChatServiceServer interface {
	mustEmbedUnimplementedChittyChatServiceServer()
}

func RegisterChittyChatServiceServer(s grpc.ServiceRegistrar, srv ChittyChatServiceServer) {
	s.RegisterService(&ChittyChatService_ServiceDesc, srv)
}

func _ChittyChatService_JoinServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).JoinServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChatService/joinServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).JoinServer(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChatService_Connected_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChittyChatServiceServer).Connected(&chittyChatServiceConnectedServer{stream})
}

type ChittyChatService_ConnectedServer interface {
	Send(*PostMessage) error
	Recv() (*PostMessage, error)
	grpc.ServerStream
}

type chittyChatServiceConnectedServer struct {
	grpc.ServerStream
}

func (x *chittyChatServiceConnectedServer) Send(m *PostMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chittyChatServiceConnectedServer) Recv() (*PostMessage, error) {
	m := new(PostMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittyChatService_ServiceDesc is the grpc.ServiceDesc for ChittyChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChittyChatService",
	HandlerType: (*ChittyChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "joinServer",
			Handler:    _ChittyChatService_JoinServer_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "connected",
			Handler:       _ChittyChatService_Connected_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/proto.proto",
}
