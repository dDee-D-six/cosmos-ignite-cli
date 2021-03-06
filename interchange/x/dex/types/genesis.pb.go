// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the dex module's genesis state.
type GenesisState struct {
	Params            Params          `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId            string          `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	SellOrderBookList []SellOrderBook `protobuf:"bytes,3,rep,name=sellOrderBookList,proto3" json:"sellOrderBookList"`
	BuyOrderBookList  []BuyOrderBook  `protobuf:"bytes,4,rep,name=buyOrderBookList,proto3" json:"buyOrderBookList"`
	DenomTraceList    []DenomTrace    `protobuf:"bytes,5,rep,name=denomTraceList,proto3" json:"denomTraceList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a803aaabd08db59d, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetSellOrderBookList() []SellOrderBook {
	if m != nil {
		return m.SellOrderBookList
	}
	return nil
}

func (m *GenesisState) GetBuyOrderBookList() []BuyOrderBook {
	if m != nil {
		return m.BuyOrderBookList
	}
	return nil
}

func (m *GenesisState) GetDenomTraceList() []DenomTrace {
	if m != nil {
		return m.DenomTraceList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "username.interchange.dex.GenesisState")
}

func init() { proto.RegisterFile("dex/genesis.proto", fileDescriptor_a803aaabd08db59d) }

var fileDescriptor_a803aaabd08db59d = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x93, 0xb6, 0x56, 0xdc, 0x8a, 0xb4, 0x41, 0x31, 0xf6, 0x10, 0x83, 0x88, 0x16, 0x84,
	0x04, 0xea, 0xdd, 0x43, 0x50, 0x44, 0x10, 0x94, 0xd6, 0x83, 0xe8, 0x21, 0x24, 0xcd, 0x90, 0x86,
	0x36, 0x99, 0xb0, 0xd9, 0x40, 0xfb, 0x16, 0x3e, 0x56, 0x8f, 0x3d, 0x7a, 0x12, 0x69, 0xdf, 0xc0,
	0x27, 0x90, 0xdd, 0xac, 0x52, 0x0d, 0xb9, 0xed, 0xce, 0xff, 0xcf, 0x37, 0x3f, 0x33, 0xa4, 0x13,
	0xc0, 0xcc, 0x0e, 0x21, 0x81, 0x2c, 0xca, 0xac, 0x94, 0x22, 0x43, 0x4d, 0xcf, 0x33, 0xa0, 0x89,
	0x17, 0x83, 0x15, 0x25, 0x0c, 0xe8, 0x68, 0xec, 0x25, 0x21, 0x58, 0x01, 0xcc, 0xba, 0xfb, 0x21,
	0x86, 0x28, 0x4c, 0x36, 0x7f, 0x15, 0xfe, 0x6e, 0x9b, 0x23, 0x52, 0x8f, 0x7a, 0xb1, 0x24, 0x74,
	0x8f, 0x78, 0x25, 0x83, 0xe9, 0xd4, 0x45, 0x1a, 0x00, 0x75, 0x7d, 0xc4, 0x89, 0x94, 0x74, 0x2e,
	0xf9, 0xf9, 0xbc, 0xac, 0x1c, 0x70, 0x25, 0x80, 0x04, 0x63, 0x97, 0x51, 0x6f, 0x04, 0x45, 0xf9,
	0xe4, 0xab, 0x46, 0x76, 0x6f, 0x8b, 0x7c, 0x43, 0xe6, 0x31, 0xd0, 0xae, 0x48, 0xb3, 0x18, 0xa6,
	0xab, 0xa6, 0xda, 0x6b, 0xf5, 0x4d, 0xab, 0x2a, 0xaf, 0xf5, 0x28, 0x7c, 0x4e, 0x63, 0xf1, 0x71,
	0xac, 0x0c, 0x64, 0x97, 0x76, 0x48, 0xb6, 0x53, 0xa4, 0xcc, 0x8d, 0x02, 0xbd, 0x66, 0xaa, 0xbd,
	0x9d, 0x41, 0x93, 0x7f, 0xef, 0x02, 0xed, 0x95, 0x74, 0x78, 0xe6, 0x07, 0x1e, 0xcc, 0x41, 0x9c,
	0xdc, 0x47, 0x19, 0xd3, 0xeb, 0x66, 0xbd, 0xd7, 0xea, 0x9f, 0x57, 0xcf, 0x18, 0x6e, 0xb6, 0xc8,
	0x51, 0x65, 0x8e, 0xf6, 0x4c, 0xda, 0x7e, 0x3e, 0xff, 0xcb, 0x6e, 0x08, 0xf6, 0x59, 0x35, 0xdb,
	0xd9, 0xe8, 0x90, 0xe8, 0x12, 0x45, 0x1b, 0x90, 0x3d, 0xb1, 0xb5, 0x27, 0xbe, 0x34, 0xc1, 0xdd,
	0x12, 0xdc, 0xd3, 0x6a, 0xee, 0xf5, 0xaf, 0x5f, 0x52, 0xff, 0x11, 0x9c, 0x9b, 0xc5, 0xca, 0x50,
	0x97, 0x2b, 0x43, 0xfd, 0x5c, 0x19, 0xea, 0xdb, 0xda, 0x50, 0x96, 0x6b, 0x43, 0x79, 0x5f, 0x1b,
	0xca, 0xcb, 0x45, 0x18, 0xb1, 0x71, 0xee, 0x5b, 0x23, 0x8c, 0xed, 0x1f, 0xbe, 0xbd, 0xc1, 0xb7,
	0xf9, 0x0d, 0x67, 0x36, 0x9b, 0xa7, 0x90, 0xf9, 0x4d, 0x71, 0xc2, 0xcb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc5, 0xe8, 0x58, 0x67, 0x65, 0x02, 0x00, 0x00,
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
	if len(m.DenomTraceList) > 0 {
		for iNdEx := len(m.DenomTraceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomTraceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.BuyOrderBookList) > 0 {
		for iNdEx := len(m.BuyOrderBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BuyOrderBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.SellOrderBookList) > 0 {
		for iNdEx := len(m.SellOrderBookList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SellOrderBookList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.SellOrderBookList) > 0 {
		for _, e := range m.SellOrderBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.BuyOrderBookList) > 0 {
		for _, e := range m.BuyOrderBookList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DenomTraceList) > 0 {
		for _, e := range m.DenomTraceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
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
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOrderBookList", wireType)
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
			m.SellOrderBookList = append(m.SellOrderBookList, SellOrderBook{})
			if err := m.SellOrderBookList[len(m.SellOrderBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuyOrderBookList", wireType)
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
			m.BuyOrderBookList = append(m.BuyOrderBookList, BuyOrderBook{})
			if err := m.BuyOrderBookList[len(m.BuyOrderBookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomTraceList", wireType)
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
			m.DenomTraceList = append(m.DenomTraceList, DenomTrace{})
			if err := m.DenomTraceList[len(m.DenomTraceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
