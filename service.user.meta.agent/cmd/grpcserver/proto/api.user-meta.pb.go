// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: service.user.meta.agent/cmd/grpcserver/proto/api.user-meta.proto

package proto

import (
	common "github.com/KonstantinGasser/datalab/common"
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

// RPC CREATE USER
type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string           `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	User       *common.UserInfo `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *CreateRequest) GetUser() *common.UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{1}
}

func (x *CreateResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string         `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	CallerUuid string         `protobuf:"bytes,2,opt,name=caller_uuid,json=callerUuid,proto3" json:"caller_uuid,omitempty"`
	User       *UpdatableUser `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *UpdateRequest) GetCallerUuid() string {
	if x != nil {
		return x.CallerUuid
	}
	return ""
}

func (x *UpdateRequest) GetUser() *UpdatableUser {
	if x != nil {
		return x.User
	}
	return nil
}

type UpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateResponse) Reset() {
	*x = UpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateResponse) ProtoMessage() {}

func (x *UpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateResponse.ProtoReflect.Descriptor instead.
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *UpdateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	CallerUuid string `protobuf:"bytes,2,opt,name=caller_uuid,json=callerUuid,proto3" json:"caller_uuid,omitempty"`
	ForUuid    string `protobuf:"bytes,3,opt,name=for_uuid,json=forUuid,proto3" json:"for_uuid,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *GetRequest) GetCallerUuid() string {
	if x != nil {
		return x.CallerUuid
	}
	return ""
}

func (x *GetRequest) GetForUuid() string {
	if x != nil {
		return x.ForUuid
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32            `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string           `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	User       *common.UserInfo `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{5}
}

func (x *GetResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetResponse) GetUser() *common.UserInfo {
	if x != nil {
		return x.User
	}
	return nil
}

type GetColleaguesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID   string `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	UserUuid     string `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Organization string `protobuf:"bytes,3,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *GetColleaguesRequest) Reset() {
	*x = GetColleaguesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetColleaguesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetColleaguesRequest) ProtoMessage() {}

func (x *GetColleaguesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetColleaguesRequest.ProtoReflect.Descriptor instead.
func (*GetColleaguesRequest) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{6}
}

func (x *GetColleaguesRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *GetColleaguesRequest) GetUserUuid() string {
	if x != nil {
		return x.UserUuid
	}
	return ""
}

func (x *GetColleaguesRequest) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

type GetColleaguesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32              `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string             `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Colleagues []*common.UserInfo `protobuf:"bytes,3,rep,name=colleagues,proto3" json:"colleagues,omitempty"`
}

func (x *GetColleaguesResponse) Reset() {
	*x = GetColleaguesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetColleaguesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetColleaguesResponse) ProtoMessage() {}

func (x *GetColleaguesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetColleaguesResponse.ProtoReflect.Descriptor instead.
func (*GetColleaguesResponse) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{7}
}

func (x *GetColleaguesResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetColleaguesResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetColleaguesResponse) GetColleagues() []*common.UserInfo {
	if x != nil {
		return x.Colleagues
	}
	return nil
}

type UpdatableUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName     string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	OrgnPosition  string `protobuf:"bytes,3,opt,name=orgn_position,json=orgnPosition,proto3" json:"orgn_position,omitempty"`
	ProfileImgUrl string `protobuf:"bytes,4,opt,name=profile_img_url,json=profileImgUrl,proto3" json:"profile_img_url,omitempty"`
}

func (x *UpdatableUser) Reset() {
	*x = UpdatableUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatableUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatableUser) ProtoMessage() {}

func (x *UpdatableUser) ProtoReflect() protoreflect.Message {
	mi := &file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatableUser.ProtoReflect.Descriptor instead.
func (*UpdatableUser) Descriptor() ([]byte, []int) {
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP(), []int{8}
}

func (x *UpdatableUser) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdatableUser) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdatableUser) GetOrgnPosition() string {
	if x != nil {
		return x.OrgnPosition
	}
	return ""
}

func (x *UpdatableUser) GetProfileImgUrl() string {
	if x != nil {
		return x.ProfileImgUrl
	}
	return ""
}

var File_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto protoreflect.FileDescriptor

var file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDesc = []byte{
	0x0a, 0x40, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e,
	0x67, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x0e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x7e,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12,
	0x2d, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x43,
	0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x67, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x44,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6f, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x6f, 0x72, 0x55, 0x75, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x24,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x76, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x61, 0x67, 0x75, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7c, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x0d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x6e,
	0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x72, 0x67, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a,
	0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x6d, 0x67, 0x55, 0x72, 0x6c, 0x32, 0xa2, 0x02, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x41, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67,
	0x75, 0x65, 0x73, 0x12, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x52, 0x5a, 0x50, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6c, 0x61,
	0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescOnce sync.Once
	file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescData = file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDesc
)

func file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescGZIP() []byte {
	file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescOnce.Do(func() {
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescData)
	})
	return file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDescData
}

var file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),         // 0: user_proto.CreateRequest
	(*CreateResponse)(nil),        // 1: user_proto.CreateResponse
	(*UpdateRequest)(nil),         // 2: user_proto.UpdateRequest
	(*UpdateResponse)(nil),        // 3: user_proto.UpdateResponse
	(*GetRequest)(nil),            // 4: user_proto.GetRequest
	(*GetResponse)(nil),           // 5: user_proto.GetResponse
	(*GetColleaguesRequest)(nil),  // 6: user_proto.GetColleaguesRequest
	(*GetColleaguesResponse)(nil), // 7: user_proto.GetColleaguesResponse
	(*UpdatableUser)(nil),         // 8: user_proto.UpdatableUser
	(*common.UserInfo)(nil),       // 9: common.UserInfo
}
var file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_depIdxs = []int32{
	9, // 0: user_proto.CreateRequest.user:type_name -> common.UserInfo
	8, // 1: user_proto.UpdateRequest.user:type_name -> user_proto.UpdatableUser
	9, // 2: user_proto.GetResponse.user:type_name -> common.UserInfo
	9, // 3: user_proto.GetColleaguesResponse.colleagues:type_name -> common.UserInfo
	0, // 4: user_proto.UserMeta.Create:input_type -> user_proto.CreateRequest
	2, // 5: user_proto.UserMeta.Update:input_type -> user_proto.UpdateRequest
	4, // 6: user_proto.UserMeta.Get:input_type -> user_proto.GetRequest
	6, // 7: user_proto.UserMeta.GetColleagues:input_type -> user_proto.GetColleaguesRequest
	1, // 8: user_proto.UserMeta.Create:output_type -> user_proto.CreateResponse
	3, // 9: user_proto.UserMeta.Update:output_type -> user_proto.UpdateResponse
	5, // 10: user_proto.UserMeta.Get:output_type -> user_proto.GetResponse
	7, // 11: user_proto.UserMeta.GetColleagues:output_type -> user_proto.GetColleaguesResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_init() }
func file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_init() {
	if File_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateResponse); i {
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
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetColleaguesRequest); i {
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
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetColleaguesResponse); i {
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
		file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatableUser); i {
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
			RawDescriptor: file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_goTypes,
		DependencyIndexes: file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_depIdxs,
		MessageInfos:      file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_msgTypes,
	}.Build()
	File_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto = out.File
	file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_rawDesc = nil
	file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_goTypes = nil
	file_service_user_meta_agent_cmd_grpcserver_proto_api_user_meta_proto_depIdxs = nil
}
