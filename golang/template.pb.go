// Code generated by protoc-gen-go. DO NOT EDIT.
// source: template.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// AnalyzeTemplate represents the analyze service template definition
type AnalyzeTemplate struct {
	Fields               []*FieldTypes `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AnalyzeTemplate) Reset()         { *m = AnalyzeTemplate{} }
func (m *AnalyzeTemplate) String() string { return proto.CompactTextString(m) }
func (*AnalyzeTemplate) ProtoMessage()    {}
func (*AnalyzeTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{0}
}
func (m *AnalyzeTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnalyzeTemplate.Unmarshal(m, b)
}
func (m *AnalyzeTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnalyzeTemplate.Marshal(b, m, deterministic)
}
func (dst *AnalyzeTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnalyzeTemplate.Merge(dst, src)
}
func (m *AnalyzeTemplate) XXX_Size() int {
	return xxx_messageInfo_AnalyzeTemplate.Size(m)
}
func (m *AnalyzeTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_AnalyzeTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_AnalyzeTemplate proto.InternalMessageInfo

func (m *AnalyzeTemplate) GetFields() []*FieldTypes {
	if m != nil {
		return m.Fields
	}
	return nil
}

// AnonymizeTemplate represents the anonymize service template definition
type AnonymizeTemplate struct {
	Name                     string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName              string                     `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
	Description              string                     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	CreateTime               string                     `protobuf:"bytes,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	ModifiedTime             string                     `protobuf:"bytes,5,opt,name=modifiedTime,proto3" json:"modifiedTime,omitempty"`
	FieldTypeTransformations []*FieldTypeTransformation `protobuf:"bytes,6,rep,name=fieldTypeTransformations,proto3" json:"fieldTypeTransformations,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}                   `json:"-"`
	XXX_unrecognized         []byte                     `json:"-"`
	XXX_sizecache            int32                      `json:"-"`
}

func (m *AnonymizeTemplate) Reset()         { *m = AnonymizeTemplate{} }
func (m *AnonymizeTemplate) String() string { return proto.CompactTextString(m) }
func (*AnonymizeTemplate) ProtoMessage()    {}
func (*AnonymizeTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{1}
}
func (m *AnonymizeTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnonymizeTemplate.Unmarshal(m, b)
}
func (m *AnonymizeTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnonymizeTemplate.Marshal(b, m, deterministic)
}
func (dst *AnonymizeTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnonymizeTemplate.Merge(dst, src)
}
func (m *AnonymizeTemplate) XXX_Size() int {
	return xxx_messageInfo_AnonymizeTemplate.Size(m)
}
func (m *AnonymizeTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_AnonymizeTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_AnonymizeTemplate proto.InternalMessageInfo

func (m *AnonymizeTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnonymizeTemplate) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *AnonymizeTemplate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AnonymizeTemplate) GetCreateTime() string {
	if m != nil {
		return m.CreateTime
	}
	return ""
}

func (m *AnonymizeTemplate) GetModifiedTime() string {
	if m != nil {
		return m.ModifiedTime
	}
	return ""
}

func (m *AnonymizeTemplate) GetFieldTypeTransformations() []*FieldTypeTransformation {
	if m != nil {
		return m.FieldTypeTransformations
	}
	return nil
}

// FieldTypeTransformation represents the transformation for array of fields types
type FieldTypeTransformation struct {
	Fields               []*FieldTypes   `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty"`
	Transformation       *Transformation `protobuf:"bytes,2,opt,name=transformation,proto3" json:"transformation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FieldTypeTransformation) Reset()         { *m = FieldTypeTransformation{} }
func (m *FieldTypeTransformation) String() string { return proto.CompactTextString(m) }
func (*FieldTypeTransformation) ProtoMessage()    {}
func (*FieldTypeTransformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{2}
}
func (m *FieldTypeTransformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldTypeTransformation.Unmarshal(m, b)
}
func (m *FieldTypeTransformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldTypeTransformation.Marshal(b, m, deterministic)
}
func (dst *FieldTypeTransformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldTypeTransformation.Merge(dst, src)
}
func (m *FieldTypeTransformation) XXX_Size() int {
	return xxx_messageInfo_FieldTypeTransformation.Size(m)
}
func (m *FieldTypeTransformation) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldTypeTransformation.DiscardUnknown(m)
}

var xxx_messageInfo_FieldTypeTransformation proto.InternalMessageInfo

func (m *FieldTypeTransformation) GetFields() []*FieldTypes {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *FieldTypeTransformation) GetTransformation() *Transformation {
	if m != nil {
		return m.Transformation
	}
	return nil
}

// Transformation represents the transformation type
type Transformation struct {
	NewValue             string        `protobuf:"bytes,1,opt,name=newValue,proto3" json:"newValue,omitempty"`
	ReplaceValue         *ReplaceValue `protobuf:"bytes,2,opt,name=replaceValue,proto3" json:"replaceValue,omitempty"`
	RedactValue          *RedactValue  `protobuf:"bytes,3,opt,name=redactValue,proto3" json:"redactValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Transformation) Reset()         { *m = Transformation{} }
func (m *Transformation) String() string { return proto.CompactTextString(m) }
func (*Transformation) ProtoMessage()    {}
func (*Transformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{3}
}
func (m *Transformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transformation.Unmarshal(m, b)
}
func (m *Transformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transformation.Marshal(b, m, deterministic)
}
func (dst *Transformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transformation.Merge(dst, src)
}
func (m *Transformation) XXX_Size() int {
	return xxx_messageInfo_Transformation.Size(m)
}
func (m *Transformation) XXX_DiscardUnknown() {
	xxx_messageInfo_Transformation.DiscardUnknown(m)
}

var xxx_messageInfo_Transformation proto.InternalMessageInfo

func (m *Transformation) GetNewValue() string {
	if m != nil {
		return m.NewValue
	}
	return ""
}

func (m *Transformation) GetReplaceValue() *ReplaceValue {
	if m != nil {
		return m.ReplaceValue
	}
	return nil
}

func (m *Transformation) GetRedactValue() *RedactValue {
	if m != nil {
		return m.RedactValue
	}
	return nil
}

type ReplaceValue struct {
	NewValue             string   `protobuf:"bytes,1,opt,name=newValue,proto3" json:"newValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplaceValue) Reset()         { *m = ReplaceValue{} }
func (m *ReplaceValue) String() string { return proto.CompactTextString(m) }
func (*ReplaceValue) ProtoMessage()    {}
func (*ReplaceValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{4}
}
func (m *ReplaceValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplaceValue.Unmarshal(m, b)
}
func (m *ReplaceValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplaceValue.Marshal(b, m, deterministic)
}
func (dst *ReplaceValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplaceValue.Merge(dst, src)
}
func (m *ReplaceValue) XXX_Size() int {
	return xxx_messageInfo_ReplaceValue.Size(m)
}
func (m *ReplaceValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplaceValue.DiscardUnknown(m)
}

var xxx_messageInfo_ReplaceValue proto.InternalMessageInfo

func (m *ReplaceValue) GetNewValue() string {
	if m != nil {
		return m.NewValue
	}
	return ""
}

type RedactValue struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedactValue) Reset()         { *m = RedactValue{} }
func (m *RedactValue) String() string { return proto.CompactTextString(m) }
func (*RedactValue) ProtoMessage()    {}
func (*RedactValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{5}
}
func (m *RedactValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedactValue.Unmarshal(m, b)
}
func (m *RedactValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedactValue.Marshal(b, m, deterministic)
}
func (dst *RedactValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedactValue.Merge(dst, src)
}
func (m *RedactValue) XXX_Size() int {
	return xxx_messageInfo_RedactValue.Size(m)
}
func (m *RedactValue) XXX_DiscardUnknown() {
	xxx_messageInfo_RedactValue.DiscardUnknown(m)
}

var xxx_messageInfo_RedactValue proto.InternalMessageInfo

type Databinder struct {
	BindType             string   `protobuf:"bytes,1,opt,name=bindType,proto3" json:"bindType,omitempty"`
	ConnectionString     string   `protobuf:"bytes,2,opt,name=connectionString,proto3" json:"connectionString,omitempty"`
	TableName            string   `protobuf:"bytes,3,opt,name=tableName,proto3" json:"tableName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Databinder) Reset()         { *m = Databinder{} }
func (m *Databinder) String() string { return proto.CompactTextString(m) }
func (*Databinder) ProtoMessage()    {}
func (*Databinder) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{6}
}
func (m *Databinder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Databinder.Unmarshal(m, b)
}
func (m *Databinder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Databinder.Marshal(b, m, deterministic)
}
func (dst *Databinder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Databinder.Merge(dst, src)
}
func (m *Databinder) XXX_Size() int {
	return xxx_messageInfo_Databinder.Size(m)
}
func (m *Databinder) XXX_DiscardUnknown() {
	xxx_messageInfo_Databinder.DiscardUnknown(m)
}

var xxx_messageInfo_Databinder proto.InternalMessageInfo

func (m *Databinder) GetBindType() string {
	if m != nil {
		return m.BindType
	}
	return ""
}

func (m *Databinder) GetConnectionString() string {
	if m != nil {
		return m.ConnectionString
	}
	return ""
}

func (m *Databinder) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

// DatabinderTemplate represents the analyzer service outputs definition
type DatabinderTemplate struct {
	Databinder           []*Databinder `protobuf:"bytes,1,rep,name=databinder,proto3" json:"databinder,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *DatabinderTemplate) Reset()         { *m = DatabinderTemplate{} }
func (m *DatabinderTemplate) String() string { return proto.CompactTextString(m) }
func (*DatabinderTemplate) ProtoMessage()    {}
func (*DatabinderTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{7}
}
func (m *DatabinderTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DatabinderTemplate.Unmarshal(m, b)
}
func (m *DatabinderTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DatabinderTemplate.Marshal(b, m, deterministic)
}
func (dst *DatabinderTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DatabinderTemplate.Merge(dst, src)
}
func (m *DatabinderTemplate) XXX_Size() int {
	return xxx_messageInfo_DatabinderTemplate.Size(m)
}
func (m *DatabinderTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_DatabinderTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_DatabinderTemplate proto.InternalMessageInfo

func (m *DatabinderTemplate) GetDatabinder() []*Databinder {
	if m != nil {
		return m.Databinder
	}
	return nil
}

type BlobStorageConfig struct {
	AccountName          string   `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	AccountKey           string   `protobuf:"bytes,2,opt,name=accountKey,proto3" json:"accountKey,omitempty"`
	ContainerName        string   `protobuf:"bytes,3,opt,name=containerName,proto3" json:"containerName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlobStorageConfig) Reset()         { *m = BlobStorageConfig{} }
func (m *BlobStorageConfig) String() string { return proto.CompactTextString(m) }
func (*BlobStorageConfig) ProtoMessage()    {}
func (*BlobStorageConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{8}
}
func (m *BlobStorageConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlobStorageConfig.Unmarshal(m, b)
}
func (m *BlobStorageConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlobStorageConfig.Marshal(b, m, deterministic)
}
func (dst *BlobStorageConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlobStorageConfig.Merge(dst, src)
}
func (m *BlobStorageConfig) XXX_Size() int {
	return xxx_messageInfo_BlobStorageConfig.Size(m)
}
func (m *BlobStorageConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_BlobStorageConfig.DiscardUnknown(m)
}

var xxx_messageInfo_BlobStorageConfig proto.InternalMessageInfo

func (m *BlobStorageConfig) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *BlobStorageConfig) GetAccountKey() string {
	if m != nil {
		return m.AccountKey
	}
	return ""
}

func (m *BlobStorageConfig) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

type S3Config struct {
	AccessId             string   `protobuf:"bytes,1,opt,name=accessId,proto3" json:"accessId,omitempty"`
	AccessKey            string   `protobuf:"bytes,2,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	Region               string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	BucketName           string   `protobuf:"bytes,4,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S3Config) Reset()         { *m = S3Config{} }
func (m *S3Config) String() string { return proto.CompactTextString(m) }
func (*S3Config) ProtoMessage()    {}
func (*S3Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{9}
}
func (m *S3Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3Config.Unmarshal(m, b)
}
func (m *S3Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3Config.Marshal(b, m, deterministic)
}
func (dst *S3Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3Config.Merge(dst, src)
}
func (m *S3Config) XXX_Size() int {
	return xxx_messageInfo_S3Config.Size(m)
}
func (m *S3Config) XXX_DiscardUnknown() {
	xxx_messageInfo_S3Config.DiscardUnknown(m)
}

var xxx_messageInfo_S3Config proto.InternalMessageInfo

func (m *S3Config) GetAccessId() string {
	if m != nil {
		return m.AccessId
	}
	return ""
}

func (m *S3Config) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *S3Config) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *S3Config) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

type InputConfig struct {
	BlobStorageConfig    *BlobStorageConfig `protobuf:"bytes,1,opt,name=blobStorageConfig,proto3" json:"blobStorageConfig,omitempty"`
	S3Config             *S3Config          `protobuf:"bytes,2,opt,name=s3Config,proto3" json:"s3Config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *InputConfig) Reset()         { *m = InputConfig{} }
func (m *InputConfig) String() string { return proto.CompactTextString(m) }
func (*InputConfig) ProtoMessage()    {}
func (*InputConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{10}
}
func (m *InputConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputConfig.Unmarshal(m, b)
}
func (m *InputConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputConfig.Marshal(b, m, deterministic)
}
func (dst *InputConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputConfig.Merge(dst, src)
}
func (m *InputConfig) XXX_Size() int {
	return xxx_messageInfo_InputConfig.Size(m)
}
func (m *InputConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_InputConfig.DiscardUnknown(m)
}

var xxx_messageInfo_InputConfig proto.InternalMessageInfo

func (m *InputConfig) GetBlobStorageConfig() *BlobStorageConfig {
	if m != nil {
		return m.BlobStorageConfig
	}
	return nil
}

func (m *InputConfig) GetS3Config() *S3Config {
	if m != nil {
		return m.S3Config
	}
	return nil
}

type ScanTemplate struct {
	Kind                 string       `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	InputConfig          *InputConfig `protobuf:"bytes,2,opt,name=inputConfig,proto3" json:"inputConfig,omitempty"`
	MinProbability       string       `protobuf:"bytes,3,opt,name=minProbability,proto3" json:"minProbability,omitempty"`
	AnalyzeTemplateId    string       `protobuf:"bytes,4,opt,name=analyzeTemplateId,proto3" json:"analyzeTemplateId,omitempty"`
	AnonymizeTemplateId  string       `protobuf:"bytes,5,opt,name=anonymizeTemplateId,proto3" json:"anonymizeTemplateId,omitempty"`
	DatabinderTemplateId string       `protobuf:"bytes,6,opt,name=databinderTemplateId,proto3" json:"databinderTemplateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ScanTemplate) Reset()         { *m = ScanTemplate{} }
func (m *ScanTemplate) String() string { return proto.CompactTextString(m) }
func (*ScanTemplate) ProtoMessage()    {}
func (*ScanTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{11}
}
func (m *ScanTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanTemplate.Unmarshal(m, b)
}
func (m *ScanTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanTemplate.Marshal(b, m, deterministic)
}
func (dst *ScanTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanTemplate.Merge(dst, src)
}
func (m *ScanTemplate) XXX_Size() int {
	return xxx_messageInfo_ScanTemplate.Size(m)
}
func (m *ScanTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_ScanTemplate proto.InternalMessageInfo

func (m *ScanTemplate) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *ScanTemplate) GetInputConfig() *InputConfig {
	if m != nil {
		return m.InputConfig
	}
	return nil
}

func (m *ScanTemplate) GetMinProbability() string {
	if m != nil {
		return m.MinProbability
	}
	return ""
}

func (m *ScanTemplate) GetAnalyzeTemplateId() string {
	if m != nil {
		return m.AnalyzeTemplateId
	}
	return ""
}

func (m *ScanTemplate) GetAnonymizeTemplateId() string {
	if m != nil {
		return m.AnonymizeTemplateId
	}
	return ""
}

func (m *ScanTemplate) GetDatabinderTemplateId() string {
	if m != nil {
		return m.DatabinderTemplateId
	}
	return ""
}

type JobTemplate struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Trigger              *Trigger `protobuf:"bytes,3,opt,name=trigger,proto3" json:"trigger,omitempty"`
	ScanTemplateId       string   `protobuf:"bytes,6,opt,name=scanTemplateId,proto3" json:"scanTemplateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobTemplate) Reset()         { *m = JobTemplate{} }
func (m *JobTemplate) String() string { return proto.CompactTextString(m) }
func (*JobTemplate) ProtoMessage()    {}
func (*JobTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{12}
}
func (m *JobTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobTemplate.Unmarshal(m, b)
}
func (m *JobTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobTemplate.Marshal(b, m, deterministic)
}
func (dst *JobTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobTemplate.Merge(dst, src)
}
func (m *JobTemplate) XXX_Size() int {
	return xxx_messageInfo_JobTemplate.Size(m)
}
func (m *JobTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_JobTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_JobTemplate proto.InternalMessageInfo

func (m *JobTemplate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JobTemplate) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *JobTemplate) GetTrigger() *Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *JobTemplate) GetScanTemplateId() string {
	if m != nil {
		return m.ScanTemplateId
	}
	return ""
}

type Trigger struct {
	Schedule             *Schedule `protobuf:"bytes,1,opt,name=schedule,proto3" json:"schedule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{13}
}
func (m *Trigger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trigger.Unmarshal(m, b)
}
func (m *Trigger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trigger.Marshal(b, m, deterministic)
}
func (dst *Trigger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trigger.Merge(dst, src)
}
func (m *Trigger) XXX_Size() int {
	return xxx_messageInfo_Trigger.Size(m)
}
func (m *Trigger) XXX_DiscardUnknown() {
	xxx_messageInfo_Trigger.DiscardUnknown(m)
}

var xxx_messageInfo_Trigger proto.InternalMessageInfo

func (m *Trigger) GetSchedule() *Schedule {
	if m != nil {
		return m.Schedule
	}
	return nil
}

type Schedule struct {
	RecurrencePeriodDuration string   `protobuf:"bytes,1,opt,name=recurrencePeriodDuration,proto3" json:"recurrencePeriodDuration,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_template_506a62ccf9ffbdeb, []int{14}
}
func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule.Unmarshal(m, b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
}
func (dst *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(dst, src)
}
func (m *Schedule) XXX_Size() int {
	return xxx_messageInfo_Schedule.Size(m)
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

func (m *Schedule) GetRecurrencePeriodDuration() string {
	if m != nil {
		return m.RecurrencePeriodDuration
	}
	return ""
}

func init() {
	proto.RegisterType((*AnalyzeTemplate)(nil), "types.AnalyzeTemplate")
	proto.RegisterType((*AnonymizeTemplate)(nil), "types.AnonymizeTemplate")
	proto.RegisterType((*FieldTypeTransformation)(nil), "types.FieldTypeTransformation")
	proto.RegisterType((*Transformation)(nil), "types.Transformation")
	proto.RegisterType((*ReplaceValue)(nil), "types.ReplaceValue")
	proto.RegisterType((*RedactValue)(nil), "types.RedactValue")
	proto.RegisterType((*Databinder)(nil), "types.Databinder")
	proto.RegisterType((*DatabinderTemplate)(nil), "types.DatabinderTemplate")
	proto.RegisterType((*BlobStorageConfig)(nil), "types.BlobStorageConfig")
	proto.RegisterType((*S3Config)(nil), "types.S3Config")
	proto.RegisterType((*InputConfig)(nil), "types.InputConfig")
	proto.RegisterType((*ScanTemplate)(nil), "types.ScanTemplate")
	proto.RegisterType((*JobTemplate)(nil), "types.JobTemplate")
	proto.RegisterType((*Trigger)(nil), "types.Trigger")
	proto.RegisterType((*Schedule)(nil), "types.Schedule")
}

func init() { proto.RegisterFile("template.proto", fileDescriptor_template_506a62ccf9ffbdeb) }

var fileDescriptor_template_506a62ccf9ffbdeb = []byte{
	// 737 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6f, 0xe3, 0x36,
	0x10, 0x85, 0x9d, 0xc4, 0x71, 0x46, 0x8e, 0x53, 0x33, 0xfd, 0x10, 0x82, 0x22, 0x08, 0x88, 0xa2,
	0x48, 0xd3, 0x22, 0x68, 0x9d, 0xa2, 0x05, 0x8a, 0xf6, 0x90, 0x36, 0x48, 0xe1, 0x16, 0x08, 0x02,
	0xd9, 0xe8, 0xa1, 0x37, 0x8a, 0x1a, 0xbb, 0x44, 0x24, 0xd2, 0xa0, 0x68, 0x14, 0x6e, 0xf7, 0xb4,
	0x7b, 0xdc, 0xeb, 0xee, 0x69, 0xff, 0xec, 0x42, 0x14, 0x25, 0xd1, 0x72, 0xb2, 0x8b, 0xbd, 0x91,
	0x6f, 0xde, 0x23, 0x67, 0x1e, 0x47, 0x23, 0x18, 0x1a, 0xcc, 0x96, 0x29, 0x33, 0x78, 0xb9, 0xd4,
	0xca, 0x28, 0xb2, 0x67, 0xd6, 0x4b, 0xcc, 0x4f, 0x06, 0x5c, 0x65, 0x99, 0x92, 0x25, 0x48, 0x7f,
	0x86, 0xa3, 0x6b, 0xc9, 0xd2, 0xf5, 0x7f, 0x38, 0x73, 0x6c, 0xf2, 0x15, 0xf4, 0xe6, 0x02, 0xd3,
	0x24, 0x0f, 0x3b, 0x67, 0x3b, 0xe7, 0xc1, 0x78, 0x74, 0x69, 0x85, 0x97, 0xb7, 0x05, 0x38, 0x2b,
	0x96, 0x91, 0x23, 0xd0, 0x97, 0x5d, 0x18, 0x5d, 0x4b, 0x25, 0xd7, 0x99, 0xf0, 0x0e, 0x20, 0xb0,
	0x2b, 0x59, 0x86, 0x61, 0xe7, 0xac, 0x73, 0x7e, 0x10, 0xd9, 0x35, 0x39, 0x83, 0x20, 0x11, 0xf9,
	0x32, 0x65, 0xeb, 0xbb, 0x22, 0xd4, 0xb5, 0x21, 0x1f, 0xb2, 0x0c, 0xcc, 0xb9, 0x16, 0x4b, 0x23,
	0x94, 0x0c, 0x77, 0x1c, 0xa3, 0x81, 0xc8, 0x29, 0x00, 0xd7, 0xc8, 0x0c, 0xce, 0x44, 0x86, 0xe1,
	0xae, 0x25, 0x78, 0x08, 0xa1, 0x30, 0xc8, 0x54, 0x22, 0xe6, 0x02, 0x13, 0xcb, 0xd8, 0xb3, 0x8c,
	0x0d, 0x8c, 0xfc, 0x0d, 0xe1, 0xbc, 0xaa, 0x63, 0xa6, 0x99, 0xcc, 0xe7, 0x4a, 0x67, 0xac, 0x38,
	0x3e, 0x0f, 0x7b, 0xb6, 0xdc, 0xd3, 0x76, 0xb9, 0x9b, 0xb4, 0xe8, 0x49, 0x3d, 0x7d, 0xd1, 0x81,
	0xcf, 0x9e, 0x50, 0x7d, 0x80, 0xa9, 0xe4, 0x17, 0x18, 0x9a, 0x0d, 0xb1, 0x75, 0x2b, 0x18, 0x7f,
	0xe2, 0x24, 0xad, 0x7c, 0x5a, 0x64, 0xfa, 0xa6, 0x03, 0xc3, 0xd6, 0xe5, 0x27, 0xd0, 0x97, 0xf8,
	0xef, 0x5f, 0x2c, 0x5d, 0x55, 0x8f, 0x52, 0xef, 0xc9, 0x8f, 0x30, 0xd0, 0xb8, 0x4c, 0x19, 0xc7,
	0x32, 0x5e, 0xde, 0x75, 0xec, 0xee, 0x8a, 0xbc, 0x50, 0xb4, 0x41, 0x24, 0xdf, 0x43, 0xa0, 0x31,
	0x61, 0xdc, 0x94, 0xba, 0x1d, 0xab, 0x23, 0xb5, 0xae, 0x8e, 0x44, 0x3e, 0x8d, 0x5e, 0xc0, 0xc0,
	0x3f, 0xf3, 0x5d, 0xa9, 0xd1, 0x43, 0x08, 0xbc, 0x73, 0xa8, 0x06, 0xb8, 0x61, 0x86, 0xc5, 0x42,
	0x26, 0xa8, 0x0b, 0x61, 0xb1, 0x2a, 0xac, 0xab, 0x84, 0xd5, 0x9e, 0x5c, 0xc0, 0x47, 0x5c, 0x49,
	0x89, 0xbc, 0xa8, 0x7e, 0x6a, 0xb4, 0x90, 0x0b, 0xd7, 0x71, 0x5b, 0x38, 0xf9, 0x1c, 0x0e, 0x0c,
	0x8b, 0x53, 0xb4, 0x6d, 0x59, 0x36, 0x5d, 0x03, 0xd0, 0xdf, 0x81, 0x34, 0x77, 0xd6, 0x0d, 0xfe,
	0x1d, 0x40, 0x52, 0xa3, 0xad, 0x07, 0x6d, 0xe8, 0x91, 0x47, 0xa2, 0xff, 0xc3, 0xe8, 0xd7, 0x54,
	0xc5, 0x53, 0xa3, 0x34, 0x5b, 0xe0, 0x6f, 0x4a, 0xce, 0xc5, 0xa2, 0x68, 0x79, 0xc6, 0xb9, 0x5a,
	0x49, 0x73, 0xd7, 0x7c, 0x2f, 0x3e, 0x54, 0xb4, 0xbc, 0xdb, 0xfe, 0x89, 0x6b, 0x57, 0x83, 0x87,
	0x90, 0x2f, 0xe0, 0x90, 0x2b, 0x69, 0x98, 0x90, 0xa8, 0xbd, 0x0a, 0x36, 0x41, 0xfa, 0x0c, 0xfa,
	0xd3, 0x2b, 0x77, 0xe7, 0x09, 0xf4, 0x19, 0xe7, 0x98, 0xe7, 0x93, 0xa4, 0xf2, 0xad, 0xda, 0x17,
	0x5e, 0x94, 0xeb, 0xe6, 0xb2, 0x06, 0x20, 0x9f, 0x42, 0x4f, 0xe3, 0xa2, 0xf9, 0x36, 0xdd, 0xae,
	0xc8, 0x31, 0x5e, 0xf1, 0x07, 0x2c, 0x8b, 0x70, 0x9f, 0x65, 0x83, 0xd0, 0xe7, 0x1d, 0x08, 0x26,
	0x72, 0xb9, 0x32, 0x2e, 0x83, 0x5b, 0x18, 0xc5, 0x6d, 0x2b, 0x6c, 0x2a, 0xc1, 0x38, 0x74, 0x26,
	0x6e, 0x59, 0x15, 0x6d, 0x4b, 0xc8, 0xd7, 0xd0, 0xcf, 0x5d, 0x55, 0xae, 0x6b, 0x8f, 0x9c, 0xbc,
	0x2a, 0x36, 0xaa, 0x09, 0xf4, 0x55, 0x17, 0x06, 0x53, 0xce, 0xa4, 0x3f, 0xa4, 0x1e, 0x84, 0xac,
	0x3c, 0xb0, 0xeb, 0xa2, 0xa5, 0x45, 0x93, 0xa8, 0x3b, 0xb4, 0x6a, 0x69, 0xaf, 0x84, 0xc8, 0xa7,
	0x91, 0x2f, 0x61, 0x98, 0x09, 0x79, 0xaf, 0x55, 0xcc, 0x62, 0x91, 0x0a, 0xb3, 0x76, 0xfe, 0xb4,
	0x50, 0xf2, 0x0d, 0x8c, 0xd8, 0xe6, 0xa8, 0x9d, 0x24, 0xce, 0xae, 0xed, 0x00, 0xf9, 0x16, 0x8e,
	0x59, 0x7b, 0xb2, 0x4e, 0x12, 0x37, 0xd3, 0x1e, 0x0b, 0x91, 0x31, 0x7c, 0x9c, 0x6c, 0xf5, 0xea,
	0x24, 0x09, 0x7b, 0x56, 0xf2, 0x68, 0x8c, 0xbe, 0xee, 0x40, 0xf0, 0x87, 0x8a, 0xdf, 0x3b, 0xba,
	0xbd, 0xc1, 0xdc, 0xdd, 0x1e, 0xcc, 0xe7, 0xb0, 0x6f, 0xb4, 0x58, 0x2c, 0x50, 0xbb, 0x31, 0x30,
	0xac, 0x47, 0x95, 0x45, 0xa3, 0x2a, 0x5c, 0x78, 0x95, 0x7b, 0xaf, 0x50, 0x67, 0xd7, 0x42, 0xe9,
	0x0f, 0xb0, 0xef, 0xb4, 0xf6, 0x99, 0xf9, 0x3f, 0x98, 0xac, 0x52, 0x74, 0x5d, 0x52, 0x3f, 0xb3,
	0x83, 0xa3, 0x9a, 0x40, 0x6f, 0xa1, 0x5f, 0xa1, 0xe4, 0x27, 0x08, 0x35, 0xf2, 0x95, 0xd6, 0x28,
	0x39, 0xde, 0xa3, 0x16, 0x2a, 0xb9, 0x59, 0xe9, 0x72, 0xa2, 0x96, 0xf5, 0x3d, 0x19, 0x8f, 0x7b,
	0xf6, 0xef, 0x78, 0xf5, 0x36, 0x00, 0x00, 0xff, 0xff, 0x26, 0x3f, 0xdc, 0xf9, 0x44, 0x07, 0x00,
	0x00,
}
