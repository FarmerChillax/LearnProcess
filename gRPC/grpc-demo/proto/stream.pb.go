// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stream.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type StreamPoint struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamPoint) Reset()         { *m = StreamPoint{} }
func (m *StreamPoint) String() string { return proto.CompactTextString(m) }
func (*StreamPoint) ProtoMessage()    {}
func (*StreamPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func (m *StreamPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamPoint.Unmarshal(m, b)
}
func (m *StreamPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamPoint.Marshal(b, m, deterministic)
}
func (m *StreamPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamPoint.Merge(m, src)
}
func (m *StreamPoint) XXX_Size() int {
	return xxx_messageInfo_StreamPoint.Size(m)
}
func (m *StreamPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamPoint.DiscardUnknown(m)
}

var xxx_messageInfo_StreamPoint proto.InternalMessageInfo

func (m *StreamPoint) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StreamPoint) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StreamRequest struct {
	Pt                   *StreamPoint `protobuf:"bytes,1,opt,name=pt,proto3" json:"pt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{1}
}

func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (m *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(m, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetPt() *StreamPoint {
	if m != nil {
		return m.Pt
	}
	return nil
}

type StreamResponse struct {
	Pt                   *StreamPoint `protobuf:"bytes,1,opt,name=pt,proto3" json:"pt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{2}
}

func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (m *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(m, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetPt() *StreamPoint {
	if m != nil {
		return m.Pt
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamPoint)(nil), "proto.StreamPoint")
	proto.RegisterType((*StreamRequest)(nil), "proto.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "proto.StreamResponse")
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54) }

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xce, 0xb1, 0x4e, 0x87, 0x30,
	0x10, 0x06, 0xf0, 0x94, 0xfc, 0x4b, 0xe2, 0xa1, 0x0e, 0x17, 0x4c, 0x88, 0x13, 0xe9, 0xc4, 0x44,
	0x08, 0x68, 0xf0, 0x21, 0x1c, 0x4c, 0x79, 0x02, 0xc4, 0x1b, 0x48, 0x84, 0xd6, 0xf6, 0xe0, 0xdd,
	0x7c, 0x3b, 0x93, 0x56, 0x12, 0x75, 0x92, 0xa9, 0xed, 0xd7, 0xfc, 0xee, 0x3b, 0xb8, 0xf6, 0xec,
	0x68, 0x5c, 0x6a, 0xeb, 0x0c, 0x1b, 0x94, 0xe1, 0x50, 0x3d, 0x64, 0x43, 0x88, 0x5f, 0xcc, 0xbc,
	0x32, 0x22, 0x5c, 0xd6, 0x71, 0xa1, 0x42, 0x94, 0xa2, 0xba, 0xd2, 0xe1, 0x8e, 0x39, 0xc8, 0x7d,
	0x7c, 0xdf, 0xa8, 0x48, 0x4a, 0x51, 0x49, 0x1d, 0x1f, 0xaa, 0x83, 0x9b, 0x08, 0x35, 0x7d, 0x6c,
	0xe4, 0x19, 0x15, 0x24, 0x96, 0x03, 0xcc, 0x5a, 0x8c, 0x25, 0xf5, 0x8f, 0xd1, 0x3a, 0xb1, 0xac,
	0x1e, 0xe0, 0xf6, 0x40, 0xde, 0x9a, 0xd5, 0xd3, 0x7f, 0x54, 0xfb, 0x29, 0x8e, 0xae, 0x81, 0xdc,
	0x3e, 0x4f, 0x84, 0x8f, 0x70, 0x79, 0x9e, 0x3d, 0x63, 0xfe, 0x4b, 0x7c, 0x6f, 0x72, 0x7f, 0xf7,
	0x27, 0x8d, 0x55, 0x8d, 0xc0, 0x1e, 0x52, 0x4d, 0x93, 0x71, 0x6f, 0xa7, 0x60, 0x25, 0xf0, 0x09,
	0xa4, 0x36, 0x1b, 0xd3, 0x49, 0xd7, 0x88, 0xd7, 0x34, 0xfc, 0x74, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x45, 0x89, 0x89, 0x8e, 0x7d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamServiceClient interface {
	List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_ListClient, error)
	Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error)
	Route(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteClient, error)
}

type streamServiceClient struct {
	cc *grpc.ClientConn
}

func NewStreamServiceClient(cc *grpc.ClientConn) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) List(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (StreamService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[0], "/proto.StreamService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamService_ListClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceListClient struct {
	grpc.ClientStream
}

func (x *streamServiceListClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Record(ctx context.Context, opts ...grpc.CallOption) (StreamService_RecordClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[1], "/proto.StreamService/Record", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRecordClient{stream}
	return x, nil
}

type StreamService_RecordClient interface {
	Send(*StreamRequest) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceRecordClient struct {
	grpc.ClientStream
}

func (x *streamServiceRecordClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRecordClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamServiceClient) Route(ctx context.Context, opts ...grpc.CallOption) (StreamService_RouteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamService_serviceDesc.Streams[2], "/proto.StreamService/Route", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamServiceRouteClient{stream}
	return x, nil
}

type StreamService_RouteClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type streamServiceRouteClient struct {
	grpc.ClientStream
}

func (x *streamServiceRouteClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamServiceRouteClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServiceServer is the server API for StreamService service.
type StreamServiceServer interface {
	List(*StreamRequest, StreamService_ListServer) error
	Record(StreamService_RecordServer) error
	Route(StreamService_RouteServer) error
}

// UnimplementedStreamServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServiceServer struct {
}

func (*UnimplementedStreamServiceServer) List(req *StreamRequest, srv StreamService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedStreamServiceServer) Record(srv StreamService_RecordServer) error {
	return status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (*UnimplementedStreamServiceServer) Route(srv StreamService_RouteServer) error {
	return status.Errorf(codes.Unimplemented, "method Route not implemented")
}

func RegisterStreamServiceServer(s *grpc.Server, srv StreamServiceServer) {
	s.RegisterService(&_StreamService_serviceDesc, srv)
}

func _StreamService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServiceServer).List(m, &streamServiceListServer{stream})
}

type StreamService_ListServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type streamServiceListServer struct {
	grpc.ServerStream
}

func (x *streamServiceListServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamService_Record_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Record(&streamServiceRecordServer{stream})
}

type StreamService_RecordServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRecordServer struct {
	grpc.ServerStream
}

func (x *streamServiceRecordServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRecordServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamService_Route_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamServiceServer).Route(&streamServiceRouteServer{stream})
}

type StreamService_RouteServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type streamServiceRouteServer struct {
	grpc.ServerStream
}

func (x *streamServiceRouteServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamServiceRouteServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _StreamService_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Record",
			Handler:       _StreamService_Record_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Route",
			Handler:       _StreamService_Route_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stream.proto",
}
