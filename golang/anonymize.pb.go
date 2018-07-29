// Code generated by protoc-gen-go. DO NOT EDIT.
// source: anonymize.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AnonymizeApiRequest represents the request to the API HTTP service
type AnonymizeApiRequest struct {
	Text                 string             `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	AnalyzeTemplateId    string             `protobuf:"bytes,2,opt,name=analyzeTemplateId" json:"analyzeTemplateId,omitempty"`
	AnonymizeTemplateId  string             `protobuf:"bytes,3,opt,name=anonymizeTemplateId" json:"anonymizeTemplateId,omitempty"`
	AnalyzeTemplate      *AnalyzeTemplate   `protobuf:"bytes,4,opt,name=analyzeTemplate" json:"analyzeTemplate,omitempty"`
	AnonymizeTemplate    *AnonymizeTemplate `protobuf:"bytes,5,opt,name=anonymizeTemplate" json:"anonymizeTemplate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AnonymizeApiRequest) Reset()         { *m = AnonymizeApiRequest{} }
func (m *AnonymizeApiRequest) String() string { return proto.CompactTextString(m) }
func (*AnonymizeApiRequest) ProtoMessage()    {}
func (*AnonymizeApiRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_anonymize_0071d4a443a1279f, []int{0}
}
func (m *AnonymizeApiRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnonymizeApiRequest.Unmarshal(m, b)
}
func (m *AnonymizeApiRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnonymizeApiRequest.Marshal(b, m, deterministic)
}
func (dst *AnonymizeApiRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnonymizeApiRequest.Merge(dst, src)
}
func (m *AnonymizeApiRequest) XXX_Size() int {
	return xxx_messageInfo_AnonymizeApiRequest.Size(m)
}
func (m *AnonymizeApiRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnonymizeApiRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnonymizeApiRequest proto.InternalMessageInfo

func (m *AnonymizeApiRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *AnonymizeApiRequest) GetAnalyzeTemplateId() string {
	if m != nil {
		return m.AnalyzeTemplateId
	}
	return ""
}

func (m *AnonymizeApiRequest) GetAnonymizeTemplateId() string {
	if m != nil {
		return m.AnonymizeTemplateId
	}
	return ""
}

func (m *AnonymizeApiRequest) GetAnalyzeTemplate() *AnalyzeTemplate {
	if m != nil {
		return m.AnalyzeTemplate
	}
	return nil
}

func (m *AnonymizeApiRequest) GetAnonymizeTemplate() *AnonymizeTemplate {
	if m != nil {
		return m.AnonymizeTemplate
	}
	return nil
}

// AnonymizeRequest represents the request to the anonymize service via GRPC
type AnonymizeRequest struct {
	Text                 string             `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Template             *AnonymizeTemplate `protobuf:"bytes,2,opt,name=template" json:"template,omitempty"`
	AnalyzeResults       []*AnalyzeResult   `protobuf:"bytes,3,rep,name=analyzeResults" json:"analyzeResults,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AnonymizeRequest) Reset()         { *m = AnonymizeRequest{} }
func (m *AnonymizeRequest) String() string { return proto.CompactTextString(m) }
func (*AnonymizeRequest) ProtoMessage()    {}
func (*AnonymizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_anonymize_0071d4a443a1279f, []int{1}
}
func (m *AnonymizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnonymizeRequest.Unmarshal(m, b)
}
func (m *AnonymizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnonymizeRequest.Marshal(b, m, deterministic)
}
func (dst *AnonymizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnonymizeRequest.Merge(dst, src)
}
func (m *AnonymizeRequest) XXX_Size() int {
	return xxx_messageInfo_AnonymizeRequest.Size(m)
}
func (m *AnonymizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnonymizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnonymizeRequest proto.InternalMessageInfo

func (m *AnonymizeRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *AnonymizeRequest) GetTemplate() *AnonymizeTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

func (m *AnonymizeRequest) GetAnalyzeResults() []*AnalyzeResult {
	if m != nil {
		return m.AnalyzeResults
	}
	return nil
}

// AnonymizeResponse represents the anonymize service response
type AnonymizeResponse struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnonymizeResponse) Reset()         { *m = AnonymizeResponse{} }
func (m *AnonymizeResponse) String() string { return proto.CompactTextString(m) }
func (*AnonymizeResponse) ProtoMessage()    {}
func (*AnonymizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_anonymize_0071d4a443a1279f, []int{2}
}
func (m *AnonymizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnonymizeResponse.Unmarshal(m, b)
}
func (m *AnonymizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnonymizeResponse.Marshal(b, m, deterministic)
}
func (dst *AnonymizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnonymizeResponse.Merge(dst, src)
}
func (m *AnonymizeResponse) XXX_Size() int {
	return xxx_messageInfo_AnonymizeResponse.Size(m)
}
func (m *AnonymizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnonymizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnonymizeResponse proto.InternalMessageInfo

func (m *AnonymizeResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*AnonymizeApiRequest)(nil), "types.AnonymizeApiRequest")
	proto.RegisterType((*AnonymizeRequest)(nil), "types.AnonymizeRequest")
	proto.RegisterType((*AnonymizeResponse)(nil), "types.AnonymizeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AnonymizeServiceClient is the client API for AnonymizeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnonymizeServiceClient interface {
	Apply(ctx context.Context, in *AnonymizeRequest, opts ...grpc.CallOption) (*AnonymizeResponse, error)
}

type anonymizeServiceClient struct {
	cc *grpc.ClientConn
}

func NewAnonymizeServiceClient(cc *grpc.ClientConn) AnonymizeServiceClient {
	return &anonymizeServiceClient{cc}
}

func (c *anonymizeServiceClient) Apply(ctx context.Context, in *AnonymizeRequest, opts ...grpc.CallOption) (*AnonymizeResponse, error) {
	out := new(AnonymizeResponse)
	err := c.cc.Invoke(ctx, "/types.AnonymizeService/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnonymizeServiceServer is the server API for AnonymizeService service.
type AnonymizeServiceServer interface {
	Apply(context.Context, *AnonymizeRequest) (*AnonymizeResponse, error)
}

func RegisterAnonymizeServiceServer(s *grpc.Server, srv AnonymizeServiceServer) {
	s.RegisterService(&_AnonymizeService_serviceDesc, srv)
}

func _AnonymizeService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnonymizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnonymizeServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.AnonymizeService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnonymizeServiceServer).Apply(ctx, req.(*AnonymizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnonymizeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.AnonymizeService",
	HandlerType: (*AnonymizeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _AnonymizeService_Apply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "anonymize.proto",
}

func init() { proto.RegisterFile("anonymize.proto", fileDescriptor_anonymize_0071d4a443a1279f) }

var fileDescriptor_anonymize_0071d4a443a1279f = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x4f, 0x83, 0x40,
	0x10, 0x85, 0xa5, 0x14, 0xa3, 0x5b, 0x53, 0xed, 0xd4, 0xe8, 0x86, 0x13, 0xe1, 0x22, 0x07, 0x43,
	0x0c, 0x7a, 0xf4, 0x20, 0x17, 0x13, 0x6f, 0x66, 0xf5, 0x0f, 0x60, 0x9d, 0x03, 0x09, 0xec, 0xae,
	0xdd, 0xad, 0x91, 0xfe, 0x0d, 0x6f, 0xfe, 0x5a, 0xc3, 0x42, 0x09, 0x2c, 0x4d, 0x6f, 0x84, 0xf7,
	0xcd, 0x9b, 0xf7, 0x26, 0x4b, 0xce, 0x33, 0x2e, 0x78, 0x55, 0xe6, 0x5b, 0x8c, 0xe5, 0x5a, 0x68,
	0x01, 0x9e, 0xae, 0x24, 0x2a, 0xff, 0x6c, 0x25, 0xca, 0x52, 0xf0, 0xe6, 0xa7, 0x3f, 0xd7, 0x58,
	0xca, 0x22, 0xd3, 0x2d, 0x14, 0xfe, 0x4e, 0xc8, 0x32, 0xdd, 0x0d, 0xa6, 0x32, 0x67, 0xf8, 0xb5,
	0x41, 0xa5, 0x01, 0xc8, 0x54, 0xe3, 0x8f, 0xa6, 0x4e, 0xe0, 0x44, 0xa7, 0xcc, 0x7c, 0xc3, 0x2d,
	0x59, 0x64, 0x3c, 0x2b, 0xaa, 0x2d, 0xbe, 0xb7, 0x26, 0x2f, 0x9f, 0x74, 0x62, 0x80, 0xb1, 0x00,
	0x77, 0x64, 0xd9, 0x25, 0xea, 0xf1, 0xae, 0xe1, 0xf7, 0x49, 0xf0, 0x54, 0x77, 0x18, 0xd8, 0xd0,
	0x69, 0xe0, 0x44, 0xb3, 0xe4, 0x2a, 0x36, 0x55, 0xe2, 0x74, 0xa8, 0x32, 0x1b, 0x87, 0xe7, 0x3a,
	0xa1, 0x65, 0x4c, 0x3d, 0xe3, 0x41, 0x3b, 0x0f, 0x4b, 0x67, 0xe3, 0x91, 0xf0, 0xcf, 0x21, 0x17,
	0x1d, 0x78, 0xe8, 0x24, 0x0f, 0xe4, 0x64, 0x77, 0x50, 0x73, 0x89, 0x43, 0x7b, 0x3a, 0x12, 0x1e,
	0xc9, 0xbc, 0x4d, 0xce, 0x50, 0x6d, 0x0a, 0xad, 0xa8, 0x1b, 0xb8, 0xd1, 0x2c, 0xb9, 0x1c, 0xf6,
	0x6c, 0x44, 0x66, 0xb1, 0xe1, 0x0d, 0x59, 0xf4, 0xb2, 0x29, 0x29, 0xb8, 0xc2, 0x7d, 0xe1, 0x92,
	0xd7, 0x5e, 0x89, 0x37, 0x5c, 0x7f, 0xe7, 0xab, 0x7a, 0xb5, 0x97, 0x4a, 0x59, 0x54, 0x70, 0x6d,
	0xe7, 0x6c, 0x6b, 0xfa, 0x74, 0x2c, 0x34, 0x3b, 0xc2, 0xa3, 0x8f, 0x63, 0xf3, 0x68, 0xee, 0xff,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xec, 0x29, 0x2d, 0x75, 0x6c, 0x02, 0x00, 0x00,
}
