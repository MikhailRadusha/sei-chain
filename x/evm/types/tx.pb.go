// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type SignerVersion int32

const (
	SignerVersion_LONDON SignerVersion = 0
	SignerVersion_CANCUN SignerVersion = 1
)

var SignerVersion_name = map[int32]string{
	0: "LONDON",
	1: "CANCUN",
}

var SignerVersion_value = map[string]int32{
	"LONDON": 0,
	"CANCUN": 1,
}

func (x SignerVersion) String() string {
	return proto.EnumName(SignerVersion_name, int32(x))
}

func (SignerVersion) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d72e73a3d1d93781, []int{0}
}

type MsgEVMTransaction struct {
	Data    *types.Any   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Derived *DerivedData `protobuf:"bytes,2,opt,name=derived,proto3" json:"derived,omitempty"`
}

func (m *MsgEVMTransaction) Reset()         { *m = MsgEVMTransaction{} }
func (m *MsgEVMTransaction) String() string { return proto.CompactTextString(m) }
func (*MsgEVMTransaction) ProtoMessage()    {}
func (*MsgEVMTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d72e73a3d1d93781, []int{0}
}
func (m *MsgEVMTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEVMTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEVMTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEVMTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEVMTransaction.Merge(m, src)
}
func (m *MsgEVMTransaction) XXX_Size() int {
	return m.Size()
}
func (m *MsgEVMTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEVMTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEVMTransaction proto.InternalMessageInfo

func (m *MsgEVMTransaction) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MsgEVMTransaction) GetDerived() *DerivedData {
	if m != nil {
		return m.Derived
	}
	return nil
}

type MsgEVMTransactionResponse struct {
	GasUsed    uint64 `protobuf:"varint,1,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	VmError    string `protobuf:"bytes,2,opt,name=vm_error,json=vmError,proto3" json:"vm_error,omitempty"`
	ReturnData []byte `protobuf:"bytes,3,opt,name=return_data,json=returnData,proto3" json:"return_data,omitempty"`
	Hash       string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *MsgEVMTransactionResponse) Reset()         { *m = MsgEVMTransactionResponse{} }
func (m *MsgEVMTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEVMTransactionResponse) ProtoMessage()    {}
func (*MsgEVMTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d72e73a3d1d93781, []int{1}
}
func (m *MsgEVMTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgEVMTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgEVMTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgEVMTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEVMTransactionResponse.Merge(m, src)
}
func (m *MsgEVMTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgEVMTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEVMTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEVMTransactionResponse proto.InternalMessageInfo

func (m *MsgEVMTransactionResponse) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *MsgEVMTransactionResponse) GetVmError() string {
	if m != nil {
		return m.VmError
	}
	return ""
}

func (m *MsgEVMTransactionResponse) GetReturnData() []byte {
	if m != nil {
		return m.ReturnData
	}
	return nil
}

func (m *MsgEVMTransactionResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

// derived from signature
type DerivedData struct {
	SenderEVMAddr []byte        `protobuf:"bytes,1,opt,name=senderEVMAddr,proto3" json:"senderEVMAddr,omitempty"`
	SenderSeiAddr []byte        `protobuf:"bytes,2,opt,name=senderSeiAddr,proto3" json:"senderSeiAddr,omitempty"`
	Pubkey        []byte        `protobuf:"bytes,3,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	IsAssociate   bool          `protobuf:"varint,4,opt,name=isAssociate,proto3" json:"isAssociate,omitempty"`
	Version       SignerVersion `protobuf:"varint,5,opt,name=version,proto3,enum=seiprotocol.seichain.evm.SignerVersion" json:"version,omitempty"`
}

func (m *DerivedData) Reset()         { *m = DerivedData{} }
func (m *DerivedData) String() string { return proto.CompactTextString(m) }
func (*DerivedData) ProtoMessage()    {}
func (*DerivedData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d72e73a3d1d93781, []int{2}
}
func (m *DerivedData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DerivedData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DerivedData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DerivedData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DerivedData.Merge(m, src)
}
func (m *DerivedData) XXX_Size() int {
	return m.Size()
}
func (m *DerivedData) XXX_DiscardUnknown() {
	xxx_messageInfo_DerivedData.DiscardUnknown(m)
}

var xxx_messageInfo_DerivedData proto.InternalMessageInfo

func (m *DerivedData) GetSenderEVMAddr() []byte {
	if m != nil {
		return m.SenderEVMAddr
	}
	return nil
}

func (m *DerivedData) GetSenderSeiAddr() []byte {
	if m != nil {
		return m.SenderSeiAddr
	}
	return nil
}

func (m *DerivedData) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *DerivedData) GetIsAssociate() bool {
	if m != nil {
		return m.IsAssociate
	}
	return false
}

func (m *DerivedData) GetVersion() SignerVersion {
	if m != nil {
		return m.Version
	}
	return SignerVersion_LONDON
}

func init() {
	proto.RegisterEnum("seiprotocol.seichain.evm.SignerVersion", SignerVersion_name, SignerVersion_value)
	proto.RegisterType((*MsgEVMTransaction)(nil), "seiprotocol.seichain.evm.MsgEVMTransaction")
	proto.RegisterType((*MsgEVMTransactionResponse)(nil), "seiprotocol.seichain.evm.MsgEVMTransactionResponse")
	proto.RegisterType((*DerivedData)(nil), "seiprotocol.seichain.evm.DerivedData")
}

func init() { proto.RegisterFile("evm/tx.proto", fileDescriptor_d72e73a3d1d93781) }

var fileDescriptor_d72e73a3d1d93781 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x6f, 0xda, 0x40,
	0x18, 0xc6, 0xb9, 0x84, 0x42, 0xfa, 0x42, 0xa2, 0xf4, 0x54, 0x55, 0x90, 0xc1, 0x45, 0xa8, 0x55,
	0x50, 0xab, 0xd8, 0x12, 0xf9, 0x00, 0x15, 0x0d, 0xa8, 0x4b, 0x21, 0x92, 0xd3, 0x30, 0x74, 0x41,
	0x07, 0x7e, 0x6b, 0x4e, 0x8d, 0xef, 0xd0, 0xbd, 0xb6, 0x15, 0x2f, 0x1d, 0x3b, 0xf7, 0x63, 0x75,
	0xcc, 0x54, 0x75, 0xac, 0xe0, 0x8b, 0x54, 0x9c, 0xb1, 0x4a, 0x1a, 0x31, 0x64, 0x7b, 0xff, 0xfc,
	0x9e, 0x7b, 0x1e, 0xfb, 0x6c, 0xa8, 0x63, 0x1a, 0x79, 0xf1, 0xad, 0xbb, 0x30, 0x3a, 0xd6, 0xbc,
	0x41, 0x28, 0x6d, 0x35, 0xd3, 0x37, 0x2e, 0xa1, 0x9c, 0xcd, 0x85, 0x54, 0x2e, 0xa6, 0xd1, 0x49,
	0x33, 0xd4, 0x3a, 0xbc, 0x41, 0xcf, 0x6e, 0xa7, 0xc9, 0x17, 0x4f, 0xa8, 0x2c, 0x17, 0xb5, 0xbf,
	0xc1, 0xb3, 0x21, 0x85, 0x83, 0xf1, 0xf0, 0x93, 0x11, 0x8a, 0xc4, 0x2c, 0x96, 0x5a, 0xf1, 0x0e,
	0x94, 0x03, 0x11, 0x8b, 0x06, 0x6b, 0xb1, 0x4e, 0xad, 0xfb, 0xdc, 0xcd, 0xe5, 0x6e, 0x21, 0x77,
	0x7b, 0x2a, 0xf3, 0x2d, 0xc1, 0xdf, 0x41, 0x35, 0x40, 0x23, 0x53, 0x0c, 0x1a, 0x7b, 0x16, 0x7e,
	0xed, 0xee, 0x4a, 0xe1, 0xf6, 0x73, 0xb0, 0x2f, 0x62, 0xe1, 0x17, 0xaa, 0xf6, 0x77, 0x06, 0xcd,
	0x07, 0x01, 0x7c, 0xa4, 0x85, 0x56, 0x84, 0xbc, 0x09, 0x07, 0xa1, 0xa0, 0x49, 0x42, 0x18, 0xd8,
	0x30, 0x65, 0xbf, 0x1a, 0x0a, 0xba, 0x26, 0x0c, 0xd6, 0xab, 0x34, 0x9a, 0xa0, 0x31, 0xda, 0x58,
	0xeb, 0xa7, 0x7e, 0x35, 0x8d, 0x06, 0xeb, 0x96, 0xbf, 0x84, 0x9a, 0xc1, 0x38, 0x31, 0x6a, 0x62,
	0x9f, 0x62, 0xbf, 0xc5, 0x3a, 0x75, 0x1f, 0xf2, 0xd1, 0xda, 0x9d, 0x73, 0x28, 0xcf, 0x05, 0xcd,
	0x1b, 0x65, 0xab, 0xb3, 0x75, 0xfb, 0x17, 0x83, 0xda, 0x56, 0x42, 0xfe, 0x0a, 0x0e, 0x09, 0x55,
	0x80, 0x66, 0x30, 0x1e, 0xf6, 0x82, 0xc0, 0x58, 0xff, 0xba, 0x7f, 0x7f, 0xf8, 0x8f, 0xba, 0x42,
	0x69, 0xa9, 0xbd, 0x6d, 0x6a, 0x33, 0xe4, 0x2f, 0xa0, 0xb2, 0x48, 0xa6, 0x5f, 0x31, 0xdb, 0x64,
	0xd9, 0x74, 0xbc, 0x05, 0x35, 0x49, 0x3d, 0x22, 0x3d, 0x93, 0x22, 0x46, 0x1b, 0xe7, 0xc0, 0xdf,
	0x1e, 0xf1, 0x1e, 0x54, 0x53, 0x34, 0x24, 0xb5, 0x6a, 0x3c, 0x69, 0xb1, 0xce, 0x51, 0xf7, 0x74,
	0xf7, 0xfb, 0xbd, 0x92, 0xa1, 0x42, 0x33, 0xce, 0x71, 0xbf, 0xd0, 0xbd, 0x39, 0x85, 0xc3, 0x7b,
	0x1b, 0x0e, 0x50, 0xf9, 0x78, 0x39, 0xea, 0x5f, 0x8e, 0x8e, 0x4b, 0xeb, 0xfa, 0xa2, 0x37, 0xba,
	0xb8, 0x1e, 0x1d, 0xb3, 0x6e, 0x06, 0xfb, 0x43, 0x0a, 0xb9, 0x81, 0xa3, 0xff, 0x3e, 0x87, 0xb7,
	0xbb, 0x3d, 0x1f, 0x5c, 0xdd, 0xc9, 0xf9, 0x23, 0xe0, 0xe2, 0x9e, 0xdf, 0x7f, 0xf8, 0xb9, 0x74,
	0xd8, 0xdd, 0xd2, 0x61, 0x7f, 0x96, 0x0e, 0xfb, 0xb1, 0x72, 0x4a, 0x77, 0x2b, 0xa7, 0xf4, 0x7b,
	0xe5, 0x94, 0x3e, 0x9f, 0x85, 0x32, 0x9e, 0x27, 0x53, 0x77, 0xa6, 0x23, 0x8f, 0x50, 0x9e, 0x15,
	0x27, 0xdb, 0xc6, 0x1e, 0xed, 0xdd, 0x7a, 0xf6, 0x37, 0xc8, 0x16, 0x48, 0xd3, 0x8a, 0xdd, 0x9f,
	0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x2d, 0x23, 0xc4, 0x1a, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	EVMTransaction(ctx context.Context, in *MsgEVMTransaction, opts ...grpc.CallOption) (*MsgEVMTransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) EVMTransaction(ctx context.Context, in *MsgEVMTransaction, opts ...grpc.CallOption) (*MsgEVMTransactionResponse, error) {
	out := new(MsgEVMTransactionResponse)
	err := c.cc.Invoke(ctx, "/seiprotocol.seichain.evm.Msg/EVMTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	EVMTransaction(context.Context, *MsgEVMTransaction) (*MsgEVMTransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) EVMTransaction(ctx context.Context, req *MsgEVMTransaction) (*MsgEVMTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EVMTransaction not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_EVMTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEVMTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EVMTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seiprotocol.seichain.evm.Msg/EVMTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EVMTransaction(ctx, req.(*MsgEVMTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "seiprotocol.seichain.evm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EVMTransaction",
			Handler:    _Msg_EVMTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evm/tx.proto",
}

func (m *MsgEVMTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEVMTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEVMTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Derived != nil {
		{
			size, err := m.Derived.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgEVMTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgEVMTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgEVMTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ReturnData) > 0 {
		i -= len(m.ReturnData)
		copy(dAtA[i:], m.ReturnData)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ReturnData)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VmError) > 0 {
		i -= len(m.VmError)
		copy(dAtA[i:], m.VmError)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VmError)))
		i--
		dAtA[i] = 0x12
	}
	if m.GasUsed != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.GasUsed))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DerivedData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DerivedData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DerivedData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x28
	}
	if m.IsAssociate {
		i--
		if m.IsAssociate {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SenderSeiAddr) > 0 {
		i -= len(m.SenderSeiAddr)
		copy(dAtA[i:], m.SenderSeiAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SenderSeiAddr)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SenderEVMAddr) > 0 {
		i -= len(m.SenderEVMAddr)
		copy(dAtA[i:], m.SenderEVMAddr)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SenderEVMAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgEVMTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Derived != nil {
		l = m.Derived.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgEVMTransactionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GasUsed != 0 {
		n += 1 + sovTx(uint64(m.GasUsed))
	}
	l = len(m.VmError)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ReturnData)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *DerivedData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SenderEVMAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SenderSeiAddr)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.IsAssociate {
		n += 2
	}
	if m.Version != 0 {
		n += 1 + sovTx(uint64(m.Version))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgEVMTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEVMTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEVMTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &types.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Derived", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Derived == nil {
				m.Derived = &DerivedData{}
			}
			if err := m.Derived.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgEVMTransactionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgEVMTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgEVMTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasUsed", wireType)
			}
			m.GasUsed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GasUsed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VmError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VmError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReturnData = append(m.ReturnData[:0], dAtA[iNdEx:postIndex]...)
			if m.ReturnData == nil {
				m.ReturnData = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *DerivedData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: DerivedData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DerivedData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderEVMAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderEVMAddr = append(m.SenderEVMAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.SenderEVMAddr == nil {
				m.SenderEVMAddr = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SenderSeiAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SenderSeiAddr = append(m.SenderSeiAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.SenderSeiAddr == nil {
				m.SenderSeiAddr = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsAssociate", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.IsAssociate = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= SignerVersion(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)