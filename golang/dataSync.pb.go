// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataSync.proto

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

type DataSyncTypesEnum int32

const (
	DataSyncTypesEnum_mysql         DataSyncTypesEnum = 0
	DataSyncTypesEnum_mssql         DataSyncTypesEnum = 1
	DataSyncTypesEnum_postgres      DataSyncTypesEnum = 2
	DataSyncTypesEnum_sqlite3       DataSyncTypesEnum = 3
	DataSyncTypesEnum_oracle        DataSyncTypesEnum = 4
	DataSyncTypesEnum_kafka         DataSyncTypesEnum = 5
	DataSyncTypesEnum_eventhub      DataSyncTypesEnum = 6
	DataSyncTypesEnum_s3            DataSyncTypesEnum = 7
	DataSyncTypesEnum_azureblob     DataSyncTypesEnum = 8
	DataSyncTypesEnum_googlestorage DataSyncTypesEnum = 9
)

var DataSyncTypesEnum_name = map[int32]string{
	0: "mysql",
	1: "mssql",
	2: "postgres",
	3: "sqlite3",
	4: "oracle",
	5: "kafka",
	6: "eventhub",
	7: "s3",
	8: "azureblob",
	9: "googlestorage",
}
var DataSyncTypesEnum_value = map[string]int32{
	"mysql":         0,
	"mssql":         1,
	"postgres":      2,
	"sqlite3":       3,
	"oracle":        4,
	"kafka":         5,
	"eventhub":      6,
	"s3":            7,
	"azureblob":     8,
	"googlestorage": 9,
}

func (x DataSyncTypesEnum) String() string {
	return proto.EnumName(DataSyncTypesEnum_name, int32(x))
}
func (DataSyncTypesEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dataSync_9aa8f5aac228fc07, []int{0}
}

// DataSyncRequest represents the request to the data-sync service via GRPC
type DataSyncRequest struct {
	AnalyzeResults       []*AnalyzeResult   `protobuf:"bytes,1,rep,name=analyzeResults,proto3" json:"analyzeResults,omitempty"`
	AnonymizeResult      *AnonymizeResponse `protobuf:"bytes,2,opt,name=anonymizeResult,proto3" json:"anonymizeResult,omitempty"`
	Path                 string             `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataSyncRequest) Reset()         { *m = DataSyncRequest{} }
func (m *DataSyncRequest) String() string { return proto.CompactTextString(m) }
func (*DataSyncRequest) ProtoMessage()    {}
func (*DataSyncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataSync_9aa8f5aac228fc07, []int{0}
}
func (m *DataSyncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSyncRequest.Unmarshal(m, b)
}
func (m *DataSyncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSyncRequest.Marshal(b, m, deterministic)
}
func (dst *DataSyncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSyncRequest.Merge(dst, src)
}
func (m *DataSyncRequest) XXX_Size() int {
	return xxx_messageInfo_DataSyncRequest.Size(m)
}
func (m *DataSyncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSyncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataSyncRequest proto.InternalMessageInfo

func (m *DataSyncRequest) GetAnalyzeResults() []*AnalyzeResult {
	if m != nil {
		return m.AnalyzeResults
	}
	return nil
}

func (m *DataSyncRequest) GetAnonymizeResult() *AnonymizeResponse {
	if m != nil {
		return m.AnonymizeResult
	}
	return nil
}

func (m *DataSyncRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// DataSyncResponse represents the response from the data-sync service via GRPC
type DataSyncResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataSyncResponse) Reset()         { *m = DataSyncResponse{} }
func (m *DataSyncResponse) String() string { return proto.CompactTextString(m) }
func (*DataSyncResponse) ProtoMessage()    {}
func (*DataSyncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataSync_9aa8f5aac228fc07, []int{1}
}
func (m *DataSyncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSyncResponse.Unmarshal(m, b)
}
func (m *DataSyncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSyncResponse.Marshal(b, m, deterministic)
}
func (dst *DataSyncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSyncResponse.Merge(dst, src)
}
func (m *DataSyncResponse) XXX_Size() int {
	return xxx_messageInfo_DataSyncResponse.Size(m)
}
func (m *DataSyncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSyncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataSyncResponse proto.InternalMessageInfo

type CompletionMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompletionMessage) Reset()         { *m = CompletionMessage{} }
func (m *CompletionMessage) String() string { return proto.CompactTextString(m) }
func (*CompletionMessage) ProtoMessage()    {}
func (*CompletionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataSync_9aa8f5aac228fc07, []int{2}
}
func (m *CompletionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompletionMessage.Unmarshal(m, b)
}
func (m *CompletionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompletionMessage.Marshal(b, m, deterministic)
}
func (dst *CompletionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompletionMessage.Merge(dst, src)
}
func (m *CompletionMessage) XXX_Size() int {
	return xxx_messageInfo_CompletionMessage.Size(m)
}
func (m *CompletionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CompletionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CompletionMessage proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DataSyncRequest)(nil), "types.DataSyncRequest")
	proto.RegisterType((*DataSyncResponse)(nil), "types.DataSyncResponse")
	proto.RegisterType((*CompletionMessage)(nil), "types.CompletionMessage")
	proto.RegisterEnum("types.DataSyncTypesEnum", DataSyncTypesEnum_name, DataSyncTypesEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataSyncServiceClient is the client API for DataSyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataSyncServiceClient interface {
	Apply(ctx context.Context, in *DataSyncRequest, opts ...grpc.CallOption) (*DataSyncResponse, error)
	Init(ctx context.Context, in *DataSyncTemplate, opts ...grpc.CallOption) (*DataSyncResponse, error)
	Completion(ctx context.Context, in *CompletionMessage, opts ...grpc.CallOption) (*DataSyncResponse, error)
}

type dataSyncServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataSyncServiceClient(cc *grpc.ClientConn) DataSyncServiceClient {
	return &dataSyncServiceClient{cc}
}

func (c *dataSyncServiceClient) Apply(ctx context.Context, in *DataSyncRequest, opts ...grpc.CallOption) (*DataSyncResponse, error) {
	out := new(DataSyncResponse)
	err := c.cc.Invoke(ctx, "/types.DataSyncService/Apply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSyncServiceClient) Init(ctx context.Context, in *DataSyncTemplate, opts ...grpc.CallOption) (*DataSyncResponse, error) {
	out := new(DataSyncResponse)
	err := c.cc.Invoke(ctx, "/types.DataSyncService/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataSyncServiceClient) Completion(ctx context.Context, in *CompletionMessage, opts ...grpc.CallOption) (*DataSyncResponse, error) {
	out := new(DataSyncResponse)
	err := c.cc.Invoke(ctx, "/types.DataSyncService/Completion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataSyncServiceServer is the server API for DataSyncService service.
type DataSyncServiceServer interface {
	Apply(context.Context, *DataSyncRequest) (*DataSyncResponse, error)
	Init(context.Context, *DataSyncTemplate) (*DataSyncResponse, error)
	Completion(context.Context, *CompletionMessage) (*DataSyncResponse, error)
}

func RegisterDataSyncServiceServer(s *grpc.Server, srv DataSyncServiceServer) {
	s.RegisterService(&_DataSyncService_serviceDesc, srv)
}

func _DataSyncService_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSyncServiceServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.DataSyncService/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSyncServiceServer).Apply(ctx, req.(*DataSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSyncService_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataSyncTemplate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSyncServiceServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.DataSyncService/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSyncServiceServer).Init(ctx, req.(*DataSyncTemplate))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataSyncService_Completion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataSyncServiceServer).Completion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.DataSyncService/Completion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataSyncServiceServer).Completion(ctx, req.(*CompletionMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataSyncService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.DataSyncService",
	HandlerType: (*DataSyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Apply",
			Handler:    _DataSyncService_Apply_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _DataSyncService_Init_Handler,
		},
		{
			MethodName: "Completion",
			Handler:    _DataSyncService_Completion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataSync.proto",
}

func init() { proto.RegisterFile("dataSync.proto", fileDescriptor_dataSync_9aa8f5aac228fc07) }

var fileDescriptor_dataSync_9aa8f5aac228fc07 = []byte{
	// 376 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x8e, 0xda, 0x30,
	0x10, 0xc6, 0x37, 0x40, 0xb2, 0x9b, 0x61, 0x97, 0x35, 0xd3, 0xaa, 0x8d, 0x38, 0x21, 0x4e, 0xa8,
	0x07, 0x0e, 0x70, 0x5b, 0xf5, 0x42, 0xff, 0x1c, 0x7a, 0xe8, 0x25, 0xcb, 0x0b, 0x38, 0xe9, 0x34,
	0x44, 0x38, 0xb6, 0x89, 0x1d, 0xa4, 0xf0, 0x1a, 0x7d, 0x88, 0xbe, 0x4c, 0x1f, 0xaa, 0x4a, 0x08,
	0x1b, 0x14, 0x24, 0x6e, 0xe3, 0x6f, 0xe6, 0x37, 0x63, 0x7f, 0x63, 0x18, 0xfd, 0xe2, 0x96, 0xbf,
	0x96, 0x32, 0x5e, 0xe8, 0x5c, 0x59, 0x85, 0xae, 0x2d, 0x35, 0x99, 0xc9, 0x63, 0xac, 0xb2, 0x4c,
	0xc9, 0x93, 0x38, 0x19, 0x59, 0xca, 0xb4, 0xe0, 0x96, 0x9a, 0xf3, 0x33, 0x97, 0x4a, 0x96, 0x59,
	0x7a, 0x6c, 0x84, 0xd9, 0x5f, 0x07, 0x9e, 0xbf, 0x35, 0x8d, 0x42, 0xda, 0x17, 0x64, 0x2c, 0x7e,
	0x86, 0x11, 0x97, 0x5c, 0x94, 0x47, 0x0a, 0xc9, 0x14, 0xc2, 0x9a, 0xc0, 0x99, 0xf6, 0xe7, 0xc3,
	0xe5, 0xfb, 0x45, 0x3d, 0x62, 0xb1, 0xbe, 0x4c, 0x86, 0x9d, 0x5a, 0xfc, 0x02, 0xed, 0x90, 0x93,
	0x16, 0xf4, 0xa6, 0xce, 0x7c, 0xb8, 0x0c, 0xde, 0xf0, 0x36, 0xab, 0x95, 0x34, 0x14, 0x76, 0x01,
	0x44, 0x18, 0x68, 0x6e, 0xb7, 0x41, 0x7f, 0xea, 0xcc, 0xfd, 0xb0, 0x8e, 0x67, 0x08, 0xac, 0xbd,
	0xe8, 0x09, 0x9c, 0xbd, 0x83, 0xf1, 0x57, 0x95, 0x69, 0x41, 0x36, 0x55, 0xf2, 0x27, 0x19, 0xc3,
	0x13, 0xfa, 0xf4, 0xc7, 0x81, 0xf1, 0xb9, 0x72, 0x53, 0x4d, 0xfc, 0x2e, 0x8b, 0x0c, 0x7d, 0x70,
	0xb3, 0xd2, 0xec, 0x05, 0xbb, 0xab, 0x43, 0x53, 0x85, 0x0e, 0x3e, 0xc2, 0x83, 0x56, 0xc6, 0x26,
	0x39, 0x19, 0xd6, 0xc3, 0x21, 0xdc, 0x9b, 0xbd, 0x48, 0x2d, 0xad, 0x58, 0x1f, 0x01, 0x3c, 0x95,
	0xf3, 0x58, 0x10, 0x1b, 0x54, 0xc4, 0x8e, 0xff, 0xde, 0x71, 0xe6, 0x56, 0x04, 0x1d, 0x48, 0xda,
	0x6d, 0x11, 0x31, 0x0f, 0x3d, 0xe8, 0x99, 0x15, 0xbb, 0xc7, 0x27, 0xf0, 0xf9, 0xb1, 0xc8, 0x29,
	0x12, 0x2a, 0x62, 0x0f, 0x38, 0x86, 0xa7, 0x44, 0xa9, 0x44, 0x90, 0xb1, 0x2a, 0xe7, 0x09, 0x31,
	0x7f, 0xf9, 0xef, 0xc2, 0xe8, 0x57, 0xca, 0x0f, 0x69, 0x4c, 0xf8, 0x02, 0xee, 0x5a, 0x6b, 0x51,
	0xe2, 0x87, 0xc6, 0x9a, 0xce, 0x26, 0x26, 0x1f, 0xaf, 0xf4, 0xe6, 0xe1, 0x77, 0xf8, 0x02, 0x83,
	0x1f, 0x32, 0xb5, 0xd8, 0x2d, 0xd9, 0x34, 0x0b, 0xbf, 0xc5, 0xae, 0x01, 0x5a, 0xdb, 0xf0, 0xbc,
	0x97, 0x2b, 0x27, 0x6f, 0xb4, 0x88, 0xbc, 0xfa, 0xfb, 0xac, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x22, 0x1c, 0xf4, 0xcd, 0x86, 0x02, 0x00, 0x00,
}
