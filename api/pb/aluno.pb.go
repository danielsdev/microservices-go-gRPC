// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/aluno.proto

package pb

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

type StudentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cpf                  string   `protobuf:"bytes,2,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Rg                   string   `protobuf:"bytes,3,opt,name=rg,proto3" json:"rg,omitempty"`
	Matricula            string   `protobuf:"bytes,4,opt,name=matricula,proto3" json:"matricula,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StudentRequest) Reset()         { *m = StudentRequest{} }
func (m *StudentRequest) String() string { return proto.CompactTextString(m) }
func (*StudentRequest) ProtoMessage()    {}
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{0}
}

func (m *StudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StudentRequest.Unmarshal(m, b)
}
func (m *StudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StudentRequest.Marshal(b, m, deterministic)
}
func (m *StudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentRequest.Merge(m, src)
}
func (m *StudentRequest) XXX_Size() int {
	return xxx_messageInfo_StudentRequest.Size(m)
}
func (m *StudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StudentRequest proto.InternalMessageInfo

func (m *StudentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StudentRequest) GetCpf() string {
	if m != nil {
		return m.Cpf
	}
	return ""
}

func (m *StudentRequest) GetRg() string {
	if m != nil {
		return m.Rg
	}
	return ""
}

func (m *StudentRequest) GetMatricula() string {
	if m != nil {
		return m.Matricula
	}
	return ""
}

type StudentResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cpf                  string   `protobuf:"bytes,3,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Rg                   string   `protobuf:"bytes,4,opt,name=rg,proto3" json:"rg,omitempty"`
	Matricula            string   `protobuf:"bytes,5,opt,name=matricula,proto3" json:"matricula,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StudentResponse) Reset()         { *m = StudentResponse{} }
func (m *StudentResponse) String() string { return proto.CompactTextString(m) }
func (*StudentResponse) ProtoMessage()    {}
func (*StudentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{1}
}

func (m *StudentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StudentResponse.Unmarshal(m, b)
}
func (m *StudentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StudentResponse.Marshal(b, m, deterministic)
}
func (m *StudentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentResponse.Merge(m, src)
}
func (m *StudentResponse) XXX_Size() int {
	return xxx_messageInfo_StudentResponse.Size(m)
}
func (m *StudentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StudentResponse proto.InternalMessageInfo

func (m *StudentResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StudentResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StudentResponse) GetCpf() string {
	if m != nil {
		return m.Cpf
	}
	return ""
}

func (m *StudentResponse) GetRg() string {
	if m != nil {
		return m.Rg
	}
	return ""
}

func (m *StudentResponse) GetMatricula() string {
	if m != nil {
		return m.Matricula
	}
	return ""
}

type EditStudentRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Cpf                  string   `protobuf:"bytes,3,opt,name=cpf,proto3" json:"cpf,omitempty"`
	Rg                   string   `protobuf:"bytes,4,opt,name=rg,proto3" json:"rg,omitempty"`
	Matricula            string   `protobuf:"bytes,5,opt,name=matricula,proto3" json:"matricula,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EditStudentRequest) Reset()         { *m = EditStudentRequest{} }
func (m *EditStudentRequest) String() string { return proto.CompactTextString(m) }
func (*EditStudentRequest) ProtoMessage()    {}
func (*EditStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{2}
}

func (m *EditStudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EditStudentRequest.Unmarshal(m, b)
}
func (m *EditStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EditStudentRequest.Marshal(b, m, deterministic)
}
func (m *EditStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditStudentRequest.Merge(m, src)
}
func (m *EditStudentRequest) XXX_Size() int {
	return xxx_messageInfo_EditStudentRequest.Size(m)
}
func (m *EditStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EditStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EditStudentRequest proto.InternalMessageInfo

func (m *EditStudentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EditStudentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EditStudentRequest) GetCpf() string {
	if m != nil {
		return m.Cpf
	}
	return ""
}

func (m *EditStudentRequest) GetRg() string {
	if m != nil {
		return m.Rg
	}
	return ""
}

func (m *EditStudentRequest) GetMatricula() string {
	if m != nil {
		return m.Matricula
	}
	return ""
}

type GetStudentRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStudentRequest) Reset()         { *m = GetStudentRequest{} }
func (m *GetStudentRequest) String() string { return proto.CompactTextString(m) }
func (*GetStudentRequest) ProtoMessage()    {}
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{3}
}

func (m *GetStudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStudentRequest.Unmarshal(m, b)
}
func (m *GetStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStudentRequest.Marshal(b, m, deterministic)
}
func (m *GetStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentRequest.Merge(m, src)
}
func (m *GetStudentRequest) XXX_Size() int {
	return xxx_messageInfo_GetStudentRequest.Size(m)
}
func (m *GetStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentRequest proto.InternalMessageInfo

func (m *GetStudentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{4}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{5}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ListStudentsRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListStudentsRequest) Reset()         { *m = ListStudentsRequest{} }
func (m *ListStudentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListStudentsRequest) ProtoMessage()    {}
func (*ListStudentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{6}
}

func (m *ListStudentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStudentsRequest.Unmarshal(m, b)
}
func (m *ListStudentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStudentsRequest.Marshal(b, m, deterministic)
}
func (m *ListStudentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStudentsRequest.Merge(m, src)
}
func (m *ListStudentsRequest) XXX_Size() int {
	return xxx_messageInfo_ListStudentsRequest.Size(m)
}
func (m *ListStudentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStudentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListStudentsRequest proto.InternalMessageInfo

func (m *ListStudentsRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListStudentsResponse struct {
	Student              []*StudentResponse `protobuf:"bytes,1,rep,name=student,proto3" json:"student,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListStudentsResponse) Reset()         { *m = ListStudentsResponse{} }
func (m *ListStudentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListStudentsResponse) ProtoMessage()    {}
func (*ListStudentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{7}
}

func (m *ListStudentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListStudentsResponse.Unmarshal(m, b)
}
func (m *ListStudentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListStudentsResponse.Marshal(b, m, deterministic)
}
func (m *ListStudentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListStudentsResponse.Merge(m, src)
}
func (m *ListStudentsResponse) XXX_Size() int {
	return xxx_messageInfo_ListStudentsResponse.Size(m)
}
func (m *ListStudentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListStudentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListStudentsResponse proto.InternalMessageInfo

func (m *ListStudentsResponse) GetStudent() []*StudentResponse {
	if m != nil {
		return m.Student
	}
	return nil
}

type DeleteStudentRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteStudentRequest) Reset()         { *m = DeleteStudentRequest{} }
func (m *DeleteStudentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteStudentRequest) ProtoMessage()    {}
func (*DeleteStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{8}
}

func (m *DeleteStudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStudentRequest.Unmarshal(m, b)
}
func (m *DeleteStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStudentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStudentRequest.Merge(m, src)
}
func (m *DeleteStudentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteStudentRequest.Size(m)
}
func (m *DeleteStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStudentRequest proto.InternalMessageInfo

func (m *DeleteStudentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteStudentResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteStudentResponse) Reset()         { *m = DeleteStudentResponse{} }
func (m *DeleteStudentResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteStudentResponse) ProtoMessage()    {}
func (*DeleteStudentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ea9864cc77e8211, []int{9}
}

func (m *DeleteStudentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStudentResponse.Unmarshal(m, b)
}
func (m *DeleteStudentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStudentResponse.Marshal(b, m, deterministic)
}
func (m *DeleteStudentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStudentResponse.Merge(m, src)
}
func (m *DeleteStudentResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteStudentResponse.Size(m)
}
func (m *DeleteStudentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStudentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStudentResponse proto.InternalMessageInfo

func (m *DeleteStudentResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*StudentRequest)(nil), "StudentRequest")
	proto.RegisterType((*StudentResponse)(nil), "StudentResponse")
	proto.RegisterType((*EditStudentRequest)(nil), "EditStudentRequest")
	proto.RegisterType((*GetStudentRequest)(nil), "GetStudentRequest")
	proto.RegisterType((*HelloRequest)(nil), "HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "HelloResponse")
	proto.RegisterType((*ListStudentsRequest)(nil), "ListStudentsRequest")
	proto.RegisterType((*ListStudentsResponse)(nil), "ListStudentsResponse")
	proto.RegisterType((*DeleteStudentRequest)(nil), "DeleteStudentRequest")
	proto.RegisterType((*DeleteStudentResponse)(nil), "DeleteStudentResponse")
}

func init() {
	proto.RegisterFile("proto/aluno.proto", fileDescriptor_0ea9864cc77e8211)
}

var fileDescriptor_0ea9864cc77e8211 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcb, 0x6f, 0xda, 0x40,
	0x10, 0xc6, 0xc1, 0x36, 0x50, 0x06, 0xcc, 0x63, 0x78, 0xc8, 0xb2, 0x7a, 0x40, 0x5b, 0xa9, 0x82,
	0x56, 0x5a, 0x24, 0x5a, 0xf5, 0x56, 0xa9, 0xa2, 0xad, 0x92, 0x43, 0x4e, 0x70, 0xcb, 0x6d, 0xc1,
	0x1b, 0x64, 0xc9, 0xd8, 0x8e, 0x77, 0x2d, 0x25, 0x7f, 0x6a, 0xfe, 0x9b, 0x88, 0xc5, 0x06, 0xfc,
	0x10, 0xdc, 0x72, 0x9b, 0x1d, 0x7d, 0xde, 0xdf, 0x7e, 0xdf, 0x8c, 0xa1, 0x1f, 0x46, 0x81, 0x0c,
	0xe6, 0xcc, 0x8b, 0xfd, 0x80, 0xaa, 0x9a, 0x38, 0xd0, 0x59, 0xcb, 0xd8, 0xe1, 0xbe, 0x5c, 0xf1,
	0xe7, 0x98, 0x0b, 0x89, 0x08, 0x86, 0xcf, 0xf6, 0xdc, 0xaa, 0x4e, 0xaa, 0xd3, 0xe6, 0x4a, 0xd5,
	0xd8, 0x03, 0x7d, 0x1b, 0x3e, 0x59, 0x9a, 0x6a, 0x1d, 0x4a, 0xec, 0x80, 0x16, 0xed, 0x2c, 0x5d,
	0x35, 0xb4, 0x68, 0x87, 0x9f, 0xa1, 0xb9, 0x67, 0x32, 0x72, 0xb7, 0xb1, 0xc7, 0x2c, 0x43, 0xb5,
	0xcf, 0x0d, 0x12, 0x43, 0xf7, 0x44, 0x11, 0x61, 0xe0, 0x0b, 0x7e, 0xb8, 0xc0, 0x75, 0x12, 0x88,
	0xe6, 0x3a, 0x27, 0xac, 0x56, 0xc4, 0xea, 0x79, 0xac, 0x51, 0x8e, 0xad, 0xe5, 0xb1, 0x2f, 0x80,
	0xff, 0x1d, 0x57, 0xe6, 0x0c, 0x7e, 0x04, 0xf9, 0x0b, 0xf4, 0xef, 0xf8, 0x0d, 0x30, 0x21, 0xd0,
	0xbe, 0xe7, 0x9e, 0x17, 0x5c, 0x49, 0x9e, 0xcc, 0xc0, 0x4c, 0x34, 0x49, 0x6e, 0x16, 0x34, 0xf6,
	0x5c, 0x08, 0xb6, 0x4b, 0x75, 0xe9, 0x91, 0xcc, 0x60, 0xf0, 0xe0, 0x8a, 0x14, 0x2a, 0xae, 0xdd,
	0xba, 0x84, 0x61, 0x56, 0x9a, 0x5c, 0xfe, 0x0d, 0x1a, 0xe2, 0xd8, 0xb3, 0xaa, 0x13, 0x7d, 0xda,
	0x5a, 0xf4, 0x68, 0x6e, 0x6e, 0xab, 0x54, 0x40, 0xbe, 0xc2, 0xf0, 0x1f, 0xf7, 0xb8, 0xe4, 0x37,
	0x5c, 0xce, 0x61, 0x94, 0xd3, 0x25, 0xb0, 0x31, 0xd4, 0x85, 0x64, 0x32, 0x16, 0x89, 0x38, 0x39,
	0x2d, 0xde, 0x34, 0xa8, 0x29, 0xcf, 0xf8, 0x1d, 0x3e, 0xad, 0xd9, 0xeb, 0xb1, 0x36, 0xe9, 0x65,
	0x56, 0x76, 0x87, 0x66, 0x62, 0x21, 0x15, 0xfc, 0x09, 0xe6, 0xdf, 0x88, 0xb3, 0x13, 0x07, 0xbb,
	0x34, 0xfb, 0x32, 0xbb, 0x60, 0x86, 0x54, 0xf0, 0x17, 0xb4, 0x2e, 0x56, 0x04, 0x07, 0xb4, 0xb8,
	0x30, 0xa5, 0xdf, 0xfd, 0x01, 0x33, 0xe3, 0x0a, 0x47, 0xb4, 0x2c, 0x0d, 0x7b, 0x4c, 0x4b, 0xcd,
	0x93, 0x0a, 0xfe, 0x86, 0xf6, 0xe5, 0x0c, 0x70, 0x48, 0x4b, 0xa6, 0x67, 0x8f, 0x68, 0xd9, 0xa0,
	0x94, 0x5d, 0x38, 0x6f, 0x18, 0x22, 0x2d, 0xac, 0x5b, 0xd9, 0xb3, 0x97, 0xf5, 0x47, 0x83, 0xce,
	0xc3, 0xcd, 0xa6, 0xae, 0xfe, 0xfe, 0x1f, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0xd4, 0x62,
	0x16, 0x12, 0x04, 0x00, 0x00,
}
