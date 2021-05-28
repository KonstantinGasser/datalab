// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: service.app.token.agent/cmd/grpcserver/proto/api.app-token.proto

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

type InitializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	AppRefUuid string `protobuf:"bytes,2,opt,name=app_ref_uuid,json=appRefUuid,proto3" json:"app_ref_uuid,omitempty"`
	AppOwner   string `protobuf:"bytes,3,opt,name=app_owner,json=appOwner,proto3" json:"app_owner,omitempty"`
	AppHash    string `protobuf:"bytes,4,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	AppOrigin  string `protobuf:"bytes,5,opt,name=app_origin,json=appOrigin,proto3" json:"app_origin,omitempty"`
}

func (x *InitializeRequest) Reset() {
	*x = InitializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest) ProtoMessage() {}

func (x *InitializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest.ProtoReflect.Descriptor instead.
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *InitializeRequest) GetAppRefUuid() string {
	if x != nil {
		return x.AppRefUuid
	}
	return ""
}

func (x *InitializeRequest) GetAppOwner() string {
	if x != nil {
		return x.AppOwner
	}
	return ""
}

func (x *InitializeRequest) GetAppHash() string {
	if x != nil {
		return x.AppHash
	}
	return ""
}

func (x *InitializeRequest) GetAppOrigin() string {
	if x != nil {
		return x.AppOrigin
	}
	return ""
}

type InitializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *InitializeResponse) Reset() {
	*x = InitializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeResponse) ProtoMessage() {}

func (x *InitializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeResponse.ProtoReflect.Descriptor instead.
func (*InitializeResponse) Descriptor() ([]byte, []int) {
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP(), []int{1}
}

func (x *InitializeResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *InitializeResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type IssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	CallerUuid string `protobuf:"bytes,2,opt,name=caller_uuid,json=callerUuid,proto3" json:"caller_uuid,omitempty"`
	AppUuid    string `protobuf:"bytes,3,opt,name=app_uuid,json=appUuid,proto3" json:"app_uuid,omitempty"`
}

func (x *IssueRequest) Reset() {
	*x = IssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueRequest) ProtoMessage() {}

func (x *IssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueRequest.ProtoReflect.Descriptor instead.
func (*IssueRequest) Descriptor() ([]byte, []int) {
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP(), []int{2}
}

func (x *IssueRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *IssueRequest) GetCallerUuid() string {
	if x != nil {
		return x.CallerUuid
	}
	return ""
}

func (x *IssueRequest) GetAppUuid() string {
	if x != nil {
		return x.AppUuid
	}
	return ""
}

type IssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32                  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token      *common.AppAccessToken `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *IssueResponse) Reset() {
	*x = IssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueResponse) ProtoMessage() {}

func (x *IssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueResponse.ProtoReflect.Descriptor instead.
func (*IssueResponse) Descriptor() ([]byte, []int) {
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP(), []int{3}
}

func (x *IssueResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *IssueResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *IssueResponse) GetToken() *common.AppAccessToken {
	if x != nil {
		return x.Token
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string             `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	AuthedUser *common.AuthedUser `protobuf:"bytes,2,opt,name=authed_user,json=authedUser,proto3" json:"authed_user,omitempty"`
	AppUuid    string             `protobuf:"bytes,3,opt,name=app_uuid,json=appUuid,proto3" json:"app_uuid,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[4]
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
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *GetRequest) GetAuthedUser() *common.AuthedUser {
	if x != nil {
		return x.AuthedUser
	}
	return nil
}

func (x *GetRequest) GetAppUuid() string {
	if x != nil {
		return x.AppUuid
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32                  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token      *common.AppAccessToken `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[5]
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
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP(), []int{5}
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

func (x *GetResponse) GetToken() *common.AppAccessToken {
	if x != nil {
		return x.Token
	}
	return nil
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string             `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	AuthedUser *common.AuthedUser `protobuf:"bytes,2,opt,name=authed_user,json=authedUser,proto3" json:"authed_user,omitempty"`
	AppToken   string             `protobuf:"bytes,3,opt,name=app_token,json=appToken,proto3" json:"app_token,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateRequest) GetTracing_ID() string {
	if x != nil {
		return x.Tracing_ID
	}
	return ""
}

func (x *ValidateRequest) GetAuthedUser() *common.AuthedUser {
	if x != nil {
		return x.AuthedUser
	}
	return nil
}

func (x *ValidateRequest) GetAppToken() string {
	if x != nil {
		return x.AppToken
	}
	return ""
}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	AppUuid    string `protobuf:"bytes,3,opt,name=app_uuid,json=appUuid,proto3" json:"app_uuid,omitempty"`
	AppOrigin  string `protobuf:"bytes,4,opt,name=app_origin,json=appOrigin,proto3" json:"app_origin,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP(), []int{7}
}

func (x *ValidateResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ValidateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ValidateResponse) GetAppUuid() string {
	if x != nil {
		return x.AppUuid
	}
	return ""
}

func (x *ValidateResponse) GetAppOrigin() string {
	if x != nil {
		return x.AppOrigin
	}
	return ""
}

var File_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto protoreflect.FileDescriptor

var file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDesc = []byte{
	0x0a, 0x40, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x2d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x70,
	0x70, 0x5f, 0x72, 0x65, 0x66, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x61, 0x70, 0x70, 0x52, 0x65, 0x66, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x61, 0x70, 0x70, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x70, 0x70, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x22, 0x47, 0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x69, 0x0a, 0x0c,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x70, 0x70, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x70, 0x70, 0x55, 0x75, 0x69, 0x64, 0x22, 0x70, 0x0a, 0x0d, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7b, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x69,
	0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x65, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x0a, 0x61, 0x75, 0x74, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x61,
	0x70, 0x70, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x70, 0x70, 0x55, 0x75, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2c, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x70, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0b, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x70, 0x70, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x61, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x7f, 0x0a, 0x10, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x70, 0x70, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x32, 0xac, 0x02, 0x0a,
	0x08, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x51, 0x0a, 0x0a, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x05,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x1a, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4b, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x52, 0x5a, 0x50, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x6f, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x6c,
	0x61, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescOnce sync.Once
	file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescData = file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDesc
)

func file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescGZIP() []byte {
	file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescOnce.Do(func() {
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescData)
	})
	return file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDescData
}

var file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_goTypes = []interface{}{
	(*InitializeRequest)(nil),     // 0: issuer_proto.InitializeRequest
	(*InitializeResponse)(nil),    // 1: issuer_proto.InitializeResponse
	(*IssueRequest)(nil),          // 2: issuer_proto.IssueRequest
	(*IssueResponse)(nil),         // 3: issuer_proto.IssueResponse
	(*GetRequest)(nil),            // 4: issuer_proto.GetRequest
	(*GetResponse)(nil),           // 5: issuer_proto.GetResponse
	(*ValidateRequest)(nil),       // 6: issuer_proto.ValidateRequest
	(*ValidateResponse)(nil),      // 7: issuer_proto.ValidateResponse
	(*common.AppAccessToken)(nil), // 8: common.AppAccessToken
	(*common.AuthedUser)(nil),     // 9: common.AuthedUser
}
var file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_depIdxs = []int32{
	8, // 0: issuer_proto.IssueResponse.token:type_name -> common.AppAccessToken
	9, // 1: issuer_proto.GetRequest.authed_user:type_name -> common.AuthedUser
	8, // 2: issuer_proto.GetResponse.token:type_name -> common.AppAccessToken
	9, // 3: issuer_proto.ValidateRequest.authed_user:type_name -> common.AuthedUser
	0, // 4: issuer_proto.AppToken.Initialize:input_type -> issuer_proto.InitializeRequest
	2, // 5: issuer_proto.AppToken.Issue:input_type -> issuer_proto.IssueRequest
	6, // 6: issuer_proto.AppToken.Validate:input_type -> issuer_proto.ValidateRequest
	4, // 7: issuer_proto.AppToken.Get:input_type -> issuer_proto.GetRequest
	1, // 8: issuer_proto.AppToken.Initialize:output_type -> issuer_proto.InitializeResponse
	3, // 9: issuer_proto.AppToken.Issue:output_type -> issuer_proto.IssueResponse
	7, // 10: issuer_proto.AppToken.Validate:output_type -> issuer_proto.ValidateResponse
	5, // 11: issuer_proto.AppToken.Get:output_type -> issuer_proto.GetResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_init() }
func file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_init() {
	if File_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest); i {
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
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeResponse); i {
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
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueRequest); i {
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
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueResponse); i {
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
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
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
		file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
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
			RawDescriptor: file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_goTypes,
		DependencyIndexes: file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_depIdxs,
		MessageInfos:      file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_msgTypes,
	}.Build()
	File_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto = out.File
	file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_rawDesc = nil
	file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_goTypes = nil
	file_service_app_token_agent_cmd_grpcserver_proto_api_app_token_proto_depIdxs = nil
}