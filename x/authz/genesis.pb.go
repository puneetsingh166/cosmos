// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/authz/v1beta1/genesis.proto

package authz

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/puneetsingh166/tm-load-test/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GenesisState defines the authz module's genesis state.
type GenesisState struct {
	Authorization []GrantAuthorization `protobuf:"bytes,1,rep,name=authorization,proto3" json:"authorization"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2fbb971da7c892, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAuthorization() []GrantAuthorization {
	if m != nil {
		return m.Authorization
	}
	return nil
}

// GrantAuthorization defines the GenesisState/GrantAuthorization type.
type GrantAuthorization struct {
	Granter       string     `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee       string     `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	Authorization *types.Any `protobuf:"bytes,3,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Expiration    time.Time  `protobuf:"bytes,4,opt,name=expiration,proto3,stdtime" json:"expiration"`
}

func (m *GrantAuthorization) Reset()         { *m = GrantAuthorization{} }
func (m *GrantAuthorization) String() string { return proto.CompactTextString(m) }
func (*GrantAuthorization) ProtoMessage()    {}
func (*GrantAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c2fbb971da7c892, []int{1}
}
func (m *GrantAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrantAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrantAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrantAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrantAuthorization.Merge(m, src)
}
func (m *GrantAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *GrantAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_GrantAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_GrantAuthorization proto.InternalMessageInfo

func (m *GrantAuthorization) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *GrantAuthorization) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *GrantAuthorization) GetAuthorization() *types.Any {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func (m *GrantAuthorization) GetExpiration() time.Time {
	if m != nil {
		return m.Expiration
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.authz.v1beta1.GenesisState")
	proto.RegisterType((*GrantAuthorization)(nil), "cosmos.authz.v1beta1.GrantAuthorization")
}

func init() {
	proto.RegisterFile("cosmos/authz/v1beta1/genesis.proto", fileDescriptor_4c2fbb971da7c892)
}

var fileDescriptor_4c2fbb971da7c892 = []byte{
	// 358 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbd, 0x6e, 0xea, 0x30,
	0x1c, 0xc5, 0x63, 0x40, 0xf7, 0xc3, 0x5c, 0x86, 0x1b, 0x31, 0xa4, 0x0c, 0x01, 0xa1, 0x0e, 0x59,
	0xb0, 0x05, 0xdd, 0x2b, 0x11, 0x55, 0x62, 0xea, 0x02, 0x4c, 0x5d, 0x2a, 0x87, 0xb8, 0xc6, 0x6a,
	0x13, 0x23, 0xdb, 0x54, 0xc0, 0x53, 0xd0, 0x77, 0xe1, 0x21, 0x50, 0x27, 0xd4, 0xa9, 0x53, 0x5b,
	0xc1, 0x8b, 0x54, 0x89, 0x13, 0x95, 0x8f, 0x0e, 0x9d, 0xe2, 0xc4, 0xbf, 0x73, 0xfe, 0xc7, 0xc7,
	0x81, 0xcd, 0x91, 0x50, 0x91, 0x50, 0x98, 0x4c, 0xf5, 0x78, 0x81, 0x1f, 0xdb, 0x01, 0xd5, 0xa4,
	0x8d, 0x19, 0x8d, 0xa9, 0xe2, 0x0a, 0x4d, 0xa4, 0xd0, 0xc2, 0xae, 0x1a, 0x06, 0xa5, 0x0c, 0xca,
	0x98, 0x5a, 0x9d, 0x09, 0xc1, 0x1e, 0x28, 0x4e, 0x99, 0x60, 0x7a, 0x87, 0x35, 0x8f, 0xa8, 0xd2,
	0x24, 0x9a, 0x18, 0x59, 0xed, 0xec, 0x18, 0x20, 0xf1, 0x3c, 0xdb, 0xaa, 0x32, 0xc1, 0x44, 0xba,
	0xc4, 0xc9, 0x2a, 0x17, 0x98, 0x39, 0xb7, 0x66, 0x23, 0x1b, 0x9a, 0xbe, 0x34, 0x43, 0xf8, 0xaf,
	0x67, 0x32, 0x0d, 0x34, 0xd1, 0xd4, 0x1e, 0xc2, 0x4a, 0x92, 0x46, 0x48, 0xbe, 0x20, 0x9a, 0x8b,
	0xd8, 0x01, 0x8d, 0xa2, 0x57, 0xee, 0x78, 0xe8, 0xbb, 0xa8, 0xa8, 0x27, 0x49, 0xac, 0xbb, 0xfb,
	0xbc, 0x5f, 0x5a, 0xbf, 0xd5, 0xad, 0xfe, 0xa1, 0x49, 0xf3, 0xa9, 0x00, 0xed, 0x53, 0xd6, 0xee,
	0xc0, 0xdf, 0x2c, 0xf9, 0x4a, 0xa5, 0x03, 0x1a, 0xc0, 0xfb, 0xeb, 0x3b, 0x2f, 0xab, 0x56, 0x5e,
	0x4a, 0x37, 0x0c, 0x25, 0x55, 0x6a, 0xa0, 0x25, 0x8f, 0x59, 0x3f, 0x07, 0xbf, 0x34, 0xd4, 0x29,
	0xfc, 0x4c, 0x43, 0xed, 0xeb, 0xe3, 0x43, 0x15, 0x1b, 0xc0, 0x2b, 0x77, 0xaa, 0xc8, 0x14, 0x89,
	0xf2, 0x22, 0x51, 0x37, 0x9e, 0xfb, 0xff, 0x9f, 0x57, 0xad, 0xca, 0x41, 0xce, 0xa3, 0xd3, 0xd8,
	0x57, 0x10, 0xd2, 0xd9, 0x84, 0x4b, 0xe3, 0x55, 0x4a, 0xbd, 0x6a, 0x27, 0x5e, 0xc3, 0xfc, 0xd6,
	0xfc, 0x3f, 0x49, 0x25, 0xcb, 0xf7, 0x3a, 0xe8, 0xef, 0xe9, 0xfc, 0xcb, 0xf5, 0xd6, 0x05, 0x9b,
	0xad, 0x0b, 0x3e, 0xb6, 0x2e, 0x58, 0xee, 0x5c, 0x6b, 0xb3, 0x73, 0xad, 0xd7, 0x9d, 0x6b, 0xdd,
	0x9c, 0x33, 0xae, 0xc7, 0xd3, 0x00, 0x8d, 0x44, 0x94, 0x5d, 0x56, 0xf6, 0x68, 0xa9, 0xf0, 0x1e,
	0xcf, 0xcc, 0x2f, 0x15, 0xfc, 0x4a, 0x27, 0x5d, 0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xaf,
	0x1e, 0x75, 0x69, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authorization) > 0 {
		for iNdEx := len(m.Authorization) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Authorization[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GrantAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrantAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrantAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Authorization) > 0 {
		for _, e := range m.Authorization {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *GrantAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authorization = append(m.Authorization, GrantAuthorization{})
			if err := m.Authorization[len(m.Authorization)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GrantAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GrantAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrantAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &types.Any{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
