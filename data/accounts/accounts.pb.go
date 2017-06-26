// Code generated by protoc-gen-go.
// source: github.com/appcelerator/amp/data/accounts/accounts.proto
// DO NOT EDIT!

/*
Package accounts is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/data/accounts/accounts.proto

It has these top-level messages:
	User
	TeamResource
	Team
	OrganizationMember
	Organization
	Account
*/
package accounts

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

type TeamPermissionLevel int32

const (
	TeamPermissionLevel_TEAM_READ  TeamPermissionLevel = 0
	TeamPermissionLevel_TEAM_WRITE TeamPermissionLevel = 1
	TeamPermissionLevel_TEAM_ADMIN TeamPermissionLevel = 2
)

var TeamPermissionLevel_name = map[int32]string{
	0: "TEAM_READ",
	1: "TEAM_WRITE",
	2: "TEAM_ADMIN",
}
var TeamPermissionLevel_value = map[string]int32{
	"TEAM_READ":  0,
	"TEAM_WRITE": 1,
	"TEAM_ADMIN": 2,
}

func (x TeamPermissionLevel) String() string {
	return proto.EnumName(TeamPermissionLevel_name, int32(x))
}
func (TeamPermissionLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type OrganizationRole int32

const (
	OrganizationRole_ORGANIZATION_MEMBER OrganizationRole = 0
	OrganizationRole_ORGANIZATION_OWNER  OrganizationRole = 1
)

var OrganizationRole_name = map[int32]string{
	0: "ORGANIZATION_MEMBER",
	1: "ORGANIZATION_OWNER",
}
var OrganizationRole_value = map[string]int32{
	"ORGANIZATION_MEMBER": 0,
	"ORGANIZATION_OWNER":  1,
}

func (x OrganizationRole) String() string {
	return proto.EnumName(OrganizationRole_name, int32(x))
}
func (OrganizationRole) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type User struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email        string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	PasswordHash string `protobuf:"bytes,3,opt,name=password_hash,json=passwordHash" json:"password_hash,omitempty"`
	IsVerified   bool   `protobuf:"varint,4,opt,name=is_verified,json=isVerified" json:"is_verified,omitempty"`
	CreateDt     int64  `protobuf:"varint,5,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
	TokenUsed    bool   `protobuf:"varint,6,opt,name=token_used,json=tokenUsed" json:"token_used,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPasswordHash() string {
	if m != nil {
		return m.PasswordHash
	}
	return ""
}

func (m *User) GetIsVerified() bool {
	if m != nil {
		return m.IsVerified
	}
	return false
}

func (m *User) GetCreateDt() int64 {
	if m != nil {
		return m.CreateDt
	}
	return 0
}

func (m *User) GetTokenUsed() bool {
	if m != nil {
		return m.TokenUsed
	}
	return false
}

type TeamResource struct {
	Id              string              `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PermissionLevel TeamPermissionLevel `protobuf:"varint,2,opt,name=permission_level,json=permissionLevel,enum=accounts.TeamPermissionLevel" json:"permission_level,omitempty"`
}

func (m *TeamResource) Reset()                    { *m = TeamResource{} }
func (m *TeamResource) String() string            { return proto.CompactTextString(m) }
func (*TeamResource) ProtoMessage()               {}
func (*TeamResource) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TeamResource) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TeamResource) GetPermissionLevel() TeamPermissionLevel {
	if m != nil {
		return m.PermissionLevel
	}
	return TeamPermissionLevel_TEAM_READ
}

type Team struct {
	Name      string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CreateDt  int64           `protobuf:"varint,2,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
	Owner     *Account        `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	Members   []string        `protobuf:"bytes,4,rep,name=members" json:"members,omitempty"`
	Resources []*TeamResource `protobuf:"bytes,5,rep,name=resources" json:"resources,omitempty"`
}

func (m *Team) Reset()                    { *m = Team{} }
func (m *Team) String() string            { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()               {}
func (*Team) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetCreateDt() int64 {
	if m != nil {
		return m.CreateDt
	}
	return 0
}

func (m *Team) GetOwner() *Account {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Team) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Team) GetResources() []*TeamResource {
	if m != nil {
		return m.Resources
	}
	return nil
}

type OrganizationMember struct {
	Name string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Role OrganizationRole `protobuf:"varint,2,opt,name=role,enum=accounts.OrganizationRole" json:"role,omitempty"`
}

func (m *OrganizationMember) Reset()                    { *m = OrganizationMember{} }
func (m *OrganizationMember) String() string            { return proto.CompactTextString(m) }
func (*OrganizationMember) ProtoMessage()               {}
func (*OrganizationMember) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *OrganizationMember) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrganizationMember) GetRole() OrganizationRole {
	if m != nil {
		return m.Role
	}
	return OrganizationRole_ORGANIZATION_MEMBER
}

type Organization struct {
	Name     string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Email    string                `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	CreateDt int64                 `protobuf:"varint,3,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
	Members  []*OrganizationMember `protobuf:"bytes,4,rep,name=members" json:"members,omitempty"`
	Teams    []*Team               `protobuf:"bytes,5,rep,name=teams" json:"teams,omitempty"`
}

func (m *Organization) Reset()                    { *m = Organization{} }
func (m *Organization) String() string            { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()               {}
func (*Organization) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Organization) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Organization) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Organization) GetCreateDt() int64 {
	if m != nil {
		return m.CreateDt
	}
	return 0
}

func (m *Organization) GetMembers() []*OrganizationMember {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Organization) GetTeams() []*Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type Account struct {
	User         string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Organization string `protobuf:"bytes,2,opt,name=organization" json:"organization,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Account) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Account) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "accounts.User")
	proto.RegisterType((*TeamResource)(nil), "accounts.TeamResource")
	proto.RegisterType((*Team)(nil), "accounts.Team")
	proto.RegisterType((*OrganizationMember)(nil), "accounts.OrganizationMember")
	proto.RegisterType((*Organization)(nil), "accounts.Organization")
	proto.RegisterType((*Account)(nil), "accounts.Account")
	proto.RegisterEnum("accounts.TeamPermissionLevel", TeamPermissionLevel_name, TeamPermissionLevel_value)
	proto.RegisterEnum("accounts.OrganizationRole", OrganizationRole_name, OrganizationRole_value)
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/data/accounts/accounts.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xeb, 0xc4, 0x69, 0xeb, 0x49, 0x9a, 0xd7, 0xef, 0x16, 0x15, 0x0b, 0xa8, 0x88, 0x0c,
	0x12, 0x51, 0x0f, 0x89, 0x14, 0x10, 0xe2, 0x6a, 0x88, 0x45, 0x23, 0x91, 0x04, 0xad, 0x52, 0x8a,
	0xb8, 0x58, 0x1b, 0x7b, 0x68, 0x56, 0xc4, 0x5e, 0x6b, 0x77, 0xd3, 0x4a, 0x7c, 0x2a, 0x04, 0x5f,
	0x10, 0xc5, 0x7f, 0x92, 0x38, 0xca, 0x85, 0xdb, 0xcc, 0x33, 0xb3, 0xb3, 0xcf, 0xfc, 0xb4, 0x0b,
	0xef, 0xee, 0xb8, 0x5e, 0xac, 0xe6, 0xbd, 0x50, 0xc4, 0x7d, 0x96, 0xa6, 0x21, 0x2e, 0x51, 0x32,
	0x2d, 0x64, 0x9f, 0xc5, 0x69, 0x3f, 0x62, 0x9a, 0xf5, 0x59, 0x18, 0x8a, 0x55, 0xa2, 0xd5, 0x26,
	0xe8, 0xa5, 0x52, 0x68, 0x41, 0x4e, 0xcb, 0xdc, 0xfd, 0x6d, 0x80, 0x79, 0xa3, 0x50, 0x12, 0x02,
	0x66, 0xc2, 0x62, 0x74, 0x8c, 0x8e, 0xd1, 0xb5, 0x68, 0x16, 0x93, 0x47, 0xd0, 0xc0, 0x98, 0xf1,
	0xa5, 0x53, 0xcb, 0xc4, 0x3c, 0x21, 0x2f, 0xe0, 0x2c, 0x65, 0x4a, 0x3d, 0x08, 0x19, 0x05, 0x0b,
	0xa6, 0x16, 0x4e, 0x3d, 0xab, 0xb6, 0x4a, 0xf1, 0x9a, 0xa9, 0x05, 0x79, 0x0e, 0x4d, 0xae, 0x82,
	0x7b, 0x94, 0xfc, 0x3b, 0xc7, 0xc8, 0x31, 0x3b, 0x46, 0xf7, 0x94, 0x02, 0x57, 0x5f, 0x0a, 0x85,
	0x3c, 0x05, 0x2b, 0x94, 0xc8, 0x34, 0x06, 0x91, 0x76, 0x1a, 0x1d, 0xa3, 0x5b, 0xa7, 0xa7, 0xb9,
	0x30, 0xd4, 0xe4, 0x12, 0x40, 0x8b, 0x1f, 0x98, 0x04, 0x2b, 0x85, 0x91, 0x73, 0x9c, 0x1d, 0xb6,
	0x32, 0xe5, 0x46, 0x61, 0xe4, 0x2e, 0xa0, 0x35, 0x43, 0x16, 0x53, 0x54, 0x62, 0x25, 0x43, 0x24,
	0x6d, 0xa8, 0xf1, 0xa8, 0x70, 0x5e, 0xe3, 0x11, 0xb9, 0x06, 0x3b, 0x45, 0x19, 0x73, 0xa5, 0xb8,
	0x48, 0x82, 0x25, 0xde, 0x63, 0xbe, 0x42, 0x7b, 0x70, 0xd9, 0xdb, 0x90, 0x58, 0x4f, 0xf8, 0xbc,
	0xe9, 0xfa, 0xb4, 0x6e, 0xa2, 0xff, 0xa5, 0x55, 0xc1, 0xfd, 0x65, 0x80, 0xb9, 0x6e, 0x3c, 0x88,
	0xa7, 0xb2, 0x42, 0x6d, 0x6f, 0x85, 0x57, 0xd0, 0x10, 0x0f, 0x09, 0xca, 0x8c, 0x4e, 0x73, 0xf0,
	0xff, 0xf6, 0x62, 0x2f, 0x0f, 0x68, 0x5e, 0x27, 0x0e, 0x9c, 0xc4, 0x18, 0xcf, 0x51, 0x2a, 0xc7,
	0xec, 0xd4, 0xbb, 0x16, 0x2d, 0x53, 0xf2, 0x06, 0x2c, 0x59, 0xac, 0xa8, 0x9c, 0x46, 0xa7, 0xde,
	0x6d, 0x0e, 0x2e, 0xaa, 0xfe, 0x4b, 0x02, 0x74, 0xdb, 0xe8, 0x7e, 0x05, 0x32, 0x95, 0x77, 0x2c,
	0xe1, 0x3f, 0x99, 0xe6, 0x22, 0x19, 0x67, 0xc3, 0x0e, 0xfa, 0xef, 0x81, 0x29, 0xc5, 0x12, 0x0b,
	0x34, 0x4f, 0xb6, 0xa3, 0x77, 0xcf, 0x53, 0xb1, 0x44, 0x9a, 0xf5, 0xb9, 0x7f, 0x0c, 0x68, 0xed,
	0x96, 0xfe, 0xe1, 0xcd, 0x54, 0x50, 0xd5, 0xf7, 0x50, 0xbd, 0xad, 0x12, 0x68, 0x0e, 0x9e, 0x1d,
	0xb6, 0x92, 0xaf, 0xb2, 0xe5, 0xf3, 0x12, 0x1a, 0x1a, 0x59, 0x5c, 0xb2, 0x69, 0xef, 0xb1, 0xc9,
	0x8b, 0xae, 0x07, 0x27, 0x05, 0xf1, 0xb5, 0xdf, 0x95, 0x42, 0x59, 0xfa, 0x5d, 0xc7, 0xc4, 0x85,
	0x96, 0xd8, 0xb9, 0xa3, 0xb0, 0x5d, 0xd1, 0xae, 0x86, 0x70, 0x7e, 0xe0, 0xb5, 0x90, 0x33, 0xb0,
	0x66, 0xbe, 0x37, 0x0e, 0xa8, 0xef, 0x0d, 0xed, 0x23, 0xd2, 0x06, 0xc8, 0xd2, 0x5b, 0x3a, 0x9a,
	0xf9, 0xb6, 0xb1, 0xc9, 0xbd, 0xe1, 0x78, 0x34, 0xb1, 0x6b, 0x57, 0x1f, 0xc0, 0xde, 0x07, 0x4b,
	0x1e, 0xc3, 0xf9, 0x94, 0x7e, 0xf4, 0x26, 0xa3, 0x6f, 0xde, 0x6c, 0x34, 0x9d, 0x04, 0x63, 0x7f,
	0xfc, 0xde, 0xa7, 0xf6, 0x11, 0xb9, 0x00, 0x52, 0x29, 0x4c, 0x6f, 0x27, 0x3e, 0xb5, 0x8d, 0xf9,
	0x71, 0xf6, 0x81, 0x5f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x6a, 0xce, 0xc1, 0x61, 0xfc, 0x03,
	0x00, 0x00,
}
