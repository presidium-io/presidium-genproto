// Code generated by protoc-gen-go. DO NOT EDIT.
// source: job.proto

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

type Job struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Trigger              *Trigger `protobuf:"bytes,3,opt,name=trigger" json:"trigger,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_d286af0f3f3ad29e, []int{0}
}
func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (dst *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(dst, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Job) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Job) GetTrigger() *Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

type Trigger struct {
	Schedule             *Schedule `protobuf:"bytes,1,opt,name=schedule" json:"schedule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Trigger) Reset()         { *m = Trigger{} }
func (m *Trigger) String() string { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()    {}
func (*Trigger) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_d286af0f3f3ad29e, []int{1}
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
	RecurrencePeriodDuration string   `protobuf:"bytes,1,opt,name=recurrencePeriodDuration" json:"recurrencePeriodDuration,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_job_d286af0f3f3ad29e, []int{2}
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
	proto.RegisterType((*Job)(nil), "types.Job")
	proto.RegisterType((*Trigger)(nil), "types.Trigger")
	proto.RegisterType((*Schedule)(nil), "types.Schedule")
}

func init() { proto.RegisterFile("job.proto", fileDescriptor_job_d286af0f3f3ad29e) }

var fileDescriptor_job_d286af0f3f3ad29e = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0xca, 0x4f, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x4a, 0xe5, 0x62,
	0xf6, 0xca, 0x4f, 0x12, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a,
	0x32, 0xf3, 0xf3, 0x24, 0x98, 0xc0, 0x52, 0xc8, 0x42, 0x42, 0x1a, 0x5c, 0xec, 0x25, 0x45, 0x99,
	0xe9, 0xe9, 0xa9, 0x45, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x7c, 0x7a, 0x60, 0x53, 0xf5,
	0x42, 0x20, 0xa2, 0x41, 0x30, 0x69, 0x25, 0x33, 0x2e, 0x76, 0xa8, 0x98, 0x90, 0x36, 0x17, 0x47,
	0x71, 0x72, 0x46, 0x6a, 0x4a, 0x69, 0x0e, 0xc4, 0x3a, 0x6e, 0x23, 0x7e, 0xa8, 0xae, 0x60, 0xa8,
	0x70, 0x10, 0x5c, 0x81, 0x92, 0x1b, 0x17, 0x07, 0x4c, 0x54, 0xc8, 0x8a, 0x4b, 0xa2, 0x28, 0x35,
	0xb9, 0xb4, 0xa8, 0x28, 0x35, 0x2f, 0x39, 0x35, 0x20, 0xb5, 0x28, 0x33, 0x3f, 0xc5, 0xa5, 0xb4,
	0x28, 0x11, 0xec, 0x38, 0x88, 0xbb, 0x71, 0xca, 0x27, 0xb1, 0x81, 0x3d, 0x6d, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x65, 0x10, 0xa0, 0x47, 0x01, 0x01, 0x00, 0x00,
}