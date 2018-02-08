// Code generated by protoc-gen-go. DO NOT EDIT.
// source: student.proto

/*
Package Model is a generated protocol buffer package.

It is generated from these files:
	student.proto

It has these top-level messages:
	StudentResponse
	StudentRequest
	ResultResponse
*/
package Model

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

type StudentResponse struct {
	Sid       int32  `protobuf:"varint,1,opt,name=sid" json:"sid,omitempty"`
	Age       int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Telephone string `protobuf:"bytes,4,opt,name=telephone" json:"telephone,omitempty"`
	Address   string `protobuf:"bytes,5,opt,name=address" json:"address,omitempty"`
}

func (m *StudentResponse) Reset()                    { *m = StudentResponse{} }
func (m *StudentResponse) String() string            { return proto.CompactTextString(m) }
func (*StudentResponse) ProtoMessage()               {}
func (*StudentResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StudentResponse) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

func (m *StudentResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *StudentResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StudentResponse) GetTelephone() string {
	if m != nil {
		return m.Telephone
	}
	return ""
}

func (m *StudentResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type StudentRequest struct {
	Sid int32 `protobuf:"varint,1,opt,name=sid" json:"sid,omitempty"`
}

func (m *StudentRequest) Reset()                    { *m = StudentRequest{} }
func (m *StudentRequest) String() string            { return proto.CompactTextString(m) }
func (*StudentRequest) ProtoMessage()               {}
func (*StudentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StudentRequest) GetSid() int32 {
	if m != nil {
		return m.Sid
	}
	return 0
}

type ResultResponse struct {
	Affectrows int32  `protobuf:"varint,1,opt,name=affectrows" json:"affectrows,omitempty"`
	Errorinfo  string `protobuf:"bytes,2,opt,name=errorinfo" json:"errorinfo,omitempty"`
}

func (m *ResultResponse) Reset()                    { *m = ResultResponse{} }
func (m *ResultResponse) String() string            { return proto.CompactTextString(m) }
func (*ResultResponse) ProtoMessage()               {}
func (*ResultResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ResultResponse) GetAffectrows() int32 {
	if m != nil {
		return m.Affectrows
	}
	return 0
}

func (m *ResultResponse) GetErrorinfo() string {
	if m != nil {
		return m.Errorinfo
	}
	return ""
}

func init() {
	proto.RegisterType((*StudentResponse)(nil), "Model.StudentResponse")
	proto.RegisterType((*StudentRequest)(nil), "Model.StudentRequest")
	proto.RegisterType((*ResultResponse)(nil), "Model.ResultResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SelfManage service

type SelfManageClient interface {
	StudentAdd(ctx context.Context, in *StudentResponse, opts ...grpc.CallOption) (*ResultResponse, error)
	GetHelloTest(ctx context.Context, opts ...grpc.CallOption) (SelfManage_GetHelloTestClient, error)
	GetNxinMethod(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (SelfManage_GetNxinMethodClient, error)
	GetStudent(ctx context.Context, opts ...grpc.CallOption) (SelfManage_GetStudentClient, error)
}

type selfManageClient struct {
	cc *grpc.ClientConn
}

func NewSelfManageClient(cc *grpc.ClientConn) SelfManageClient {
	return &selfManageClient{cc}
}

func (c *selfManageClient) StudentAdd(ctx context.Context, in *StudentResponse, opts ...grpc.CallOption) (*ResultResponse, error) {
	out := new(ResultResponse)
	err := grpc.Invoke(ctx, "/Model.SelfManage/StudentAdd", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *selfManageClient) GetHelloTest(ctx context.Context, opts ...grpc.CallOption) (SelfManage_GetHelloTestClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SelfManage_serviceDesc.Streams[0], c.cc, "/Model.SelfManage/GetHelloTest", opts...)
	if err != nil {
		return nil, err
	}
	x := &selfManageGetHelloTestClient{stream}
	return x, nil
}

type SelfManage_GetHelloTestClient interface {
	Send(*StudentRequest) error
	CloseAndRecv() (*ResultResponse, error)
	grpc.ClientStream
}

type selfManageGetHelloTestClient struct {
	grpc.ClientStream
}

func (x *selfManageGetHelloTestClient) Send(m *StudentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *selfManageGetHelloTestClient) CloseAndRecv() (*ResultResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ResultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *selfManageClient) GetNxinMethod(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (SelfManage_GetNxinMethodClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SelfManage_serviceDesc.Streams[1], c.cc, "/Model.SelfManage/GetNxinMethod", opts...)
	if err != nil {
		return nil, err
	}
	x := &selfManageGetNxinMethodClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SelfManage_GetNxinMethodClient interface {
	Recv() (*ResultResponse, error)
	grpc.ClientStream
}

type selfManageGetNxinMethodClient struct {
	grpc.ClientStream
}

func (x *selfManageGetNxinMethodClient) Recv() (*ResultResponse, error) {
	m := new(ResultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *selfManageClient) GetStudent(ctx context.Context, opts ...grpc.CallOption) (SelfManage_GetStudentClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SelfManage_serviceDesc.Streams[2], c.cc, "/Model.SelfManage/GetStudent", opts...)
	if err != nil {
		return nil, err
	}
	x := &selfManageGetStudentClient{stream}
	return x, nil
}

type SelfManage_GetStudentClient interface {
	Send(*StudentRequest) error
	Recv() (*ResultResponse, error)
	grpc.ClientStream
}

type selfManageGetStudentClient struct {
	grpc.ClientStream
}

func (x *selfManageGetStudentClient) Send(m *StudentRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *selfManageGetStudentClient) Recv() (*ResultResponse, error) {
	m := new(ResultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SelfManage service

type SelfManageServer interface {
	StudentAdd(context.Context, *StudentResponse) (*ResultResponse, error)
	GetHelloTest(SelfManage_GetHelloTestServer) error
	GetNxinMethod(*StudentRequest, SelfManage_GetNxinMethodServer) error
	GetStudent(SelfManage_GetStudentServer) error
}

func RegisterSelfManageServer(s *grpc.Server, srv SelfManageServer) {
	s.RegisterService(&_SelfManage_serviceDesc, srv)
}

func _SelfManage_StudentAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelfManageServer).StudentAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Model.SelfManage/StudentAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelfManageServer).StudentAdd(ctx, req.(*StudentResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _SelfManage_GetHelloTest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SelfManageServer).GetHelloTest(&selfManageGetHelloTestServer{stream})
}

type SelfManage_GetHelloTestServer interface {
	SendAndClose(*ResultResponse) error
	Recv() (*StudentRequest, error)
	grpc.ServerStream
}

type selfManageGetHelloTestServer struct {
	grpc.ServerStream
}

func (x *selfManageGetHelloTestServer) SendAndClose(m *ResultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *selfManageGetHelloTestServer) Recv() (*StudentRequest, error) {
	m := new(StudentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SelfManage_GetNxinMethod_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StudentRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SelfManageServer).GetNxinMethod(m, &selfManageGetNxinMethodServer{stream})
}

type SelfManage_GetNxinMethodServer interface {
	Send(*ResultResponse) error
	grpc.ServerStream
}

type selfManageGetNxinMethodServer struct {
	grpc.ServerStream
}

func (x *selfManageGetNxinMethodServer) Send(m *ResultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SelfManage_GetStudent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SelfManageServer).GetStudent(&selfManageGetStudentServer{stream})
}

type SelfManage_GetStudentServer interface {
	Send(*ResultResponse) error
	Recv() (*StudentRequest, error)
	grpc.ServerStream
}

type selfManageGetStudentServer struct {
	grpc.ServerStream
}

func (x *selfManageGetStudentServer) Send(m *ResultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *selfManageGetStudentServer) Recv() (*StudentRequest, error) {
	m := new(StudentRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SelfManage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Model.SelfManage",
	HandlerType: (*SelfManageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StudentAdd",
			Handler:    _SelfManage_StudentAdd_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetHelloTest",
			Handler:       _SelfManage_GetHelloTest_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetNxinMethod",
			Handler:       _SelfManage_GetNxinMethod_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetStudent",
			Handler:       _SelfManage_GetStudent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "student.proto",
}

func init() { proto.RegisterFile("student.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xcd, 0xb6, 0x2a, 0x7d, 0xb8, 0x29, 0x81, 0x49, 0x10, 0x91, 0xd1, 0x53, 0x4f, 0x65,
	0xe8, 0x59, 0xd8, 0x4e, 0xf5, 0xd2, 0x1d, 0x3a, 0xbf, 0x40, 0x34, 0xaf, 0x5b, 0x21, 0x26, 0x35,
	0x49, 0xd1, 0xab, 0x5f, 0xc1, 0x4f, 0x2c, 0xc9, 0xaa, 0xd5, 0x89, 0x87, 0xdd, 0x5e, 0x7f, 0xff,
	0xd7, 0xc7, 0xef, 0x3d, 0x02, 0x63, 0xeb, 0x5a, 0x81, 0xca, 0x65, 0x8d, 0xd1, 0x4e, 0xd3, 0xa8,
	0xd0, 0x02, 0x65, 0xf2, 0x4e, 0xe0, 0x6c, 0xbd, 0x0b, 0x4a, 0xb4, 0x8d, 0x56, 0x16, 0xe9, 0x39,
	0x0c, 0x6d, 0x2d, 0x18, 0x99, 0x91, 0x34, 0x2a, 0x7d, 0xe9, 0x09, 0xdf, 0x20, 0x1b, 0xec, 0x08,
	0xdf, 0x20, 0xa5, 0x30, 0x52, 0xfc, 0x19, 0xd9, 0x70, 0x46, 0xd2, 0xb8, 0x0c, 0x35, 0xbd, 0x82,
	0xd8, 0xa1, 0xc4, 0x66, 0xab, 0x15, 0xb2, 0x51, 0x08, 0x7a, 0x40, 0x19, 0x9c, 0x70, 0x21, 0x0c,
	0x5a, 0xcb, 0xa2, 0x90, 0x7d, 0x7d, 0x26, 0x09, 0x4c, 0xbe, 0x15, 0x5e, 0x5a, 0xb4, 0xee, 0xaf,
	0x41, 0xb2, 0x82, 0x49, 0x89, 0xb6, 0x95, 0xbd, 0xe5, 0x35, 0x00, 0xaf, 0x2a, 0x7c, 0x72, 0x46,
	0xbf, 0xda, 0xae, 0xf5, 0x07, 0xf1, 0x36, 0x68, 0x8c, 0x36, 0xb5, 0xaa, 0x74, 0x30, 0x8f, 0xcb,
	0x1e, 0xdc, 0x7c, 0x0c, 0x00, 0xd6, 0x28, 0xab, 0x82, 0x2b, 0xbf, 0xce, 0x1d, 0x40, 0xa7, 0xb0,
	0x14, 0x82, 0x5e, 0x64, 0xe1, 0x38, 0xd9, 0xde, 0x61, 0x2e, 0xa7, 0x1d, 0xff, 0x6d, 0x92, 0x1c,
	0xd1, 0x05, 0x9c, 0xe6, 0xe8, 0xee, 0x51, 0x4a, 0xfd, 0xe0, 0xfd, 0xa7, 0xfb, 0x03, 0xc2, 0x5a,
	0xff, 0xfe, 0x9f, 0x12, 0xba, 0x84, 0x71, 0x8e, 0x6e, 0xf5, 0x56, 0xab, 0x02, 0xdd, 0x56, 0x8b,
	0x43, 0x47, 0xcc, 0x09, 0x5d, 0x00, 0xe4, 0xe8, 0xba, 0xfe, 0xc3, 0x15, 0xe6, 0xe4, 0xf1, 0x38,
	0x3c, 0x8d, 0xdb, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x3c, 0x0a, 0x6f, 0x2b, 0x02, 0x00,
	0x00,
}