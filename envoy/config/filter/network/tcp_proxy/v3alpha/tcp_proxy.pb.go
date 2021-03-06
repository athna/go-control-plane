// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/tcp_proxy/v3alpha/tcp_proxy.proto

package envoy_config_filter_network_tcp_proxy_v3alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v3alpha"
	v3alpha1 "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type TcpProxy struct {
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*TcpProxy_Cluster
	//	*TcpProxy_WeightedClusters
	ClusterSpecifier                  isTcpProxy_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	MetadataMatch                     *core.Metadata              `protobuf:"bytes,9,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	IdleTimeout                       *duration.Duration          `protobuf:"bytes,8,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	DownstreamIdleTimeout             *duration.Duration          `protobuf:"bytes,3,opt,name=downstream_idle_timeout,json=downstreamIdleTimeout,proto3" json:"downstream_idle_timeout,omitempty"`
	UpstreamIdleTimeout               *duration.Duration          `protobuf:"bytes,4,opt,name=upstream_idle_timeout,json=upstreamIdleTimeout,proto3" json:"upstream_idle_timeout,omitempty"`
	AccessLog                         []*v3alpha.AccessLog        `protobuf:"bytes,5,rep,name=access_log,json=accessLog,proto3" json:"access_log,omitempty"`
	HiddenEnvoyDeprecatedDeprecatedV1 *TcpProxy_DeprecatedV1      `protobuf:"bytes,6,opt,name=hidden_envoy_deprecated_deprecated_v1,json=hiddenEnvoyDeprecatedDeprecatedV1,proto3" json:"hidden_envoy_deprecated_deprecated_v1,omitempty"` // Deprecated: Do not use.
	MaxConnectAttempts                *wrappers.UInt32Value       `protobuf:"bytes,7,opt,name=max_connect_attempts,json=maxConnectAttempts,proto3" json:"max_connect_attempts,omitempty"`
	HashPolicy                        []*v3alpha1.HashPolicy      `protobuf:"bytes,11,rep,name=hash_policy,json=hashPolicy,proto3" json:"hash_policy,omitempty"`
	XXX_NoUnkeyedLiteral              struct{}                    `json:"-"`
	XXX_unrecognized                  []byte                      `json:"-"`
	XXX_sizecache                     int32                       `json:"-"`
}

func (m *TcpProxy) Reset()         { *m = TcpProxy{} }
func (m *TcpProxy) String() string { return proto.CompactTextString(m) }
func (*TcpProxy) ProtoMessage()    {}
func (*TcpProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cada114fb4da110, []int{0}
}

func (m *TcpProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy.Unmarshal(m, b)
}
func (m *TcpProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy.Marshal(b, m, deterministic)
}
func (m *TcpProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy.Merge(m, src)
}
func (m *TcpProxy) XXX_Size() int {
	return xxx_messageInfo_TcpProxy.Size(m)
}
func (m *TcpProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy proto.InternalMessageInfo

func (m *TcpProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

type isTcpProxy_ClusterSpecifier interface {
	isTcpProxy_ClusterSpecifier()
}

type TcpProxy_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

type TcpProxy_WeightedClusters struct {
	WeightedClusters *TcpProxy_WeightedCluster `protobuf:"bytes,10,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*TcpProxy_Cluster) isTcpProxy_ClusterSpecifier() {}

func (*TcpProxy_WeightedClusters) isTcpProxy_ClusterSpecifier() {}

func (m *TcpProxy) GetClusterSpecifier() isTcpProxy_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *TcpProxy) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *TcpProxy) GetWeightedClusters() *TcpProxy_WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*TcpProxy_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *TcpProxy) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *TcpProxy) GetIdleTimeout() *duration.Duration {
	if m != nil {
		return m.IdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetDownstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.DownstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetUpstreamIdleTimeout() *duration.Duration {
	if m != nil {
		return m.UpstreamIdleTimeout
	}
	return nil
}

func (m *TcpProxy) GetAccessLog() []*v3alpha.AccessLog {
	if m != nil {
		return m.AccessLog
	}
	return nil
}

// Deprecated: Do not use.
func (m *TcpProxy) GetHiddenEnvoyDeprecatedDeprecatedV1() *TcpProxy_DeprecatedV1 {
	if m != nil {
		return m.HiddenEnvoyDeprecatedDeprecatedV1
	}
	return nil
}

func (m *TcpProxy) GetMaxConnectAttempts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxConnectAttempts
	}
	return nil
}

func (m *TcpProxy) GetHashPolicy() []*v3alpha1.HashPolicy {
	if m != nil {
		return m.HashPolicy
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TcpProxy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TcpProxy_Cluster)(nil),
		(*TcpProxy_WeightedClusters)(nil),
	}
}

// Deprecated: Do not use.
type TcpProxy_DeprecatedV1 struct {
	Routes               []*TcpProxy_DeprecatedV1_TCPRoute `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1) Reset()         { *m = TcpProxy_DeprecatedV1{} }
func (m *TcpProxy_DeprecatedV1) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cada114fb4da110, []int{0, 0}
}

func (m *TcpProxy_DeprecatedV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1.Size(m)
}
func (m *TcpProxy_DeprecatedV1) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1 proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1) GetRoutes() []*TcpProxy_DeprecatedV1_TCPRoute {
	if m != nil {
		return m.Routes
	}
	return nil
}

type TcpProxy_DeprecatedV1_TCPRoute struct {
	Cluster              string            `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	DestinationIpList    []*core.CidrRange `protobuf:"bytes,2,rep,name=destination_ip_list,json=destinationIpList,proto3" json:"destination_ip_list,omitempty"`
	DestinationPorts     string            `protobuf:"bytes,3,opt,name=destination_ports,json=destinationPorts,proto3" json:"destination_ports,omitempty"`
	SourceIpList         []*core.CidrRange `protobuf:"bytes,4,rep,name=source_ip_list,json=sourceIpList,proto3" json:"source_ip_list,omitempty"`
	SourcePorts          string            `protobuf:"bytes,5,opt,name=source_ports,json=sourcePorts,proto3" json:"source_ports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) Reset()         { *m = TcpProxy_DeprecatedV1_TCPRoute{} }
func (m *TcpProxy_DeprecatedV1_TCPRoute) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_DeprecatedV1_TCPRoute) ProtoMessage()    {}
func (*TcpProxy_DeprecatedV1_TCPRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cada114fb4da110, []int{0, 0, 0}
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Unmarshal(m, b)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Marshal(b, m, deterministic)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Merge(m, src)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.Size(m)
}
func (m *TcpProxy_DeprecatedV1_TCPRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_DeprecatedV1_TCPRoute proto.InternalMessageInfo

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationIpList() []*core.CidrRange {
	if m != nil {
		return m.DestinationIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetDestinationPorts() string {
	if m != nil {
		return m.DestinationPorts
	}
	return ""
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourceIpList() []*core.CidrRange {
	if m != nil {
		return m.SourceIpList
	}
	return nil
}

func (m *TcpProxy_DeprecatedV1_TCPRoute) GetSourcePorts() string {
	if m != nil {
		return m.SourcePorts
	}
	return ""
}

type TcpProxy_WeightedCluster struct {
	Clusters             []*TcpProxy_WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *TcpProxy_WeightedCluster) Reset()         { *m = TcpProxy_WeightedCluster{} }
func (m *TcpProxy_WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cada114fb4da110, []int{0, 1}
}

func (m *TcpProxy_WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster.Size(m)
}
func (m *TcpProxy_WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster) GetClusters() []*TcpProxy_WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type TcpProxy_WeightedCluster_ClusterWeight struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               uint32         `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	MetadataMatch        *core.Metadata `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) Reset() {
	*m = TcpProxy_WeightedCluster_ClusterWeight{}
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*TcpProxy_WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*TcpProxy_WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cada114fb4da110, []int{0, 1, 0}
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.Size(m)
}
func (m *TcpProxy_WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProxy_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *TcpProxy_WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*TcpProxy)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy")
	proto.RegisterType((*TcpProxy_DeprecatedV1)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy.DeprecatedV1")
	proto.RegisterType((*TcpProxy_DeprecatedV1_TCPRoute)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy.DeprecatedV1.TCPRoute")
	proto.RegisterType((*TcpProxy_WeightedCluster)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy.WeightedCluster")
	proto.RegisterType((*TcpProxy_WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.tcp_proxy.v3alpha.TcpProxy.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/tcp_proxy/v3alpha/tcp_proxy.proto", fileDescriptor_0cada114fb4da110)
}

var fileDescriptor_0cada114fb4da110 = []byte{
	// 924 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0x27, 0x69, 0x9b, 0x9e, 0xb4, 0x6c, 0xeb, 0x65, 0xb5, 0x26, 0xa0, 0x92, 0x20, 0x90,
	0x22, 0x10, 0x36, 0x6d, 0x85, 0x84, 0xc2, 0x8f, 0x68, 0x52, 0xd4, 0x46, 0xda, 0xee, 0xa6, 0x56,
	0xb7, 0x7b, 0x69, 0x4d, 0xed, 0x49, 0x32, 0xc2, 0xf1, 0x8c, 0x66, 0xc6, 0x49, 0x2b, 0x5e, 0x80,
	0x6b, 0x2e, 0xf7, 0x39, 0x78, 0x05, 0x2e, 0x78, 0x05, 0xde, 0x80, 0x1b, 0xae, 0xd1, 0x5e, 0x21,
	0xcf, 0x8c, 0x1d, 0xb7, 0xb4, 0xdb, 0xed, 0xee, 0x5e, 0xc5, 0x3e, 0xe7, 0x7c, 0xdf, 0x77, 0x7c,
	0xfe, 0x02, 0xdf, 0xe3, 0x64, 0x46, 0x2f, 0xbc, 0x90, 0x26, 0x23, 0x32, 0xf6, 0x46, 0x24, 0x96,
	0x98, 0x7b, 0x09, 0x96, 0x73, 0xca, 0x7f, 0xf6, 0x64, 0xc8, 0x02, 0xc6, 0xe9, 0xf9, 0x85, 0x37,
	0xdb, 0x45, 0x31, 0x9b, 0xa0, 0x85, 0xc5, 0x65, 0x9c, 0x4a, 0x6a, 0x7f, 0xa9, 0xe0, 0xae, 0x86,
	0xbb, 0x1a, 0xee, 0x1a, 0xb8, 0xbb, 0x08, 0x36, 0xf0, 0xe6, 0xa7, 0x5a, 0x0d, 0x31, 0x52, 0x30,
	0x86, 0x94, 0x63, 0x0f, 0x45, 0x11, 0xc7, 0x42, 0x68, 0xd2, 0x66, 0xfb, 0x86, 0xa8, 0x33, 0x24,
	0xb0, 0x09, 0xf9, 0xfa, 0xba, 0xb4, 0x51, 0x18, 0x62, 0x21, 0x62, 0x3a, 0x2e, 0x60, 0x85, 0xc5,
	0xc0, 0x8c, 0xbe, 0xbc, 0x60, 0xb8, 0x88, 0x99, 0x20, 0x31, 0x09, 0x18, 0x8d, 0x49, 0x68, 0x3e,
	0xaa, 0xb9, 0x35, 0xa6, 0x74, 0x1c, 0x63, 0x4f, 0xbd, 0x9d, 0xa5, 0x23, 0x2f, 0x4a, 0x39, 0x92,
	0x84, 0x26, 0x37, 0xf9, 0xe7, 0x1c, 0x31, 0x86, 0x79, 0x91, 0x7f, 0x1a, 0x31, 0xe4, 0xa1, 0x24,
	0xa1, 0x52, 0xc1, 0x84, 0x37, 0xc3, 0x5c, 0x10, 0x9a, 0x90, 0x24, 0x4f, 0xe4, 0xd1, 0x0c, 0xc5,
	0x24, 0x42, 0x12, 0x7b, 0xf9, 0x83, 0x76, 0x7c, 0xf2, 0xd7, 0x7d, 0xa8, 0x9f, 0x84, 0x6c, 0x98,
	0x95, 0xcd, 0xee, 0x40, 0x43, 0x48, 0x24, 0x03, 0xc6, 0xf1, 0x88, 0x9c, 0x3b, 0x56, 0xcb, 0xea,
	0xac, 0xf6, 0x56, 0x5e, 0xf6, 0x6a, 0xbc, 0xd2, 0xb2, 0x7c, 0xc8, 0x7c, 0x43, 0xe5, 0xb2, 0x9b,
	0xb0, 0x12, 0xc6, 0xa9, 0x90, 0x98, 0x3b, 0x95, 0x2c, 0xea, 0xf0, 0x9e, 0x9f, 0x1b, 0xec, 0x19,
	0x6c, 0xce, 0x31, 0x19, 0x4f, 0x24, 0x8e, 0x02, 0x63, 0x13, 0x0e, 0xb4, 0xac, 0x4e, 0x63, 0xe7,
	0xc0, 0xbd, 0x53, 0xff, 0xdc, 0x3c, 0x33, 0xf7, 0xb9, 0x21, 0xec, 0x6b, 0xbe, 0xc3, 0x7b, 0xfe,
	0xc6, 0xfc, 0xb2, 0x49, 0xd8, 0x07, 0xf0, 0xde, 0x14, 0x4b, 0x14, 0x21, 0x89, 0x82, 0x29, 0x92,
	0xe1, 0xc4, 0x59, 0x55, 0xa2, 0x2d, 0x23, 0x8a, 0x18, 0x29, 0x88, 0xb3, 0xfe, 0xba, 0x47, 0x26,
	0xda, 0x5f, 0xcf, 0x71, 0x47, 0x19, 0xcc, 0xfe, 0x0e, 0xd6, 0x48, 0x14, 0xe3, 0x40, 0x92, 0x29,
	0xa6, 0xa9, 0x74, 0xea, 0x8a, 0xe6, 0x03, 0x57, 0xb7, 0xc1, 0xcd, 0xdb, 0xe0, 0xee, 0x9b, 0x36,
	0xf9, 0x8d, 0x2c, 0xfc, 0x44, 0x47, 0xdb, 0xc7, 0xf0, 0x28, 0xa2, 0xf3, 0x44, 0x48, 0x8e, 0xd1,
	0x34, 0xb8, 0x44, 0x54, 0xbd, 0x8d, 0xe8, 0xe1, 0x02, 0x39, 0x28, 0x51, 0x1e, 0xc1, 0xc3, 0x94,
	0x5d, 0x47, 0x58, 0xbb, 0x8d, 0xf0, 0x41, 0x8e, 0x2b, 0xd3, 0x3d, 0x05, 0xd0, 0x83, 0x1a, 0xc4,
	0x74, 0xec, 0x2c, 0xb5, 0xaa, 0x9d, 0xc6, 0xce, 0x57, 0xd7, 0x76, 0x66, 0x31, 0xcf, 0x79, 0xe1,
	0xf6, 0x94, 0xe5, 0x31, 0x1d, 0xfb, 0xab, 0x28, 0x7f, 0xb4, 0x5f, 0x58, 0xf0, 0xd9, 0x84, 0x44,
	0x11, 0x4e, 0x02, 0xc5, 0x12, 0x44, 0x98, 0x71, 0x1c, 0xa2, 0x6c, 0x02, 0x4a, 0x8f, 0xb3, 0x6d,
	0x67, 0x59, 0x25, 0xbc, 0xff, 0xa6, 0x63, 0xb0, 0x5f, 0x90, 0x9d, 0x6e, 0xf7, 0x2a, 0x8e, 0xe5,
	0xb7, 0xb5, 0xec, 0x4f, 0x19, 0xdd, 0xc2, 0x59, 0x0e, 0xb3, 0x9f, 0xc3, 0xfb, 0x53, 0x74, 0x1e,
	0x84, 0x34, 0x49, 0x70, 0x28, 0x03, 0x24, 0x25, 0x9e, 0x32, 0x29, 0x9c, 0x15, 0x95, 0xca, 0x47,
	0xff, 0xab, 0xdd, 0xb3, 0x41, 0x22, 0x77, 0x77, 0x4e, 0x51, 0x9c, 0x62, 0x35, 0xfb, 0x9f, 0x57,
	0x3a, 0x96, 0x6f, 0x4f, 0xd1, 0x79, 0x5f, 0x33, 0xec, 0x19, 0x02, 0x7b, 0x00, 0x8d, 0xd2, 0x2e,
	0x3b, 0x0d, 0x55, 0xc7, 0x2d, 0xf3, 0x69, 0xd9, 0xca, 0x17, 0xf9, 0x1f, 0x22, 0x31, 0x19, 0xaa,
	0xa8, 0x5e, 0xfd, 0x65, 0x6f, 0xe9, 0x37, 0xab, 0xb2, 0x61, 0xf9, 0x30, 0x29, 0xac, 0xcd, 0xdf,
	0x6b, 0xb0, 0x76, 0x29, 0x69, 0x0a, 0xcb, 0x9c, 0xa6, 0x12, 0x0b, 0xc7, 0x52, 0xb4, 0x47, 0xef,
	0xa2, 0x62, 0xee, 0x49, 0x7f, 0xe8, 0x67, 0xac, 0x79, 0x16, 0x75, 0xcb, 0x37, 0x32, 0xcd, 0xbf,
	0x2b, 0x50, 0xcf, 0xdd, 0x76, 0x7b, 0xb1, 0xdd, 0x57, 0x6e, 0x40, 0xb1, 0xe4, 0xc7, 0xf0, 0x20,
	0xc2, 0x42, 0x92, 0x44, 0xcd, 0x59, 0x40, 0x58, 0x10, 0x13, 0x21, 0x9d, 0x8a, 0xca, 0xb6, 0x7d,
	0xd3, 0xc6, 0xf5, 0x49, 0xc4, 0x7d, 0x94, 0x8c, 0xb1, 0xbf, 0x59, 0x42, 0x0f, 0xd8, 0x63, 0x22,
	0xa4, 0xfd, 0x05, 0x94, 0x8d, 0x01, 0xa3, 0x5c, 0x0a, 0xb5, 0x32, 0xab, 0xfe, 0x46, 0xc9, 0x31,
	0xcc, 0xec, 0xd9, 0xb2, 0x0b, 0x9a, 0xf2, 0x10, 0x17, 0xd2, 0xb5, 0xd7, 0x95, 0x5e, 0xd3, 0x40,
	0xa3, 0xda, 0x06, 0xf3, 0x6e, 0x04, 0x97, 0x94, 0x60, 0x43, 0xdb, 0x94, 0x56, 0xf7, 0xc9, 0x8b,
	0x3f, 0x7e, 0xdd, 0x1a, 0xc0, 0xc1, 0x6b, 0xb6, 0x60, 0xe7, 0x96, 0xea, 0x77, 0x0f, 0x32, 0xbe,
	0x1f, 0xe1, 0x87, 0xb7, 0xe3, 0x73, 0xac, 0xe6, 0x9f, 0x55, 0xb8, 0x7f, 0xe5, 0x32, 0xda, 0xbf,
	0x40, 0xbd, 0x38, 0xba, 0x7a, 0x76, 0x9e, 0xbd, 0xa3, 0xa3, 0xeb, 0x9a, 0x5f, 0x6d, 0x2e, 0xcd,
	0x50, 0x21, 0xd8, 0xfc, 0xc7, 0x82, 0xf5, 0x4b, 0x51, 0xf6, 0x87, 0x50, 0x4b, 0xd0, 0x14, 0x5f,
	0x9d, 0x23, 0x65, 0xb4, 0x3f, 0x86, 0x65, 0x7d, 0xc5, 0xd5, 0x9f, 0xc8, 0xfa, 0x62, 0xdd, 0x8c,
	0xf9, 0x9a, 0x93, 0x5e, 0x7d, 0xa3, 0x93, 0xde, 0x3d, 0xcd, 0x4a, 0x7e, 0x0c, 0x4f, 0xef, 0x5e,
	0xf2, 0x57, 0x16, 0xa1, 0x7b, 0x98, 0xf1, 0xf6, 0x61, 0xef, 0xad, 0x79, 0xbb, 0xdf, 0x64, 0x4c,
	0xbb, 0xb0, 0x7d, 0x67, 0xa6, 0x9e, 0x03, 0x9b, 0xa6, 0x01, 0x81, 0x60, 0x38, 0x24, 0x23, 0x82,
	0xb9, 0x5d, 0xfd, 0xb7, 0x67, 0xf5, 0x9e, 0xc0, 0xb7, 0x84, 0xea, 0x52, 0x69, 0xd8, 0x9d, 0x06,
	0xa1, 0xb7, 0x9e, 0x4b, 0x0c, 0xb3, 0xdb, 0x38, 0xb4, 0xce, 0x96, 0xd5, 0x91, 0xdc, 0xfd, 0x2f,
	0x00, 0x00, 0xff, 0xff, 0x15, 0x59, 0xc2, 0x9c, 0xc5, 0x09, 0x00, 0x00,
}
