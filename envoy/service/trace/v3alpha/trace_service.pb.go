// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/trace/v3alpha/trace_service.proto

package envoy_service_trace_v3alpha

import (
	context "context"
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type StreamTracesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamTracesResponse) Reset()         { *m = StreamTracesResponse{} }
func (m *StreamTracesResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTracesResponse) ProtoMessage()    {}
func (*StreamTracesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_838555a39f11343b, []int{0}
}

func (m *StreamTracesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesResponse.Unmarshal(m, b)
}
func (m *StreamTracesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesResponse.Marshal(b, m, deterministic)
}
func (m *StreamTracesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesResponse.Merge(m, src)
}
func (m *StreamTracesResponse) XXX_Size() int {
	return xxx_messageInfo_StreamTracesResponse.Size(m)
}
func (m *StreamTracesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesResponse proto.InternalMessageInfo

type StreamTracesMessage struct {
	Identifier           *StreamTracesMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Spans                []*v1.Span                      `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *StreamTracesMessage) Reset()         { *m = StreamTracesMessage{} }
func (m *StreamTracesMessage) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage) ProtoMessage()    {}
func (*StreamTracesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_838555a39f11343b, []int{1}
}

func (m *StreamTracesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesMessage.Unmarshal(m, b)
}
func (m *StreamTracesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesMessage.Marshal(b, m, deterministic)
}
func (m *StreamTracesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage.Merge(m, src)
}
func (m *StreamTracesMessage) XXX_Size() int {
	return xxx_messageInfo_StreamTracesMessage.Size(m)
}
func (m *StreamTracesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage proto.InternalMessageInfo

func (m *StreamTracesMessage) GetIdentifier() *StreamTracesMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamTracesMessage) GetSpans() []*v1.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type StreamTracesMessage_Identifier struct {
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamTracesMessage_Identifier) Reset()         { *m = StreamTracesMessage_Identifier{} }
func (m *StreamTracesMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage_Identifier) ProtoMessage()    {}
func (*StreamTracesMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_838555a39f11343b, []int{1, 0}
}

func (m *StreamTracesMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamTracesMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamTracesMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage_Identifier.Merge(m, src)
}
func (m *StreamTracesMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Size(m)
}
func (m *StreamTracesMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage_Identifier proto.InternalMessageInfo

func (m *StreamTracesMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamTracesResponse)(nil), "envoy.service.trace.v3alpha.StreamTracesResponse")
	proto.RegisterType((*StreamTracesMessage)(nil), "envoy.service.trace.v3alpha.StreamTracesMessage")
	proto.RegisterType((*StreamTracesMessage_Identifier)(nil), "envoy.service.trace.v3alpha.StreamTracesMessage.Identifier")
}

func init() {
	proto.RegisterFile("envoy/service/trace/v3alpha/trace_service.proto", fileDescriptor_838555a39f11343b)
}

var fileDescriptor_838555a39f11343b = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xdf, 0xea, 0xd3, 0x30,
	0x14, 0xc7, 0xcd, 0xf4, 0x27, 0x92, 0xed, 0x42, 0xab, 0xe0, 0xa8, 0x43, 0xb7, 0x81, 0x30, 0xff,
	0x25, 0xb6, 0x63, 0x08, 0xd5, 0xab, 0x7a, 0xa5, 0xa0, 0x8c, 0xce, 0x3b, 0x2f, 0x24, 0x6b, 0x8f,
	0x35, 0x30, 0x93, 0x90, 0x74, 0xd5, 0xbd, 0x80, 0x0e, 0x1f, 0xc1, 0xf7, 0xf1, 0x51, 0x7c, 0x09,
	0xaf, 0xa4, 0x4d, 0xda, 0x15, 0x1c, 0x63, 0xde, 0x25, 0xe7, 0x7c, 0x3f, 0x5f, 0xce, 0xf9, 0x1e,
	0x4c, 0x41, 0x94, 0x72, 0x47, 0x0d, 0xe8, 0x92, 0xa7, 0x40, 0x0b, 0xcd, 0x52, 0xa0, 0xe5, 0x9c,
	0x6d, 0xd4, 0x27, 0x66, 0x7f, 0x1f, 0x5c, 0x8f, 0x28, 0x2d, 0x0b, 0xe9, 0xdd, 0xa9, 0x01, 0xd2,
	0x14, 0x6b, 0x09, 0x71, 0x80, 0x3f, 0xb1, 0x6e, 0x4c, 0xf1, 0xd6, 0x23, 0x95, 0x1a, 0xe8, 0x9a,
	0x19, 0xc7, 0xfb, 0xa3, 0x5c, 0xca, 0x7c, 0x03, 0xb5, 0x86, 0x09, 0x21, 0x0b, 0x56, 0x70, 0x29,
	0x8c, 0xeb, 0xde, 0x97, 0x0a, 0x44, 0x0a, 0xc2, 0x6c, 0x0d, 0xad, 0x2b, 0xcd, 0x44, 0x81, 0x7d,
	0x38, 0xd9, 0x64, 0x9b, 0x29, 0xd6, 0xc5, 0x69, 0x09, 0xda, 0x70, 0x29, 0xb8, 0xc8, 0x9d, 0xe4,
	0x76, 0xc9, 0x36, 0x3c, 0x63, 0x05, 0xd0, 0xe6, 0x61, 0x1b, 0xd3, 0xd7, 0xf8, 0xd6, 0xaa, 0xd0,
	0xc0, 0x3e, 0xbf, 0xab, 0x0c, 0x4d, 0x02, 0x46, 0x49, 0x61, 0x20, 0x0a, 0x7f, 0xfe, 0xda, 0xdf,
	0x7d, 0x82, 0x1f, 0x1d, 0xdd, 0x2f, 0x24, 0xc7, 0x98, 0xe9, 0xef, 0x1e, 0xbe, 0xd9, 0x6d, 0xbc,
	0x01, 0x63, 0x58, 0x0e, 0xde, 0x7b, 0x8c, 0x79, 0x06, 0xa2, 0xe0, 0x1f, 0x39, 0xe8, 0x21, 0x1a,
	0xa3, 0x59, 0x3f, 0x7c, 0x4e, 0x4e, 0x24, 0x47, 0x8e, 0xb8, 0x90, 0x57, 0xad, 0x45, 0xd2, 0xb1,
	0xf3, 0x16, 0xf8, 0xc2, 0x28, 0x26, 0xcc, 0xb0, 0x37, 0xbe, 0x3c, 0xeb, 0x87, 0xf7, 0xc8, 0x21,
	0x33, 0xbb, 0x62, 0x63, 0x1d, 0x90, 0x95, 0x62, 0x22, 0xb1, 0x6a, 0xff, 0x1b, 0xc2, 0xf8, 0xe0,
	0xe8, 0x45, 0xf8, 0x8a, 0x90, 0x19, 0xb8, 0xe1, 0x46, 0x6e, 0x38, 0xa6, 0x78, 0x3b, 0x52, 0x75,
	0x39, 0xf2, 0x56, 0x66, 0x10, 0x5f, 0xfb, 0x13, 0x5f, 0xfc, 0x40, 0xbd, 0xeb, 0x28, 0xa9, 0x99,
	0xe8, 0x45, 0x15, 0xd5, 0x33, 0xbc, 0x38, 0x23, 0xaa, 0x7f, 0x77, 0x89, 0x82, 0x8a, 0x7e, 0x8c,
	0x1f, 0x9e, 0x4f, 0x87, 0xdf, 0x11, 0x1e, 0xd4, 0x95, 0x95, 0x55, 0x7b, 0x5f, 0xf0, 0xa0, 0xab,
	0xf3, 0x9e, 0xfe, 0x6f, 0xb8, 0x7e, 0x70, 0x36, 0xd1, 0x5e, 0xfb, 0xd2, 0x0c, 0xc5, 0x2f, 0xf1,
	0x03, 0x2e, 0x2d, 0xaa, 0xb4, 0xfc, 0xba, 0x3b, 0xe5, 0x12, 0xdf, 0xe8, 0xce, 0xbc, 0xac, 0x4e,
	0xb3, 0x44, 0x7b, 0x84, 0xd6, 0x57, 0xeb, 0x33, 0xcd, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x22,
	0x98, 0xfb, 0x4e, 0x7d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[0], "/envoy.service.trace.v3alpha.TraceService/StreamTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceStreamTracesClient{stream}
	return x, nil
}

type TraceService_StreamTracesClient interface {
	Send(*StreamTracesMessage) error
	CloseAndRecv() (*StreamTracesResponse, error)
	grpc.ClientStream
}

type traceServiceStreamTracesClient struct {
	grpc.ClientStream
}

func (x *traceServiceStreamTracesClient) Send(m *StreamTracesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceStreamTracesClient) CloseAndRecv() (*StreamTracesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTracesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	StreamTraces(TraceService_StreamTracesServer) error
}

// UnimplementedTraceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTraceServiceServer struct {
}

func (*UnimplementedTraceServiceServer) StreamTraces(srv TraceService_StreamTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamTraces not implemented")
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_StreamTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).StreamTraces(&traceServiceStreamTracesServer{stream})
}

type TraceService_StreamTracesServer interface {
	SendAndClose(*StreamTracesResponse) error
	Recv() (*StreamTracesMessage, error)
	grpc.ServerStream
}

type traceServiceStreamTracesServer struct {
	grpc.ServerStream
}

func (x *traceServiceStreamTracesServer) SendAndClose(m *StreamTracesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceStreamTracesServer) Recv() (*StreamTracesMessage, error) {
	m := new(StreamTracesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.trace.v3alpha.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTraces",
			Handler:       _TraceService_StreamTraces_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/trace/v3alpha/trace_service.proto",
}
