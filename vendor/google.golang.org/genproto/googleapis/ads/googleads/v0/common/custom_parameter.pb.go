// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/custom_parameter.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v0/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A mapping that can be used by custom parameter tags in a
// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
type CustomParameter struct {
	// The key matching the parameter tag name.
	Key *wrappers.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value to be substituted.
	Value                *wrappers.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomParameter) Reset()         { *m = CustomParameter{} }
func (m *CustomParameter) String() string { return proto.CompactTextString(m) }
func (*CustomParameter) ProtoMessage()    {}
func (*CustomParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_custom_parameter_67efbc4679d7ce99, []int{0}
}
func (m *CustomParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomParameter.Unmarshal(m, b)
}
func (m *CustomParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomParameter.Marshal(b, m, deterministic)
}
func (dst *CustomParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomParameter.Merge(dst, src)
}
func (m *CustomParameter) XXX_Size() int {
	return xxx_messageInfo_CustomParameter.Size(m)
}
func (m *CustomParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomParameter.DiscardUnknown(m)
}

var xxx_messageInfo_CustomParameter proto.InternalMessageInfo

func (m *CustomParameter) GetKey() *wrappers.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CustomParameter) GetValue() *wrappers.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomParameter)(nil), "google.ads.googleads.v0.common.CustomParameter")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/custom_parameter.proto", fileDescriptor_custom_parameter_67efbc4679d7ce99)
}

var fileDescriptor_custom_parameter_67efbc4679d7ce99 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xb1, 0x4a, 0xec, 0x40,
	0x14, 0x86, 0x49, 0x96, 0x7b, 0x8b, 0x58, 0x08, 0xc1, 0x62, 0x11, 0x59, 0x24, 0x95, 0xd5, 0x99,
	0xb0, 0x62, 0x33, 0x56, 0xd9, 0x15, 0xb6, 0x0d, 0x0a, 0x29, 0x24, 0x20, 0xb3, 0xc9, 0x38, 0x2c,
	0x26, 0x39, 0x61, 0x26, 0x89, 0xf8, 0x3a, 0x96, 0x3e, 0x8a, 0x8f, 0x62, 0xe3, 0x2b, 0xc8, 0xcc,
	0x49, 0x52, 0x08, 0x8a, 0x55, 0xfe, 0xcc, 0x7c, 0xff, 0x37, 0x87, 0x13, 0x5c, 0x29, 0x44, 0x55,
	0x49, 0x26, 0x4a, 0xc3, 0x28, 0xda, 0x34, 0xc4, 0xac, 0xc0, 0xba, 0xc6, 0x86, 0x15, 0xbd, 0xe9,
	0xb0, 0x7e, 0x68, 0x85, 0x16, 0xb5, 0xec, 0xa4, 0x86, 0x56, 0x63, 0x87, 0xe1, 0x8a, 0x58, 0x10,
	0xa5, 0x81, 0xb9, 0x06, 0x43, 0x0c, 0x54, 0x3b, 0x1d, 0xef, 0x99, 0xa3, 0xf7, 0xfd, 0x23, 0x7b,
	0xd6, 0xa2, 0x6d, 0xa5, 0x36, 0xd4, 0x8f, 0xfa, 0xe0, 0x78, 0xeb, 0xcc, 0xe9, 0x24, 0x0e, 0x21,
	0x58, 0x3c, 0xc9, 0x97, 0xa5, 0x77, 0xee, 0x5d, 0x1c, 0xad, 0xcf, 0x46, 0x2b, 0x4c, 0x02, 0xb8,
	0xeb, 0xf4, 0xa1, 0x51, 0x99, 0xa8, 0x7a, 0x79, 0x6b, 0xc1, 0x70, 0x1d, 0xfc, 0x1b, 0xec, 0xdf,
	0xd2, 0xff, 0x43, 0x83, 0xd0, 0xcd, 0xa7, 0x17, 0x44, 0x05, 0xd6, 0xf0, 0xfb, 0xf4, 0x9b, 0x93,
	0x6f, 0xb3, 0xa5, 0x56, 0x99, 0x7a, 0xf7, 0x37, 0x63, 0x4f, 0x61, 0x25, 0x1a, 0x05, 0xa8, 0x15,
	0x53, 0xb2, 0x71, 0x0f, 0x4e, 0xcb, 0x6b, 0x0f, 0xe6, 0xa7, 0x5d, 0x5e, 0xd3, 0xe7, 0xd5, 0x5f,
	0xec, 0x92, 0xe4, 0xcd, 0x5f, 0xed, 0x48, 0x96, 0x94, 0x06, 0x28, 0xda, 0x94, 0xc5, 0xb0, 0x75,
	0xd8, 0xfb, 0x04, 0xe4, 0x49, 0x69, 0xf2, 0x19, 0xc8, 0xb3, 0x38, 0x27, 0xe0, 0xc3, 0x8f, 0xe8,
	0x94, 0xf3, 0xa4, 0x34, 0x9c, 0xcf, 0x08, 0xe7, 0x59, 0xcc, 0x39, 0x41, 0xfb, 0xff, 0x6e, 0xba,
	0xcb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x98, 0x2d, 0x13, 0xe8, 0x01, 0x00, 0x00,
}