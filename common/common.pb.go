// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: common/common.proto

package common

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

type AppRole int32

const (
	AppRole_OWNER  AppRole = 0
	AppRole_EDITOR AppRole = 1
	AppRole_VIEWER AppRole = 2
)

// Enum value maps for AppRole.
var (
	AppRole_name = map[int32]string{
		0: "OWNER",
		1: "EDITOR",
		2: "VIEWER",
	}
	AppRole_value = map[string]int32{
		"OWNER":  0,
		"EDITOR": 1,
		"VIEWER": 2,
	}
)

func (x AppRole) Enum() *AppRole {
	p := new(AppRole)
	*p = x
	return p
}

func (x AppRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppRole) Descriptor() protoreflect.EnumDescriptor {
	return file_common_common_proto_enumTypes[0].Descriptor()
}

func (AppRole) Type() protoreflect.EnumType {
	return &file_common_common_proto_enumTypes[0]
}

func (x AppRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppRole.Descriptor instead.
func (AppRole) EnumDescriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg  string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{0}
}

func (x *Hello) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Hello) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type AppInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string       `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name        string       `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	URL         string       `protobuf:"bytes,3,opt,name=URL,proto3" json:"URL,omitempty"`
	Description string       `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Owner       string       `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Member      []*AppMember `protobuf:"bytes,6,rep,name=member,proto3" json:"member,omitempty"`
}

func (x *AppInfo) Reset() {
	*x = AppInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfo) ProtoMessage() {}

func (x *AppInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfo.ProtoReflect.Descriptor instead.
func (*AppInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{1}
}

func (x *AppInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AppInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppInfo) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *AppInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AppInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *AppInfo) GetMember() []*AppMember {
	if x != nil {
		return x.Member
	}
	return nil
}

type AppMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid   string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Role   string `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Status int32  `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *AppMember) Reset() {
	*x = AppMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMember) ProtoMessage() {}

func (x *AppMember) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMember.ProtoReflect.Descriptor instead.
func (*AppMember) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{2}
}

func (x *AppMember) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *AppMember) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *AppMember) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type AppMetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Uuid string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *AppMetaInfo) Reset() {
	*x = AppMetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMetaInfo) ProtoMessage() {}

func (x *AppMetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMetaInfo.ProtoReflect.Descriptor instead.
func (*AppMetaInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{3}
}

func (x *AppMetaInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppMetaInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type AppConfigInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Funnel   []*Funnel   `protobuf:"bytes,1,rep,name=funnel,proto3" json:"funnel,omitempty"`
	Campaign []*Campaign `protobuf:"bytes,2,rep,name=campaign,proto3" json:"campaign,omitempty"`
	BtnTime  []*BtnTime  `protobuf:"bytes,3,rep,name=btn_time,json=btnTime,proto3" json:"btn_time,omitempty"`
}

func (x *AppConfigInfo) Reset() {
	*x = AppConfigInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppConfigInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppConfigInfo) ProtoMessage() {}

func (x *AppConfigInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppConfigInfo.ProtoReflect.Descriptor instead.
func (*AppConfigInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{4}
}

func (x *AppConfigInfo) GetFunnel() []*Funnel {
	if x != nil {
		return x.Funnel
	}
	return nil
}

func (x *AppConfigInfo) GetCampaign() []*Campaign {
	if x != nil {
		return x.Campaign
	}
	return nil
}

func (x *AppConfigInfo) GetBtnTime() []*BtnTime {
	if x != nil {
		return x.BtnTime
	}
	return nil
}

type Funnel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Transition string `protobuf:"bytes,3,opt,name=transition,proto3" json:"transition,omitempty"`
}

func (x *Funnel) Reset() {
	*x = Funnel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Funnel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Funnel) ProtoMessage() {}

func (x *Funnel) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Funnel.ProtoReflect.Descriptor instead.
func (*Funnel) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{5}
}

func (x *Funnel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Funnel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Funnel) GetTransition() string {
	if x != nil {
		return x.Transition
	}
	return ""
}

type Campaign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Prefix string `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *Campaign) Reset() {
	*x = Campaign{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Campaign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Campaign) ProtoMessage() {}

func (x *Campaign) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Campaign.ProtoReflect.Descriptor instead.
func (*Campaign) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{6}
}

func (x *Campaign) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Campaign) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Campaign) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type BtnTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BtnName string `protobuf:"bytes,3,opt,name=btn_name,json=btnName,proto3" json:"btn_name,omitempty"`
}

func (x *BtnTime) Reset() {
	*x = BtnTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BtnTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BtnTime) ProtoMessage() {}

func (x *BtnTime) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BtnTime.ProtoReflect.Descriptor instead.
func (*BtnTime) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{7}
}

func (x *BtnTime) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BtnTime) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BtnTime) GetBtnName() string {
	if x != nil {
		return x.BtnName
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName     string `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	OrgnDomain    string `protobuf:"bytes,5,opt,name=orgn_domain,json=orgnDomain,proto3" json:"orgn_domain,omitempty"`
	OrgnPosition  string `protobuf:"bytes,6,opt,name=orgn_position,json=orgnPosition,proto3" json:"orgn_position,omitempty"`
	ProfileImgUrl string `protobuf:"bytes,7,opt,name=profile_img_url,json=profileImgUrl,proto3" json:"profile_img_url,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{8}
}

func (x *UserInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserInfo) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserInfo) GetOrgnDomain() string {
	if x != nil {
		return x.OrgnDomain
	}
	return ""
}

func (x *UserInfo) GetOrgnPosition() string {
	if x != nil {
		return x.OrgnPosition
	}
	return ""
}

func (x *UserInfo) GetProfileImgUrl() string {
	if x != nil {
		return x.ProfileImgUrl
	}
	return ""
}

type UserMetaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *UserMetaInfo) Reset() {
	*x = UserMetaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMetaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMetaInfo) ProtoMessage() {}

func (x *UserMetaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMetaInfo.ProtoReflect.Descriptor instead.
func (*UserMetaInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{9}
}

func (x *UserMetaInfo) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserMetaInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type AppTokenInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Exp   int64  `protobuf:"varint,2,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (x *AppTokenInfo) Reset() {
	*x = AppTokenInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppTokenInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppTokenInfo) ProtoMessage() {}

func (x *AppTokenInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppTokenInfo.ProtoReflect.Descriptor instead.
func (*AppTokenInfo) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{10}
}

func (x *AppTokenInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *AppTokenInfo) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

type TokenClaims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string           `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Organization string           `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
	Permissions  *UserPermissions `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *TokenClaims) Reset() {
	*x = TokenClaims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenClaims) ProtoMessage() {}

func (x *TokenClaims) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenClaims.ProtoReflect.Descriptor instead.
func (*TokenClaims) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{11}
}

func (x *TokenClaims) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *TokenClaims) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *TokenClaims) GetPermissions() *UserPermissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type UserPermissions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Apps []*AppPermission `protobuf:"bytes,1,rep,name=apps,proto3" json:"apps,omitempty"`
}

func (x *UserPermissions) Reset() {
	*x = UserPermissions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPermissions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPermissions) ProtoMessage() {}

func (x *UserPermissions) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPermissions.ProtoReflect.Descriptor instead.
func (*UserPermissions) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{12}
}

func (x *UserPermissions) GetApps() []*AppPermission {
	if x != nil {
		return x.Apps
	}
	return nil
}

type AppPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppUuid string  `protobuf:"bytes,1,opt,name=app_uuid,json=appUuid,proto3" json:"app_uuid,omitempty"`
	Role    AppRole `protobuf:"varint,2,opt,name=role,proto3,enum=common.AppRole" json:"role,omitempty"`
}

func (x *AppPermission) Reset() {
	*x = AppPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_common_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppPermission) ProtoMessage() {}

func (x *AppPermission) ProtoReflect() protoreflect.Message {
	mi := &file_common_common_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppPermission.ProtoReflect.Descriptor instead.
func (*AppPermission) Descriptor() ([]byte, []int) {
	return file_common_common_proto_rawDescGZIP(), []int{13}
}

func (x *AppPermission) GetAppUuid() string {
	if x != nil {
		return x.AppUuid
	}
	return ""
}

func (x *AppPermission) GetRole() AppRole {
	if x != nil {
		return x.Role
	}
	return AppRole_OWNER
}

var File_common_common_proto protoreflect.FileDescriptor

var file_common_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x2d, 0x0a,
	0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0xa6, 0x01, 0x0a,
	0x07, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55,
	0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x35, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x0d, 0x41, 0x70,
	0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x06, 0x66,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x06, 0x66, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67,
	0x6e, 0x12, 0x2a, 0x0a, 0x08, 0x62, 0x74, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x74, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x07, 0x62, 0x74, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x4c, 0x0a,
	0x06, 0x46, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x46, 0x0a, 0x08, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x22, 0x48, 0x0a, 0x07, 0x42, 0x74, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x74, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x74, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xe4, 0x01,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x67, 0x6e, 0x5f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x67,
	0x6e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x6e, 0x5f,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x72, 0x67, 0x6e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6d, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d,
	0x67, 0x55, 0x72, 0x6c, 0x22, 0x3e, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78,
	0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x65, 0x78, 0x70, 0x22, 0x80, 0x01, 0x0a,
	0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22,
	0x3c, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x61, 0x70, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x61, 0x70, 0x70, 0x73, 0x22, 0x4f, 0x0a,
	0x0d, 0x41, 0x70, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x70, 0x70, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x70, 0x55, 0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x2a, 0x2c,
	0x0a, 0x07, 0x41, 0x70, 0x70, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e,
	0x45, 0x52, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x44, 0x49, 0x54, 0x4f, 0x52, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x56, 0x49, 0x45, 0x57, 0x45, 0x52, 0x10, 0x02, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4b, 0x6f, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x73, 0x65, 0x72, 0x2f, 0x64, 0x61, 0x74, 0x61,
	0x6c, 0x61, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_common_proto_rawDescOnce sync.Once
	file_common_common_proto_rawDescData = file_common_common_proto_rawDesc
)

func file_common_common_proto_rawDescGZIP() []byte {
	file_common_common_proto_rawDescOnce.Do(func() {
		file_common_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_common_proto_rawDescData)
	})
	return file_common_common_proto_rawDescData
}

var file_common_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_common_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_common_common_proto_goTypes = []interface{}{
	(AppRole)(0),            // 0: common.AppRole
	(*Hello)(nil),           // 1: common.Hello
	(*AppInfo)(nil),         // 2: common.AppInfo
	(*AppMember)(nil),       // 3: common.AppMember
	(*AppMetaInfo)(nil),     // 4: common.AppMetaInfo
	(*AppConfigInfo)(nil),   // 5: common.AppConfigInfo
	(*Funnel)(nil),          // 6: common.Funnel
	(*Campaign)(nil),        // 7: common.Campaign
	(*BtnTime)(nil),         // 8: common.BtnTime
	(*UserInfo)(nil),        // 9: common.UserInfo
	(*UserMetaInfo)(nil),    // 10: common.UserMetaInfo
	(*AppTokenInfo)(nil),    // 11: common.AppTokenInfo
	(*TokenClaims)(nil),     // 12: common.TokenClaims
	(*UserPermissions)(nil), // 13: common.UserPermissions
	(*AppPermission)(nil),   // 14: common.AppPermission
}
var file_common_common_proto_depIdxs = []int32{
	3,  // 0: common.AppInfo.member:type_name -> common.AppMember
	6,  // 1: common.AppConfigInfo.funnel:type_name -> common.Funnel
	7,  // 2: common.AppConfigInfo.campaign:type_name -> common.Campaign
	8,  // 3: common.AppConfigInfo.btn_time:type_name -> common.BtnTime
	13, // 4: common.TokenClaims.permissions:type_name -> common.UserPermissions
	14, // 5: common.UserPermissions.apps:type_name -> common.AppPermission
	0,  // 6: common.AppPermission.role:type_name -> common.AppRole
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_common_common_proto_init() }
func file_common_common_proto_init() {
	if File_common_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_common_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfo); i {
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
		file_common_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMember); i {
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
		file_common_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMetaInfo); i {
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
		file_common_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppConfigInfo); i {
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
		file_common_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Funnel); i {
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
		file_common_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Campaign); i {
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
		file_common_common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BtnTime); i {
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
		file_common_common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_common_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMetaInfo); i {
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
		file_common_common_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppTokenInfo); i {
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
		file_common_common_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenClaims); i {
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
		file_common_common_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPermissions); i {
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
		file_common_common_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppPermission); i {
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
			RawDescriptor: file_common_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_common_proto_goTypes,
		DependencyIndexes: file_common_common_proto_depIdxs,
		EnumInfos:         file_common_common_proto_enumTypes,
		MessageInfos:      file_common_common_proto_msgTypes,
	}.Build()
	File_common_common_proto = out.File
	file_common_common_proto_rawDesc = nil
	file_common_common_proto_goTypes = nil
	file_common_common_proto_depIdxs = nil
}
