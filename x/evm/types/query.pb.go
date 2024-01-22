// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type QuerySeiAddressByEVMAddressRequest struct {
	EvmAddress string `protobuf:"bytes,1,opt,name=evm_address,json=evmAddress,proto3" json:"evm_address,omitempty"`
}

func (m *QuerySeiAddressByEVMAddressRequest) Reset()         { *m = QuerySeiAddressByEVMAddressRequest{} }
func (m *QuerySeiAddressByEVMAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySeiAddressByEVMAddressRequest) ProtoMessage()    {}
func (*QuerySeiAddressByEVMAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{0}
}
func (m *QuerySeiAddressByEVMAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySeiAddressByEVMAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySeiAddressByEVMAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySeiAddressByEVMAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySeiAddressByEVMAddressRequest.Merge(m, src)
}
func (m *QuerySeiAddressByEVMAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySeiAddressByEVMAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySeiAddressByEVMAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySeiAddressByEVMAddressRequest proto.InternalMessageInfo

func (m *QuerySeiAddressByEVMAddressRequest) GetEvmAddress() string {
	if m != nil {
		return m.EvmAddress
	}
	return ""
}

type QuerySeiAddressByEVMAddressResponse struct {
	SeiAddress string `protobuf:"bytes,1,opt,name=sei_address,json=seiAddress,proto3" json:"sei_address,omitempty"`
	Associated bool   `protobuf:"varint,2,opt,name=associated,proto3" json:"associated,omitempty"`
}

func (m *QuerySeiAddressByEVMAddressResponse) Reset()         { *m = QuerySeiAddressByEVMAddressResponse{} }
func (m *QuerySeiAddressByEVMAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySeiAddressByEVMAddressResponse) ProtoMessage()    {}
func (*QuerySeiAddressByEVMAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{1}
}
func (m *QuerySeiAddressByEVMAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySeiAddressByEVMAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySeiAddressByEVMAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySeiAddressByEVMAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySeiAddressByEVMAddressResponse.Merge(m, src)
}
func (m *QuerySeiAddressByEVMAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySeiAddressByEVMAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySeiAddressByEVMAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySeiAddressByEVMAddressResponse proto.InternalMessageInfo

func (m *QuerySeiAddressByEVMAddressResponse) GetSeiAddress() string {
	if m != nil {
		return m.SeiAddress
	}
	return ""
}

func (m *QuerySeiAddressByEVMAddressResponse) GetAssociated() bool {
	if m != nil {
		return m.Associated
	}
	return false
}

type QueryEVMAddressBySeiAddressRequest struct {
	SeiAddress string `protobuf:"bytes,1,opt,name=sei_address,json=seiAddress,proto3" json:"sei_address,omitempty"`
}

func (m *QueryEVMAddressBySeiAddressRequest) Reset()         { *m = QueryEVMAddressBySeiAddressRequest{} }
func (m *QueryEVMAddressBySeiAddressRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEVMAddressBySeiAddressRequest) ProtoMessage()    {}
func (*QueryEVMAddressBySeiAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{2}
}
func (m *QueryEVMAddressBySeiAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEVMAddressBySeiAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEVMAddressBySeiAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEVMAddressBySeiAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEVMAddressBySeiAddressRequest.Merge(m, src)
}
func (m *QueryEVMAddressBySeiAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEVMAddressBySeiAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEVMAddressBySeiAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEVMAddressBySeiAddressRequest proto.InternalMessageInfo

func (m *QueryEVMAddressBySeiAddressRequest) GetSeiAddress() string {
	if m != nil {
		return m.SeiAddress
	}
	return ""
}

type QueryEVMAddressBySeiAddressResponse struct {
	EvmAddress string `protobuf:"bytes,1,opt,name=evm_address,json=evmAddress,proto3" json:"evm_address,omitempty"`
	Associated bool   `protobuf:"varint,2,opt,name=associated,proto3" json:"associated,omitempty"`
}

func (m *QueryEVMAddressBySeiAddressResponse) Reset()         { *m = QueryEVMAddressBySeiAddressResponse{} }
func (m *QueryEVMAddressBySeiAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEVMAddressBySeiAddressResponse) ProtoMessage()    {}
func (*QueryEVMAddressBySeiAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c0d37eed5339f7, []int{3}
}
func (m *QueryEVMAddressBySeiAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEVMAddressBySeiAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEVMAddressBySeiAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEVMAddressBySeiAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEVMAddressBySeiAddressResponse.Merge(m, src)
}
func (m *QueryEVMAddressBySeiAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEVMAddressBySeiAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEVMAddressBySeiAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEVMAddressBySeiAddressResponse proto.InternalMessageInfo

func (m *QueryEVMAddressBySeiAddressResponse) GetEvmAddress() string {
	if m != nil {
		return m.EvmAddress
	}
	return ""
}

func (m *QueryEVMAddressBySeiAddressResponse) GetAssociated() bool {
	if m != nil {
		return m.Associated
	}
	return false
}

func init() {
	proto.RegisterType((*QuerySeiAddressByEVMAddressRequest)(nil), "seiprotocol.seichain.evm.QuerySeiAddressByEVMAddressRequest")
	proto.RegisterType((*QuerySeiAddressByEVMAddressResponse)(nil), "seiprotocol.seichain.evm.QuerySeiAddressByEVMAddressResponse")
	proto.RegisterType((*QueryEVMAddressBySeiAddressRequest)(nil), "seiprotocol.seichain.evm.QueryEVMAddressBySeiAddressRequest")
	proto.RegisterType((*QueryEVMAddressBySeiAddressResponse)(nil), "seiprotocol.seichain.evm.QueryEVMAddressBySeiAddressResponse")
}

func init() { proto.RegisterFile("evm/query.proto", fileDescriptor_11c0d37eed5339f7) }

var fileDescriptor_11c0d37eed5339f7 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x2d, 0xcb, 0xd5,
	0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0x4e, 0xcd,
	0x04, 0xb3, 0x92, 0xf3, 0x73, 0xf4, 0x8a, 0x53, 0x33, 0x93, 0x33, 0x12, 0x33, 0xf3, 0xf4, 0x52,
	0xcb, 0x72, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13,
	0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xfa, 0x94, 0x5c, 0xb9, 0x94,
	0x02, 0x41, 0xc6, 0x04, 0xa7, 0x66, 0x3a, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x3b, 0x55, 0xba,
	0x86, 0xf9, 0x42, 0xd9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xf2, 0x5c, 0xdc, 0xa9,
	0x65, 0xb9, 0xf1, 0x89, 0x10, 0x51, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xae, 0xd4, 0xb2,
	0x5c, 0xa8, 0x3a, 0xa5, 0x34, 0x2e, 0x65, 0xbc, 0xc6, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x82,
	0xcc, 0x29, 0x4e, 0xcd, 0x44, 0x37, 0xa7, 0x18, 0xae, 0x49, 0x48, 0x8e, 0x8b, 0x2b, 0xb1, 0xb8,
	0x38, 0x3f, 0x39, 0x33, 0xb1, 0x24, 0x35, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x23, 0x08, 0x49,
	0x04, 0xee, 0x5c, 0x84, 0xd9, 0x4e, 0x48, 0x76, 0x22, 0x39, 0x17, 0xaf, 0x35, 0x70, 0xe7, 0xe2,
	0x32, 0x06, 0xe1, 0x5c, 0xbc, 0xde, 0x26, 0xe4, 0x5c, 0xa3, 0xe9, 0xcc, 0x5c, 0xac, 0x60, 0x8b,
	0x84, 0x8e, 0x32, 0x72, 0x89, 0x61, 0x0f, 0x1c, 0x21, 0x1b, 0x3d, 0x5c, 0x71, 0xa7, 0x47, 0x38,
	0x6a, 0xa4, 0x6c, 0xc9, 0xd4, 0x0d, 0xf1, 0xa2, 0x92, 0x5e, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x34,
	0x84, 0xd4, 0xf4, 0x8b, 0x53, 0x33, 0x75, 0x61, 0xe6, 0xe8, 0xc3, 0xcc, 0xd1, 0x07, 0xa5, 0x33,
	0xa4, 0xb0, 0x04, 0xfb, 0x03, 0x7b, 0xa8, 0x11, 0xf4, 0x07, 0xde, 0x38, 0x23, 0xe8, 0x0f, 0xfc,
	0x51, 0x45, 0x94, 0x3f, 0x90, 0xe2, 0xd2, 0xc9, 0xfd, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4,
	0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f,
	0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0x31, 0xcc,
	0xd2, 0x85, 0x18, 0x56, 0x01, 0x36, 0xae, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x2c, 0x6f,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x29, 0xad, 0x26, 0xce, 0x92, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	SeiAddressByEVMAddress(ctx context.Context, in *QuerySeiAddressByEVMAddressRequest, opts ...grpc.CallOption) (*QuerySeiAddressByEVMAddressResponse, error)
	EVMAddressBySeiAddress(ctx context.Context, in *QueryEVMAddressBySeiAddressRequest, opts ...grpc.CallOption) (*QueryEVMAddressBySeiAddressResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) SeiAddressByEVMAddress(ctx context.Context, in *QuerySeiAddressByEVMAddressRequest, opts ...grpc.CallOption) (*QuerySeiAddressByEVMAddressResponse, error) {
	out := new(QuerySeiAddressByEVMAddressResponse)
	err := c.cc.Invoke(ctx, "/seiprotocol.seichain.evm.Query/SeiAddressByEVMAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EVMAddressBySeiAddress(ctx context.Context, in *QueryEVMAddressBySeiAddressRequest, opts ...grpc.CallOption) (*QueryEVMAddressBySeiAddressResponse, error) {
	out := new(QueryEVMAddressBySeiAddressResponse)
	err := c.cc.Invoke(ctx, "/seiprotocol.seichain.evm.Query/EVMAddressBySeiAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	SeiAddressByEVMAddress(context.Context, *QuerySeiAddressByEVMAddressRequest) (*QuerySeiAddressByEVMAddressResponse, error)
	EVMAddressBySeiAddress(context.Context, *QueryEVMAddressBySeiAddressRequest) (*QueryEVMAddressBySeiAddressResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) SeiAddressByEVMAddress(ctx context.Context, req *QuerySeiAddressByEVMAddressRequest) (*QuerySeiAddressByEVMAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SeiAddressByEVMAddress not implemented")
}
func (*UnimplementedQueryServer) EVMAddressBySeiAddress(ctx context.Context, req *QueryEVMAddressBySeiAddressRequest) (*QueryEVMAddressBySeiAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EVMAddressBySeiAddress not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_SeiAddressByEVMAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySeiAddressByEVMAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SeiAddressByEVMAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seiprotocol.seichain.evm.Query/SeiAddressByEVMAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SeiAddressByEVMAddress(ctx, req.(*QuerySeiAddressByEVMAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EVMAddressBySeiAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEVMAddressBySeiAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EVMAddressBySeiAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seiprotocol.seichain.evm.Query/EVMAddressBySeiAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EVMAddressBySeiAddress(ctx, req.(*QueryEVMAddressBySeiAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seiprotocol.seichain.evm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SeiAddressByEVMAddress",
			Handler:    _Query_SeiAddressByEVMAddress_Handler,
		},
		{
			MethodName: "EVMAddressBySeiAddress",
			Handler:    _Query_EVMAddressBySeiAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evm/query.proto",
}

func (m *QuerySeiAddressByEVMAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySeiAddressByEVMAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySeiAddressByEVMAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EvmAddress) > 0 {
		i -= len(m.EvmAddress)
		copy(dAtA[i:], m.EvmAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EvmAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QuerySeiAddressByEVMAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySeiAddressByEVMAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySeiAddressByEVMAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Associated {
		i--
		if m.Associated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.SeiAddress) > 0 {
		i -= len(m.SeiAddress)
		copy(dAtA[i:], m.SeiAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.SeiAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEVMAddressBySeiAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEVMAddressBySeiAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEVMAddressBySeiAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SeiAddress) > 0 {
		i -= len(m.SeiAddress)
		copy(dAtA[i:], m.SeiAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.SeiAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryEVMAddressBySeiAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEVMAddressBySeiAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEVMAddressBySeiAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Associated {
		i--
		if m.Associated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.EvmAddress) > 0 {
		i -= len(m.EvmAddress)
		copy(dAtA[i:], m.EvmAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.EvmAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuerySeiAddressByEVMAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EvmAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QuerySeiAddressByEVMAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SeiAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Associated {
		n += 2
	}
	return n
}

func (m *QueryEVMAddressBySeiAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SeiAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryEVMAddressBySeiAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EvmAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Associated {
		n += 2
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySeiAddressByEVMAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySeiAddressByEVMAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySeiAddressByEVMAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvmAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvmAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySeiAddressByEVMAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySeiAddressByEVMAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySeiAddressByEVMAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeiAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SeiAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Associated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Associated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryEVMAddressBySeiAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryEVMAddressBySeiAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEVMAddressBySeiAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeiAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SeiAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryEVMAddressBySeiAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryEVMAddressBySeiAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEVMAddressBySeiAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvmAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvmAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Associated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Associated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)