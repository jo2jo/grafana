// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: secretsmanager.proto

package secretsmanagerplugin

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

type SecretsGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyDescriptor *Key `protobuf:"bytes,1,opt,name=keyDescriptor,proto3" json:"keyDescriptor,omitempty"`
}

func (x *SecretsGetRequest) Reset() {
	*x = SecretsGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsGetRequest) ProtoMessage() {}

func (x *SecretsGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsGetRequest.ProtoReflect.Descriptor instead.
func (*SecretsGetRequest) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{0}
}

func (x *SecretsGetRequest) GetKeyDescriptor() *Key {
	if x != nil {
		return x.KeyDescriptor
	}
	return nil
}

type SecretsSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyDescriptor *Key   `protobuf:"bytes,1,opt,name=keyDescriptor,proto3" json:"keyDescriptor,omitempty"`
	Value         string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SecretsSetRequest) Reset() {
	*x = SecretsSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsSetRequest) ProtoMessage() {}

func (x *SecretsSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsSetRequest.ProtoReflect.Descriptor instead.
func (*SecretsSetRequest) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{1}
}

func (x *SecretsSetRequest) GetKeyDescriptor() *Key {
	if x != nil {
		return x.KeyDescriptor
	}
	return nil
}

func (x *SecretsSetRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SecretsDelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyDescriptor *Key `protobuf:"bytes,1,opt,name=keyDescriptor,proto3" json:"keyDescriptor,omitempty"`
}

func (x *SecretsDelRequest) Reset() {
	*x = SecretsDelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsDelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsDelRequest) ProtoMessage() {}

func (x *SecretsDelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsDelRequest.ProtoReflect.Descriptor instead.
func (*SecretsDelRequest) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{2}
}

func (x *SecretsDelRequest) GetKeyDescriptor() *Key {
	if x != nil {
		return x.KeyDescriptor
	}
	return nil
}

type SecretsKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyDescriptor    *Key `protobuf:"bytes,1,opt,name=keyDescriptor,proto3" json:"keyDescriptor,omitempty"`
	AllOrganizations bool `protobuf:"varint,2,opt,name=allOrganizations,proto3" json:"allOrganizations,omitempty"`
}

func (x *SecretsKeysRequest) Reset() {
	*x = SecretsKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsKeysRequest) ProtoMessage() {}

func (x *SecretsKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsKeysRequest.ProtoReflect.Descriptor instead.
func (*SecretsKeysRequest) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{3}
}

func (x *SecretsKeysRequest) GetKeyDescriptor() *Key {
	if x != nil {
		return x.KeyDescriptor
	}
	return nil
}

func (x *SecretsKeysRequest) GetAllOrganizations() bool {
	if x != nil {
		return x.AllOrganizations
	}
	return false
}

type SecretsRenameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyDescriptor *Key   `protobuf:"bytes,1,opt,name=keyDescriptor,proto3" json:"keyDescriptor,omitempty"`
	NewNamespace  string `protobuf:"bytes,4,opt,name=newNamespace,proto3" json:"newNamespace,omitempty"`
}

func (x *SecretsRenameRequest) Reset() {
	*x = SecretsRenameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsRenameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsRenameRequest) ProtoMessage() {}

func (x *SecretsRenameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsRenameRequest.ProtoReflect.Descriptor instead.
func (*SecretsRenameRequest) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{4}
}

func (x *SecretsRenameRequest) GetKeyDescriptor() *Key {
	if x != nil {
		return x.KeyDescriptor
	}
	return nil
}

func (x *SecretsRenameRequest) GetNewNamespace() string {
	if x != nil {
		return x.NewNamespace
	}
	return ""
}

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgId     int64  `protobuf:"varint,1,opt,name=orgId,proto3" json:"orgId,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Type      string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{5}
}

func (x *Key) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *Key) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Key) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type SecretsErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SecretsErrorResponse) Reset() {
	*x = SecretsErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsErrorResponse) ProtoMessage() {}

func (x *SecretsErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsErrorResponse.ProtoReflect.Descriptor instead.
func (*SecretsErrorResponse) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{6}
}

func (x *SecretsErrorResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type SecretsGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error          string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	DecryptedValue string `protobuf:"bytes,2,opt,name=decryptedValue,proto3" json:"decryptedValue,omitempty"`
	Exists         bool   `protobuf:"varint,3,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *SecretsGetResponse) Reset() {
	*x = SecretsGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsGetResponse) ProtoMessage() {}

func (x *SecretsGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsGetResponse.ProtoReflect.Descriptor instead.
func (*SecretsGetResponse) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{7}
}

func (x *SecretsGetResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SecretsGetResponse) GetDecryptedValue() string {
	if x != nil {
		return x.DecryptedValue
	}
	return ""
}

func (x *SecretsGetResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type SecretsKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Keys  []*Key `protobuf:"bytes,2,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *SecretsKeysResponse) Reset() {
	*x = SecretsKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_secretsmanager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretsKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretsKeysResponse) ProtoMessage() {}

func (x *SecretsKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_secretsmanager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretsKeysResponse.ProtoReflect.Descriptor instead.
func (*SecretsKeysResponse) Descriptor() ([]byte, []int) {
	return file_secretsmanager_proto_rawDescGZIP(), []int{8}
}

func (x *SecretsKeysResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *SecretsKeysResponse) GetKeys() []*Key {
	if x != nil {
		return x.Keys
	}
	return nil
}

var File_secretsmanager_proto protoreflect.FileDescriptor

var file_secretsmanager_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x54, 0x0a, 0x11,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x4b, 0x65, 0x79, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x22, 0x6a, 0x0a, 0x11, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x53, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x54,
	0x0a, 0x11, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x6f, 0x72, 0x22, 0x81, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x0d, 0x6b,
	0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x0d, 0x6b,
	0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x2a, 0x0a, 0x10,
	0x61, 0x6c, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x4f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7b, 0x0a, 0x14, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3f, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x4b,
	0x65, 0x79, 0x52, 0x0d, 0x6b, 0x65, 0x79, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f,
	0x72, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x2c, 0x0a, 0x14, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x22, 0x6a, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x26,
	0x0a, 0x0e, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x5a,
	0x0a, 0x13, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x6b,
	0x65, 0x79, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x4b, 0x65, 0x79, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x32, 0xe7, 0x03, 0x0a, 0x14, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a,
	0x03, 0x53, 0x65, 0x74, 0x12, 0x27, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x03, 0x44, 0x65, 0x6c,
	0x12, 0x27, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x44,
	0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x28, 0x2e,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x53,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x60, 0x0a, 0x06, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x2e, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17, 0x2e, 0x2f, 0x3b, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_secretsmanager_proto_rawDescOnce sync.Once
	file_secretsmanager_proto_rawDescData = file_secretsmanager_proto_rawDesc
)

func file_secretsmanager_proto_rawDescGZIP() []byte {
	file_secretsmanager_proto_rawDescOnce.Do(func() {
		file_secretsmanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_secretsmanager_proto_rawDescData)
	})
	return file_secretsmanager_proto_rawDescData
}

var file_secretsmanager_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_secretsmanager_proto_goTypes = []interface{}{
	(*SecretsGetRequest)(nil),    // 0: secretsmanagerplugin.SecretsGetRequest
	(*SecretsSetRequest)(nil),    // 1: secretsmanagerplugin.SecretsSetRequest
	(*SecretsDelRequest)(nil),    // 2: secretsmanagerplugin.SecretsDelRequest
	(*SecretsKeysRequest)(nil),   // 3: secretsmanagerplugin.SecretsKeysRequest
	(*SecretsRenameRequest)(nil), // 4: secretsmanagerplugin.SecretsRenameRequest
	(*Key)(nil),                  // 5: secretsmanagerplugin.Key
	(*SecretsErrorResponse)(nil), // 6: secretsmanagerplugin.SecretsErrorResponse
	(*SecretsGetResponse)(nil),   // 7: secretsmanagerplugin.SecretsGetResponse
	(*SecretsKeysResponse)(nil),  // 8: secretsmanagerplugin.SecretsKeysResponse
}
var file_secretsmanager_proto_depIdxs = []int32{
	5,  // 0: secretsmanagerplugin.SecretsGetRequest.keyDescriptor:type_name -> secretsmanagerplugin.Key
	5,  // 1: secretsmanagerplugin.SecretsSetRequest.keyDescriptor:type_name -> secretsmanagerplugin.Key
	5,  // 2: secretsmanagerplugin.SecretsDelRequest.keyDescriptor:type_name -> secretsmanagerplugin.Key
	5,  // 3: secretsmanagerplugin.SecretsKeysRequest.keyDescriptor:type_name -> secretsmanagerplugin.Key
	5,  // 4: secretsmanagerplugin.SecretsRenameRequest.keyDescriptor:type_name -> secretsmanagerplugin.Key
	5,  // 5: secretsmanagerplugin.SecretsKeysResponse.keys:type_name -> secretsmanagerplugin.Key
	0,  // 6: secretsmanagerplugin.RemoteSecretsManager.Get:input_type -> secretsmanagerplugin.SecretsGetRequest
	1,  // 7: secretsmanagerplugin.RemoteSecretsManager.Set:input_type -> secretsmanagerplugin.SecretsSetRequest
	2,  // 8: secretsmanagerplugin.RemoteSecretsManager.Del:input_type -> secretsmanagerplugin.SecretsDelRequest
	3,  // 9: secretsmanagerplugin.RemoteSecretsManager.Keys:input_type -> secretsmanagerplugin.SecretsKeysRequest
	4,  // 10: secretsmanagerplugin.RemoteSecretsManager.Rename:input_type -> secretsmanagerplugin.SecretsRenameRequest
	7,  // 11: secretsmanagerplugin.RemoteSecretsManager.Get:output_type -> secretsmanagerplugin.SecretsGetResponse
	6,  // 12: secretsmanagerplugin.RemoteSecretsManager.Set:output_type -> secretsmanagerplugin.SecretsErrorResponse
	6,  // 13: secretsmanagerplugin.RemoteSecretsManager.Del:output_type -> secretsmanagerplugin.SecretsErrorResponse
	8,  // 14: secretsmanagerplugin.RemoteSecretsManager.Keys:output_type -> secretsmanagerplugin.SecretsKeysResponse
	6,  // 15: secretsmanagerplugin.RemoteSecretsManager.Rename:output_type -> secretsmanagerplugin.SecretsErrorResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_secretsmanager_proto_init() }
func file_secretsmanager_proto_init() {
	if File_secretsmanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_secretsmanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsGetRequest); i {
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
		file_secretsmanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsSetRequest); i {
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
		file_secretsmanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsDelRequest); i {
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
		file_secretsmanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsKeysRequest); i {
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
		file_secretsmanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsRenameRequest); i {
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
		file_secretsmanager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_secretsmanager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsErrorResponse); i {
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
		file_secretsmanager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsGetResponse); i {
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
		file_secretsmanager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretsKeysResponse); i {
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
			RawDescriptor: file_secretsmanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_secretsmanager_proto_goTypes,
		DependencyIndexes: file_secretsmanager_proto_depIdxs,
		MessageInfos:      file_secretsmanager_proto_msgTypes,
	}.Build()
	File_secretsmanager_proto = out.File
	file_secretsmanager_proto_rawDesc = nil
	file_secretsmanager_proto_goTypes = nil
	file_secretsmanager_proto_depIdxs = nil
}
