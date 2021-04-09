// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.6
// source: proto/manage.proto

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

type ConfigRequest_Field int32

const (
	ConfigRequest_ELECTRONIC_VOTING ConfigRequest_Field = 0
)

// Enum value maps for ConfigRequest_Field.
var (
	ConfigRequest_Field_name = map[int32]string{
		0: "ELECTRONIC_VOTING",
	}
	ConfigRequest_Field_value = map[string]int32{
		"ELECTRONIC_VOTING": 0,
	}
)

func (x ConfigRequest_Field) Enum() *ConfigRequest_Field {
	p := new(ConfigRequest_Field)
	*p = x
	return p
}

func (x ConfigRequest_Field) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigRequest_Field) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_manage_proto_enumTypes[0].Descriptor()
}

func (ConfigRequest_Field) Type() protoreflect.EnumType {
	return &file_proto_manage_proto_enumTypes[0]
}

func (x ConfigRequest_Field) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigRequest_Field.Descriptor instead.
func (ConfigRequest_Field) EnumDescriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{6, 0}
}

type CheckServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SkipClient bool `protobuf:"varint,1,opt,name=skipClient,proto3" json:"skipClient,omitempty"`
}

func (x *CheckServerRequest) Reset() {
	*x = CheckServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckServerRequest) ProtoMessage() {}

func (x *CheckServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckServerRequest.ProtoReflect.Descriptor instead.
func (*CheckServerRequest) Descriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{0}
}

func (x *CheckServerRequest) GetSkipClient() bool {
	if x != nil {
		return x.SkipClient
	}
	return false
}

type CheckServerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CheckServerResponse) Reset() {
	*x = CheckServerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckServerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckServerResponse) ProtoMessage() {}

func (x *CheckServerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckServerResponse.ProtoReflect.Descriptor instead.
func (*CheckServerResponse) Descriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{1}
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username                    string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password                    string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	OrganisationManagementLevel string `protobuf:"bytes,3,opt,name=organisation_management_level,json=organisationManagementLevel,proto3" json:"organisation_management_level,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetOrganisationManagementLevel() string {
	if x != nil {
		return x.OrganisationManagementLevel
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int64 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserResponse) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type SetPasswordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SetPasswordRequest) Reset() {
	*x = SetPasswordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordRequest) ProtoMessage() {}

func (x *SetPasswordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordRequest.ProtoReflect.Descriptor instead.
func (*SetPasswordRequest) Descriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{4}
}

func (x *SetPasswordRequest) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *SetPasswordRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SetPasswordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPasswordResponse) Reset() {
	*x = SetPasswordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordResponse) ProtoMessage() {}

func (x *SetPasswordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordResponse.ProtoReflect.Descriptor instead.
func (*SetPasswordResponse) Descriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{5}
}

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field    ConfigRequest_Field `protobuf:"varint,1,opt,name=field,proto3,enum=ConfigRequest_Field" json:"field,omitempty"`
	NewValue string              `protobuf:"bytes,2,opt,name=newValue,proto3" json:"newValue,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{6}
}

func (x *ConfigRequest) GetField() ConfigRequest_Field {
	if x != nil {
		return x.Field
	}
	return ConfigRequest_ELECTRONIC_VOTING
}

func (x *ConfigRequest) GetNewValue() string {
	if x != nil {
		return x.NewValue
	}
	return ""
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_manage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_manage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_manage_proto_rawDescGZIP(), []int{7}
}

func (x *ConfigResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_proto_manage_proto protoreflect.FileDescriptor

var file_proto_manage_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6b,
	0x69, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x73, 0x6b, 0x69, 0x70, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x42, 0x0a, 0x1d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x22, 0x2c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x22, 0x48, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x53,
	0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x77, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x05, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52, 0x4f, 0x4e,
	0x49, 0x43, 0x5f, 0x56, 0x4f, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x22, 0x26, 0x0a, 0x0e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0xde, 0x01, 0x0a, 0x06, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x12, 0x38,
	0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x13, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x38, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x13,
	0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x0e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x53, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x2f, 0x6f, 0x70,
	0x65, 0x6e, 0x73, 0x6c, 0x69, 0x64, 0x65, 0x73, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_manage_proto_rawDescOnce sync.Once
	file_proto_manage_proto_rawDescData = file_proto_manage_proto_rawDesc
)

func file_proto_manage_proto_rawDescGZIP() []byte {
	file_proto_manage_proto_rawDescOnce.Do(func() {
		file_proto_manage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_manage_proto_rawDescData)
	})
	return file_proto_manage_proto_rawDescData
}

var file_proto_manage_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_manage_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_manage_proto_goTypes = []interface{}{
	(ConfigRequest_Field)(0),    // 0: ConfigRequest.Field
	(*CheckServerRequest)(nil),  // 1: CheckServerRequest
	(*CheckServerResponse)(nil), // 2: CheckServerResponse
	(*CreateUserRequest)(nil),   // 3: CreateUserRequest
	(*CreateUserResponse)(nil),  // 4: CreateUserResponse
	(*SetPasswordRequest)(nil),  // 5: SetPasswordRequest
	(*SetPasswordResponse)(nil), // 6: SetPasswordResponse
	(*ConfigRequest)(nil),       // 7: ConfigRequest
	(*ConfigResponse)(nil),      // 8: ConfigResponse
}
var file_proto_manage_proto_depIdxs = []int32{
	0, // 0: ConfigRequest.field:type_name -> ConfigRequest.Field
	1, // 1: Manage.CheckServer:input_type -> CheckServerRequest
	3, // 2: Manage.CreateUser:input_type -> CreateUserRequest
	5, // 3: Manage.SetPassword:input_type -> SetPasswordRequest
	7, // 4: Manage.Config:input_type -> ConfigRequest
	2, // 5: Manage.CheckServer:output_type -> CheckServerResponse
	4, // 6: Manage.CreateUser:output_type -> CreateUserResponse
	6, // 7: Manage.SetPassword:output_type -> SetPasswordResponse
	8, // 8: Manage.Config:output_type -> ConfigResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_manage_proto_init() }
func file_proto_manage_proto_init() {
	if File_proto_manage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_manage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckServerRequest); i {
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
		file_proto_manage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckServerResponse); i {
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
		file_proto_manage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_proto_manage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_proto_manage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordRequest); i {
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
		file_proto_manage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordResponse); i {
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
		file_proto_manage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_proto_manage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
			RawDescriptor: file_proto_manage_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_manage_proto_goTypes,
		DependencyIndexes: file_proto_manage_proto_depIdxs,
		EnumInfos:         file_proto_manage_proto_enumTypes,
		MessageInfos:      file_proto_manage_proto_msgTypes,
	}.Build()
	File_proto_manage_proto = out.File
	file_proto_manage_proto_rawDesc = nil
	file_proto_manage_proto_goTypes = nil
	file_proto_manage_proto_depIdxs = nil
}
