// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CalcRequest struct {
	No1                  int32    `protobuf:"varint,1,opt,name=no1,proto3" json:"no1,omitempty"`
	No2                  int32    `protobuf:"varint,2,opt,name=no2,proto3" json:"no2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcRequest) Reset()         { *m = CalcRequest{} }
func (m *CalcRequest) String() string { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()    {}
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{0}
}

func (m *CalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcRequest.Unmarshal(m, b)
}
func (m *CalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcRequest.Marshal(b, m, deterministic)
}
func (m *CalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcRequest.Merge(m, src)
}
func (m *CalcRequest) XXX_Size() int {
	return xxx_messageInfo_CalcRequest.Size(m)
}
func (m *CalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcRequest proto.InternalMessageInfo

func (m *CalcRequest) GetNo1() int32 {
	if m != nil {
		return m.No1
	}
	return 0
}

func (m *CalcRequest) GetNo2() int32 {
	if m != nil {
		return m.No2
	}
	return 0
}

type CalcPrimeRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcPrimeRequest) Reset()         { *m = CalcPrimeRequest{} }
func (m *CalcPrimeRequest) String() string { return proto.CompactTextString(m) }
func (*CalcPrimeRequest) ProtoMessage()    {}
func (*CalcPrimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{1}
}

func (m *CalcPrimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcPrimeRequest.Unmarshal(m, b)
}
func (m *CalcPrimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcPrimeRequest.Marshal(b, m, deterministic)
}
func (m *CalcPrimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcPrimeRequest.Merge(m, src)
}
func (m *CalcPrimeRequest) XXX_Size() int {
	return xxx_messageInfo_CalcPrimeRequest.Size(m)
}
func (m *CalcPrimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcPrimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcPrimeRequest proto.InternalMessageInfo

func (m *CalcPrimeRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type CalcResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcResponse) Reset()         { *m = CalcResponse{} }
func (m *CalcResponse) String() string { return proto.CompactTextString(m) }
func (*CalcResponse) ProtoMessage()    {}
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{2}
}

func (m *CalcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcResponse.Unmarshal(m, b)
}
func (m *CalcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcResponse.Marshal(b, m, deterministic)
}
func (m *CalcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcResponse.Merge(m, src)
}
func (m *CalcResponse) XXX_Size() int {
	return xxx_messageInfo_CalcResponse.Size(m)
}
func (m *CalcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalcResponse proto.InternalMessageInfo

func (m *CalcResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ComputeAverageRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageRequest) Reset()         { *m = ComputeAverageRequest{} }
func (m *ComputeAverageRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageRequest) ProtoMessage()    {}
func (*ComputeAverageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{3}
}

func (m *ComputeAverageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageRequest.Unmarshal(m, b)
}
func (m *ComputeAverageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAverageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageRequest.Merge(m, src)
}
func (m *ComputeAverageRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageRequest.Size(m)
}
func (m *ComputeAverageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageRequest proto.InternalMessageInfo

func (m *ComputeAverageRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAverageResponse struct {
	Result               float32  `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAverageResponse) Reset()         { *m = ComputeAverageResponse{} }
func (m *ComputeAverageResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAverageResponse) ProtoMessage()    {}
func (*ComputeAverageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{4}
}

func (m *ComputeAverageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAverageResponse.Unmarshal(m, b)
}
func (m *ComputeAverageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAverageResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAverageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAverageResponse.Merge(m, src)
}
func (m *ComputeAverageResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAverageResponse.Size(m)
}
func (m *ComputeAverageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAverageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAverageResponse proto.InternalMessageInfo

func (m *ComputeAverageResponse) GetResult() float32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type FindMaxRequest struct {
	Number               int32    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxRequest) Reset()         { *m = FindMaxRequest{} }
func (m *FindMaxRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaxRequest) ProtoMessage()    {}
func (*FindMaxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{5}
}

func (m *FindMaxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxRequest.Unmarshal(m, b)
}
func (m *FindMaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxRequest.Marshal(b, m, deterministic)
}
func (m *FindMaxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxRequest.Merge(m, src)
}
func (m *FindMaxRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaxRequest.Size(m)
}
func (m *FindMaxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxRequest proto.InternalMessageInfo

func (m *FindMaxRequest) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

type FindMaxResponse struct {
	Result               int32    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxResponse) Reset()         { *m = FindMaxResponse{} }
func (m *FindMaxResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaxResponse) ProtoMessage()    {}
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f42938f8c8365cf, []int{6}
}

func (m *FindMaxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxResponse.Unmarshal(m, b)
}
func (m *FindMaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxResponse.Marshal(b, m, deterministic)
}
func (m *FindMaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxResponse.Merge(m, src)
}
func (m *FindMaxResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaxResponse.Size(m)
}
func (m *FindMaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxResponse proto.InternalMessageInfo

func (m *FindMaxResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalcRequest)(nil), "calculator.CalcRequest")
	proto.RegisterType((*CalcPrimeRequest)(nil), "calculator.CalcPrimeRequest")
	proto.RegisterType((*CalcResponse)(nil), "calculator.CalcResponse")
	proto.RegisterType((*ComputeAverageRequest)(nil), "calculator.ComputeAverageRequest")
	proto.RegisterType((*ComputeAverageResponse)(nil), "calculator.ComputeAverageResponse")
	proto.RegisterType((*FindMaxRequest)(nil), "calculator.FindMaxRequest")
	proto.RegisterType((*FindMaxResponse)(nil), "calculator.FindMaxResponse")
}

func init() {
	proto.RegisterFile("calculator/calculatorpb/calculator.proto", fileDescriptor_7f42938f8c8365cf)
}

var fileDescriptor_7f42938f8c8365cf = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4f, 0xb3, 0x40,
	0x10, 0xc6, 0x5f, 0x68, 0xda, 0x37, 0x19, 0x1b, 0xac, 0x9b, 0x58, 0x09, 0x7a, 0xd0, 0x3d, 0x18,
	0xf4, 0xd0, 0x52, 0xbc, 0x79, 0x53, 0x12, 0xe3, 0xc5, 0xc4, 0xb4, 0x37, 0x3d, 0x01, 0x4e, 0x0c,
	0x09, 0xb0, 0xb8, 0xec, 0x36, 0x7e, 0x37, 0xbf, 0x9c, 0xa1, 0xfc, 0x5b, 0x2a, 0xa1, 0xb7, 0x9d,
	0x99, 0xe7, 0x79, 0x66, 0xf3, 0xcb, 0x80, 0x1d, 0xfa, 0x71, 0x28, 0x63, 0x5f, 0x30, 0xbe, 0x6c,
	0x9f, 0x59, 0xa0, 0x14, 0x8b, 0x8c, 0x33, 0xc1, 0x08, 0xb4, 0x1d, 0xba, 0x82, 0x23, 0xcf, 0x8f,
	0xc3, 0x35, 0x7e, 0x49, 0xcc, 0x05, 0x99, 0xc1, 0x28, 0x65, 0x2b, 0x53, 0xbb, 0xd4, 0xec, 0xf1,
	0xba, 0x78, 0x96, 0x1d, 0xd7, 0xd4, 0xeb, 0x8e, 0x4b, 0x6f, 0x61, 0x56, 0x58, 0x5e, 0x79, 0x94,
	0x60, 0xed, 0x9b, 0xc3, 0x24, 0x95, 0x49, 0x80, 0xbc, 0xb2, 0x56, 0x15, 0xbd, 0x86, 0x69, 0x19,
	0x9f, 0x67, 0x2c, 0xcd, 0xb1, 0xd0, 0x71, 0xcc, 0x65, 0x2c, 0x6a, 0x5d, 0x59, 0xd1, 0x25, 0x9c,
	0x7a, 0x2c, 0xc9, 0xa4, 0xc0, 0x87, 0x2d, 0x72, 0xff, 0xf3, 0x60, 0xb0, 0x03, 0xf3, 0x7d, 0x43,
	0xef, 0x0a, 0xbd, 0x59, 0x61, 0x83, 0xf1, 0x14, 0xa5, 0x1f, 0x2f, 0xfe, 0xf7, 0xa1, 0xec, 0x1b,
	0x38, 0x6e, 0x94, 0xc3, 0xff, 0x76, 0x7f, 0x74, 0x38, 0xf1, 0x1a, 0x9a, 0x1b, 0xe4, 0xdb, 0x28,
	0x44, 0x72, 0x0f, 0xa3, 0x8d, 0x4c, 0xc8, 0xd9, 0x42, 0x41, 0xaf, 0x50, 0xb6, 0xcc, 0xbf, 0x83,
	0x72, 0x0f, 0xfd, 0x47, 0x3c, 0x18, 0xef, 0xc8, 0x92, 0x8b, 0x7d, 0x91, 0x0a, 0x7c, 0x28, 0xc2,
	0xd1, 0xc8, 0x3b, 0x18, 0x5d, 0x3a, 0xe4, 0xaa, 0xa3, 0xef, 0x43, 0x6d, 0xd1, 0x21, 0x49, 0x1d,
	0x6e, 0x6b, 0xe4, 0x19, 0xfe, 0x57, 0x78, 0x88, 0xa5, 0x5a, 0xba, 0x74, 0xad, 0xf3, 0xde, 0x59,
	0x9b, 0xe3, 0x68, 0x8f, 0xc6, 0xdb, 0x54, 0xbd, 0xd4, 0x60, 0xb2, 0xbb, 0xcf, 0xbb, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x14, 0x34, 0x41, 0x65, 0xcb, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	//Unary
	Sum(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	Prime(ctx context.Context, in *CalcPrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error)
	FindMax(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaxClient, error)
}

type calculatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorServiceClient(cc *grpc.ClientConn) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Prime(ctx context.Context, in *CalcPrimeRequest, opts ...grpc.CallOption) (CalculatorService_PrimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[0], "/calculator.CalculatorService/Prime", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServicePrimeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_PrimeClient interface {
	Recv() (*CalcResponse, error)
	grpc.ClientStream
}

type calculatorServicePrimeClient struct {
	grpc.ClientStream
}

func (x *calculatorServicePrimeClient) Recv() (*CalcResponse, error) {
	m := new(CalcResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[1], "/calculator.CalculatorService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceComputeAverageClient{stream}
	return x, nil
}

type CalculatorService_ComputeAverageClient interface {
	Send(*ComputeAverageRequest) error
	CloseAndRecv() (*ComputeAverageResponse, error)
	grpc.ClientStream
}

type calculatorServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceComputeAverageClient) Send(m *ComputeAverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageClient) CloseAndRecv() (*ComputeAverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_FindMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalculatorService_serviceDesc.Streams[2], "/calculator.CalculatorService/FindMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceFindMaxClient{stream}
	return x, nil
}

type CalculatorService_FindMaxClient interface {
	Send(*FindMaxRequest) error
	Recv() (*FindMaxResponse, error)
	grpc.ClientStream
}

type calculatorServiceFindMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceFindMaxClient) Send(m *FindMaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceFindMaxClient) Recv() (*FindMaxResponse, error) {
	m := new(FindMaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
type CalculatorServiceServer interface {
	//Unary
	Sum(context.Context, *CalcRequest) (*CalcResponse, error)
	Prime(*CalcPrimeRequest, CalculatorService_PrimeServer) error
	ComputeAverage(CalculatorService_ComputeAverageServer) error
	FindMax(CalculatorService_FindMaxServer) error
}

// UnimplementedCalculatorServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (*UnimplementedCalculatorServiceServer) Sum(ctx context.Context, req *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalculatorServiceServer) Prime(req *CalcPrimeRequest, srv CalculatorService_PrimeServer) error {
	return status.Errorf(codes.Unimplemented, "method Prime not implemented")
}
func (*UnimplementedCalculatorServiceServer) ComputeAverage(srv CalculatorService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalculatorServiceServer) FindMax(srv CalculatorService_FindMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}

func RegisterCalculatorServiceServer(s *grpc.Server, srv CalculatorServiceServer) {
	s.RegisterService(&_CalculatorService_serviceDesc, srv)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Prime_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CalcPrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).Prime(m, &calculatorServicePrimeServer{stream})
}

type CalculatorService_PrimeServer interface {
	Send(*CalcResponse) error
	grpc.ServerStream
}

type calculatorServicePrimeServer struct {
	grpc.ServerStream
}

func (x *calculatorServicePrimeServer) Send(m *CalcResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).ComputeAverage(&calculatorServiceComputeAverageServer{stream})
}

type CalculatorService_ComputeAverageServer interface {
	SendAndClose(*ComputeAverageResponse) error
	Recv() (*ComputeAverageRequest, error)
	grpc.ServerStream
}

type calculatorServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceComputeAverageServer) SendAndClose(m *ComputeAverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceComputeAverageServer) Recv() (*ComputeAverageRequest, error) {
	m := new(ComputeAverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).FindMax(&calculatorServiceFindMaxServer{stream})
}

type CalculatorService_FindMaxServer interface {
	Send(*FindMaxResponse) error
	Recv() (*FindMaxRequest, error)
	grpc.ServerStream
}

type calculatorServiceFindMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceFindMaxServer) Send(m *FindMaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceFindMaxServer) Recv() (*FindMaxRequest, error) {
	m := new(FindMaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalculatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Prime",
			Handler:       _CalculatorService_Prime_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalculatorService_ComputeAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMax",
			Handler:       _CalculatorService_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
