// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact.proto

package puppet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ContactGender int32

const (
	ContactGender_CONTACT_GENDER_UNSPECIFIED ContactGender = 0
	ContactGender_CONTACT_GENDER_MALE        ContactGender = 1
	ContactGender_CONTACT_GENDER_FEMALE      ContactGender = 2
)

var ContactGender_name = map[int32]string{
	0: "CONTACT_GENDER_UNSPECIFIED",
	1: "CONTACT_GENDER_MALE",
	2: "CONTACT_GENDER_FEMALE",
}

var ContactGender_value = map[string]int32{
	"CONTACT_GENDER_UNSPECIFIED": 0,
	"CONTACT_GENDER_MALE":        1,
	"CONTACT_GENDER_FEMALE":      2,
}

func (x ContactGender) String() string {
	return proto.EnumName(ContactGender_name, int32(x))
}

func (ContactGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}

type ContactType int32

const (
	ContactType_CONTACT_TYPE_UNSPECIFIED ContactType = 0
	ContactType_CONTACT_TYPE_PERSONAL    ContactType = 1
	ContactType_CONTACT_TYPE_OFFICIAL    ContactType = 2
)

var ContactType_name = map[int32]string{
	0: "CONTACT_TYPE_UNSPECIFIED",
	1: "CONTACT_TYPE_PERSONAL",
	2: "CONTACT_TYPE_OFFICIAL",
}

var ContactType_value = map[string]int32{
	"CONTACT_TYPE_UNSPECIFIED": 0,
	"CONTACT_TYPE_PERSONAL":    1,
	"CONTACT_TYPE_OFFICIAL":    2,
}

func (x ContactType) String() string {
	return proto.EnumName(ContactType_name, int32(x))
}

func (ContactType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}

type ContactListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactListRequest) Reset()         { *m = ContactListRequest{} }
func (m *ContactListRequest) String() string { return proto.CompactTextString(m) }
func (*ContactListRequest) ProtoMessage()    {}
func (*ContactListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}

func (m *ContactListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactListRequest.Unmarshal(m, b)
}
func (m *ContactListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactListRequest.Marshal(b, m, deterministic)
}
func (m *ContactListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactListRequest.Merge(m, src)
}
func (m *ContactListRequest) XXX_Size() int {
	return xxx_messageInfo_ContactListRequest.Size(m)
}
func (m *ContactListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactListRequest proto.InternalMessageInfo

type ContactListResponse struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactListResponse) Reset()         { *m = ContactListResponse{} }
func (m *ContactListResponse) String() string { return proto.CompactTextString(m) }
func (*ContactListResponse) ProtoMessage()    {}
func (*ContactListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}

func (m *ContactListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactListResponse.Unmarshal(m, b)
}
func (m *ContactListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactListResponse.Marshal(b, m, deterministic)
}
func (m *ContactListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactListResponse.Merge(m, src)
}
func (m *ContactListResponse) XXX_Size() int {
	return xxx_messageInfo_ContactListResponse.Size(m)
}
func (m *ContactListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactListResponse proto.InternalMessageInfo

func (m *ContactListResponse) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type ContactPayloadRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactPayloadRequest) Reset()         { *m = ContactPayloadRequest{} }
func (m *ContactPayloadRequest) String() string { return proto.CompactTextString(m) }
func (*ContactPayloadRequest) ProtoMessage()    {}
func (*ContactPayloadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}

func (m *ContactPayloadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPayloadRequest.Unmarshal(m, b)
}
func (m *ContactPayloadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPayloadRequest.Marshal(b, m, deterministic)
}
func (m *ContactPayloadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPayloadRequest.Merge(m, src)
}
func (m *ContactPayloadRequest) XXX_Size() int {
	return xxx_messageInfo_ContactPayloadRequest.Size(m)
}
func (m *ContactPayloadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPayloadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPayloadRequest proto.InternalMessageInfo

func (m *ContactPayloadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ContactPayloadResponse struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Gender               ContactGender `protobuf:"varint,2,opt,name=gender,proto3,enum=wechaty.puppet.ContactGender" json:"gender,omitempty"`
	Type                 ContactType   `protobuf:"varint,3,opt,name=type,proto3,enum=wechaty.puppet.ContactType" json:"type,omitempty"`
	Name                 string        `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Avatar               string        `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Address              string        `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Alias                string        `protobuf:"bytes,7,opt,name=alias,proto3" json:"alias,omitempty"`
	City                 string        `protobuf:"bytes,8,opt,name=city,proto3" json:"city,omitempty"`
	Friend               bool          `protobuf:"varint,9,opt,name=friend,proto3" json:"friend,omitempty"`
	Province             string        `protobuf:"bytes,10,opt,name=province,proto3" json:"province,omitempty"`
	Signature            string        `protobuf:"bytes,11,opt,name=signature,proto3" json:"signature,omitempty"`
	Star                 bool          `protobuf:"varint,12,opt,name=star,proto3" json:"star,omitempty"`
	Weixin               string        `protobuf:"bytes,13,opt,name=weixin,proto3" json:"weixin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ContactPayloadResponse) Reset()         { *m = ContactPayloadResponse{} }
func (m *ContactPayloadResponse) String() string { return proto.CompactTextString(m) }
func (*ContactPayloadResponse) ProtoMessage()    {}
func (*ContactPayloadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{3}
}

func (m *ContactPayloadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactPayloadResponse.Unmarshal(m, b)
}
func (m *ContactPayloadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactPayloadResponse.Marshal(b, m, deterministic)
}
func (m *ContactPayloadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactPayloadResponse.Merge(m, src)
}
func (m *ContactPayloadResponse) XXX_Size() int {
	return xxx_messageInfo_ContactPayloadResponse.Size(m)
}
func (m *ContactPayloadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactPayloadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactPayloadResponse proto.InternalMessageInfo

func (m *ContactPayloadResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContactPayloadResponse) GetGender() ContactGender {
	if m != nil {
		return m.Gender
	}
	return ContactGender_CONTACT_GENDER_UNSPECIFIED
}

func (m *ContactPayloadResponse) GetType() ContactType {
	if m != nil {
		return m.Type
	}
	return ContactType_CONTACT_TYPE_UNSPECIFIED
}

func (m *ContactPayloadResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContactPayloadResponse) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *ContactPayloadResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContactPayloadResponse) GetAlias() string {
	if m != nil {
		return m.Alias
	}
	return ""
}

func (m *ContactPayloadResponse) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *ContactPayloadResponse) GetFriend() bool {
	if m != nil {
		return m.Friend
	}
	return false
}

func (m *ContactPayloadResponse) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *ContactPayloadResponse) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ContactPayloadResponse) GetStar() bool {
	if m != nil {
		return m.Star
	}
	return false
}

func (m *ContactPayloadResponse) GetWeixin() string {
	if m != nil {
		return m.Weixin
	}
	return ""
}

type ContactSelfQRCodeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactSelfQRCodeRequest) Reset()         { *m = ContactSelfQRCodeRequest{} }
func (m *ContactSelfQRCodeRequest) String() string { return proto.CompactTextString(m) }
func (*ContactSelfQRCodeRequest) ProtoMessage()    {}
func (*ContactSelfQRCodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{4}
}

func (m *ContactSelfQRCodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactSelfQRCodeRequest.Unmarshal(m, b)
}
func (m *ContactSelfQRCodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactSelfQRCodeRequest.Marshal(b, m, deterministic)
}
func (m *ContactSelfQRCodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactSelfQRCodeRequest.Merge(m, src)
}
func (m *ContactSelfQRCodeRequest) XXX_Size() int {
	return xxx_messageInfo_ContactSelfQRCodeRequest.Size(m)
}
func (m *ContactSelfQRCodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactSelfQRCodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactSelfQRCodeRequest proto.InternalMessageInfo

type ContactSelfQRCodeResponse struct {
	Qrcode               string   `protobuf:"bytes,1,opt,name=qrcode,proto3" json:"qrcode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactSelfQRCodeResponse) Reset()         { *m = ContactSelfQRCodeResponse{} }
func (m *ContactSelfQRCodeResponse) String() string { return proto.CompactTextString(m) }
func (*ContactSelfQRCodeResponse) ProtoMessage()    {}
func (*ContactSelfQRCodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{5}
}

func (m *ContactSelfQRCodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactSelfQRCodeResponse.Unmarshal(m, b)
}
func (m *ContactSelfQRCodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactSelfQRCodeResponse.Marshal(b, m, deterministic)
}
func (m *ContactSelfQRCodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactSelfQRCodeResponse.Merge(m, src)
}
func (m *ContactSelfQRCodeResponse) XXX_Size() int {
	return xxx_messageInfo_ContactSelfQRCodeResponse.Size(m)
}
func (m *ContactSelfQRCodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactSelfQRCodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactSelfQRCodeResponse proto.InternalMessageInfo

func (m *ContactSelfQRCodeResponse) GetQrcode() string {
	if m != nil {
		return m.Qrcode
	}
	return ""
}

type ContactSelfNameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactSelfNameRequest) Reset()         { *m = ContactSelfNameRequest{} }
func (m *ContactSelfNameRequest) String() string { return proto.CompactTextString(m) }
func (*ContactSelfNameRequest) ProtoMessage()    {}
func (*ContactSelfNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{6}
}

func (m *ContactSelfNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactSelfNameRequest.Unmarshal(m, b)
}
func (m *ContactSelfNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactSelfNameRequest.Marshal(b, m, deterministic)
}
func (m *ContactSelfNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactSelfNameRequest.Merge(m, src)
}
func (m *ContactSelfNameRequest) XXX_Size() int {
	return xxx_messageInfo_ContactSelfNameRequest.Size(m)
}
func (m *ContactSelfNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactSelfNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactSelfNameRequest proto.InternalMessageInfo

func (m *ContactSelfNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ContactSelfNameResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactSelfNameResponse) Reset()         { *m = ContactSelfNameResponse{} }
func (m *ContactSelfNameResponse) String() string { return proto.CompactTextString(m) }
func (*ContactSelfNameResponse) ProtoMessage()    {}
func (*ContactSelfNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{7}
}

func (m *ContactSelfNameResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactSelfNameResponse.Unmarshal(m, b)
}
func (m *ContactSelfNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactSelfNameResponse.Marshal(b, m, deterministic)
}
func (m *ContactSelfNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactSelfNameResponse.Merge(m, src)
}
func (m *ContactSelfNameResponse) XXX_Size() int {
	return xxx_messageInfo_ContactSelfNameResponse.Size(m)
}
func (m *ContactSelfNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactSelfNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactSelfNameResponse proto.InternalMessageInfo

type ContactSelfSignatureRequest struct {
	Signature            string   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactSelfSignatureRequest) Reset()         { *m = ContactSelfSignatureRequest{} }
func (m *ContactSelfSignatureRequest) String() string { return proto.CompactTextString(m) }
func (*ContactSelfSignatureRequest) ProtoMessage()    {}
func (*ContactSelfSignatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{8}
}

func (m *ContactSelfSignatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactSelfSignatureRequest.Unmarshal(m, b)
}
func (m *ContactSelfSignatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactSelfSignatureRequest.Marshal(b, m, deterministic)
}
func (m *ContactSelfSignatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactSelfSignatureRequest.Merge(m, src)
}
func (m *ContactSelfSignatureRequest) XXX_Size() int {
	return xxx_messageInfo_ContactSelfSignatureRequest.Size(m)
}
func (m *ContactSelfSignatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactSelfSignatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactSelfSignatureRequest proto.InternalMessageInfo

func (m *ContactSelfSignatureRequest) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

type ContactSelfSignatureResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactSelfSignatureResponse) Reset()         { *m = ContactSelfSignatureResponse{} }
func (m *ContactSelfSignatureResponse) String() string { return proto.CompactTextString(m) }
func (*ContactSelfSignatureResponse) ProtoMessage()    {}
func (*ContactSelfSignatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{9}
}

func (m *ContactSelfSignatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactSelfSignatureResponse.Unmarshal(m, b)
}
func (m *ContactSelfSignatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactSelfSignatureResponse.Marshal(b, m, deterministic)
}
func (m *ContactSelfSignatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactSelfSignatureResponse.Merge(m, src)
}
func (m *ContactSelfSignatureResponse) XXX_Size() int {
	return xxx_messageInfo_ContactSelfSignatureResponse.Size(m)
}
func (m *ContactSelfSignatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactSelfSignatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactSelfSignatureResponse proto.InternalMessageInfo

type ContactAliasRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// nullable
	Alias                *wrappers.StringValue `protobuf:"bytes,2,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ContactAliasRequest) Reset()         { *m = ContactAliasRequest{} }
func (m *ContactAliasRequest) String() string { return proto.CompactTextString(m) }
func (*ContactAliasRequest) ProtoMessage()    {}
func (*ContactAliasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{10}
}

func (m *ContactAliasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactAliasRequest.Unmarshal(m, b)
}
func (m *ContactAliasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactAliasRequest.Marshal(b, m, deterministic)
}
func (m *ContactAliasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactAliasRequest.Merge(m, src)
}
func (m *ContactAliasRequest) XXX_Size() int {
	return xxx_messageInfo_ContactAliasRequest.Size(m)
}
func (m *ContactAliasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactAliasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactAliasRequest proto.InternalMessageInfo

func (m *ContactAliasRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContactAliasRequest) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

type ContactAliasResponse struct {
	Alias                *wrappers.StringValue `protobuf:"bytes,1,opt,name=alias,proto3" json:"alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ContactAliasResponse) Reset()         { *m = ContactAliasResponse{} }
func (m *ContactAliasResponse) String() string { return proto.CompactTextString(m) }
func (*ContactAliasResponse) ProtoMessage()    {}
func (*ContactAliasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{11}
}

func (m *ContactAliasResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactAliasResponse.Unmarshal(m, b)
}
func (m *ContactAliasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactAliasResponse.Marshal(b, m, deterministic)
}
func (m *ContactAliasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactAliasResponse.Merge(m, src)
}
func (m *ContactAliasResponse) XXX_Size() int {
	return xxx_messageInfo_ContactAliasResponse.Size(m)
}
func (m *ContactAliasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactAliasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactAliasResponse proto.InternalMessageInfo

func (m *ContactAliasResponse) GetAlias() *wrappers.StringValue {
	if m != nil {
		return m.Alias
	}
	return nil
}

type ContactAvatarRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// nullable
	Filebox              *wrappers.StringValue `protobuf:"bytes,2,opt,name=filebox,proto3" json:"filebox,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ContactAvatarRequest) Reset()         { *m = ContactAvatarRequest{} }
func (m *ContactAvatarRequest) String() string { return proto.CompactTextString(m) }
func (*ContactAvatarRequest) ProtoMessage()    {}
func (*ContactAvatarRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{12}
}

func (m *ContactAvatarRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactAvatarRequest.Unmarshal(m, b)
}
func (m *ContactAvatarRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactAvatarRequest.Marshal(b, m, deterministic)
}
func (m *ContactAvatarRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactAvatarRequest.Merge(m, src)
}
func (m *ContactAvatarRequest) XXX_Size() int {
	return xxx_messageInfo_ContactAvatarRequest.Size(m)
}
func (m *ContactAvatarRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactAvatarRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ContactAvatarRequest proto.InternalMessageInfo

func (m *ContactAvatarRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContactAvatarRequest) GetFilebox() *wrappers.StringValue {
	if m != nil {
		return m.Filebox
	}
	return nil
}

type ContactAvatarResponse struct {
	Filebox              *wrappers.StringValue `protobuf:"bytes,1,opt,name=filebox,proto3" json:"filebox,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ContactAvatarResponse) Reset()         { *m = ContactAvatarResponse{} }
func (m *ContactAvatarResponse) String() string { return proto.CompactTextString(m) }
func (*ContactAvatarResponse) ProtoMessage()    {}
func (*ContactAvatarResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{13}
}

func (m *ContactAvatarResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactAvatarResponse.Unmarshal(m, b)
}
func (m *ContactAvatarResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactAvatarResponse.Marshal(b, m, deterministic)
}
func (m *ContactAvatarResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactAvatarResponse.Merge(m, src)
}
func (m *ContactAvatarResponse) XXX_Size() int {
	return xxx_messageInfo_ContactAvatarResponse.Size(m)
}
func (m *ContactAvatarResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactAvatarResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContactAvatarResponse proto.InternalMessageInfo

func (m *ContactAvatarResponse) GetFilebox() *wrappers.StringValue {
	if m != nil {
		return m.Filebox
	}
	return nil
}

func init() {
	proto.RegisterEnum("wechaty.puppet.ContactGender", ContactGender_name, ContactGender_value)
	proto.RegisterEnum("wechaty.puppet.ContactType", ContactType_name, ContactType_value)
	proto.RegisterType((*ContactListRequest)(nil), "wechaty.puppet.ContactListRequest")
	proto.RegisterType((*ContactListResponse)(nil), "wechaty.puppet.ContactListResponse")
	proto.RegisterType((*ContactPayloadRequest)(nil), "wechaty.puppet.ContactPayloadRequest")
	proto.RegisterType((*ContactPayloadResponse)(nil), "wechaty.puppet.ContactPayloadResponse")
	proto.RegisterType((*ContactSelfQRCodeRequest)(nil), "wechaty.puppet.ContactSelfQRCodeRequest")
	proto.RegisterType((*ContactSelfQRCodeResponse)(nil), "wechaty.puppet.ContactSelfQRCodeResponse")
	proto.RegisterType((*ContactSelfNameRequest)(nil), "wechaty.puppet.ContactSelfNameRequest")
	proto.RegisterType((*ContactSelfNameResponse)(nil), "wechaty.puppet.ContactSelfNameResponse")
	proto.RegisterType((*ContactSelfSignatureRequest)(nil), "wechaty.puppet.ContactSelfSignatureRequest")
	proto.RegisterType((*ContactSelfSignatureResponse)(nil), "wechaty.puppet.ContactSelfSignatureResponse")
	proto.RegisterType((*ContactAliasRequest)(nil), "wechaty.puppet.ContactAliasRequest")
	proto.RegisterType((*ContactAliasResponse)(nil), "wechaty.puppet.ContactAliasResponse")
	proto.RegisterType((*ContactAvatarRequest)(nil), "wechaty.puppet.ContactAvatarRequest")
	proto.RegisterType((*ContactAvatarResponse)(nil), "wechaty.puppet.ContactAvatarResponse")
}

func init() {
	proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15)
}

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 640 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x5f, 0x4f, 0xdb, 0x3c,
	0x14, 0xc6, 0xdf, 0x14, 0x28, 0xf4, 0xf0, 0x82, 0x2a, 0xf3, 0xcf, 0x94, 0x0e, 0x55, 0xb9, 0xa1,
	0x42, 0x53, 0x22, 0x81, 0xb6, 0x9b, 0x5d, 0x75, 0x25, 0x45, 0x9d, 0xba, 0xb6, 0x4b, 0xbb, 0x49,
	0xec, 0x62, 0xc8, 0x4d, 0xdc, 0x60, 0xa9, 0xc4, 0xc1, 0x71, 0x81, 0x7e, 0xf9, 0x69, 0x8a, 0xe3,
	0x94, 0x36, 0xa3, 0x12, 0xbb, 0xf3, 0xf1, 0xf3, 0xf8, 0x77, 0xec, 0x73, 0x6c, 0xc3, 0x8e, 0xc7,
	0x43, 0x49, 0x3c, 0x69, 0x45, 0x82, 0x4b, 0x8e, 0x76, 0x9f, 0xa8, 0x77, 0x47, 0xe4, 0xcc, 0x8a,
	0xa6, 0x51, 0x44, 0x65, 0xe5, 0x34, 0xe0, 0x3c, 0x98, 0x50, 0x5b, 0xa9, 0xa3, 0xe9, 0xd8, 0x7e,
	0x12, 0x24, 0x8a, 0xa8, 0x88, 0x53, 0xbf, 0xb9, 0x0f, 0xa8, 0x99, 0x02, 0x3a, 0x2c, 0x96, 0x2e,
	0x7d, 0x98, 0xd2, 0x58, 0x9a, 0x67, 0xb0, 0xb7, 0x34, 0x1b, 0x47, 0x3c, 0x8c, 0x29, 0x2a, 0xc3,
	0x1a, 0xf3, 0x63, 0x6c, 0xd4, 0xd6, 0xea, 0x25, 0x37, 0x19, 0x9a, 0x67, 0x70, 0xa0, 0x8d, 0x7d,
	0x32, 0x9b, 0x70, 0xe2, 0x6b, 0x02, 0xda, 0x85, 0x02, 0xf3, 0xb1, 0x51, 0x33, 0xea, 0x25, 0xb7,
	0xc0, 0x7c, 0xf3, 0x77, 0x01, 0x0e, 0xf3, 0x4e, 0x4d, 0xcd, 0x59, 0xd1, 0x07, 0x28, 0x06, 0x34,
	0xf4, 0xa9, 0xc0, 0x85, 0x9a, 0x51, 0xdf, 0xbd, 0x78, 0x67, 0x2d, 0x9f, 0xc9, 0xd2, 0x9c, 0x6b,
	0x65, 0x72, 0xb5, 0x19, 0xd9, 0xb0, 0x2e, 0x67, 0x11, 0xc5, 0x6b, 0x6a, 0xd1, 0xc9, 0x8a, 0x45,
	0xc3, 0x59, 0x44, 0x5d, 0x65, 0x44, 0x08, 0xd6, 0x43, 0x72, 0x4f, 0xf1, 0xba, 0xca, 0xac, 0xc6,
	0xe8, 0x10, 0x8a, 0xe4, 0x91, 0x48, 0x22, 0xf0, 0x86, 0x9a, 0xd5, 0x11, 0xc2, 0xb0, 0x49, 0x7c,
	0x5f, 0xd0, 0x38, 0xc6, 0x45, 0x25, 0x64, 0x21, 0xda, 0x87, 0x0d, 0x32, 0x61, 0x24, 0xc6, 0x9b,
	0x6a, 0x3e, 0x0d, 0x12, 0xb6, 0xc7, 0xe4, 0x0c, 0x6f, 0xa5, 0xec, 0x64, 0x9c, 0xb0, 0xc7, 0x82,
	0xd1, 0xd0, 0xc7, 0xa5, 0x9a, 0x51, 0xdf, 0x72, 0x75, 0x84, 0x2a, 0xb0, 0x15, 0x09, 0xfe, 0xc8,
	0x42, 0x8f, 0x62, 0x50, 0xfe, 0x79, 0x8c, 0xaa, 0x50, 0x8a, 0x59, 0x10, 0x12, 0x39, 0x15, 0x14,
	0x6f, 0x2b, 0xf1, 0x65, 0x22, 0xc9, 0x12, 0x27, 0x7b, 0xfd, 0x5f, 0xf1, 0xd4, 0x38, 0xc9, 0xf2,
	0x44, 0xd9, 0x33, 0x0b, 0xf1, 0x4e, 0x7a, 0x82, 0x34, 0x32, 0x2b, 0x80, 0x75, 0x09, 0x06, 0x74,
	0x32, 0xfe, 0xe6, 0x36, 0xb9, 0x4f, 0xb3, 0x76, 0x5f, 0xc2, 0xf1, 0x2b, 0x9a, 0x6e, 0xcf, 0x21,
	0x14, 0x1f, 0x84, 0xc7, 0x7d, 0xaa, 0x5b, 0xa4, 0x23, 0xf3, 0xfd, 0xbc, 0xa1, 0xc9, 0xa2, 0x2e,
	0xb9, 0xcf, 0x70, 0xf3, 0xc2, 0x1a, 0x2f, 0x85, 0x35, 0x8f, 0xe1, 0xe8, 0x2f, 0x77, 0x9a, 0xc0,
	0xfc, 0x04, 0x27, 0x0b, 0xd2, 0x20, 0x3b, 0x5d, 0x46, 0x5b, 0x2a, 0x81, 0x91, 0x2b, 0x81, 0x79,
	0x0a, 0xd5, 0xd7, 0x17, 0x6b, 0xf8, 0xcd, 0xfc, 0x26, 0x37, 0x92, 0xc6, 0xac, 0xb8, 0x9e, 0xe8,
	0x22, 0xeb, 0x62, 0x72, 0xe5, 0xb6, 0x2f, 0xaa, 0x56, 0xfa, 0x6c, 0xac, 0xec, 0xd9, 0x58, 0x03,
	0x29, 0x58, 0x18, 0xfc, 0x20, 0x93, 0x29, 0xd5, 0x3d, 0x36, 0xbf, 0xc0, 0xfe, 0x32, 0x5a, 0x17,
	0x6c, 0xce, 0x32, 0xde, 0xce, 0xfa, 0xf5, 0xc2, 0x52, 0x17, 0x6e, 0xd5, 0x3e, 0x3f, 0xc2, 0xe6,
	0x98, 0x4d, 0xe8, 0x88, 0x3f, 0xbf, 0x69, 0xa7, 0x99, 0xd9, 0xec, 0xcd, 0xdf, 0x69, 0xc6, 0xd7,
	0x9b, 0x5d, 0x00, 0x1a, 0xff, 0x00, 0x3c, 0xf7, 0x60, 0x67, 0xe9, 0x19, 0xa2, 0x53, 0xa8, 0x34,
	0x7b, 0xdd, 0x61, 0xa3, 0x39, 0xbc, 0xbd, 0x76, 0xba, 0x57, 0x8e, 0x7b, 0xfb, 0xbd, 0x3b, 0xe8,
	0x3b, 0xcd, 0x76, 0xab, 0xed, 0x5c, 0x95, 0xff, 0x43, 0x47, 0xb0, 0x97, 0xd3, 0xbf, 0x36, 0x3a,
	0x4e, 0xd9, 0x40, 0xc7, 0x70, 0x90, 0x13, 0x5a, 0x8e, 0x92, 0x0a, 0xe7, 0x04, 0xb6, 0x17, 0x9e,
	0x2d, 0xaa, 0x02, 0xce, 0x9c, 0xc3, 0x9b, 0xbe, 0x93, 0x4b, 0xb0, 0xc0, 0x51, 0x6a, 0xdf, 0x71,
	0x07, 0xbd, 0x6e, 0xa3, 0xb3, 0x9c, 0x42, 0x49, 0xbd, 0x56, 0xab, 0xdd, 0x6c, 0x37, 0x3a, 0xe5,
	0xc2, 0xe7, 0xf3, 0x9f, 0xf5, 0x80, 0xc9, 0xbb, 0xe9, 0xc8, 0xf2, 0xf8, 0xbd, 0x9d, 0xfc, 0x18,
	0x8c, 0xda, 0x81, 0x88, 0x3c, 0x3b, 0xe0, 0xb6, 0xfe, 0x42, 0xec, 0xf4, 0x0b, 0x19, 0x15, 0x55,
	0x49, 0x2e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x72, 0xb3, 0x5a, 0x73, 0x05, 0x00, 0x00,
}