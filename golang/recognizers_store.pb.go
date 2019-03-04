// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recognizers_store.proto

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

// PatternRecognizer is the object stored in the store
type PatternRecognizer struct {
	// The name of the recognizer, unique
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Regex used by this recognizer
	Pattern string `protobuf:"bytes,2,opt,name=pattern,proto3" json:"pattern,omitempty"`
	// The confidence of the result
	Score float32 `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	// The entity name which this recognizer can detect
	Entity string `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	// The supported language code, in ISO-639 format, https://en.wikipedia.org/wiki/ISO_639-1
	Language string `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	// A list of words that are considered PII and should always be detected
	// e.g. ["Mr", "Mrs", "account"]
	Blacklist []string `protobuf:"bytes,6,rep,name=blacklist,proto3" json:"blacklist,omitempty"`
	// A list of strings that if found in the surroundings of the entity, increase the score of the result
	// e.g. ["name", "address", ]
	ContextPhrases       []string `protobuf:"bytes,7,rep,name=contextPhrases,proto3" json:"contextPhrases,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PatternRecognizer) Reset()         { *m = PatternRecognizer{} }
func (m *PatternRecognizer) String() string { return proto.CompactTextString(m) }
func (*PatternRecognizer) ProtoMessage()    {}
func (*PatternRecognizer) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{0}
}
func (m *PatternRecognizer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PatternRecognizer.Unmarshal(m, b)
}
func (m *PatternRecognizer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PatternRecognizer.Marshal(b, m, deterministic)
}
func (dst *PatternRecognizer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PatternRecognizer.Merge(dst, src)
}
func (m *PatternRecognizer) XXX_Size() int {
	return xxx_messageInfo_PatternRecognizer.Size(m)
}
func (m *PatternRecognizer) XXX_DiscardUnknown() {
	xxx_messageInfo_PatternRecognizer.DiscardUnknown(m)
}

var xxx_messageInfo_PatternRecognizer proto.InternalMessageInfo

func (m *PatternRecognizer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PatternRecognizer) GetPattern() string {
	if m != nil {
		return m.Pattern
	}
	return ""
}

func (m *PatternRecognizer) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *PatternRecognizer) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *PatternRecognizer) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *PatternRecognizer) GetBlacklist() []string {
	if m != nil {
		return m.Blacklist
	}
	return nil
}

func (m *PatternRecognizer) GetContextPhrases() []string {
	if m != nil {
		return m.ContextPhrases
	}
	return nil
}

// RecognizerInsertOrUpdateRequest represents the request for
// inserting a new or updating an existing pattern recognizer
type RecognizerInsertOrUpdateRequest struct {
	// The recognizer to be inserted or updated
	Value                *PatternRecognizer `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RecognizerInsertOrUpdateRequest) Reset()         { *m = RecognizerInsertOrUpdateRequest{} }
func (m *RecognizerInsertOrUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*RecognizerInsertOrUpdateRequest) ProtoMessage()    {}
func (*RecognizerInsertOrUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{1}
}
func (m *RecognizerInsertOrUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizerInsertOrUpdateRequest.Unmarshal(m, b)
}
func (m *RecognizerInsertOrUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizerInsertOrUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *RecognizerInsertOrUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizerInsertOrUpdateRequest.Merge(dst, src)
}
func (m *RecognizerInsertOrUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_RecognizerInsertOrUpdateRequest.Size(m)
}
func (m *RecognizerInsertOrUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizerInsertOrUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizerInsertOrUpdateRequest proto.InternalMessageInfo

func (m *RecognizerInsertOrUpdateRequest) GetValue() *PatternRecognizer {
	if m != nil {
		return m.Value
	}
	return nil
}

// RecognizerDeleteRequest represents the request for deleting an existing recognizer
type RecognizerDeleteRequest struct {
	// Name of the requested recognizer
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizerDeleteRequest) Reset()         { *m = RecognizerDeleteRequest{} }
func (m *RecognizerDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*RecognizerDeleteRequest) ProtoMessage()    {}
func (*RecognizerDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{2}
}
func (m *RecognizerDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizerDeleteRequest.Unmarshal(m, b)
}
func (m *RecognizerDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizerDeleteRequest.Marshal(b, m, deterministic)
}
func (dst *RecognizerDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizerDeleteRequest.Merge(dst, src)
}
func (m *RecognizerDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_RecognizerDeleteRequest.Size(m)
}
func (m *RecognizerDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizerDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizerDeleteRequest proto.InternalMessageInfo

func (m *RecognizerDeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// RecognizerGetRequest represents the request for retrieving a recognizer
type RecognizerGetRequest struct {
	// Name of the requested recognizer
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizerGetRequest) Reset()         { *m = RecognizerGetRequest{} }
func (m *RecognizerGetRequest) String() string { return proto.CompactTextString(m) }
func (*RecognizerGetRequest) ProtoMessage()    {}
func (*RecognizerGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{3}
}
func (m *RecognizerGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizerGetRequest.Unmarshal(m, b)
}
func (m *RecognizerGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizerGetRequest.Marshal(b, m, deterministic)
}
func (dst *RecognizerGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizerGetRequest.Merge(dst, src)
}
func (m *RecognizerGetRequest) XXX_Size() int {
	return xxx_messageInfo_RecognizerGetRequest.Size(m)
}
func (m *RecognizerGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizerGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizerGetRequest proto.InternalMessageInfo

func (m *RecognizerGetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// RecognizersGetAllRequest represents the request for retrieving all recognizers
type RecognizersGetAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizersGetAllRequest) Reset()         { *m = RecognizersGetAllRequest{} }
func (m *RecognizersGetAllRequest) String() string { return proto.CompactTextString(m) }
func (*RecognizersGetAllRequest) ProtoMessage()    {}
func (*RecognizersGetAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{4}
}
func (m *RecognizersGetAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizersGetAllRequest.Unmarshal(m, b)
}
func (m *RecognizersGetAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizersGetAllRequest.Marshal(b, m, deterministic)
}
func (dst *RecognizersGetAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizersGetAllRequest.Merge(dst, src)
}
func (m *RecognizersGetAllRequest) XXX_Size() int {
	return xxx_messageInfo_RecognizersGetAllRequest.Size(m)
}
func (m *RecognizersGetAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizersGetAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizersGetAllRequest proto.InternalMessageInfo

// RecognizerGetTimestampRequest represents the request for getting the timestamp when the recognizers store was last updated
type RecognizerGetTimestampRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizerGetTimestampRequest) Reset()         { *m = RecognizerGetTimestampRequest{} }
func (m *RecognizerGetTimestampRequest) String() string { return proto.CompactTextString(m) }
func (*RecognizerGetTimestampRequest) ProtoMessage()    {}
func (*RecognizerGetTimestampRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{5}
}
func (m *RecognizerGetTimestampRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizerGetTimestampRequest.Unmarshal(m, b)
}
func (m *RecognizerGetTimestampRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizerGetTimestampRequest.Marshal(b, m, deterministic)
}
func (dst *RecognizerGetTimestampRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizerGetTimestampRequest.Merge(dst, src)
}
func (m *RecognizerGetTimestampRequest) XXX_Size() int {
	return xxx_messageInfo_RecognizerGetTimestampRequest.Size(m)
}
func (m *RecognizerGetTimestampRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizerGetTimestampRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizerGetTimestampRequest proto.InternalMessageInfo

// RecognizersStoreResponse represents a store response
type RecognizersStoreResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizersStoreResponse) Reset()         { *m = RecognizersStoreResponse{} }
func (m *RecognizersStoreResponse) String() string { return proto.CompactTextString(m) }
func (*RecognizersStoreResponse) ProtoMessage()    {}
func (*RecognizersStoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{6}
}
func (m *RecognizersStoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizersStoreResponse.Unmarshal(m, b)
}
func (m *RecognizersStoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizersStoreResponse.Marshal(b, m, deterministic)
}
func (dst *RecognizersStoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizersStoreResponse.Merge(dst, src)
}
func (m *RecognizersStoreResponse) XXX_Size() int {
	return xxx_messageInfo_RecognizersStoreResponse.Size(m)
}
func (m *RecognizersStoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizersStoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizersStoreResponse proto.InternalMessageInfo

// RecognizersGetResponse represents a response to a get request
type RecognizersGetResponse struct {
	// An array of recognizers
	Recognizers          []*PatternRecognizer `protobuf:"bytes,1,rep,name=recognizers,proto3" json:"recognizers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RecognizersGetResponse) Reset()         { *m = RecognizersGetResponse{} }
func (m *RecognizersGetResponse) String() string { return proto.CompactTextString(m) }
func (*RecognizersGetResponse) ProtoMessage()    {}
func (*RecognizersGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{7}
}
func (m *RecognizersGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizersGetResponse.Unmarshal(m, b)
}
func (m *RecognizersGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizersGetResponse.Marshal(b, m, deterministic)
}
func (dst *RecognizersGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizersGetResponse.Merge(dst, src)
}
func (m *RecognizersGetResponse) XXX_Size() int {
	return xxx_messageInfo_RecognizersGetResponse.Size(m)
}
func (m *RecognizersGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizersGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizersGetResponse proto.InternalMessageInfo

func (m *RecognizersGetResponse) GetRecognizers() []*PatternRecognizer {
	if m != nil {
		return m.Recognizers
	}
	return nil
}

// RecognizerTimestampResponse represents the response of getting the time of last store update
type RecognizerTimestampResponse struct {
	// The time (in seconds) of the last store update
	UnixTimestamp        uint64   `protobuf:"varint,1,opt,name=unixTimestamp,proto3" json:"unixTimestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecognizerTimestampResponse) Reset()         { *m = RecognizerTimestampResponse{} }
func (m *RecognizerTimestampResponse) String() string { return proto.CompactTextString(m) }
func (*RecognizerTimestampResponse) ProtoMessage()    {}
func (*RecognizerTimestampResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_recognizers_store_c643c2e6a6efce33, []int{8}
}
func (m *RecognizerTimestampResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecognizerTimestampResponse.Unmarshal(m, b)
}
func (m *RecognizerTimestampResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecognizerTimestampResponse.Marshal(b, m, deterministic)
}
func (dst *RecognizerTimestampResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecognizerTimestampResponse.Merge(dst, src)
}
func (m *RecognizerTimestampResponse) XXX_Size() int {
	return xxx_messageInfo_RecognizerTimestampResponse.Size(m)
}
func (m *RecognizerTimestampResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecognizerTimestampResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecognizerTimestampResponse proto.InternalMessageInfo

func (m *RecognizerTimestampResponse) GetUnixTimestamp() uint64 {
	if m != nil {
		return m.UnixTimestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*PatternRecognizer)(nil), "types.PatternRecognizer")
	proto.RegisterType((*RecognizerInsertOrUpdateRequest)(nil), "types.RecognizerInsertOrUpdateRequest")
	proto.RegisterType((*RecognizerDeleteRequest)(nil), "types.RecognizerDeleteRequest")
	proto.RegisterType((*RecognizerGetRequest)(nil), "types.RecognizerGetRequest")
	proto.RegisterType((*RecognizersGetAllRequest)(nil), "types.RecognizersGetAllRequest")
	proto.RegisterType((*RecognizerGetTimestampRequest)(nil), "types.RecognizerGetTimestampRequest")
	proto.RegisterType((*RecognizersStoreResponse)(nil), "types.RecognizersStoreResponse")
	proto.RegisterType((*RecognizersGetResponse)(nil), "types.RecognizersGetResponse")
	proto.RegisterType((*RecognizerTimestampResponse)(nil), "types.RecognizerTimestampResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecognizersStoreServiceClient is the client API for RecognizersStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecognizersStoreServiceClient interface {
	ApplyInsert(ctx context.Context, in *RecognizerInsertOrUpdateRequest, opts ...grpc.CallOption) (*RecognizersStoreResponse, error)
	ApplyUpdate(ctx context.Context, in *RecognizerInsertOrUpdateRequest, opts ...grpc.CallOption) (*RecognizersStoreResponse, error)
	ApplyDelete(ctx context.Context, in *RecognizerDeleteRequest, opts ...grpc.CallOption) (*RecognizersStoreResponse, error)
	ApplyGet(ctx context.Context, in *RecognizerGetRequest, opts ...grpc.CallOption) (*RecognizersGetResponse, error)
	ApplyGetAll(ctx context.Context, in *RecognizersGetAllRequest, opts ...grpc.CallOption) (*RecognizersGetResponse, error)
	ApplyGetTimestamp(ctx context.Context, in *RecognizerGetTimestampRequest, opts ...grpc.CallOption) (*RecognizerTimestampResponse, error)
}

type recognizersStoreServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecognizersStoreServiceClient(cc *grpc.ClientConn) RecognizersStoreServiceClient {
	return &recognizersStoreServiceClient{cc}
}

func (c *recognizersStoreServiceClient) ApplyInsert(ctx context.Context, in *RecognizerInsertOrUpdateRequest, opts ...grpc.CallOption) (*RecognizersStoreResponse, error) {
	out := new(RecognizersStoreResponse)
	err := c.cc.Invoke(ctx, "/types.RecognizersStoreService/ApplyInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recognizersStoreServiceClient) ApplyUpdate(ctx context.Context, in *RecognizerInsertOrUpdateRequest, opts ...grpc.CallOption) (*RecognizersStoreResponse, error) {
	out := new(RecognizersStoreResponse)
	err := c.cc.Invoke(ctx, "/types.RecognizersStoreService/ApplyUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recognizersStoreServiceClient) ApplyDelete(ctx context.Context, in *RecognizerDeleteRequest, opts ...grpc.CallOption) (*RecognizersStoreResponse, error) {
	out := new(RecognizersStoreResponse)
	err := c.cc.Invoke(ctx, "/types.RecognizersStoreService/ApplyDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recognizersStoreServiceClient) ApplyGet(ctx context.Context, in *RecognizerGetRequest, opts ...grpc.CallOption) (*RecognizersGetResponse, error) {
	out := new(RecognizersGetResponse)
	err := c.cc.Invoke(ctx, "/types.RecognizersStoreService/ApplyGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recognizersStoreServiceClient) ApplyGetAll(ctx context.Context, in *RecognizersGetAllRequest, opts ...grpc.CallOption) (*RecognizersGetResponse, error) {
	out := new(RecognizersGetResponse)
	err := c.cc.Invoke(ctx, "/types.RecognizersStoreService/ApplyGetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recognizersStoreServiceClient) ApplyGetTimestamp(ctx context.Context, in *RecognizerGetTimestampRequest, opts ...grpc.CallOption) (*RecognizerTimestampResponse, error) {
	out := new(RecognizerTimestampResponse)
	err := c.cc.Invoke(ctx, "/types.RecognizersStoreService/ApplyGetTimestamp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecognizersStoreServiceServer is the server API for RecognizersStoreService service.
type RecognizersStoreServiceServer interface {
	ApplyInsert(context.Context, *RecognizerInsertOrUpdateRequest) (*RecognizersStoreResponse, error)
	ApplyUpdate(context.Context, *RecognizerInsertOrUpdateRequest) (*RecognizersStoreResponse, error)
	ApplyDelete(context.Context, *RecognizerDeleteRequest) (*RecognizersStoreResponse, error)
	ApplyGet(context.Context, *RecognizerGetRequest) (*RecognizersGetResponse, error)
	ApplyGetAll(context.Context, *RecognizersGetAllRequest) (*RecognizersGetResponse, error)
	ApplyGetTimestamp(context.Context, *RecognizerGetTimestampRequest) (*RecognizerTimestampResponse, error)
}

func RegisterRecognizersStoreServiceServer(s *grpc.Server, srv RecognizersStoreServiceServer) {
	s.RegisterService(&_RecognizersStoreService_serviceDesc, srv)
}

func _RecognizersStoreService_ApplyInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizerInsertOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecognizersStoreServiceServer).ApplyInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.RecognizersStoreService/ApplyInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecognizersStoreServiceServer).ApplyInsert(ctx, req.(*RecognizerInsertOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecognizersStoreService_ApplyUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizerInsertOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecognizersStoreServiceServer).ApplyUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.RecognizersStoreService/ApplyUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecognizersStoreServiceServer).ApplyUpdate(ctx, req.(*RecognizerInsertOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecognizersStoreService_ApplyDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecognizersStoreServiceServer).ApplyDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.RecognizersStoreService/ApplyDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecognizersStoreServiceServer).ApplyDelete(ctx, req.(*RecognizerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecognizersStoreService_ApplyGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizerGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecognizersStoreServiceServer).ApplyGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.RecognizersStoreService/ApplyGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecognizersStoreServiceServer).ApplyGet(ctx, req.(*RecognizerGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecognizersStoreService_ApplyGetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizersGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecognizersStoreServiceServer).ApplyGetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.RecognizersStoreService/ApplyGetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecognizersStoreServiceServer).ApplyGetAll(ctx, req.(*RecognizersGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RecognizersStoreService_ApplyGetTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecognizerGetTimestampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecognizersStoreServiceServer).ApplyGetTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.RecognizersStoreService/ApplyGetTimestamp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecognizersStoreServiceServer).ApplyGetTimestamp(ctx, req.(*RecognizerGetTimestampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecognizersStoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.RecognizersStoreService",
	HandlerType: (*RecognizersStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyInsert",
			Handler:    _RecognizersStoreService_ApplyInsert_Handler,
		},
		{
			MethodName: "ApplyUpdate",
			Handler:    _RecognizersStoreService_ApplyUpdate_Handler,
		},
		{
			MethodName: "ApplyDelete",
			Handler:    _RecognizersStoreService_ApplyDelete_Handler,
		},
		{
			MethodName: "ApplyGet",
			Handler:    _RecognizersStoreService_ApplyGet_Handler,
		},
		{
			MethodName: "ApplyGetAll",
			Handler:    _RecognizersStoreService_ApplyGetAll_Handler,
		},
		{
			MethodName: "ApplyGetTimestamp",
			Handler:    _RecognizersStoreService_ApplyGetTimestamp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "recognizers_store.proto",
}

func init() {
	proto.RegisterFile("recognizers_store.proto", fileDescriptor_recognizers_store_c643c2e6a6efce33)
}

var fileDescriptor_recognizers_store_c643c2e6a6efce33 = []byte{
	// 464 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe3, 0xe6, 0xa3, 0xcd, 0x44, 0x20, 0x75, 0x54, 0xb5, 0x2b, 0x97, 0x92, 0x68, 0x55,
	0x55, 0x11, 0x12, 0x39, 0x94, 0x1b, 0xb7, 0x0a, 0xa4, 0xc0, 0xa9, 0xc1, 0x2d, 0x12, 0xb7, 0x6a,
	0x6b, 0x46, 0xc1, 0xc2, 0x59, 0x9b, 0xdd, 0x4d, 0xd5, 0xf0, 0x2a, 0x3c, 0x15, 0x6f, 0x84, 0xd8,
	0x75, 0xbc, 0x76, 0x1c, 0x35, 0x1c, 0xb8, 0x79, 0x66, 0x7e, 0xfb, 0xd7, 0x7c, 0xfc, 0x65, 0x38,
	0x51, 0x14, 0x67, 0x73, 0x99, 0xfc, 0x24, 0xa5, 0xef, 0xb4, 0xc9, 0x14, 0x4d, 0x72, 0x95, 0x99,
	0x0c, 0xbb, 0x66, 0x95, 0x93, 0xe6, 0xbf, 0x03, 0x38, 0x9c, 0x09, 0x63, 0x48, 0xc9, 0xa8, 0x24,
	0x11, 0xa1, 0x23, 0xc5, 0x82, 0x58, 0x30, 0x0a, 0xc6, 0xfd, 0xc8, 0x7e, 0x23, 0x83, 0xfd, 0xdc,
	0x81, 0x6c, 0xcf, 0xa6, 0xd7, 0x21, 0x1e, 0x41, 0x57, 0xc7, 0x99, 0x22, 0xd6, 0x1e, 0x05, 0xe3,
	0xbd, 0xc8, 0x05, 0x78, 0x0c, 0x3d, 0x92, 0x26, 0x31, 0x2b, 0xd6, 0xb1, 0x78, 0x11, 0x61, 0x08,
	0x07, 0xa9, 0x90, 0xf3, 0xa5, 0x98, 0x13, 0xeb, 0xda, 0x4a, 0x19, 0xe3, 0x0b, 0xe8, 0xdf, 0xa7,
	0x22, 0xfe, 0x9e, 0x26, 0xda, 0xb0, 0xde, 0xa8, 0x3d, 0xee, 0x47, 0x3e, 0x81, 0x17, 0xf0, 0x3c,
	0xce, 0xa4, 0xa1, 0x47, 0x33, 0xfb, 0xa6, 0x84, 0x26, 0xcd, 0xf6, 0x2d, 0xb2, 0x91, 0xe5, 0x9f,
	0x60, 0xe8, 0x67, 0xf9, 0x28, 0x35, 0x29, 0x73, 0xad, 0x3e, 0xe7, 0x5f, 0x85, 0xa1, 0x88, 0x7e,
	0x2c, 0x49, 0x1b, 0x9c, 0x40, 0xf7, 0x41, 0xa4, 0x4b, 0x37, 0xe1, 0xe0, 0x92, 0x4d, 0xec, 0x36,
	0x26, 0x8d, 0x4d, 0x44, 0x0e, 0xe3, 0xaf, 0xe1, 0xc4, 0x27, 0xdf, 0x53, 0x4a, 0x5e, 0x6a, 0xcb,
	0xae, 0xf8, 0x2b, 0x38, 0xf2, 0xf8, 0x94, 0xcc, 0x53, 0x6c, 0x08, 0xcc, 0xb3, 0x7a, 0x4a, 0xe6,
	0x2a, 0x4d, 0x0b, 0x9e, 0x0f, 0xe1, 0xac, 0xa6, 0x73, 0x9b, 0x2c, 0x48, 0x1b, 0xb1, 0xc8, 0xd7,
	0x40, 0xfd, 0xf1, 0xcd, 0xdf, 0xfb, 0x46, 0xa4, 0xf3, 0x4c, 0x6a, 0xe2, 0xb7, 0x70, 0x5c, 0x17,
	0x5e, 0x57, 0xf0, 0x2d, 0x0c, 0x2a, 0xb6, 0x60, 0xc1, 0xa8, 0xfd, 0xe4, 0x0e, 0xaa, 0x30, 0x7f,
	0x07, 0xa7, 0xbe, 0x54, 0xe9, 0xa7, 0x90, 0x3e, 0x87, 0x67, 0x4b, 0x99, 0x3c, 0x96, 0x05, 0x3b,
	0x6a, 0x27, 0xaa, 0x27, 0x2f, 0x7f, 0x75, 0xaa, 0xfb, 0x74, 0x7d, 0xdf, 0x90, 0x7a, 0x48, 0x62,
	0xc2, 0x2f, 0x30, 0xb8, 0xca, 0xf3, 0x74, 0xe5, 0x0e, 0x87, 0x17, 0x45, 0x5b, 0x3b, 0x2e, 0x1a,
	0x0e, 0x1b, 0xdc, 0xc6, 0x3a, 0x5a, 0xa5, 0xb2, 0x7b, 0xf8, 0x3f, 0x95, 0x67, 0x85, 0xb2, 0x73,
	0x06, 0xbe, 0x6c, 0xbc, 0xa8, 0x59, 0xe6, 0x5f, 0x14, 0x3f, 0xc0, 0x81, 0x55, 0x9c, 0x92, 0xc1,
	0xd3, 0x06, 0xee, 0x2d, 0x15, 0x9e, 0x35, 0xb5, 0x2a, 0xa7, 0xe6, 0x2d, 0xbc, 0x2e, 0x7a, 0x73,
	0xce, 0xc2, 0xe1, 0x56, 0xde, 0x7b, 0x6e, 0xb7, 0xe0, 0x1d, 0x1c, 0xae, 0x05, 0xcb, 0x8b, 0xe2,
	0xf9, 0xb6, 0x1e, 0x37, 0xed, 0x1a, 0xf2, 0x06, 0xd5, 0x70, 0x10, 0x6f, 0xdd, 0xf7, 0xec, 0x1f,
	0xea, 0xcd, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x58, 0xe1, 0x28, 0xbc, 0x04, 0x00, 0x00,
}