// (C) Copyright IBM Corp. 2022.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: hypervisor.proto

package __

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

type VersionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypervisor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hypervisor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_hypervisor_proto_rawDescGZIP(), []int{0}
}

func (x *VersionRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type VersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypervisor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hypervisor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_hypervisor_proto_rawDescGZIP(), []int{1}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type CreateVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Annotations          map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NetworkNamespacePath string            `protobuf:"bytes,3,opt,name=networkNamespacePath,proto3" json:"networkNamespacePath,omitempty"`
}

func (x *CreateVMRequest) Reset() {
	*x = CreateVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypervisor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVMRequest) ProtoMessage() {}

func (x *CreateVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hypervisor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVMRequest.ProtoReflect.Descriptor instead.
func (*CreateVMRequest) Descriptor() ([]byte, []int) {
	return file_hypervisor_proto_rawDescGZIP(), []int{2}
}

func (x *CreateVMRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateVMRequest) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *CreateVMRequest) GetNetworkNamespacePath() string {
	if x != nil {
		return x.NetworkNamespacePath
	}
	return ""
}

type CreateVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentSocketPath string `protobuf:"bytes,1,opt,name=agentSocketPath,proto3" json:"agentSocketPath,omitempty"`
}

func (x *CreateVMResponse) Reset() {
	*x = CreateVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypervisor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVMResponse) ProtoMessage() {}

func (x *CreateVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hypervisor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVMResponse.ProtoReflect.Descriptor instead.
func (*CreateVMResponse) Descriptor() ([]byte, []int) {
	return file_hypervisor_proto_rawDescGZIP(), []int{3}
}

func (x *CreateVMResponse) GetAgentSocketPath() string {
	if x != nil {
		return x.AgentSocketPath
	}
	return ""
}

type StartVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StartVMRequest) Reset() {
	*x = StartVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypervisor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartVMRequest) ProtoMessage() {}

func (x *StartVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hypervisor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartVMRequest.ProtoReflect.Descriptor instead.
func (*StartVMRequest) Descriptor() ([]byte, []int) {
	return file_hypervisor_proto_rawDescGZIP(), []int{4}
}

func (x *StartVMRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StartVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartVMResponse) Reset() {
	*x = StartVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypervisor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartVMResponse) ProtoMessage() {}

func (x *StartVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hypervisor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartVMResponse.ProtoReflect.Descriptor instead.
func (*StartVMResponse) Descriptor() ([]byte, []int) {
	return file_hypervisor_proto_rawDescGZIP(), []int{5}
}

type StopVMRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StopVMRequest) Reset() {
	*x = StopVMRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypervisor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopVMRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopVMRequest) ProtoMessage() {}

func (x *StopVMRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hypervisor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopVMRequest.ProtoReflect.Descriptor instead.
func (*StopVMRequest) Descriptor() ([]byte, []int) {
	return file_hypervisor_proto_rawDescGZIP(), []int{6}
}

func (x *StopVMRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type StopVMResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopVMResponse) Reset() {
	*x = StopVMResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hypervisor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopVMResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopVMResponse) ProtoMessage() {}

func (x *StopVMResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hypervisor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopVMResponse.ProtoReflect.Descriptor instead.
func (*StopVMResponse) Descriptor() ([]byte, []int) {
	return file_hypervisor_proto_rawDescGZIP(), []int{7}
}

var File_hypervisor_proto protoreflect.FileDescriptor

var file_hypervisor_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x22, 0x2a,
	0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x0f, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe5, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4e, 0x0a, 0x0b, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2c, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x50, 0x61, 0x74, 0x68, 0x1a,
	0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x3c, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x20, 0x0a,
	0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x11, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x70, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x70, 0x56, 0x4d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa4, 0x02, 0x0a, 0x0a, 0x48, 0x79, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x12, 0x47, 0x0a, 0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x4d,
	0x12, 0x1b, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x56, 0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a,
	0x07, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x4d, 0x12, 0x1a, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x74, 0x6f, 0x70, 0x56, 0x4d, 0x12, 0x19, 0x2e,
	0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x56,
	0x4d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x56, 0x4d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x68, 0x79, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hypervisor_proto_rawDescOnce sync.Once
	file_hypervisor_proto_rawDescData = file_hypervisor_proto_rawDesc
)

func file_hypervisor_proto_rawDescGZIP() []byte {
	file_hypervisor_proto_rawDescOnce.Do(func() {
		file_hypervisor_proto_rawDescData = protoimpl.X.CompressGZIP(file_hypervisor_proto_rawDescData)
	})
	return file_hypervisor_proto_rawDescData
}

var file_hypervisor_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_hypervisor_proto_goTypes = []interface{}{
	(*VersionRequest)(nil),   // 0: hypervisor.VersionRequest
	(*VersionResponse)(nil),  // 1: hypervisor.VersionResponse
	(*CreateVMRequest)(nil),  // 2: hypervisor.CreateVMRequest
	(*CreateVMResponse)(nil), // 3: hypervisor.CreateVMResponse
	(*StartVMRequest)(nil),   // 4: hypervisor.StartVMRequest
	(*StartVMResponse)(nil),  // 5: hypervisor.StartVMResponse
	(*StopVMRequest)(nil),    // 6: hypervisor.StopVMRequest
	(*StopVMResponse)(nil),   // 7: hypervisor.StopVMResponse
	nil,                      // 8: hypervisor.CreateVMRequest.AnnotationsEntry
}
var file_hypervisor_proto_depIdxs = []int32{
	8, // 0: hypervisor.CreateVMRequest.annotations:type_name -> hypervisor.CreateVMRequest.AnnotationsEntry
	2, // 1: hypervisor.Hypervisor.CreateVM:input_type -> hypervisor.CreateVMRequest
	4, // 2: hypervisor.Hypervisor.StartVM:input_type -> hypervisor.StartVMRequest
	6, // 3: hypervisor.Hypervisor.StopVM:input_type -> hypervisor.StopVMRequest
	0, // 4: hypervisor.Hypervisor.Version:input_type -> hypervisor.VersionRequest
	3, // 5: hypervisor.Hypervisor.CreateVM:output_type -> hypervisor.CreateVMResponse
	5, // 6: hypervisor.Hypervisor.StartVM:output_type -> hypervisor.StartVMResponse
	7, // 7: hypervisor.Hypervisor.StopVM:output_type -> hypervisor.StopVMResponse
	1, // 8: hypervisor.Hypervisor.Version:output_type -> hypervisor.VersionResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_hypervisor_proto_init() }
func file_hypervisor_proto_init() {
	if File_hypervisor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hypervisor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionRequest); i {
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
		file_hypervisor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionResponse); i {
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
		file_hypervisor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVMRequest); i {
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
		file_hypervisor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVMResponse); i {
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
		file_hypervisor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartVMRequest); i {
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
		file_hypervisor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartVMResponse); i {
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
		file_hypervisor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopVMRequest); i {
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
		file_hypervisor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopVMResponse); i {
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
			RawDescriptor: file_hypervisor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hypervisor_proto_goTypes,
		DependencyIndexes: file_hypervisor_proto_depIdxs,
		MessageInfos:      file_hypervisor_proto_msgTypes,
	}.Build()
	File_hypervisor_proto = out.File
	file_hypervisor_proto_rawDesc = nil
	file_hypervisor_proto_goTypes = nil
	file_hypervisor_proto_depIdxs = nil
}
