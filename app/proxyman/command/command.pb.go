package command

import (
	v5 "github.com/v2fly/v2ray-core/v5"
	net "github.com/v2fly/v2ray-core/v5/common/net"
	protocol "github.com/v2fly/v2ray-core/v5/common/protocol"
	_ "github.com/v2fly/v2ray-core/v5/common/protoext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddUserOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *protocol.User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *AddUserOperation) Reset() {
	*x = AddUserOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserOperation) ProtoMessage() {}

func (x *AddUserOperation) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserOperation.ProtoReflect.Descriptor instead.
func (*AddUserOperation) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserOperation) GetUser() *protocol.User {
	if x != nil {
		return x.User
	}
	return nil
}

type RemoveUserOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *RemoveUserOperation) Reset() {
	*x = RemoveUserOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserOperation) ProtoMessage() {}

func (x *RemoveUserOperation) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserOperation.ProtoReflect.Descriptor instead.
func (*RemoveUserOperation) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{1}
}

func (x *RemoveUserOperation) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type AddVNextOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *protocol.ServerEndpoint `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *AddVNextOperation) Reset() {
	*x = AddVNextOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddVNextOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddVNextOperation) ProtoMessage() {}

func (x *AddVNextOperation) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddVNextOperation.ProtoReflect.Descriptor instead.
func (*AddVNextOperation) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{2}
}

func (x *AddVNextOperation) GetServer() *protocol.ServerEndpoint {
	if x != nil {
		return x.Server
	}
	return nil
}

type RemoveVNextOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *net.IPOrDomain `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32          `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *RemoveVNextOperation) Reset() {
	*x = RemoveVNextOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveVNextOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveVNextOperation) ProtoMessage() {}

func (x *RemoveVNextOperation) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveVNextOperation.ProtoReflect.Descriptor instead.
func (*RemoveVNextOperation) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{3}
}

func (x *RemoveVNextOperation) GetAddress() *net.IPOrDomain {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *RemoveVNextOperation) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type AddInboundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inbound *v5.InboundHandlerConfig `protobuf:"bytes,1,opt,name=inbound,proto3" json:"inbound,omitempty"`
}

func (x *AddInboundRequest) Reset() {
	*x = AddInboundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInboundRequest) ProtoMessage() {}

func (x *AddInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInboundRequest.ProtoReflect.Descriptor instead.
func (*AddInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{4}
}

func (x *AddInboundRequest) GetInbound() *v5.InboundHandlerConfig {
	if x != nil {
		return x.Inbound
	}
	return nil
}

type AddInboundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddInboundResponse) Reset() {
	*x = AddInboundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInboundResponse) ProtoMessage() {}

func (x *AddInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInboundResponse.ProtoReflect.Descriptor instead.
func (*AddInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{5}
}

type RemoveInboundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *RemoveInboundRequest) Reset() {
	*x = RemoveInboundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInboundRequest) ProtoMessage() {}

func (x *RemoveInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInboundRequest.ProtoReflect.Descriptor instead.
func (*RemoveInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveInboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type RemoveInboundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveInboundResponse) Reset() {
	*x = RemoveInboundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveInboundResponse) ProtoMessage() {}

func (x *RemoveInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveInboundResponse.ProtoReflect.Descriptor instead.
func (*RemoveInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{7}
}

type AlterInboundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag       string     `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Operation *anypb.Any `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *AlterInboundRequest) Reset() {
	*x = AlterInboundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterInboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterInboundRequest) ProtoMessage() {}

func (x *AlterInboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterInboundRequest.ProtoReflect.Descriptor instead.
func (*AlterInboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{8}
}

func (x *AlterInboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AlterInboundRequest) GetOperation() *anypb.Any {
	if x != nil {
		return x.Operation
	}
	return nil
}

type AlterInboundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AlterInboundResponse) Reset() {
	*x = AlterInboundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterInboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterInboundResponse) ProtoMessage() {}

func (x *AlterInboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterInboundResponse.ProtoReflect.Descriptor instead.
func (*AlterInboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{9}
}

type AddOutboundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Outbound *v5.OutboundHandlerConfig `protobuf:"bytes,1,opt,name=outbound,proto3" json:"outbound,omitempty"`
}

func (x *AddOutboundRequest) Reset() {
	*x = AddOutboundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOutboundRequest) ProtoMessage() {}

func (x *AddOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOutboundRequest.ProtoReflect.Descriptor instead.
func (*AddOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{10}
}

func (x *AddOutboundRequest) GetOutbound() *v5.OutboundHandlerConfig {
	if x != nil {
		return x.Outbound
	}
	return nil
}

type AddOutboundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddOutboundResponse) Reset() {
	*x = AddOutboundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOutboundResponse) ProtoMessage() {}

func (x *AddOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOutboundResponse.ProtoReflect.Descriptor instead.
func (*AddOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{11}
}

type RemoveOutboundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag string `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *RemoveOutboundRequest) Reset() {
	*x = RemoveOutboundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOutboundRequest) ProtoMessage() {}

func (x *RemoveOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOutboundRequest.ProtoReflect.Descriptor instead.
func (*RemoveOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{12}
}

func (x *RemoveOutboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type RemoveOutboundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveOutboundResponse) Reset() {
	*x = RemoveOutboundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveOutboundResponse) ProtoMessage() {}

func (x *RemoveOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveOutboundResponse.ProtoReflect.Descriptor instead.
func (*RemoveOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{13}
}

type AlterOutboundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tag       string     `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Operation *anypb.Any `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
}

func (x *AlterOutboundRequest) Reset() {
	*x = AlterOutboundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterOutboundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterOutboundRequest) ProtoMessage() {}

func (x *AlterOutboundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterOutboundRequest.ProtoReflect.Descriptor instead.
func (*AlterOutboundRequest) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{14}
}

func (x *AlterOutboundRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *AlterOutboundRequest) GetOperation() *anypb.Any {
	if x != nil {
		return x.Operation
	}
	return nil
}

type AlterOutboundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AlterOutboundResponse) Reset() {
	*x = AlterOutboundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterOutboundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterOutboundResponse) ProtoMessage() {}

func (x *AlterOutboundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterOutboundResponse.ProtoReflect.Descriptor instead.
func (*AlterOutboundResponse) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{15}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_proxyman_command_command_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_proxyman_command_command_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_proxyman_command_command_proto_rawDescGZIP(), []int{16}
}

var File_app_proxyman_command_command_proto protoreflect.FileDescriptor

var file_app_proxyman_command_command_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x74,
	0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x10, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x22, 0x57, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x56, 0x4e, 0x65, 0x78, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x67, 0x0a, 0x14, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x56, 0x4e, 0x65, 0x78, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3b, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x6e, 0x65, 0x74, 0x2e, 0x49, 0x50, 0x4f, 0x72, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x22, 0x4f, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x07, 0x69, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07, 0x69, 0x6e, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a, 0x14, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x22, 0x17, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5b, 0x0a, 0x13,
	0x41, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x32, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x6c, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x53, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x08, 0x6f, 0x75, 0x74, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x6f, 0x75,
	0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a,
	0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x22, 0x18, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5c, 0x0a, 0x14, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x32, 0x0a, 0x09,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x17, 0x0a, 0x15, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x3a, 0x1f, 0x82, 0xb5, 0x18, 0x0d, 0x0a, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x82, 0xb5, 0x18, 0x0a, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x6d, 0x61, 0x6e, 0x32, 0x90, 0x06, 0x0a, 0x0e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x49, 0x6e,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x32, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x49,
	0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x80, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x12, 0x35, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x0c, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x12, 0x34, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x7a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x33, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x75, 0x74, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x83,
	0x01, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x12, 0x36, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x80, 0x01, 0x0a, 0x0d, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x75,
	0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x35, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x75,
	0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e,
	0x41, 0x6c, 0x74, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7e, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x50, 0x01,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66,
	0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35,
	0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xaa, 0x02, 0x1f, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x6d, 0x61, 0x6e, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_proxyman_command_command_proto_rawDescOnce sync.Once
	file_app_proxyman_command_command_proto_rawDescData = file_app_proxyman_command_command_proto_rawDesc
)

func file_app_proxyman_command_command_proto_rawDescGZIP() []byte {
	file_app_proxyman_command_command_proto_rawDescOnce.Do(func() {
		file_app_proxyman_command_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_proxyman_command_command_proto_rawDescData)
	})
	return file_app_proxyman_command_command_proto_rawDescData
}

var file_app_proxyman_command_command_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_app_proxyman_command_command_proto_goTypes = []interface{}{
	(*AddUserOperation)(nil),         // 0: v2ray.core.app.proxyman.command.AddUserOperation
	(*RemoveUserOperation)(nil),      // 1: v2ray.core.app.proxyman.command.RemoveUserOperation
	(*AddVNextOperation)(nil),        // 2: v2ray.core.app.proxyman.command.AddVNextOperation
	(*RemoveVNextOperation)(nil),     // 3: v2ray.core.app.proxyman.command.RemoveVNextOperation
	(*AddInboundRequest)(nil),        // 4: v2ray.core.app.proxyman.command.AddInboundRequest
	(*AddInboundResponse)(nil),       // 5: v2ray.core.app.proxyman.command.AddInboundResponse
	(*RemoveInboundRequest)(nil),     // 6: v2ray.core.app.proxyman.command.RemoveInboundRequest
	(*RemoveInboundResponse)(nil),    // 7: v2ray.core.app.proxyman.command.RemoveInboundResponse
	(*AlterInboundRequest)(nil),      // 8: v2ray.core.app.proxyman.command.AlterInboundRequest
	(*AlterInboundResponse)(nil),     // 9: v2ray.core.app.proxyman.command.AlterInboundResponse
	(*AddOutboundRequest)(nil),       // 10: v2ray.core.app.proxyman.command.AddOutboundRequest
	(*AddOutboundResponse)(nil),      // 11: v2ray.core.app.proxyman.command.AddOutboundResponse
	(*RemoveOutboundRequest)(nil),    // 12: v2ray.core.app.proxyman.command.RemoveOutboundRequest
	(*RemoveOutboundResponse)(nil),   // 13: v2ray.core.app.proxyman.command.RemoveOutboundResponse
	(*AlterOutboundRequest)(nil),     // 14: v2ray.core.app.proxyman.command.AlterOutboundRequest
	(*AlterOutboundResponse)(nil),    // 15: v2ray.core.app.proxyman.command.AlterOutboundResponse
	(*Config)(nil),                   // 16: v2ray.core.app.proxyman.command.Config
	(*protocol.User)(nil),            // 17: v2ray.core.common.protocol.User
	(*protocol.ServerEndpoint)(nil),  // 18: v2ray.core.common.protocol.ServerEndpoint
	(*net.IPOrDomain)(nil),           // 19: v2ray.core.common.net.IPOrDomain
	(*v5.InboundHandlerConfig)(nil),  // 20: v2ray.core.InboundHandlerConfig
	(*anypb.Any)(nil),                // 21: google.protobuf.Any
	(*v5.OutboundHandlerConfig)(nil), // 22: v2ray.core.OutboundHandlerConfig
}
var file_app_proxyman_command_command_proto_depIdxs = []int32{
	17, // 0: v2ray.core.app.proxyman.command.AddUserOperation.user:type_name -> v2ray.core.common.protocol.User
	18, // 1: v2ray.core.app.proxyman.command.AddVNextOperation.server:type_name -> v2ray.core.common.protocol.ServerEndpoint
	19, // 2: v2ray.core.app.proxyman.command.RemoveVNextOperation.address:type_name -> v2ray.core.common.net.IPOrDomain
	20, // 3: v2ray.core.app.proxyman.command.AddInboundRequest.inbound:type_name -> v2ray.core.InboundHandlerConfig
	21, // 4: v2ray.core.app.proxyman.command.AlterInboundRequest.operation:type_name -> google.protobuf.Any
	22, // 5: v2ray.core.app.proxyman.command.AddOutboundRequest.outbound:type_name -> v2ray.core.OutboundHandlerConfig
	21, // 6: v2ray.core.app.proxyman.command.AlterOutboundRequest.operation:type_name -> google.protobuf.Any
	4,  // 7: v2ray.core.app.proxyman.command.HandlerService.AddInbound:input_type -> v2ray.core.app.proxyman.command.AddInboundRequest
	6,  // 8: v2ray.core.app.proxyman.command.HandlerService.RemoveInbound:input_type -> v2ray.core.app.proxyman.command.RemoveInboundRequest
	8,  // 9: v2ray.core.app.proxyman.command.HandlerService.AlterInbound:input_type -> v2ray.core.app.proxyman.command.AlterInboundRequest
	10, // 10: v2ray.core.app.proxyman.command.HandlerService.AddOutbound:input_type -> v2ray.core.app.proxyman.command.AddOutboundRequest
	12, // 11: v2ray.core.app.proxyman.command.HandlerService.RemoveOutbound:input_type -> v2ray.core.app.proxyman.command.RemoveOutboundRequest
	14, // 12: v2ray.core.app.proxyman.command.HandlerService.AlterOutbound:input_type -> v2ray.core.app.proxyman.command.AlterOutboundRequest
	5,  // 13: v2ray.core.app.proxyman.command.HandlerService.AddInbound:output_type -> v2ray.core.app.proxyman.command.AddInboundResponse
	7,  // 14: v2ray.core.app.proxyman.command.HandlerService.RemoveInbound:output_type -> v2ray.core.app.proxyman.command.RemoveInboundResponse
	9,  // 15: v2ray.core.app.proxyman.command.HandlerService.AlterInbound:output_type -> v2ray.core.app.proxyman.command.AlterInboundResponse
	11, // 16: v2ray.core.app.proxyman.command.HandlerService.AddOutbound:output_type -> v2ray.core.app.proxyman.command.AddOutboundResponse
	13, // 17: v2ray.core.app.proxyman.command.HandlerService.RemoveOutbound:output_type -> v2ray.core.app.proxyman.command.RemoveOutboundResponse
	15, // 18: v2ray.core.app.proxyman.command.HandlerService.AlterOutbound:output_type -> v2ray.core.app.proxyman.command.AlterOutboundResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_app_proxyman_command_command_proto_init() }
func file_app_proxyman_command_command_proto_init() {
	if File_app_proxyman_command_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_proxyman_command_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserOperation); i {
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
		file_app_proxyman_command_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveUserOperation); i {
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
		file_app_proxyman_command_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddVNextOperation); i {
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
		file_app_proxyman_command_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveVNextOperation); i {
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
		file_app_proxyman_command_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInboundRequest); i {
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
		file_app_proxyman_command_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInboundResponse); i {
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
		file_app_proxyman_command_command_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveInboundRequest); i {
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
		file_app_proxyman_command_command_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveInboundResponse); i {
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
		file_app_proxyman_command_command_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterInboundRequest); i {
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
		file_app_proxyman_command_command_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterInboundResponse); i {
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
		file_app_proxyman_command_command_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOutboundRequest); i {
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
		file_app_proxyman_command_command_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOutboundResponse); i {
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
		file_app_proxyman_command_command_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveOutboundRequest); i {
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
		file_app_proxyman_command_command_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveOutboundResponse); i {
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
		file_app_proxyman_command_command_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterOutboundRequest); i {
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
		file_app_proxyman_command_command_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterOutboundResponse); i {
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
		file_app_proxyman_command_command_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
			RawDescriptor: file_app_proxyman_command_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_proxyman_command_command_proto_goTypes,
		DependencyIndexes: file_app_proxyman_command_command_proto_depIdxs,
		MessageInfos:      file_app_proxyman_command_command_proto_msgTypes,
	}.Build()
	File_app_proxyman_command_command_proto = out.File
	file_app_proxyman_command_command_proto_rawDesc = nil
	file_app_proxyman_command_command_proto_goTypes = nil
	file_app_proxyman_command_command_proto_depIdxs = nil
}
