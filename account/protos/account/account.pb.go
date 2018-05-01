// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

/*
Package go_micro_srv_users is a generated protocol buffer package.

It is generated from these files:
	account.proto

It has these top-level messages:
	GetAccountRequest
	Account
*/
package go_micro_srv_users

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Account_Type int32

const (
	Account_ADMIN    Account_Type = 0
	Account_CUSTOMER Account_Type = 1
	Account_GUEST    Account_Type = 2
)

var Account_Type_name = map[int32]string{
	0: "ADMIN",
	1: "CUSTOMER",
	2: "GUEST",
}
var Account_Type_value = map[string]int32{
	"ADMIN":    0,
	"CUSTOMER": 1,
	"GUEST":    2,
}

func (x Account_Type) String() string {
	return proto.EnumName(Account_Type_name, int32(x))
}
func (Account_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type GetAccountRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetAccountRequest) Reset()                    { *m = GetAccountRequest{} }
func (m *GetAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAccountRequest) ProtoMessage()               {}
func (*GetAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetAccountRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Account struct {
	Guid     string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Account) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAccountRequest)(nil), "go.micro.srv.users.GetAccountRequest")
	proto.RegisterType((*Account)(nil), "go.micro.srv.users.Account")
	proto.RegisterEnum("go.micro.srv.users.Account_Type", Account_Type_name, Account_Type_value)
}

func init() { proto.RegisterFile("account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4c, 0x4e, 0xce,
	0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0xcf, 0xd7, 0xcb, 0xcd,
	0x4c, 0x2e, 0xca, 0xd7, 0x2b, 0x2e, 0x2a, 0xd3, 0x2b, 0x2d, 0x4e, 0x2d, 0x2a, 0x56, 0x52, 0xe6,
	0x12, 0x74, 0x4f, 0x2d, 0x71, 0x84, 0xa8, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2,
	0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x51, 0xea,
	0x67, 0xe4, 0x62, 0x87, 0x2a, 0x11, 0x12, 0xe2, 0x62, 0x49, 0x2f, 0x85, 0xcb, 0x82, 0xd9, 0x20,
	0xb1, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x88, 0x18, 0x88, 0x2d, 0x24, 0xc2, 0xc5, 0x9a, 0x9a,
	0x9b, 0x98, 0x99, 0x23, 0xc1, 0x0c, 0x16, 0x84, 0x70, 0x84, 0xa4, 0xb8, 0x38, 0x0a, 0x12, 0x8b,
	0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24, 0x58, 0xc0, 0x12, 0x70, 0xbe, 0x92, 0x16, 0x17, 0x4b, 0x48,
	0x65, 0x41, 0xaa, 0x10, 0x27, 0x17, 0xab, 0xa3, 0x8b, 0xaf, 0xa7, 0x9f, 0x00, 0x83, 0x10, 0x0f,
	0x17, 0x87, 0x73, 0x68, 0x70, 0x88, 0xbf, 0xaf, 0x6b, 0x90, 0x00, 0x23, 0x48, 0xc2, 0x3d, 0xd4,
	0x35, 0x38, 0x44, 0x80, 0xc9, 0x28, 0x8e, 0x8b, 0x03, 0xea, 0xa0, 0x62, 0xa1, 0x20, 0x2e, 0x2e,
	0x84, 0x17, 0x84, 0x54, 0xf5, 0x30, 0x7d, 0xa9, 0x87, 0xe1, 0x45, 0x29, 0x69, 0x6c, 0xca, 0xa0,
	0x6a, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x21, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x06,
	0x7b, 0x65, 0x42, 0x01, 0x00, 0x00,
}