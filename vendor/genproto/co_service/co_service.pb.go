// Code generated by protoc-gen-go. DO NOT EDIT.
// source: co_service.proto

package co_service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{0}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{1}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetAllRequest struct {
	Page                 uint64   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAllRequest) Reset()         { *m = GetAllRequest{} }
func (m *GetAllRequest) String() string { return proto.CompactTextString(m) }
func (*GetAllRequest) ProtoMessage()    {}
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{2}
}

func (m *GetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllRequest.Unmarshal(m, b)
}
func (m *GetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllRequest.Marshal(b, m, deterministic)
}
func (m *GetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllRequest.Merge(m, src)
}
func (m *GetAllRequest) XXX_Size() int {
	return xxx_messageInfo_GetAllRequest.Size(m)
}
func (m *GetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllRequest proto.InternalMessageInfo

func (m *GetAllRequest) GetPage() uint64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *GetAllRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetAllResponse struct {
	Owners               []*CargoOwner `protobuf:"bytes,1,rep,name=owners,proto3" json:"owners,omitempty"`
	Count                uint64        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetAllResponse) Reset()         { *m = GetAllResponse{} }
func (m *GetAllResponse) String() string { return proto.CompactTextString(m) }
func (*GetAllResponse) ProtoMessage()    {}
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{3}
}

func (m *GetAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAllResponse.Unmarshal(m, b)
}
func (m *GetAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAllResponse.Marshal(b, m, deterministic)
}
func (m *GetAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAllResponse.Merge(m, src)
}
func (m *GetAllResponse) XXX_Size() int {
	return xxx_messageInfo_GetAllResponse.Size(m)
}
func (m *GetAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAllResponse proto.InternalMessageInfo

func (m *GetAllResponse) GetOwners() []*CargoOwner {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *GetAllResponse) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type AddDistributorRequest struct {
	CoId                 string   `protobuf:"bytes,1,opt,name=co_id,json=coId,proto3" json:"co_id,omitempty"`
	DistributorId        string   `protobuf:"bytes,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDistributorRequest) Reset()         { *m = AddDistributorRequest{} }
func (m *AddDistributorRequest) String() string { return proto.CompactTextString(m) }
func (*AddDistributorRequest) ProtoMessage()    {}
func (*AddDistributorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{4}
}

func (m *AddDistributorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDistributorRequest.Unmarshal(m, b)
}
func (m *AddDistributorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDistributorRequest.Marshal(b, m, deterministic)
}
func (m *AddDistributorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDistributorRequest.Merge(m, src)
}
func (m *AddDistributorRequest) XXX_Size() int {
	return xxx_messageInfo_AddDistributorRequest.Size(m)
}
func (m *AddDistributorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDistributorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddDistributorRequest proto.InternalMessageInfo

func (m *AddDistributorRequest) GetCoId() string {
	if m != nil {
		return m.CoId
	}
	return ""
}

func (m *AddDistributorRequest) GetDistributorId() string {
	if m != nil {
		return m.DistributorId
	}
	return ""
}

type DeleteDistributorRequest struct {
	CoId                 string   `protobuf:"bytes,1,opt,name=co_id,json=coId,proto3" json:"co_id,omitempty"`
	DistributorId        string   `protobuf:"bytes,2,opt,name=distributor_id,json=distributorId,proto3" json:"distributor_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDistributorRequest) Reset()         { *m = DeleteDistributorRequest{} }
func (m *DeleteDistributorRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDistributorRequest) ProtoMessage()    {}
func (*DeleteDistributorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{5}
}

func (m *DeleteDistributorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDistributorRequest.Unmarshal(m, b)
}
func (m *DeleteDistributorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDistributorRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDistributorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDistributorRequest.Merge(m, src)
}
func (m *DeleteDistributorRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDistributorRequest.Size(m)
}
func (m *DeleteDistributorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDistributorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDistributorRequest proto.InternalMessageInfo

func (m *DeleteDistributorRequest) GetCoId() string {
	if m != nil {
		return m.CoId
	}
	return ""
}

func (m *DeleteDistributorRequest) GetDistributorId() string {
	if m != nil {
		return m.DistributorId
	}
	return ""
}

type CheckExistsRequest struct {
	Column               string   `protobuf:"bytes,1,opt,name=column,proto3" json:"column,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckExistsRequest) Reset()         { *m = CheckExistsRequest{} }
func (m *CheckExistsRequest) String() string { return proto.CompactTextString(m) }
func (*CheckExistsRequest) ProtoMessage()    {}
func (*CheckExistsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{6}
}

func (m *CheckExistsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckExistsRequest.Unmarshal(m, b)
}
func (m *CheckExistsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckExistsRequest.Marshal(b, m, deterministic)
}
func (m *CheckExistsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckExistsRequest.Merge(m, src)
}
func (m *CheckExistsRequest) XXX_Size() int {
	return xxx_messageInfo_CheckExistsRequest.Size(m)
}
func (m *CheckExistsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckExistsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckExistsRequest proto.InternalMessageInfo

func (m *CheckExistsRequest) GetColumn() string {
	if m != nil {
		return m.Column
	}
	return ""
}

func (m *CheckExistsRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CheckExistsResponse struct {
	IsExists             bool     `protobuf:"varint,1,opt,name=is_exists,json=isExists,proto3" json:"is_exists,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckExistsResponse) Reset()         { *m = CheckExistsResponse{} }
func (m *CheckExistsResponse) String() string { return proto.CompactTextString(m) }
func (*CheckExistsResponse) ProtoMessage()    {}
func (*CheckExistsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{7}
}

func (m *CheckExistsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckExistsResponse.Unmarshal(m, b)
}
func (m *CheckExistsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckExistsResponse.Marshal(b, m, deterministic)
}
func (m *CheckExistsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckExistsResponse.Merge(m, src)
}
func (m *CheckExistsResponse) XXX_Size() int {
	return xxx_messageInfo_CheckExistsResponse.Size(m)
}
func (m *CheckExistsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckExistsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckExistsResponse proto.InternalMessageInfo

func (m *CheckExistsResponse) GetIsExists() bool {
	if m != nil {
		return m.IsExists
	}
	return false
}

type LoginRequest struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_825073e23d8cf421, []int{8}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "genproto.GetRequest")
	proto.RegisterType((*DeleteRequest)(nil), "genproto.DeleteRequest")
	proto.RegisterType((*GetAllRequest)(nil), "genproto.GetAllRequest")
	proto.RegisterType((*GetAllResponse)(nil), "genproto.GetAllResponse")
	proto.RegisterType((*AddDistributorRequest)(nil), "genproto.AddDistributorRequest")
	proto.RegisterType((*DeleteDistributorRequest)(nil), "genproto.DeleteDistributorRequest")
	proto.RegisterType((*CheckExistsRequest)(nil), "genproto.CheckExistsRequest")
	proto.RegisterType((*CheckExistsResponse)(nil), "genproto.CheckExistsResponse")
	proto.RegisterType((*LoginRequest)(nil), "genproto.LoginRequest")
}

func init() { proto.RegisterFile("co_service.proto", fileDescriptor_825073e23d8cf421) }

var fileDescriptor_825073e23d8cf421 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x41, 0x6f, 0xda, 0x4c,
	0x10, 0x86, 0x81, 0x10, 0x0b, 0x86, 0x0f, 0xf4, 0x75, 0x49, 0x28, 0x72, 0x52, 0x25, 0x5a, 0xb5,
	0x52, 0x0e, 0x95, 0x91, 0xc8, 0xa5, 0x55, 0x9b, 0x43, 0x20, 0x08, 0x45, 0x8d, 0x94, 0x96, 0xa4,
	0x3d, 0xf4, 0x82, 0x8c, 0x3d, 0x75, 0x57, 0x31, 0x5e, 0xd7, 0xbb, 0x4e, 0x9a, 0xbf, 0xdb, 0x5f,
	0x52, 0x79, 0xd7, 0xc6, 0xa6, 0x40, 0x24, 0xaa, 0xde, 0x76, 0x76, 0xde, 0x79, 0x3d, 0x9e, 0x79,
	0x6c, 0xf8, 0xdf, 0xe1, 0x53, 0x81, 0xd1, 0x3d, 0x73, 0xd0, 0x0a, 0x23, 0x2e, 0x39, 0xa9, 0x79,
	0x18, 0xa8, 0x93, 0x79, 0xe0, 0x71, 0xee, 0xf9, 0xd8, 0x53, 0xd1, 0x2c, 0xfe, 0xd6, 0xc3, 0x79,
	0x28, 0x1f, 0xb5, 0xcc, 0xac, 0x39, 0x5c, 0x9f, 0xe8, 0x21, 0xc0, 0x18, 0xe5, 0x04, 0x7f, 0xc4,
	0x28, 0x24, 0x69, 0x41, 0x85, 0xb9, 0xdd, 0xf2, 0x71, 0xf9, 0xa4, 0x3e, 0xa9, 0x30, 0x97, 0x1e,
	0x41, 0xf3, 0x02, 0x7d, 0x94, 0xb8, 0x49, 0xf0, 0x16, 0x9a, 0x63, 0x94, 0xe7, 0xbe, 0x9f, 0x09,
	0x08, 0x54, 0x43, 0xdb, 0x43, 0x25, 0xa9, 0x4e, 0xd4, 0x99, 0xec, 0xc1, 0xae, 0xcf, 0xe6, 0x4c,
	0x76, 0x2b, 0xea, 0x52, 0x07, 0xf4, 0x16, 0x5a, 0x59, 0xa9, 0x08, 0x79, 0x20, 0x90, 0xbc, 0x06,
	0x83, 0x3f, 0x04, 0x18, 0x89, 0x6e, 0xf9, 0x78, 0xe7, 0xa4, 0xd1, 0xdf, 0xb3, 0xb2, 0xb7, 0xb1,
	0x86, 0x76, 0xe4, 0xf1, 0xeb, 0x24, 0x39, 0x49, 0x35, 0x89, 0xab, 0xc3, 0xe3, 0x60, 0xe1, 0xaa,
	0x02, 0x7a, 0x03, 0xfb, 0xe7, 0xae, 0x7b, 0xc1, 0x84, 0x8c, 0xd8, 0x2c, 0x96, 0x3c, 0xca, 0x1a,
	0x6b, 0x27, 0xf2, 0xe9, 0xa2, 0xf9, 0xaa, 0xc3, 0x2f, 0x5d, 0xf2, 0x0a, 0x5a, 0x6e, 0x2e, 0x4d,
	0xb2, 0x15, 0x95, 0x6d, 0x16, 0x6e, 0x2f, 0x5d, 0xfa, 0x05, 0xba, 0x7a, 0x0c, 0xff, 0xd8, 0x77,
	0x00, 0x64, 0xf8, 0x1d, 0x9d, 0xbb, 0xd1, 0x4f, 0x26, 0xa4, 0xc8, 0x1c, 0x3b, 0x60, 0x38, 0xdc,
	0x8f, 0xe7, 0x41, 0x6a, 0x99, 0x46, 0xc9, 0x0b, 0xdf, 0xdb, 0x7e, 0x8c, 0xa9, 0x97, 0x0e, 0x68,
	0x1f, 0xda, 0x4b, 0x1e, 0xe9, 0x2c, 0x0f, 0xa0, 0xce, 0xc4, 0x14, 0xd5, 0xa5, 0xf2, 0xa9, 0x4d,
	0x6a, 0x4c, 0x68, 0x11, 0x7d, 0x09, 0xff, 0x5d, 0x71, 0x8f, 0x05, 0xd9, 0x13, 0x93, 0x05, 0x25,
	0x71, 0xfa, 0x40, 0x1d, 0xf4, 0x7f, 0xed, 0x42, 0x7d, 0x78, 0x7d, 0xa3, 0xf9, 0x22, 0x6f, 0xc0,
	0x18, 0x46, 0x68, 0x4b, 0x24, 0x6b, 0xd7, 0x62, 0x76, 0x2c, 0x0d, 0x9c, 0x95, 0x01, 0x67, 0x8d,
	0x12, 0xe0, 0x68, 0x29, 0xa9, 0xfc, 0x1c, 0xba, 0x7f, 0x53, 0x79, 0x0a, 0x3b, 0x63, 0x94, 0xc5,
	0xb2, 0x9c, 0x55, 0x73, 0xad, 0x19, 0x2d, 0x91, 0x33, 0x30, 0x34, 0x57, 0xe4, 0xf9, 0x52, 0x5d,
	0x0e, 0xa9, 0xd9, 0x5d, 0x4d, 0xe8, 0xb1, 0xd1, 0x12, 0x79, 0x07, 0x86, 0xde, 0x75, 0xb1, 0x7c,
	0xe9, 0x23, 0x78, 0xa2, 0xe1, 0x0f, 0xd0, 0x5a, 0xa6, 0x8f, 0x1c, 0xe5, 0x26, 0x6b, 0xb9, 0x7c,
	0xc2, 0xec, 0x13, 0x3c, 0x5b, 0xa1, 0x8e, 0xd0, 0x3f, 0x9b, 0xda, 0xca, 0xf2, 0x0a, 0x1a, 0x05,
	0x58, 0xc8, 0x61, 0x61, 0x84, 0x2b, 0x1c, 0x9a, 0x2f, 0x36, 0x64, 0x17, 0xa3, 0x3a, 0x83, 0x86,
	0x5e, 0xec, 0x2d, 0xbf, 0xc3, 0x60, 0xeb, 0xed, 0x8e, 0x12, 0x72, 0xed, 0xc0, 0x43, 0xc5, 0xe2,
	0x47, 0x5b, 0x88, 0x07, 0x1e, 0xb9, 0x5b, 0xdb, 0xbc, 0x57, 0x7f, 0xb0, 0xc1, 0xa3, 0x72, 0x21,
	0x9d, 0xbc, 0xba, 0x88, 0xf8, 0x26, 0x5a, 0x06, 0xfb, 0x5f, 0xdb, 0x59, 0xa2, 0x97, 0xff, 0x4d,
	0x67, 0x86, 0xba, 0x39, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x4f, 0x95, 0xfe, 0x62, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// COServiceClient is the client API for COService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type COServiceClient interface {
	Create(ctx context.Context, in *CargoOwner, opts ...grpc.CallOption) (*empty.Empty, error)
	Update(ctx context.Context, in *CargoOwner, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*CargoOwner, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	AddDistributor(ctx context.Context, in *AddDistributorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteDistributor(ctx context.Context, in *DeleteDistributorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckExists(ctx context.Context, in *CheckExistsRequest, opts ...grpc.CallOption) (*CheckExistsResponse, error)
	UpdateToken(ctx context.Context, in *CargoOwner, opts ...grpc.CallOption) (*empty.Empty, error)
	ChangeLoginPassword(ctx context.Context, in *CargoOwner, opts ...grpc.CallOption) (*empty.Empty, error)
	GetByLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CargoOwner, error)
}

type cOServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCOServiceClient(cc grpc.ClientConnInterface) COServiceClient {
	return &cOServiceClient{cc}
}

func (c *cOServiceClient) Create(ctx context.Context, in *CargoOwner, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.COService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) Update(ctx context.Context, in *CargoOwner, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.COService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*CargoOwner, error) {
	out := new(CargoOwner)
	err := c.cc.Invoke(ctx, "/genproto.COService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/genproto.COService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.COService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) AddDistributor(ctx context.Context, in *AddDistributorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.COService/AddDistributor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) DeleteDistributor(ctx context.Context, in *DeleteDistributorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.COService/DeleteDistributor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) CheckExists(ctx context.Context, in *CheckExistsRequest, opts ...grpc.CallOption) (*CheckExistsResponse, error) {
	out := new(CheckExistsResponse)
	err := c.cc.Invoke(ctx, "/genproto.COService/CheckExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) UpdateToken(ctx context.Context, in *CargoOwner, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.COService/UpdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) ChangeLoginPassword(ctx context.Context, in *CargoOwner, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/genproto.COService/ChangeLoginPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cOServiceClient) GetByLogin(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*CargoOwner, error) {
	out := new(CargoOwner)
	err := c.cc.Invoke(ctx, "/genproto.COService/GetByLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// COServiceServer is the server API for COService service.
type COServiceServer interface {
	Create(context.Context, *CargoOwner) (*empty.Empty, error)
	Update(context.Context, *CargoOwner) (*empty.Empty, error)
	Get(context.Context, *GetRequest) (*CargoOwner, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
	AddDistributor(context.Context, *AddDistributorRequest) (*empty.Empty, error)
	DeleteDistributor(context.Context, *DeleteDistributorRequest) (*empty.Empty, error)
	CheckExists(context.Context, *CheckExistsRequest) (*CheckExistsResponse, error)
	UpdateToken(context.Context, *CargoOwner) (*empty.Empty, error)
	ChangeLoginPassword(context.Context, *CargoOwner) (*empty.Empty, error)
	GetByLogin(context.Context, *LoginRequest) (*CargoOwner, error)
}

// UnimplementedCOServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCOServiceServer struct {
}

func (*UnimplementedCOServiceServer) Create(ctx context.Context, req *CargoOwner) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedCOServiceServer) Update(ctx context.Context, req *CargoOwner) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCOServiceServer) Get(ctx context.Context, req *GetRequest) (*CargoOwner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCOServiceServer) GetAll(ctx context.Context, req *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedCOServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedCOServiceServer) AddDistributor(ctx context.Context, req *AddDistributorRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDistributor not implemented")
}
func (*UnimplementedCOServiceServer) DeleteDistributor(ctx context.Context, req *DeleteDistributorRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDistributor not implemented")
}
func (*UnimplementedCOServiceServer) CheckExists(ctx context.Context, req *CheckExistsRequest) (*CheckExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExists not implemented")
}
func (*UnimplementedCOServiceServer) UpdateToken(ctx context.Context, req *CargoOwner) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (*UnimplementedCOServiceServer) ChangeLoginPassword(ctx context.Context, req *CargoOwner) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeLoginPassword not implemented")
}
func (*UnimplementedCOServiceServer) GetByLogin(ctx context.Context, req *LoginRequest) (*CargoOwner, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByLogin not implemented")
}

func RegisterCOServiceServer(s *grpc.Server, srv COServiceServer) {
	s.RegisterService(&_COService_serviceDesc, srv)
}

func _COService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CargoOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).Create(ctx, req.(*CargoOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CargoOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).Update(ctx, req.(*CargoOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_AddDistributor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDistributorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).AddDistributor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/AddDistributor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).AddDistributor(ctx, req.(*AddDistributorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_DeleteDistributor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDistributorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).DeleteDistributor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/DeleteDistributor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).DeleteDistributor(ctx, req.(*DeleteDistributorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_CheckExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).CheckExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/CheckExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).CheckExists(ctx, req.(*CheckExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CargoOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).UpdateToken(ctx, req.(*CargoOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_ChangeLoginPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CargoOwner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).ChangeLoginPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/ChangeLoginPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).ChangeLoginPassword(ctx, req.(*CargoOwner))
	}
	return interceptor(ctx, in, info, handler)
}

func _COService_GetByLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(COServiceServer).GetByLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.COService/GetByLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(COServiceServer).GetByLogin(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _COService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.COService",
	HandlerType: (*COServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _COService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _COService_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _COService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _COService_GetAll_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _COService_Delete_Handler,
		},
		{
			MethodName: "AddDistributor",
			Handler:    _COService_AddDistributor_Handler,
		},
		{
			MethodName: "DeleteDistributor",
			Handler:    _COService_DeleteDistributor_Handler,
		},
		{
			MethodName: "CheckExists",
			Handler:    _COService_CheckExists_Handler,
		},
		{
			MethodName: "UpdateToken",
			Handler:    _COService_UpdateToken_Handler,
		},
		{
			MethodName: "ChangeLoginPassword",
			Handler:    _COService_ChangeLoginPassword_Handler,
		},
		{
			MethodName: "GetByLogin",
			Handler:    _COService_GetByLogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "co_service.proto",
}