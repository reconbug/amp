// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/appcelerator/amp/data/object_stores/object_stores.proto

/*
Package object_stores is a generated protocol buffer package.

It is generated from these files:
	github.com/appcelerator/amp/data/object_stores/object_stores.proto

It has these top-level messages:
	ObjectStore
*/
package object_stores

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import accounts "github.com/appcelerator/amp/data/accounts"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ObjectStore struct {
	Id       string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Owner    *accounts.Account `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
	CreateDt int64             `protobuf:"varint,4,opt,name=create_dt,json=createDt" json:"create_dt,omitempty"`
	Location string            `protobuf:"bytes,5,opt,name=location" json:"location,omitempty"`
}

func (m *ObjectStore) Reset()                    { *m = ObjectStore{} }
func (m *ObjectStore) String() string            { return proto.CompactTextString(m) }
func (*ObjectStore) ProtoMessage()               {}
func (*ObjectStore) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ObjectStore) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectStore) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectStore) GetOwner() *accounts.Account {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ObjectStore) GetCreateDt() int64 {
	if m != nil {
		return m.CreateDt
	}
	return 0
}

func (m *ObjectStore) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func init() {
	proto.RegisterType((*ObjectStore)(nil), "object_stores.ObjectStore")
}

func init() {
	proto.RegisterFile("github.com/appcelerator/amp/data/object_stores/object_stores.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0xce, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x06, 0x60, 0xd2, 0xdd, 0x95, 0xdd, 0x2c, 0x0a, 0xe6, 0x14, 0xd6, 0x4b, 0xf1, 0x62, 0x4f,
	0x0d, 0xe8, 0xc5, 0xab, 0xe2, 0x5d, 0xa8, 0x0f, 0xb0, 0x4c, 0x93, 0x41, 0x23, 0x6d, 0x26, 0xa4,
	0x53, 0x7c, 0x10, 0x5f, 0x58, 0x4c, 0xa0, 0xd0, 0xd3, 0xde, 0xfe, 0xff, 0x1f, 0xf8, 0x18, 0xf9,
	0xfa, 0xe9, 0xf9, 0x6b, 0xee, 0x5b, 0x4b, 0xa3, 0x81, 0x18, 0x2d, 0x0e, 0x98, 0x80, 0x29, 0x19,
	0x18, 0xa3, 0x71, 0xc0, 0x60, 0xa8, 0xff, 0x46, 0xcb, 0xe7, 0x89, 0x29, 0xe1, 0xb4, 0x6e, 0x6d,
	0x4c, 0xc4, 0xa4, 0xae, 0x57, 0xe3, 0xe9, 0xf9, 0x22, 0x09, 0xd6, 0xd2, 0x1c, 0x78, 0x5a, 0x42,
	0x81, 0xee, 0x7f, 0x85, 0x3c, 0xbe, 0x67, 0xeb, 0xe3, 0x9f, 0x52, 0x37, 0xb2, 0xf2, 0x4e, 0x8b,
	0x5a, 0x34, 0x87, 0xae, 0xf2, 0x4e, 0x29, 0xb9, 0x0d, 0x30, 0xa2, 0xae, 0xf2, 0x92, 0xb3, 0x7a,
	0x90, 0x3b, 0xfa, 0x09, 0x98, 0xf4, 0xa6, 0x16, 0xcd, 0xf1, 0xf1, 0xb6, 0x5d, 0xcc, 0x97, 0x12,
	0xba, 0x72, 0x57, 0x77, 0xf2, 0x60, 0x13, 0x02, 0xe3, 0xd9, 0xb1, 0xde, 0xd6, 0xa2, 0xd9, 0x74,
	0xfb, 0x32, 0xbc, 0xb1, 0x3a, 0xc9, 0xfd, 0x40, 0x16, 0xd8, 0x53, 0xd0, 0xbb, 0xac, 0x2f, 0xbd,
	0xbf, 0xca, 0xcf, 0x3d, 0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xaa, 0x8d, 0x5d, 0x2b, 0x01,
	0x00, 0x00,
}
