// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contract.proto

package contract

import (
	fmt "fmt"
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

type ResponseStatus int32

const (
	ResponseStatus_FAIL ResponseStatus = 0
	ResponseStatus_OK   ResponseStatus = 1
)

var ResponseStatus_name = map[int32]string{
	0: "FAIL",
	1: "OK",
}

var ResponseStatus_value = map[string]int32{
	"FAIL": 0,
	"OK":   1,
}

func (x ResponseStatus) String() string {
	return proto.EnumName(ResponseStatus_name, int32(x))
}

func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{0}
}

type AddNewsRequest struct {
	Status               ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=contract.ResponseStatus" json:"status,omitempty"`
	Title                string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Date                 string         `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddNewsRequest) Reset()         { *m = AddNewsRequest{} }
func (m *AddNewsRequest) String() string { return proto.CompactTextString(m) }
func (*AddNewsRequest) ProtoMessage()    {}
func (*AddNewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{0}
}

func (m *AddNewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddNewsRequest.Unmarshal(m, b)
}
func (m *AddNewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddNewsRequest.Marshal(b, m, deterministic)
}
func (m *AddNewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddNewsRequest.Merge(m, src)
}
func (m *AddNewsRequest) XXX_Size() int {
	return xxx_messageInfo_AddNewsRequest.Size(m)
}
func (m *AddNewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddNewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddNewsRequest proto.InternalMessageInfo

func (m *AddNewsRequest) GetStatus() ResponseStatus {
	if m != nil {
		return m.Status
	}
	return ResponseStatus_FAIL
}

func (m *AddNewsRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddNewsRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type OnAddNewsResponse struct {
	Status               ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=contract.ResponseStatus" json:"status,omitempty"`
	Id                   uint32         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OnAddNewsResponse) Reset()         { *m = OnAddNewsResponse{} }
func (m *OnAddNewsResponse) String() string { return proto.CompactTextString(m) }
func (*OnAddNewsResponse) ProtoMessage()    {}
func (*OnAddNewsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{1}
}

func (m *OnAddNewsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnAddNewsResponse.Unmarshal(m, b)
}
func (m *OnAddNewsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnAddNewsResponse.Marshal(b, m, deterministic)
}
func (m *OnAddNewsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnAddNewsResponse.Merge(m, src)
}
func (m *OnAddNewsResponse) XXX_Size() int {
	return xxx_messageInfo_OnAddNewsResponse.Size(m)
}
func (m *OnAddNewsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OnAddNewsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OnAddNewsResponse proto.InternalMessageInfo

func (m *OnAddNewsResponse) GetStatus() ResponseStatus {
	if m != nil {
		return m.Status
	}
	return ResponseStatus_FAIL
}

func (m *OnAddNewsResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetNewsRequest struct {
	Status               ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=contract.ResponseStatus" json:"status,omitempty"`
	Id                   uint32         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetNewsRequest) Reset()         { *m = GetNewsRequest{} }
func (m *GetNewsRequest) String() string { return proto.CompactTextString(m) }
func (*GetNewsRequest) ProtoMessage()    {}
func (*GetNewsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{2}
}

func (m *GetNewsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNewsRequest.Unmarshal(m, b)
}
func (m *GetNewsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNewsRequest.Marshal(b, m, deterministic)
}
func (m *GetNewsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNewsRequest.Merge(m, src)
}
func (m *GetNewsRequest) XXX_Size() int {
	return xxx_messageInfo_GetNewsRequest.Size(m)
}
func (m *GetNewsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNewsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetNewsRequest proto.InternalMessageInfo

func (m *GetNewsRequest) GetStatus() ResponseStatus {
	if m != nil {
		return m.Status
	}
	return ResponseStatus_FAIL
}

func (m *GetNewsRequest) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type OnGetNewsResponse struct {
	Status               ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=contract.ResponseStatus" json:"status,omitempty"`
	Id                   uint32         `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Title                string         `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Date                 string         `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *OnGetNewsResponse) Reset()         { *m = OnGetNewsResponse{} }
func (m *OnGetNewsResponse) String() string { return proto.CompactTextString(m) }
func (*OnGetNewsResponse) ProtoMessage()    {}
func (*OnGetNewsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{3}
}

func (m *OnGetNewsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnGetNewsResponse.Unmarshal(m, b)
}
func (m *OnGetNewsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnGetNewsResponse.Marshal(b, m, deterministic)
}
func (m *OnGetNewsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnGetNewsResponse.Merge(m, src)
}
func (m *OnGetNewsResponse) XXX_Size() int {
	return xxx_messageInfo_OnGetNewsResponse.Size(m)
}
func (m *OnGetNewsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OnGetNewsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OnGetNewsResponse proto.InternalMessageInfo

func (m *OnGetNewsResponse) GetStatus() ResponseStatus {
	if m != nil {
		return m.Status
	}
	return ResponseStatus_FAIL
}

func (m *OnGetNewsResponse) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OnGetNewsResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *OnGetNewsResponse) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func init() {
	proto.RegisterEnum("contract.ResponseStatus", ResponseStatus_name, ResponseStatus_value)
	proto.RegisterType((*AddNewsRequest)(nil), "contract.AddNewsRequest")
	proto.RegisterType((*OnAddNewsResponse)(nil), "contract.OnAddNewsResponse")
	proto.RegisterType((*GetNewsRequest)(nil), "contract.GetNewsRequest")
	proto.RegisterType((*OnGetNewsResponse)(nil), "contract.OnGetNewsResponse")
}

func init() { proto.RegisterFile("contract.proto", fileDescriptor_d19debeba7dea55a) }

var fileDescriptor_d19debeba7dea55a = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xce, 0xcf, 0x2b,
	0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x72,
	0xb8, 0xf8, 0x1c, 0x53, 0x52, 0xfc, 0x52, 0xcb, 0x8b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b,
	0x84, 0x0c, 0xb8, 0xd8, 0x8a, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0xf8,
	0x8c, 0x24, 0xf4, 0xe0, 0x9a, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x83, 0xc1, 0xf2,
	0x41, 0x50, 0x75, 0x42, 0x22, 0x5c, 0xac, 0x25, 0x99, 0x25, 0x39, 0xa9, 0x12, 0x4c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x10, 0x17, 0x4b, 0x4a, 0x62, 0x49, 0xaa, 0x04, 0x33, 0x58,
	0x10, 0xcc, 0x56, 0x0a, 0xe5, 0x12, 0xf4, 0xcf, 0x83, 0xdb, 0x07, 0x31, 0x8c, 0x0c, 0x0b, 0xf9,
	0xb8, 0x98, 0x32, 0x53, 0xc0, 0xb6, 0xf1, 0x06, 0x31, 0x65, 0xa6, 0x28, 0x05, 0x71, 0xf1, 0xb9,
	0xa7, 0x96, 0x50, 0xe6, 0x09, 0x74, 0x33, 0xeb, 0x41, 0x4e, 0x85, 0x9b, 0x4a, 0x2d, 0xa7, 0x22,
	0xc2, 0x8a, 0x19, 0x5b, 0x58, 0xb1, 0x20, 0xc2, 0x4a, 0x4b, 0x89, 0x8b, 0x0f, 0xd5, 0x4c, 0x21,
	0x0e, 0x2e, 0x16, 0x37, 0x47, 0x4f, 0x1f, 0x01, 0x06, 0x21, 0x36, 0x2e, 0x26, 0x7f, 0x6f, 0x01,
	0xc6, 0x24, 0x36, 0x70, 0x74, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xc6, 0xce, 0x0c,
	0xe0, 0x01, 0x00, 0x00,
}
