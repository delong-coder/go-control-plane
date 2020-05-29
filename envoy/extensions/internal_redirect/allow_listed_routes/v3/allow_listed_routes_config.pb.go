// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/internal_redirect/allow_listed_routes/v3/allow_listed_routes_config.proto

package envoy_extensions_internal_redirect_allow_listed_routes_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type AllowListedRoutesConfig struct {
	AllowedRouteNames    []string `protobuf:"bytes,1,rep,name=allowed_route_names,json=allowedRouteNames,proto3" json:"allowed_route_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllowListedRoutesConfig) Reset()         { *m = AllowListedRoutesConfig{} }
func (m *AllowListedRoutesConfig) String() string { return proto.CompactTextString(m) }
func (*AllowListedRoutesConfig) ProtoMessage()    {}
func (*AllowListedRoutesConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cf9599f79277a9d, []int{0}
}

func (m *AllowListedRoutesConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllowListedRoutesConfig.Unmarshal(m, b)
}
func (m *AllowListedRoutesConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllowListedRoutesConfig.Marshal(b, m, deterministic)
}
func (m *AllowListedRoutesConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowListedRoutesConfig.Merge(m, src)
}
func (m *AllowListedRoutesConfig) XXX_Size() int {
	return xxx_messageInfo_AllowListedRoutesConfig.Size(m)
}
func (m *AllowListedRoutesConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowListedRoutesConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AllowListedRoutesConfig proto.InternalMessageInfo

func (m *AllowListedRoutesConfig) GetAllowedRouteNames() []string {
	if m != nil {
		return m.AllowedRouteNames
	}
	return nil
}

func init() {
	proto.RegisterType((*AllowListedRoutesConfig)(nil), "envoy.extensions.internal_redirect.allow_listed_routes.v3.AllowListedRoutesConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/internal_redirect/allow_listed_routes/v3/allow_listed_routes_config.proto", fileDescriptor_0cf9599f79277a9d)
}

var fileDescriptor_0cf9599f79277a9d = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x8a, 0x4a, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0xcf, 0xcc, 0x2b,
	0x49, 0x2d, 0xca, 0x4b, 0xcc, 0x89, 0x2f, 0x4a, 0x4d, 0xc9, 0x2c, 0x4a, 0x4d, 0x2e, 0xd1, 0x4f,
	0xcc, 0xc9, 0xc9, 0x2f, 0x8f, 0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0x4d, 0x89, 0x2f, 0xca, 0x2f, 0x2d,
	0x49, 0x2d, 0xd6, 0x2f, 0x33, 0xc6, 0x26, 0x1c, 0x9f, 0x9c, 0x9f, 0x97, 0x96, 0x99, 0xae, 0x57,
	0x50, 0x94, 0x5f, 0x92, 0x2f, 0x64, 0x09, 0x36, 0x5b, 0x0f, 0x61, 0xb6, 0x1e, 0x86, 0xd9, 0x7a,
	0x58, 0x0c, 0xd1, 0x2b, 0x33, 0x96, 0x92, 0x2d, 0x4d, 0x29, 0x48, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb,
	0x2f, 0x49, 0x2c, 0x01, 0x3b, 0xab, 0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x18, 0x62, 0xb2, 0x94, 0x22,
	0x86, 0x74, 0x59, 0x6a, 0x11, 0xc8, 0x8a, 0xcc, 0x3c, 0xa8, 0xe5, 0x52, 0xe2, 0x65, 0x89, 0x39,
	0x99, 0x29, 0x89, 0x25, 0xa9, 0xfa, 0x30, 0x06, 0x44, 0x42, 0x29, 0x9c, 0x4b, 0xdc, 0x11, 0x64,
	0xa9, 0x0f, 0xd8, 0xce, 0x20, 0xb0, 0x95, 0xce, 0x60, 0x67, 0x0b, 0xd9, 0x70, 0x09, 0x83, 0xdd,
	0x03, 0x73, 0x4a, 0x7c, 0x5e, 0x62, 0x6e, 0x6a, 0xb1, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xa7, 0x13,
	0xcf, 0x2f, 0x27, 0xce, 0x49, 0x8c, 0x6c, 0x4a, 0x2c, 0x45, 0x4c, 0x02, 0x8c, 0x41, 0x82, 0x50,
	0x85, 0x60, 0xfd, 0x7e, 0x20, 0x65, 0x4e, 0x85, 0xbb, 0x1a, 0x4e, 0x5c, 0x64, 0x63, 0x12, 0x60,
	0xe2, 0x72, 0xcf, 0xcc, 0xd7, 0x03, 0xfb, 0xbd, 0xa0, 0x28, 0xbf, 0xa2, 0x52, 0x8f, 0xec, 0x60,
	0x70, 0x92, 0xc1, 0xe1, 0xd2, 0x00, 0x90, 0x4f, 0x02, 0x18, 0x93, 0xd8, 0xc0, 0x5e, 0x32, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x48, 0xeb, 0x48, 0xc6, 0x01, 0x00, 0x00,
}