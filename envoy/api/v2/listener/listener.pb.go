// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/listener/listener.proto

package envoy_api_v2_listener

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	auth "github.com/envoyproxy/go-control-plane/envoy/api/v2/auth"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type FilterChainMatch_ConnectionSourceType int32

const (
	FilterChainMatch_ANY      FilterChainMatch_ConnectionSourceType = 0
	FilterChainMatch_LOCAL    FilterChainMatch_ConnectionSourceType = 1
	FilterChainMatch_EXTERNAL FilterChainMatch_ConnectionSourceType = 2
)

var FilterChainMatch_ConnectionSourceType_name = map[int32]string{
	0: "ANY",
	1: "LOCAL",
	2: "EXTERNAL",
}

var FilterChainMatch_ConnectionSourceType_value = map[string]int32{
	"ANY":      0,
	"LOCAL":    1,
	"EXTERNAL": 2,
}

func (x FilterChainMatch_ConnectionSourceType) String() string {
	return proto.EnumName(FilterChainMatch_ConnectionSourceType_name, int32(x))
}

func (FilterChainMatch_ConnectionSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{1, 0}
}

type Filter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*Filter_Config
	//	*Filter_TypedConfig
	ConfigType           isFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isFilter_ConfigType interface {
	isFilter_ConfigType()
}

type Filter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type Filter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,4,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*Filter_Config) isFilter_ConfigType() {}

func (*Filter_TypedConfig) isFilter_ConfigType() {}

func (m *Filter) GetConfigType() isFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *Filter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*Filter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *Filter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*Filter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Filter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Filter_Config)(nil),
		(*Filter_TypedConfig)(nil),
	}
}

type FilterChainMatch struct {
	DestinationPort      *wrappers.UInt32Value                 `protobuf:"bytes,8,opt,name=destination_port,json=destinationPort,proto3" json:"destination_port,omitempty"`
	PrefixRanges         []*core.CidrRange                     `protobuf:"bytes,3,rep,name=prefix_ranges,json=prefixRanges,proto3" json:"prefix_ranges,omitempty"`
	AddressSuffix        string                                `protobuf:"bytes,4,opt,name=address_suffix,json=addressSuffix,proto3" json:"address_suffix,omitempty"`
	SuffixLen            *wrappers.UInt32Value                 `protobuf:"bytes,5,opt,name=suffix_len,json=suffixLen,proto3" json:"suffix_len,omitempty"`
	SourceType           FilterChainMatch_ConnectionSourceType `protobuf:"varint,12,opt,name=source_type,json=sourceType,proto3,enum=envoy.api.v2.listener.FilterChainMatch_ConnectionSourceType" json:"source_type,omitempty"`
	SourcePrefixRanges   []*core.CidrRange                     `protobuf:"bytes,6,rep,name=source_prefix_ranges,json=sourcePrefixRanges,proto3" json:"source_prefix_ranges,omitempty"`
	SourcePorts          []uint32                              `protobuf:"varint,7,rep,packed,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	ServerNames          []string                              `protobuf:"bytes,11,rep,name=server_names,json=serverNames,proto3" json:"server_names,omitempty"`
	TransportProtocol    string                                `protobuf:"bytes,9,opt,name=transport_protocol,json=transportProtocol,proto3" json:"transport_protocol,omitempty"`
	ApplicationProtocols []string                              `protobuf:"bytes,10,rep,name=application_protocols,json=applicationProtocols,proto3" json:"application_protocols,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *FilterChainMatch) Reset()         { *m = FilterChainMatch{} }
func (m *FilterChainMatch) String() string { return proto.CompactTextString(m) }
func (*FilterChainMatch) ProtoMessage()    {}
func (*FilterChainMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{1}
}

func (m *FilterChainMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChainMatch.Unmarshal(m, b)
}
func (m *FilterChainMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChainMatch.Marshal(b, m, deterministic)
}
func (m *FilterChainMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChainMatch.Merge(m, src)
}
func (m *FilterChainMatch) XXX_Size() int {
	return xxx_messageInfo_FilterChainMatch.Size(m)
}
func (m *FilterChainMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChainMatch.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChainMatch proto.InternalMessageInfo

func (m *FilterChainMatch) GetDestinationPort() *wrappers.UInt32Value {
	if m != nil {
		return m.DestinationPort
	}
	return nil
}

func (m *FilterChainMatch) GetPrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.PrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetAddressSuffix() string {
	if m != nil {
		return m.AddressSuffix
	}
	return ""
}

func (m *FilterChainMatch) GetSuffixLen() *wrappers.UInt32Value {
	if m != nil {
		return m.SuffixLen
	}
	return nil
}

func (m *FilterChainMatch) GetSourceType() FilterChainMatch_ConnectionSourceType {
	if m != nil {
		return m.SourceType
	}
	return FilterChainMatch_ANY
}

func (m *FilterChainMatch) GetSourcePrefixRanges() []*core.CidrRange {
	if m != nil {
		return m.SourcePrefixRanges
	}
	return nil
}

func (m *FilterChainMatch) GetSourcePorts() []uint32 {
	if m != nil {
		return m.SourcePorts
	}
	return nil
}

func (m *FilterChainMatch) GetServerNames() []string {
	if m != nil {
		return m.ServerNames
	}
	return nil
}

func (m *FilterChainMatch) GetTransportProtocol() string {
	if m != nil {
		return m.TransportProtocol
	}
	return ""
}

func (m *FilterChainMatch) GetApplicationProtocols() []string {
	if m != nil {
		return m.ApplicationProtocols
	}
	return nil
}

type FilterChain struct {
	FilterChainMatch     *FilterChainMatch          `protobuf:"bytes,1,opt,name=filter_chain_match,json=filterChainMatch,proto3" json:"filter_chain_match,omitempty"`
	TlsContext           *auth.DownstreamTlsContext `protobuf:"bytes,2,opt,name=tls_context,json=tlsContext,proto3" json:"tls_context,omitempty"` // Deprecated: Do not use.
	Filters              []*Filter                  `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	UseProxyProto        *wrappers.BoolValue        `protobuf:"bytes,4,opt,name=use_proxy_proto,json=useProxyProto,proto3" json:"use_proxy_proto,omitempty"`
	Metadata             *core.Metadata             `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	TransportSocket      *core.TransportSocket      `protobuf:"bytes,6,opt,name=transport_socket,json=transportSocket,proto3" json:"transport_socket,omitempty"`
	Name                 string                     `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *FilterChain) Reset()         { *m = FilterChain{} }
func (m *FilterChain) String() string { return proto.CompactTextString(m) }
func (*FilterChain) ProtoMessage()    {}
func (*FilterChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{2}
}

func (m *FilterChain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterChain.Unmarshal(m, b)
}
func (m *FilterChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterChain.Marshal(b, m, deterministic)
}
func (m *FilterChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterChain.Merge(m, src)
}
func (m *FilterChain) XXX_Size() int {
	return xxx_messageInfo_FilterChain.Size(m)
}
func (m *FilterChain) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterChain.DiscardUnknown(m)
}

var xxx_messageInfo_FilterChain proto.InternalMessageInfo

func (m *FilterChain) GetFilterChainMatch() *FilterChainMatch {
	if m != nil {
		return m.FilterChainMatch
	}
	return nil
}

// Deprecated: Do not use.
func (m *FilterChain) GetTlsContext() *auth.DownstreamTlsContext {
	if m != nil {
		return m.TlsContext
	}
	return nil
}

func (m *FilterChain) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *FilterChain) GetUseProxyProto() *wrappers.BoolValue {
	if m != nil {
		return m.UseProxyProto
	}
	return nil
}

func (m *FilterChain) GetMetadata() *core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *FilterChain) GetTransportSocket() *core.TransportSocket {
	if m != nil {
		return m.TransportSocket
	}
	return nil
}

func (m *FilterChain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListenerFilter struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to ConfigType:
	//	*ListenerFilter_Config
	//	*ListenerFilter_TypedConfig
	ConfigType           isListenerFilter_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ListenerFilter) Reset()         { *m = ListenerFilter{} }
func (m *ListenerFilter) String() string { return proto.CompactTextString(m) }
func (*ListenerFilter) ProtoMessage()    {}
func (*ListenerFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ced541f18749edd, []int{3}
}

func (m *ListenerFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenerFilter.Unmarshal(m, b)
}
func (m *ListenerFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenerFilter.Marshal(b, m, deterministic)
}
func (m *ListenerFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenerFilter.Merge(m, src)
}
func (m *ListenerFilter) XXX_Size() int {
	return xxx_messageInfo_ListenerFilter.Size(m)
}
func (m *ListenerFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenerFilter.DiscardUnknown(m)
}

var xxx_messageInfo_ListenerFilter proto.InternalMessageInfo

func (m *ListenerFilter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type isListenerFilter_ConfigType interface {
	isListenerFilter_ConfigType()
}

type ListenerFilter_Config struct {
	Config *_struct.Struct `protobuf:"bytes,2,opt,name=config,proto3,oneof"`
}

type ListenerFilter_TypedConfig struct {
	TypedConfig *any.Any `protobuf:"bytes,3,opt,name=typed_config,json=typedConfig,proto3,oneof"`
}

func (*ListenerFilter_Config) isListenerFilter_ConfigType() {}

func (*ListenerFilter_TypedConfig) isListenerFilter_ConfigType() {}

func (m *ListenerFilter) GetConfigType() isListenerFilter_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

// Deprecated: Do not use.
func (m *ListenerFilter) GetConfig() *_struct.Struct {
	if x, ok := m.GetConfigType().(*ListenerFilter_Config); ok {
		return x.Config
	}
	return nil
}

func (m *ListenerFilter) GetTypedConfig() *any.Any {
	if x, ok := m.GetConfigType().(*ListenerFilter_TypedConfig); ok {
		return x.TypedConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ListenerFilter) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ListenerFilter_Config)(nil),
		(*ListenerFilter_TypedConfig)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.api.v2.listener.FilterChainMatch_ConnectionSourceType", FilterChainMatch_ConnectionSourceType_name, FilterChainMatch_ConnectionSourceType_value)
	proto.RegisterType((*Filter)(nil), "envoy.api.v2.listener.Filter")
	proto.RegisterType((*FilterChainMatch)(nil), "envoy.api.v2.listener.FilterChainMatch")
	proto.RegisterType((*FilterChain)(nil), "envoy.api.v2.listener.FilterChain")
	proto.RegisterType((*ListenerFilter)(nil), "envoy.api.v2.listener.ListenerFilter")
}

func init() {
	proto.RegisterFile("envoy/api/v2/listener/listener.proto", fileDescriptor_0ced541f18749edd)
}

var fileDescriptor_0ced541f18749edd = []byte{
	// 905 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0xae, 0xf3, 0x9d, 0xd7, 0x49, 0x6b, 0x86, 0x56, 0x35, 0x6d, 0x59, 0x42, 0xf8, 0x8a, 0x90,
	0xb0, 0xa5, 0x54, 0x68, 0x85, 0xe0, 0x12, 0x87, 0xa2, 0xdd, 0x25, 0x4d, 0x23, 0x27, 0x5d, 0xc1,
	0xc9, 0x9a, 0x3a, 0x93, 0xd4, 0xe0, 0xcc, 0x58, 0x33, 0x93, 0x6c, 0x73, 0xe5, 0x0f, 0x20, 0x71,
	0xe2, 0x37, 0x20, 0xc4, 0x1f, 0xe1, 0x57, 0x70, 0xe5, 0xc8, 0xb1, 0x07, 0x16, 0x79, 0x6c, 0x67,
	0x9b, 0x34, 0x40, 0xc5, 0x81, 0xdb, 0xcc, 0xfb, 0x3c, 0xef, 0xf7, 0x07, 0xbc, 0x4b, 0xe8, 0x82,
	0x2d, 0x6d, 0x1c, 0x05, 0xf6, 0xa2, 0x6d, 0x87, 0x81, 0x90, 0x84, 0x12, 0xbe, 0x7a, 0x58, 0x11,
	0x67, 0x92, 0xa1, 0x03, 0xc5, 0xb2, 0x70, 0x14, 0x58, 0x8b, 0xb6, 0x95, 0x81, 0x47, 0x27, 0x6b,
	0xca, 0x78, 0x2e, 0xaf, 0x6d, 0x9f, 0x70, 0x99, 0x28, 0x1d, 0xbd, 0xb5, 0x86, 0xfa, 0x8c, 0x13,
	0x1b, 0x8f, 0xc7, 0x9c, 0x08, 0x91, 0x12, 0x4e, 0xee, 0x13, 0xae, 0xb0, 0x20, 0x29, 0xfa, 0xc6,
	0x94, 0xb1, 0x69, 0x48, 0x6c, 0xf5, 0xbb, 0x9a, 0x4f, 0x6c, 0x4c, 0x97, 0x99, 0xe2, 0x26, 0x24,
	0x24, 0x9f, 0xfb, 0x99, 0xdf, 0x47, 0x9b, 0xe8, 0x0b, 0x8e, 0xa3, 0x88, 0xf0, 0xcc, 0xed, 0xa3,
	0xf9, 0x38, 0xc2, 0x36, 0xa6, 0x94, 0x49, 0x2c, 0x03, 0x46, 0x85, 0x3d, 0x0b, 0xa6, 0x1c, 0xcb,
	0xcc, 0xf1, 0xe1, 0x02, 0x87, 0xc1, 0x18, 0x4b, 0x62, 0x67, 0x8f, 0x04, 0x68, 0xfe, 0xac, 0x41,
	0xe9, 0x8b, 0x20, 0x94, 0x84, 0xa3, 0x63, 0x28, 0x50, 0x3c, 0x23, 0xa6, 0xd6, 0xd0, 0x5a, 0x55,
	0xa7, 0x7c, 0xeb, 0x14, 0x78, 0xae, 0xa1, 0xb9, 0x4a, 0x88, 0x3e, 0x86, 0x92, 0xcf, 0xe8, 0x24,
	0x98, 0x9a, 0xb9, 0x86, 0xd6, 0xd2, 0xdb, 0x87, 0x56, 0x12, 0x91, 0x95, 0x45, 0x64, 0x0d, 0x55,
	0xbc, 0x4e, 0xce, 0xd4, 0x9e, 0xec, 0xb8, 0x29, 0x19, 0x7d, 0x02, 0x35, 0xb9, 0x8c, 0xc8, 0xd8,
	0x4b, 0x95, 0x0b, 0x4a, 0x79, 0xff, 0x9e, 0x72, 0x87, 0x2e, 0x9f, 0xec, 0xb8, 0xba, 0xe2, 0x76,
	0x15, 0xd5, 0xa9, 0x83, 0x9e, 0x28, 0x79, 0xb1, 0xf4, 0x59, 0xa1, 0x92, 0x37, 0x0a, 0xcd, 0xdf,
	0x8a, 0x60, 0x24, 0xe1, 0x76, 0xaf, 0x71, 0x40, 0xcf, 0xb1, 0xf4, 0xaf, 0xd1, 0x08, 0x8c, 0x31,
	0x11, 0x32, 0xa0, 0x2a, 0x75, 0x2f, 0x62, 0x5c, 0x9a, 0x15, 0xe5, 0xe8, 0xe4, 0x9e, 0xa3, 0xcb,
	0xa7, 0x54, 0x9e, 0xb6, 0x9f, 0xe3, 0x70, 0x4e, 0x1c, 0xfd, 0xd6, 0xa9, 0x7c, 0x58, 0x32, 0x5f,
	0xbe, 0xcc, 0xb7, 0x34, 0x77, 0xef, 0x8e, 0x89, 0x01, 0xe3, 0x12, 0x75, 0xa0, 0x1e, 0x71, 0x32,
	0x09, 0x6e, 0x3c, 0x8e, 0xe9, 0x94, 0x08, 0x33, 0xdf, 0xc8, 0x2b, 0x93, 0x6b, 0x73, 0x13, 0x77,
	0xd8, 0xea, 0x06, 0x63, 0xee, 0xc6, 0x24, 0xb7, 0x96, 0xa8, 0xa8, 0x8f, 0x40, 0xef, 0xc1, 0x6e,
	0x3a, 0x1d, 0x9e, 0x98, 0x4f, 0x26, 0xc1, 0x8d, 0xca, 0xbf, 0xea, 0xd6, 0x53, 0xe9, 0x50, 0x09,
	0xd1, 0xa7, 0x00, 0x09, 0xec, 0x85, 0x84, 0x9a, 0xc5, 0x7f, 0x8f, 0xdc, 0xad, 0x26, 0xfc, 0x1e,
	0xa1, 0x68, 0x0a, 0xba, 0x60, 0x73, 0xee, 0x13, 0x55, 0x26, 0xb3, 0xd6, 0xd0, 0x5a, 0xbb, 0xed,
	0xcf, 0xac, 0xad, 0xc3, 0x6d, 0x6d, 0x96, 0xce, 0xea, 0x32, 0x4a, 0x89, 0x1f, 0xe7, 0x3c, 0x54,
	0x46, 0x46, 0xcb, 0x88, 0x38, 0x95, 0x5b, 0xa7, 0xf8, 0x9d, 0x96, 0x33, 0x34, 0x17, 0xc4, 0x4a,
	0x8a, 0xfa, 0xb0, 0x9f, 0x3a, 0x5a, 0x2f, 0x4b, 0xe9, 0x01, 0x65, 0x41, 0x89, 0xe6, 0xe0, 0x6e,
	0x71, 0x4e, 0xa1, 0x96, 0xd9, 0x63, 0x5c, 0x0a, 0xb3, 0xdc, 0xc8, 0xb7, 0xea, 0x8e, 0x71, 0xeb,
	0xd4, 0x7f, 0xd0, 0xa0, 0xf9, 0xaa, 0x31, 0x69, 0x7a, 0x71, 0x4f, 0x04, 0x7a, 0x1b, 0x6a, 0x82,
	0xf0, 0x05, 0xe1, 0x5e, 0x3c, 0x95, 0xc2, 0xd4, 0x1b, 0xf9, 0x56, 0xd5, 0xd5, 0x13, 0x59, 0x3f,
	0x16, 0xa1, 0x8f, 0x00, 0x49, 0x8e, 0xa9, 0x88, 0xad, 0x7a, 0xaa, 0x7a, 0x3e, 0x0b, 0xcd, 0xaa,
	0x2a, 0xfc, 0x6b, 0x2b, 0x64, 0x90, 0x02, 0xe8, 0x14, 0x0e, 0x70, 0x14, 0x85, 0x81, 0x9f, 0x0e,
	0x4f, 0x2a, 0x17, 0x26, 0x28, 0xd3, 0xfb, 0x77, 0xc0, 0x4c, 0x47, 0x34, 0x2f, 0x61, 0x7f, 0x5b,
	0xe5, 0x50, 0x19, 0xf2, 0x9d, 0xfe, 0xd7, 0xc6, 0x0e, 0x7a, 0x1f, 0x8a, 0xbd, 0x8b, 0x6e, 0xa7,
	0x67, 0x68, 0x47, 0xc7, 0x7f, 0xfc, 0xf8, 0xe7, 0xf7, 0xc5, 0x03, 0x78, 0x7d, 0xd8, 0x39, 0x3f,
	0xf3, 0x9e, 0x0e, 0xbc, 0x0b, 0xd7, 0xeb, 0x5d, 0x5c, 0x0c, 0x9c, 0x4e, 0xf7, 0x4b, 0x54, 0x83,
	0xca, 0xd9, 0x57, 0xa3, 0x33, 0xb7, 0xdf, 0xe9, 0x19, 0xb9, 0x67, 0x85, 0x8a, 0x66, 0xe4, 0x9a,
	0xbf, 0xe6, 0x41, 0xbf, 0xd3, 0x28, 0x74, 0x09, 0x68, 0xa2, 0xbe, 0x9e, 0x1f, 0xff, 0xbd, 0x59,
	0xdc, 0x39, 0xb5, 0xa5, 0x7a, 0xfb, 0x83, 0x07, 0x36, 0xda, 0x35, 0x26, 0x9b, 0x5b, 0xd3, 0x03,
	0x5d, 0x86, 0x22, 0x5e, 0x4c, 0x49, 0x6e, 0x64, 0xba, 0xd6, 0x1b, 0xf6, 0xe2, 0xf3, 0x67, 0x7d,
	0xce, 0x5e, 0x50, 0x21, 0x39, 0xc1, 0xb3, 0x51, 0x28, 0xba, 0x09, 0x3d, 0x5e, 0x73, 0x17, 0xe4,
	0xea, 0x8f, 0x1e, 0x43, 0x39, 0xf1, 0x90, 0xed, 0xc9, 0x9b, 0xff, 0x18, 0x99, 0x9b, 0xb1, 0x91,
	0x03, 0x7b, 0x73, 0x11, 0xcf, 0x14, 0xbb, 0x59, 0x26, 0xd5, 0x4f, 0x8f, 0xc4, 0xd1, 0xbd, 0x0d,
	0x70, 0x18, 0x0b, 0x93, 0xf9, 0xaf, 0xcf, 0x05, 0x19, 0xc4, 0x1a, 0xaa, 0x25, 0xe8, 0x31, 0x54,
	0x66, 0x44, 0xe2, 0x31, 0x96, 0x38, 0x5d, 0x9f, 0xe3, 0x2d, 0xe3, 0x78, 0x9e, 0x52, 0xdc, 0x15,
	0x19, 0x9d, 0x83, 0xf1, 0x6a, 0x56, 0x04, 0xf3, 0xbf, 0x25, 0xd2, 0x2c, 0x29, 0x03, 0xcd, 0x2d,
	0x06, 0x46, 0x19, 0x75, 0xa8, 0x98, 0xee, 0x9e, 0x5c, 0x17, 0x20, 0x94, 0x5e, 0xd0, 0xb2, 0x1a,
	0x36, 0xf5, 0x6e, 0xfe, 0xa2, 0xc1, 0x6e, 0x2f, 0x4d, 0xfe, 0x7f, 0x3c, 0xb4, 0xf9, 0xff, 0x7a,
	0x68, 0x9d, 0x6f, 0xe0, 0x9d, 0x80, 0x25, 0xd9, 0xab, 0xa6, 0x6c, 0xef, 0xa3, 0x53, 0xcf, 0x92,
	0x52, 0x2d, 0x18, 0x68, 0x3f, 0xe5, 0x0e, 0xcf, 0x14, 0xb1, 0x13, 0x05, 0xd6, 0xf3, 0xb6, 0x95,
	0xc1, 0xfd, 0xe1, 0xef, 0x7f, 0x8b, 0x5c, 0x95, 0x54, 0x5c, 0xa7, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xd0, 0xe7, 0x39, 0x72, 0xb2, 0x07, 0x00, 0x00,
}
