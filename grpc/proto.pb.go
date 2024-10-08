// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: grpc/proto.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PublishRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TimeStamp int32  `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *PublishRequest) Reset() {
	*x = PublishRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishRequest) ProtoMessage() {}

func (x *PublishRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishRequest.ProtoReflect.Descriptor instead.
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{0}
}

func (x *PublishRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PublishRequest) GetTimeStamp() int32 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type PublishResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` //bool to mark the success
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`  //message marking succes or failure
}

func (x *PublishResponse) Reset() {
	*x = PublishResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishResponse) ProtoMessage() {}

func (x *PublishResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishResponse.ProtoReflect.Descriptor instead.
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{1}
}

func (x *PublishResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PublishResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type BroadcastAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BroadcastAllRequest) Reset() {
	*x = BroadcastAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastAllRequest) ProtoMessage() {}

func (x *BroadcastAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastAllRequest.ProtoReflect.Descriptor instead.
func (*BroadcastAllRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{2}
}

type BroadcastAllresponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   []string `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty"`
	TimeStamp []int32  `protobuf:"varint,2,rep,packed,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *BroadcastAllresponse) Reset() {
	*x = BroadcastAllresponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastAllresponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastAllresponse) ProtoMessage() {}

func (x *BroadcastAllresponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastAllresponse.ProtoReflect.Descriptor instead.
func (*BroadcastAllresponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{3}
}

func (x *BroadcastAllresponse) GetMessage() []string {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *BroadcastAllresponse) GetTimeStamp() []int32 {
	if x != nil {
		return x.TimeStamp
	}
	return nil
}

type NewClientJoinedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TimeStamp int32  `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *NewClientJoinedRequest) Reset() {
	*x = NewClientJoinedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewClientJoinedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewClientJoinedRequest) ProtoMessage() {}

func (x *NewClientJoinedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewClientJoinedRequest.ProtoReflect.Descriptor instead.
func (*NewClientJoinedRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{4}
}

func (x *NewClientJoinedRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewClientJoinedRequest) GetTimeStamp() int32 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type NewClientJoinedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TimeStamp int32  `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *NewClientJoinedResponse) Reset() {
	*x = NewClientJoinedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewClientJoinedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewClientJoinedResponse) ProtoMessage() {}

func (x *NewClientJoinedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewClientJoinedResponse.ProtoReflect.Descriptor instead.
func (*NewClientJoinedResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{5}
}

func (x *NewClientJoinedResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *NewClientJoinedResponse) GetTimeStamp() int32 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type ClientLeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TimeStamp int32  `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *ClientLeaveRequest) Reset() {
	*x = ClientLeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientLeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientLeaveRequest) ProtoMessage() {}

func (x *ClientLeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientLeaveRequest.ProtoReflect.Descriptor instead.
func (*ClientLeaveRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{6}
}

func (x *ClientLeaveRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientLeaveRequest) GetTimeStamp() int32 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

type ClientLeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TimeStamp int32  `protobuf:"varint,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (x *ClientLeaveResponse) Reset() {
	*x = ClientLeaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientLeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientLeaveResponse) ProtoMessage() {}

func (x *ClientLeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientLeaveResponse.ProtoReflect.Descriptor instead.
func (*ClientLeaveResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_proto_rawDescGZIP(), []int{7}
}

func (x *ClientLeaveResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ClientLeaveResponse) GetTimeStamp() int32 {
	if x != nil {
		return x.TimeStamp
	}
	return 0
}

var File_grpc_proto_proto protoreflect.FileDescriptor

var file_grpc_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x45, 0x0a, 0x0f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x14, 0x62, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4a, 0x0a, 0x16, 0x6e, 0x65,
	0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x51, 0x0a, 0x17, 0x6e, 0x65, 0x77, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x46, 0x0a, 0x12, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d,
	0x70, 0x22, 0x4d, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x32, 0x8d, 0x02, 0x0a, 0x11, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0c, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e,
	0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0f,
	0x6e, 0x65, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x12,
	0x17, 0x2e, 0x6e, 0x65, 0x77, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x65, 0x77, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x12, 0x13, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x14, 0x5a, 0x12, 0x68, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x33, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_proto_rawDescOnce sync.Once
	file_grpc_proto_proto_rawDescData = file_grpc_proto_proto_rawDesc
)

func file_grpc_proto_proto_rawDescGZIP() []byte {
	file_grpc_proto_proto_rawDescOnce.Do(func() {
		file_grpc_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_proto_rawDescData)
	})
	return file_grpc_proto_proto_rawDescData
}

var file_grpc_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_proto_proto_goTypes = []interface{}{
	(*PublishRequest)(nil),          // 0: publishRequest
	(*PublishResponse)(nil),         // 1: publishResponse
	(*BroadcastAllRequest)(nil),     // 2: broadcastAllRequest
	(*BroadcastAllresponse)(nil),    // 3: broadcastAllresponse
	(*NewClientJoinedRequest)(nil),  // 4: newClientJoinedRequest
	(*NewClientJoinedResponse)(nil), // 5: newClientJoinedResponse
	(*ClientLeaveRequest)(nil),      // 6: clientLeaveRequest
	(*ClientLeaveResponse)(nil),     // 7: clientLeaveResponse
}
var file_grpc_proto_proto_depIdxs = []int32{
	0, // 0: ChittyChatService.publishMessage:input_type -> publishRequest
	2, // 1: ChittyChatService.broadcastAll:input_type -> broadcastAllRequest
	4, // 2: ChittyChatService.newClientJoined:input_type -> newClientJoinedRequest
	6, // 3: ChittyChatService.clientLeave:input_type -> clientLeaveRequest
	1, // 4: ChittyChatService.publishMessage:output_type -> publishResponse
	3, // 5: ChittyChatService.broadcastAll:output_type -> broadcastAllresponse
	5, // 6: ChittyChatService.newClientJoined:output_type -> newClientJoinedResponse
	7, // 7: ChittyChatService.clientLeave:output_type -> clientLeaveResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_proto_init() }
func file_grpc_proto_proto_init() {
	if File_grpc_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastAllresponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewClientJoinedRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewClientJoinedResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientLeaveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientLeaveResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_proto_goTypes,
		DependencyIndexes: file_grpc_proto_proto_depIdxs,
		MessageInfos:      file_grpc_proto_proto_msgTypes,
	}.Build()
	File_grpc_proto_proto = out.File
	file_grpc_proto_proto_rawDesc = nil
	file_grpc_proto_proto_goTypes = nil
	file_grpc_proto_proto_depIdxs = nil
}