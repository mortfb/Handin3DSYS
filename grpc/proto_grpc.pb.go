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
	PublishMessage(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	BroadcastAll(ctx context.Context, in *BroadcastAllRequest, opts ...grpc.CallOption) (*BroadcastAllresponse, error)
	NewClientJoined(ctx context.Context, in *NewClientJoinedRequest, opts ...grpc.CallOption) (*NewClientJoinedResponse, error)
	ClientLeave(ctx context.Context, in *ClientLeaveRequest, opts ...grpc.CallOption) (*ClientLeaveResponse, error)
}

type chittyChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatServiceClient(cc grpc.ClientConnInterface) ChittyChatServiceClient {
	return &chittyChatServiceClient{cc}
}

func (c *chittyChatServiceClient) PublishMessage(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/ChittyChatService/publishMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatServiceClient) BroadcastAll(ctx context.Context, in *BroadcastAllRequest, opts ...grpc.CallOption) (*BroadcastAllresponse, error) {
	out := new(BroadcastAllresponse)
	err := c.cc.Invoke(ctx, "/ChittyChatService/broadcastAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatServiceClient) NewClientJoined(ctx context.Context, in *NewClientJoinedRequest, opts ...grpc.CallOption) (*NewClientJoinedResponse, error) {
	out := new(NewClientJoinedResponse)
	err := c.cc.Invoke(ctx, "/ChittyChatService/newClientJoined", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatServiceClient) ClientLeave(ctx context.Context, in *ClientLeaveRequest, opts ...grpc.CallOption) (*ClientLeaveResponse, error) {
	out := new(ClientLeaveResponse)
	err := c.cc.Invoke(ctx, "/ChittyChatService/clientLeave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittyChatServiceServer is the server API for ChittyChatService service.
// All implementations must embed UnimplementedChittyChatServiceServer
// for forward compatibility
type ChittyChatServiceServer interface {
	PublishMessage(context.Context, *PublishRequest) (*PublishResponse, error)
	BroadcastAll(context.Context, *BroadcastAllRequest) (*BroadcastAllresponse, error)
	NewClientJoined(context.Context, *NewClientJoinedRequest) (*NewClientJoinedResponse, error)
	ClientLeave(context.Context, *ClientLeaveRequest) (*ClientLeaveResponse, error)
	mustEmbedUnimplementedChittyChatServiceServer()
}

// UnimplementedChittyChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServiceServer struct {
}

func (UnimplementedChittyChatServiceServer) PublishMessage(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishMessage not implemented")
}
func (UnimplementedChittyChatServiceServer) BroadcastAll(context.Context, *BroadcastAllRequest) (*BroadcastAllresponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastAll not implemented")
}
func (UnimplementedChittyChatServiceServer) NewClientJoined(context.Context, *NewClientJoinedRequest) (*NewClientJoinedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewClientJoined not implemented")
}
func (UnimplementedChittyChatServiceServer) ClientLeave(context.Context, *ClientLeaveRequest) (*ClientLeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientLeave not implemented")
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

func _ChittyChatService_PublishMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).PublishMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChatService/publishMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).PublishMessage(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChatService_BroadcastAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).BroadcastAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChatService/broadcastAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).BroadcastAll(ctx, req.(*BroadcastAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChatService_NewClientJoined_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewClientJoinedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).NewClientJoined(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChatService/newClientJoined",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).NewClientJoined(ctx, req.(*NewClientJoinedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittyChatService_ClientLeave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientLeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).ClientLeave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChittyChatService/clientLeave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).ClientLeave(ctx, req.(*ClientLeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChittyChatService_ServiceDesc is the grpc.ServiceDesc for ChittyChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ChittyChatService",
	HandlerType: (*ChittyChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "publishMessage",
			Handler:    _ChittyChatService_PublishMessage_Handler,
		},
		{
			MethodName: "broadcastAll",
			Handler:    _ChittyChatService_BroadcastAll_Handler,
		},
		{
			MethodName: "newClientJoined",
			Handler:    _ChittyChatService_NewClientJoined_Handler,
		},
		{
			MethodName: "clientLeave",
			Handler:    _ChittyChatService_ClientLeave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto.proto",
}
