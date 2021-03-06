// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v3alpha/thrift_proxy.proto

package envoy_config_filter_network_thrift_proxy_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type TransportType int32

const (
	TransportType_AUTO_TRANSPORT TransportType = 0
	TransportType_FRAMED         TransportType = 1
	TransportType_UNFRAMED       TransportType = 2
	TransportType_HEADER         TransportType = 3
)

var TransportType_name = map[int32]string{
	0: "AUTO_TRANSPORT",
	1: "FRAMED",
	2: "UNFRAMED",
	3: "HEADER",
}

var TransportType_value = map[string]int32{
	"AUTO_TRANSPORT": 0,
	"FRAMED":         1,
	"UNFRAMED":       2,
	"HEADER":         3,
}

func (x TransportType) String() string {
	return proto.EnumName(TransportType_name, int32(x))
}

func (TransportType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5bf22ec9d8f099b6, []int{0}
}

type ProtocolType int32

const (
	ProtocolType_AUTO_PROTOCOL ProtocolType = 0
	ProtocolType_BINARY        ProtocolType = 1
	ProtocolType_LAX_BINARY    ProtocolType = 2
	ProtocolType_COMPACT       ProtocolType = 3
	ProtocolType_TWITTER       ProtocolType = 4
)

var ProtocolType_name = map[int32]string{
	0: "AUTO_PROTOCOL",
	1: "BINARY",
	2: "LAX_BINARY",
	3: "COMPACT",
	4: "TWITTER",
}

var ProtocolType_value = map[string]int32{
	"AUTO_PROTOCOL": 0,
	"BINARY":        1,
	"LAX_BINARY":    2,
	"COMPACT":       3,
	"TWITTER":       4,
}

func (x ProtocolType) String() string {
	return proto.EnumName(ProtocolType_name, int32(x))
}

func (ProtocolType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5bf22ec9d8f099b6, []int{1}
}

type ThriftProxy struct {
	Transport            TransportType       `protobuf:"varint,2,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v3alpha.TransportType" json:"transport,omitempty"`
	Protocol             ProtocolType        `protobuf:"varint,3,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v3alpha.ProtocolType" json:"protocol,omitempty"`
	StatPrefix           string              `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	RouteConfig          *RouteConfiguration `protobuf:"bytes,4,opt,name=route_config,json=routeConfig,proto3" json:"route_config,omitempty"`
	ThriftFilters        []*ThriftFilter     `protobuf:"bytes,5,rep,name=thrift_filters,json=thriftFilters,proto3" json:"thrift_filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ThriftProxy) Reset()         { *m = ThriftProxy{} }
func (m *ThriftProxy) String() string { return proto.CompactTextString(m) }
func (*ThriftProxy) ProtoMessage()    {}
func (*ThriftProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf22ec9d8f099b6, []int{0}
}

func (m *ThriftProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProxy.Unmarshal(m, b)
}
func (m *ThriftProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProxy.Marshal(b, m, deterministic)
}
func (m *ThriftProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProxy.Merge(m, src)
}
func (m *ThriftProxy) XXX_Size() int {
	return xxx_messageInfo_ThriftProxy.Size(m)
}
func (m *ThriftProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProxy.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProxy proto.InternalMessageInfo

func (m *ThriftProxy) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProxy) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func (m *ThriftProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *ThriftProxy) GetRouteConfig() *RouteConfiguration {
	if m != nil {
		return m.RouteConfig
	}
	return nil
}

func (m *ThriftProxy) GetThriftFilters() []*ThriftFilter {
	if m != nil {
		return m.ThriftFilters
	}
	return nil
}

type ThriftFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ThriftFilter_HiddenEnvoyDeprecatedConfig
	//	*ThriftFilter_TypedConfig
	ConfigType           isThriftFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ThriftFilter) Reset()         { *m = ThriftFilter{} }
func (m *ThriftFilter) String() string { return proto.CompactTextString(m) }
func (*ThriftFilter) ProtoMessage()    {}
func (*ThriftFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf22ec9d8f099b6, []int{1}
}

func (m *ThriftFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftFilter.Unmarshal(m, b)
}
func (m *ThriftFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftFilter.Marshal(b, m, deterministic)
}
func (m *ThriftFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftFilter.Merge(m, src)
}
func (m *ThriftFilter) XXX_Size() int {
	return xxx_messageInfo_ThriftFilter.Size(m)
}
func (m *ThriftFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftFilter proto.InternalMessageInfo

func (m *ThriftFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isThriftFilter_ConfigType interface {
	isThriftFilter_ConfigType()
}

type ThriftFilter_HiddenEnvoyDeprecatedConfig struct {
	HiddenEnvoyDeprecatedConfig *_struct.Struct `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_config,json=hiddenEnvoyDeprecatedConfig,proto3,oneof"`
}

type ThriftFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ThriftFilter_HiddenEnvoyDeprecatedConfig) isThriftFilter_ConfigType() {}

func (*ThriftFilter_TypedConfig) isThriftFilter_ConfigType() {}

func (m *ThriftFilter) GetConfigType() isThriftFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *ThriftFilter) GetHiddenEnvoyDeprecatedConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ThriftFilter_HiddenEnvoyDeprecatedConfig); ok {
		return x.HiddenEnvoyDeprecatedConfig
	}
	return nil
}

func (m *ThriftFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ThriftFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ThriftFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ThriftFilter_HiddenEnvoyDeprecatedConfig)(nil),
		(*ThriftFilter_TypedConfig)(nil),
	}
}

type ThriftProtocolOptions struct {
	Transport            TransportType `protobuf:"varint,1,opt,name=transport,proto3,enum=envoy.config.filter.network.thrift_proxy.v3alpha.TransportType" json:"transport,omitempty"`
	Protocol             ProtocolType  `protobuf:"varint,2,opt,name=protocol,proto3,enum=envoy.config.filter.network.thrift_proxy.v3alpha.ProtocolType" json:"protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ThriftProtocolOptions) Reset()         { *m = ThriftProtocolOptions{} }
func (m *ThriftProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*ThriftProtocolOptions) ProtoMessage()    {}
func (*ThriftProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_5bf22ec9d8f099b6, []int{2}
}

func (m *ThriftProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThriftProtocolOptions.Unmarshal(m, b)
}
func (m *ThriftProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThriftProtocolOptions.Marshal(b, m, deterministic)
}
func (m *ThriftProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThriftProtocolOptions.Merge(m, src)
}
func (m *ThriftProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_ThriftProtocolOptions.Size(m)
}
func (m *ThriftProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ThriftProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ThriftProtocolOptions proto.InternalMessageInfo

func (m *ThriftProtocolOptions) GetTransport() TransportType {
	if m != nil {
		return m.Transport
	}
	return TransportType_AUTO_TRANSPORT
}

func (m *ThriftProtocolOptions) GetProtocol() ProtocolType {
	if m != nil {
		return m.Protocol
	}
	return ProtocolType_AUTO_PROTOCOL
}

func init() {
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v3alpha.TransportType", TransportType_name, TransportType_value)
	proto.RegisterEnum("envoy.config.filter.network.thrift_proxy.v3alpha.ProtocolType", ProtocolType_name, ProtocolType_value)
	proto.RegisterType((*ThriftProxy)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.ThriftProxy")
	proto.RegisterType((*ThriftFilter)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.ThriftFilter")
	proto.RegisterType((*ThriftProtocolOptions)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.ThriftProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v3alpha/thrift_proxy.proto", fileDescriptor_5bf22ec9d8f099b6)
}

var fileDescriptor_5bf22ec9d8f099b6 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x6b, 0x27, 0xfd, 0x37, 0x4e, 0x22, 0xff, 0x56, 0x3f, 0xd4, 0xd0, 0xa2, 0x2a, 0xf4,
	0x14, 0xf5, 0x60, 0x43, 0x7b, 0xa2, 0x82, 0x16, 0x3b, 0x49, 0x69, 0xa4, 0x36, 0xb6, 0xb6, 0xae,
	0x0a, 0x27, 0xcb, 0x8d, 0x37, 0xa9, 0x45, 0xf0, 0x5a, 0xeb, 0x4d, 0x68, 0xae, 0x9c, 0x78, 0x06,
	0xc4, 0x81, 0xa7, 0xe0, 0x0d, 0x78, 0xa9, 0x9e, 0x90, 0x77, 0x9d, 0x36, 0x09, 0x17, 0x52, 0x04,
	0x37, 0x8f, 0x67, 0xf6, 0xf3, 0x9d, 0xfd, 0xee, 0x0c, 0x34, 0x48, 0x3c, 0xa2, 0x63, 0xb3, 0x4b,
	0xe3, 0x5e, 0xd4, 0x37, 0x7b, 0xd1, 0x80, 0x13, 0x66, 0xc6, 0x84, 0x7f, 0xa4, 0xec, 0xbd, 0xc9,
	0xaf, 0x59, 0xd4, 0xe3, 0x7e, 0xc2, 0xe8, 0xcd, 0xd8, 0x1c, 0xed, 0x07, 0x83, 0xe4, 0x3a, 0x98,
	0xf9, 0x69, 0x24, 0x8c, 0x72, 0x8a, 0x9e, 0x09, 0x88, 0x21, 0x21, 0x86, 0x84, 0x18, 0x39, 0xc4,
	0x98, 0xa9, 0xcf, 0x21, 0x9b, 0x2f, 0x17, 0x96, 0x65, 0x74, 0xc8, 0x89, 0xd4, 0xdb, 0x7c, 0xdc,
	0xa7, 0xb4, 0x3f, 0x20, 0xa6, 0x88, 0xae, 0x86, 0x3d, 0x33, 0x88, 0xf3, 0x56, 0x36, 0x9f, 0xcc,
	0xa7, 0x52, 0xce, 0x86, 0x5d, 0x9e, 0x67, 0x9f, 0x0e, 0xc3, 0x24, 0x30, 0x83, 0x38, 0xa6, 0x3c,
	0xe0, 0x11, 0x8d, 0x53, 0x73, 0x44, 0x58, 0x1a, 0xd1, 0x38, 0x8a, 0xfb, 0x79, 0xc9, 0xc6, 0x28,
	0x18, 0x44, 0x61, 0xc0, 0x89, 0x39, 0xf9, 0x90, 0x89, 0x9d, 0xaf, 0x45, 0xd0, 0x3c, 0xd1, 0x99,
	0x9b, 0x35, 0x86, 0xfa, 0xb0, 0xce, 0x59, 0x10, 0xa7, 0x09, 0x65, 0xbc, 0xaa, 0xd6, 0x94, 0x7a,
	0x65, 0xef, 0xc8, 0x58, 0xd4, 0x08, 0xc3, 0x9b, 0x20, 0xbc, 0x71, 0x42, 0xec, 0xb5, 0x5b, 0x7b,
	0xf9, 0x93, 0xa2, 0xea, 0x0a, 0xbe, 0x67, 0xa3, 0x10, 0xd6, 0x44, 0x07, 0x5d, 0x3a, 0xa8, 0x16,
	0x84, 0xce, 0xe1, 0xe2, 0x3a, 0x6e, 0x4e, 0x98, 0x93, 0xb9, 0x23, 0xa3, 0x3a, 0x68, 0x29, 0x0f,
	0xb2, 0x63, 0xa4, 0x17, 0xdd, 0x54, 0x95, 0x9a, 0x52, 0x5f, 0xb7, 0x57, 0x6f, 0xed, 0x22, 0x53,
	0x6b, 0x0a, 0x86, 0x2c, 0xe7, 0x8a, 0x14, 0xea, 0x43, 0x49, 0x3c, 0x86, 0x2f, 0xe5, 0xab, 0xc5,
	0x9a, 0x52, 0xd7, 0xf6, 0x9a, 0x8b, 0xf7, 0x84, 0x33, 0x4a, 0x43, 0xd4, 0x0f, 0x99, 0x78, 0x0f,
	0xac, 0xb1, 0xfb, 0x7f, 0x88, 0x40, 0x25, 0x3f, 0x27, 0x69, 0x69, 0x75, 0xb9, 0x56, 0xa8, 0x6b,
	0x0f, 0xb9, 0xbe, 0x7c, 0xb8, 0x63, 0x51, 0x89, 0xcb, 0x7c, 0x2a, 0x4a, 0x0f, 0x9a, 0x5f, 0x7e,
	0x7c, 0xde, 0x3e, 0x82, 0x57, 0xbf, 0x0f, 0xdd, 0x13, 0xd0, 0xe7, 0xc6, 0xd4, 0x38, 0xec, 0x7c,
	0x53, 0xa1, 0x34, 0xad, 0x82, 0xb6, 0xa0, 0x18, 0x07, 0x1f, 0xc8, 0xbc, 0x93, 0xe2, 0x27, 0xba,
	0x82, 0xed, 0xeb, 0x28, 0x0c, 0x49, 0xec, 0x0b, 0x55, 0x3f, 0x24, 0x09, 0x23, 0xdd, 0x80, 0x93,
	0x70, 0xe2, 0xaa, 0x2a, 0x5c, 0xdd, 0x30, 0xe4, 0x3c, 0x1b, 0x93, 0x79, 0x36, 0xce, 0xc5, 0x3c,
	0xdb, 0x6a, 0x55, 0x39, 0x59, 0xc2, 0x5b, 0x12, 0xd2, 0xca, 0x18, 0xcd, 0x3b, 0x44, 0x6e, 0xdf,
	0x0b, 0x28, 0xf1, 0x71, 0x72, 0x4f, 0x2c, 0x08, 0xe2, 0xff, 0xbf, 0x10, 0xad, 0x78, 0x7c, 0xb2,
	0x84, 0x35, 0x51, 0x2b, 0x8f, 0x1e, 0xb4, 0x32, 0x4b, 0x5e, 0xc3, 0xe1, 0x43, 0x2d, 0x91, 0x16,
	0xd8, 0x65, 0xd0, 0xe4, 0x59, 0x3f, 0x83, 0xef, 0x7c, 0x57, 0xe1, 0xd1, 0x9d, 0x65, 0x62, 0xea,
	0x9c, 0x44, 0x6c, 0xe1, 0xec, 0x2e, 0x29, 0xff, 0x68, 0x97, 0xd4, 0xbf, 0xb5, 0x4b, 0x07, 0x9d,
	0xcc, 0xbe, 0x36, 0xbc, 0xf9, 0x83, 0x89, 0x9a, 0xb6, 0x67, 0xb7, 0x0d, 0xe5, 0x99, 0xbb, 0x21,
	0x04, 0x15, 0xeb, 0xc2, 0x73, 0x7c, 0x0f, 0x5b, 0x9d, 0x73, 0xd7, 0xc1, 0x9e, 0xbe, 0x84, 0x00,
	0x56, 0x8e, 0xb1, 0x75, 0xd6, 0x6a, 0xea, 0x0a, 0x2a, 0xc1, 0xda, 0x45, 0x27, 0x8f, 0xd4, 0x2c,
	0x73, 0xd2, 0xb2, 0x9a, 0x2d, 0xac, 0x17, 0x76, 0x2f, 0xa1, 0x34, 0xdd, 0x3e, 0xfa, 0x0f, 0xca,
	0x82, 0xe4, 0x62, 0xc7, 0x73, 0x1a, 0xce, 0xa9, 0x04, 0xd9, 0xed, 0x8e, 0x85, 0xdf, 0xe9, 0x0a,
	0xaa, 0x00, 0x9c, 0x5a, 0x6f, 0xfd, 0x3c, 0x56, 0x91, 0x06, 0xab, 0x0d, 0xe7, 0xcc, 0xb5, 0x1a,
	0x9e, 0x5e, 0xc8, 0x02, 0xef, 0xb2, 0xed, 0x79, 0x2d, 0xac, 0x17, 0x6d, 0x0f, 0x0e, 0x23, 0x2a,
	0xbd, 0x94, 0x77, 0x5a, 0xd4, 0x56, 0x5b, 0x9f, 0x5a, 0x27, 0xd1, 0xa3, 0xab, 0x5c, 0xad, 0x08,
	0x4f, 0xf7, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xc8, 0xba, 0xd8, 0xa7, 0x06, 0x00, 0x00,
}
