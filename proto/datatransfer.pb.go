// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: proto/datatransfer.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Status int32

const (
	Status_SUCCESS        Status = 0
	Status_ERROR          Status = 1
	Status_ERROR_CHECKSUM Status = 2
	Status_ERROR_EXISTS   Status = 4
	Status_CHECKSUM_OK    Status = 6
	Status_IN_PROGRESS    Status = 12
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "SUCCESS",
		1:  "ERROR",
		2:  "ERROR_CHECKSUM",
		4:  "ERROR_EXISTS",
		6:  "CHECKSUM_OK",
		12: "IN_PROGRESS",
	}
	Status_value = map[string]int32{
		"SUCCESS":        0,
		"ERROR":          1,
		"ERROR_CHECKSUM": 2,
		"ERROR_EXISTS":   4,
		"CHECKSUM_OK":    6,
		"IN_PROGRESS":    12,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_datatransfer_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_proto_datatransfer_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_datatransfer_proto_rawDescGZIP(), []int{0}
}

type ClientWriteBlockProto_PipelineStatus int32

const (
	ClientWriteBlockProto_PIPELINE_SETUP_APPEND   ClientWriteBlockProto_PipelineStatus = 0 // is setting up
	ClientWriteBlockProto_PIPELINE_DATA_STREAMING ClientWriteBlockProto_PipelineStatus = 1 // work well
	ClientWriteBlockProto_PIPELINE_CLOSE          ClientWriteBlockProto_PipelineStatus = 2 // pipeline end working
)

// Enum value maps for ClientWriteBlockProto_PipelineStatus.
var (
	ClientWriteBlockProto_PipelineStatus_name = map[int32]string{
		0: "PIPELINE_SETUP_APPEND",
		1: "PIPELINE_DATA_STREAMING",
		2: "PIPELINE_CLOSE",
	}
	ClientWriteBlockProto_PipelineStatus_value = map[string]int32{
		"PIPELINE_SETUP_APPEND":   0,
		"PIPELINE_DATA_STREAMING": 1,
		"PIPELINE_CLOSE":          2,
	}
)

func (x ClientWriteBlockProto_PipelineStatus) Enum() *ClientWriteBlockProto_PipelineStatus {
	p := new(ClientWriteBlockProto_PipelineStatus)
	*p = x
	return p
}

func (x ClientWriteBlockProto_PipelineStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientWriteBlockProto_PipelineStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_datatransfer_proto_enumTypes[1].Descriptor()
}

func (ClientWriteBlockProto_PipelineStatus) Type() protoreflect.EnumType {
	return &file_proto_datatransfer_proto_enumTypes[1]
}

func (x ClientWriteBlockProto_PipelineStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientWriteBlockProto_PipelineStatus.Descriptor instead.
func (ClientWriteBlockProto_PipelineStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_datatransfer_proto_rawDescGZIP(), []int{2, 0}
}

type ClientOperationHeaderProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block    *ExtendedBlockInfoProto `protobuf:"bytes,1,opt,name=Block,proto3" json:"Block,omitempty"`
	ClientIP string                  `protobuf:"bytes,2,opt,name=ClientIP,proto3" json:"ClientIP,omitempty"`
}

func (x *ClientOperationHeaderProto) Reset() {
	*x = ClientOperationHeaderProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datatransfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientOperationHeaderProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientOperationHeaderProto) ProtoMessage() {}

func (x *ClientOperationHeaderProto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datatransfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientOperationHeaderProto.ProtoReflect.Descriptor instead.
func (*ClientOperationHeaderProto) Descriptor() ([]byte, []int) {
	return file_proto_datatransfer_proto_rawDescGZIP(), []int{0}
}

func (x *ClientOperationHeaderProto) GetBlock() *ExtendedBlockInfoProto {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *ClientOperationHeaderProto) GetClientIP() string {
	if x != nil {
		return x.ClientIP
	}
	return ""
}

type ClientReadBlockProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ClientOperationHeaderProto `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Offset uint32                      `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Length uint32                      `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *ClientReadBlockProto) Reset() {
	*x = ClientReadBlockProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datatransfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientReadBlockProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientReadBlockProto) ProtoMessage() {}

func (x *ClientReadBlockProto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datatransfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientReadBlockProto.ProtoReflect.Descriptor instead.
func (*ClientReadBlockProto) Descriptor() ([]byte, []int) {
	return file_proto_datatransfer_proto_rawDescGZIP(), []int{1}
}

func (x *ClientReadBlockProto) GetHeader() *ClientOperationHeaderProto {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ClientReadBlockProto) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ClientReadBlockProto) GetLength() uint32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ClientWriteBlockProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ClientOperationHeaderProto `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	// for Client send this proto(ClientWriteBlockProto)
	// thus, the first elem in Targets is the first DataNode
	Targets      []*DataNodeInfoProto                 `protobuf:"bytes,2,rep,name=Targets,proto3" json:"Targets,omitempty"`
	Status       ClientWriteBlockProto_PipelineStatus `protobuf:"varint,3,opt,name=Status,proto3,enum=proto.ClientWriteBlockProto_PipelineStatus" json:"Status,omitempty"`
	PipelineSize uint32                               `protobuf:"varint,4,opt,name=PipelineSize,proto3" json:"PipelineSize,omitempty"`
}

func (x *ClientWriteBlockProto) Reset() {
	*x = ClientWriteBlockProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datatransfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientWriteBlockProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientWriteBlockProto) ProtoMessage() {}

func (x *ClientWriteBlockProto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datatransfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientWriteBlockProto.ProtoReflect.Descriptor instead.
func (*ClientWriteBlockProto) Descriptor() ([]byte, []int) {
	return file_proto_datatransfer_proto_rawDescGZIP(), []int{2}
}

func (x *ClientWriteBlockProto) GetHeader() *ClientOperationHeaderProto {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ClientWriteBlockProto) GetTargets() []*DataNodeInfoProto {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *ClientWriteBlockProto) GetStatus() ClientWriteBlockProto_PipelineStatus {
	if x != nil {
		return x.Status
	}
	return ClientWriteBlockProto_PIPELINE_SETUP_APPEND
}

func (x *ClientWriteBlockProto) GetPipelineSize() uint32 {
	if x != nil {
		return x.PipelineSize
	}
	return 0
}

type DataNodeTransferBlockProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header *ClientOperationHeaderProto `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	// the first elem in Targets is the next DataNode
	Targets []*DataNodeInfoProto `protobuf:"bytes,2,rep,name=Targets,proto3" json:"Targets,omitempty"`
}

func (x *DataNodeTransferBlockProto) Reset() {
	*x = DataNodeTransferBlockProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datatransfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataNodeTransferBlockProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataNodeTransferBlockProto) ProtoMessage() {}

func (x *DataNodeTransferBlockProto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datatransfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataNodeTransferBlockProto.ProtoReflect.Descriptor instead.
func (*DataNodeTransferBlockProto) Descriptor() ([]byte, []int) {
	return file_proto_datatransfer_proto_rawDescGZIP(), []int{3}
}

func (x *DataNodeTransferBlockProto) GetHeader() *ClientOperationHeaderProto {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *DataNodeTransferBlockProto) GetTargets() []*DataNodeInfoProto {
	if x != nil {
		return x.Targets
	}
	return nil
}

type PipelineAckProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNo uint32   `protobuf:"varint,1,opt,name=SeqNo,proto3" json:"SeqNo,omitempty"`
	Reply []Status `protobuf:"varint,2,rep,packed,name=Reply,proto3,enum=proto.Status" json:"Reply,omitempty"` // append the former ACK.Reply
}

func (x *PipelineAckProto) Reset() {
	*x = PipelineAckProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_datatransfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipelineAckProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipelineAckProto) ProtoMessage() {}

func (x *PipelineAckProto) ProtoReflect() protoreflect.Message {
	mi := &file_proto_datatransfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipelineAckProto.ProtoReflect.Descriptor instead.
func (*PipelineAckProto) Descriptor() ([]byte, []int) {
	return file_proto_datatransfer_proto_rawDescGZIP(), []int{4}
}

func (x *PipelineAckProto) GetSeqNo() uint32 {
	if x != nil {
		return x.SeqNo
	}
	return 0
}

func (x *PipelineAckProto) GetReply() []Status {
	if x != nil {
		return x.Reply
	}
	return nil
}

var File_proto_datatransfer_proto protoreflect.FileDescriptor

var file_proto_datatransfer_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x64, 0x66, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x1a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x33, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52,
	0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x50, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x61,
	0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x06, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0xcd, 0x02, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x39, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x07, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12,
	0x43, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5c, 0x0a, 0x0e, 0x50, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x49,
	0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x55, 0x50, 0x5f, 0x41, 0x50, 0x50,
	0x45, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e,
	0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x54, 0x52, 0x45, 0x41, 0x4d, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x49, 0x50, 0x45, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x10, 0x02, 0x22, 0x8b, 0x01, 0x0a, 0x1a, 0x44, 0x61, 0x74, 0x61, 0x4e,
	0x6f, 0x64, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x07, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x07, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x22, 0x4d, 0x0a, 0x10, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x41, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x65, 0x71, 0x4e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x53, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x23,
	0x0a, 0x05, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2a, 0x68, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x43,
	0x48, 0x45, 0x43, 0x4b, 0x53, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x43,
	0x48, 0x45, 0x43, 0x4b, 0x53, 0x55, 0x4d, 0x5f, 0x4f, 0x4b, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b,
	0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x0c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_datatransfer_proto_rawDescOnce sync.Once
	file_proto_datatransfer_proto_rawDescData = file_proto_datatransfer_proto_rawDesc
)

func file_proto_datatransfer_proto_rawDescGZIP() []byte {
	file_proto_datatransfer_proto_rawDescOnce.Do(func() {
		file_proto_datatransfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_datatransfer_proto_rawDescData)
	})
	return file_proto_datatransfer_proto_rawDescData
}

var file_proto_datatransfer_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_datatransfer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_datatransfer_proto_goTypes = []interface{}{
	(Status)(0), // 0: proto.Status
	(ClientWriteBlockProto_PipelineStatus)(0), // 1: proto.ClientWriteBlockProto.PipelineStatus
	(*ClientOperationHeaderProto)(nil),        // 2: proto.ClientOperationHeaderProto
	(*ClientReadBlockProto)(nil),              // 3: proto.ClientReadBlockProto
	(*ClientWriteBlockProto)(nil),             // 4: proto.ClientWriteBlockProto
	(*DataNodeTransferBlockProto)(nil),        // 5: proto.DataNodeTransferBlockProto
	(*PipelineAckProto)(nil),                  // 6: proto.PipelineAckProto
	(*ExtendedBlockInfoProto)(nil),            // 7: proto.ExtendedBlockInfoProto
	(*DataNodeInfoProto)(nil),                 // 8: proto.DataNodeInfoProto
}
var file_proto_datatransfer_proto_depIdxs = []int32{
	7, // 0: proto.ClientOperationHeaderProto.Block:type_name -> proto.ExtendedBlockInfoProto
	2, // 1: proto.ClientReadBlockProto.Header:type_name -> proto.ClientOperationHeaderProto
	2, // 2: proto.ClientWriteBlockProto.Header:type_name -> proto.ClientOperationHeaderProto
	8, // 3: proto.ClientWriteBlockProto.Targets:type_name -> proto.DataNodeInfoProto
	1, // 4: proto.ClientWriteBlockProto.Status:type_name -> proto.ClientWriteBlockProto.PipelineStatus
	2, // 5: proto.DataNodeTransferBlockProto.Header:type_name -> proto.ClientOperationHeaderProto
	8, // 6: proto.DataNodeTransferBlockProto.Targets:type_name -> proto.DataNodeInfoProto
	0, // 7: proto.PipelineAckProto.Reply:type_name -> proto.Status
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_proto_datatransfer_proto_init() }
func file_proto_datatransfer_proto_init() {
	if File_proto_datatransfer_proto != nil {
		return
	}
	file_proto_hdfs_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_datatransfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientOperationHeaderProto); i {
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
		file_proto_datatransfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientReadBlockProto); i {
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
		file_proto_datatransfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientWriteBlockProto); i {
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
		file_proto_datatransfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataNodeTransferBlockProto); i {
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
		file_proto_datatransfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipelineAckProto); i {
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
			RawDescriptor: file_proto_datatransfer_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_datatransfer_proto_goTypes,
		DependencyIndexes: file_proto_datatransfer_proto_depIdxs,
		EnumInfos:         file_proto_datatransfer_proto_enumTypes,
		MessageInfos:      file_proto_datatransfer_proto_msgTypes,
	}.Build()
	File_proto_datatransfer_proto = out.File
	file_proto_datatransfer_proto_rawDesc = nil
	file_proto_datatransfer_proto_goTypes = nil
	file_proto_datatransfer_proto_depIdxs = nil
}