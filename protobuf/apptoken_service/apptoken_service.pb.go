// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: apptoken_service.proto

package apptoken_service

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

type IssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tracing_ID string `protobuf:"bytes,1,opt,name=tracing_ID,json=tracingID,proto3" json:"tracing_ID,omitempty"`
	CallerUuid string `protobuf:"bytes,2,opt,name=caller_uuid,json=callerUuid,proto3" json:"caller_uuid,omitempty"`
	AppUuid    string `protobuf:"bytes,3,opt,name=app_uuid,json=appUuid,proto3" json:"app_uuid,omitempty"`
	AppHash    string `protobuf:"bytes,4,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	AppOrigin  string `protobuf:"bytes,5,opt,name=app_origin,json=appOrigin,proto3" json:"app_origin,omitempty"`
}

func (x *IssueRequest) Reset() {
	*x = IssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apptoken_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueRequest) ProtoMessage() {}

func (x *IssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_apptoken_service_proto_msgTypes[0]
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
	return file_apptoken_service_proto_rawDescGZIP(), []int{0}
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

func (x *IssueRequest) GetAppHash() string {
	if x != nil {
		return x.AppHash
	}
	return ""
}

func (x *IssueRequest) GetAppOrigin() string {
	if x != nil {
		return x.AppOrigin
	}
	return ""
}

type IssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	Msg        string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token      string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *IssueResponse) Reset() {
	*x = IssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apptoken_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueResponse) ProtoMessage() {}

func (x *IssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_apptoken_service_proto_msgTypes[1]
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
	return file_apptoken_service_proto_rawDescGZIP(), []int{1}
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

func (x *IssueResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_apptoken_service_proto protoreflect.FileDescriptor

var file_apptoken_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x70, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x0c, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x63, 0x69, 0x6e, 0x67, 0x5f, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63,
	0x61, 0x6c, 0x6c, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x55, 0x75, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x70, 0x70, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x70, 0x70, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x22, 0x58,
	0x0a, 0x0d, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x34, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x05, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x0d, 0x2e,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apptoken_service_proto_rawDescOnce sync.Once
	file_apptoken_service_proto_rawDescData = file_apptoken_service_proto_rawDesc
)

func file_apptoken_service_proto_rawDescGZIP() []byte {
	file_apptoken_service_proto_rawDescOnce.Do(func() {
		file_apptoken_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_apptoken_service_proto_rawDescData)
	})
	return file_apptoken_service_proto_rawDescData
}

var file_apptoken_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apptoken_service_proto_goTypes = []interface{}{
	(*IssueRequest)(nil),  // 0: IssueRequest
	(*IssueResponse)(nil), // 1: IssueResponse
}
var file_apptoken_service_proto_depIdxs = []int32{
	0, // 0: AppToken.Issue:input_type -> IssueRequest
	1, // 1: AppToken.Issue:output_type -> IssueResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apptoken_service_proto_init() }
func file_apptoken_service_proto_init() {
	if File_apptoken_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apptoken_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_apptoken_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apptoken_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apptoken_service_proto_goTypes,
		DependencyIndexes: file_apptoken_service_proto_depIdxs,
		MessageInfos:      file_apptoken_service_proto_msgTypes,
	}.Build()
	File_apptoken_service_proto = out.File
	file_apptoken_service_proto_rawDesc = nil
	file_apptoken_service_proto_goTypes = nil
	file_apptoken_service_proto_depIdxs = nil
}
