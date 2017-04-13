// Code generated by protoc-gen-gogo.
// source: silence.proto
// DO NOT EDIT!

/*
	Package silencepb is a generated protocol buffer package.

	It is generated from these files:
		silence.proto

	It has these top-level messages:
		Matcher
		Comment
		Silence
		MeshSilence
*/
package silencepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import time "time"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Type specifies how the given name and pattern are matched
// against a label set.
type Matcher_Type int32

const (
	Matcher_EQUAL  Matcher_Type = 0
	Matcher_REGEXP Matcher_Type = 1
)

var Matcher_Type_name = map[int32]string{
	0: "EQUAL",
	1: "REGEXP",
}
var Matcher_Type_value = map[string]int32{
	"EQUAL":  0,
	"REGEXP": 1,
}

func (x Matcher_Type) String() string {
	return proto.EnumName(Matcher_Type_name, int32(x))
}
func (Matcher_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorSilence, []int{0, 0} }

// Matcher specifies a rule, which can match or set of labels or not.
type Matcher struct {
	Type Matcher_Type `protobuf:"varint,1,opt,name=type,proto3,enum=silencepb.Matcher_Type" json:"type,omitempty"`
	// The label name in a label set to against which the matcher
	// checks the pattern.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The pattern being checked according to the matcher's type.
	Pattern string `protobuf:"bytes,3,opt,name=pattern,proto3" json:"pattern,omitempty"`
}

func (m *Matcher) Reset()                    { *m = Matcher{} }
func (m *Matcher) String() string            { return proto.CompactTextString(m) }
func (*Matcher) ProtoMessage()               {}
func (*Matcher) Descriptor() ([]byte, []int) { return fileDescriptorSilence, []int{0} }

// A comment can be attached to a silence.
type Comment struct {
	Author    string    `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Comment   string    `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	Timestamp time.Time `protobuf:"bytes,3,opt,name=timestamp,stdtime" json:"timestamp"`
}

func (m *Comment) Reset()                    { *m = Comment{} }
func (m *Comment) String() string            { return proto.CompactTextString(m) }
func (*Comment) ProtoMessage()               {}
func (*Comment) Descriptor() ([]byte, []int) { return fileDescriptorSilence, []int{1} }

// Silence specifies an object that ignores alerts based
// on a set of matchers during a given time frame.
type Silence struct {
	// A globally unique identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A set of matchers all of which have to be true for a silence
	// to affect a given label set.
	Matchers []*Matcher `protobuf:"bytes,2,rep,name=matchers" json:"matchers,omitempty"`
	// The time range during which the silence is active.
	StartsAt time.Time `protobuf:"bytes,3,opt,name=starts_at,json=startsAt,stdtime" json:"starts_at"`
	EndsAt   time.Time `protobuf:"bytes,4,opt,name=ends_at,json=endsAt,stdtime" json:"ends_at"`
	// The last motification made to the silence.
	UpdatedAt time.Time `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at"`
	// A set of comments made on the silence.
	Comments []*Comment `protobuf:"bytes,7,rep,name=comments" json:"comments,omitempty"`
}

func (m *Silence) Reset()                    { *m = Silence{} }
func (m *Silence) String() string            { return proto.CompactTextString(m) }
func (*Silence) ProtoMessage()               {}
func (*Silence) Descriptor() ([]byte, []int) { return fileDescriptorSilence, []int{2} }

// MeshSilence wraps a regular silence with an expiration timestamp
// after which the silence may be garbage collected.
type MeshSilence struct {
	Silence   *Silence  `protobuf:"bytes,1,opt,name=silence" json:"silence,omitempty"`
	ExpiresAt time.Time `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt,stdtime" json:"expires_at"`
}

func (m *MeshSilence) Reset()                    { *m = MeshSilence{} }
func (m *MeshSilence) String() string            { return proto.CompactTextString(m) }
func (*MeshSilence) ProtoMessage()               {}
func (*MeshSilence) Descriptor() ([]byte, []int) { return fileDescriptorSilence, []int{3} }

func init() {
	proto.RegisterType((*Matcher)(nil), "silencepb.Matcher")
	proto.RegisterType((*Comment)(nil), "silencepb.Comment")
	proto.RegisterType((*Silence)(nil), "silencepb.Silence")
	proto.RegisterType((*MeshSilence)(nil), "silencepb.MeshSilence")
	proto.RegisterEnum("silencepb.Matcher_Type", Matcher_Type_name, Matcher_Type_value)
}
func (m *Matcher) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Matcher) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSilence(dAtA, i, uint64(m.Type))
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if len(m.Pattern) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Pattern)))
		i += copy(dAtA[i:], m.Pattern)
	}
	return i, nil
}

func (m *Comment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Comment) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Author) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Author)))
		i += copy(dAtA[i:], m.Author)
	}
	if len(m.Comment) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Comment)))
		i += copy(dAtA[i:], m.Comment)
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSilence(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)))
	n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *Silence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Silence) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSilence(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if len(m.Matchers) > 0 {
		for _, msg := range m.Matchers {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSilence(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSilence(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.StartsAt)))
	n2, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartsAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x22
	i++
	i = encodeVarintSilence(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.EndsAt)))
	n3, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndsAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x2a
	i++
	i = encodeVarintSilence(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)))
	n4, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Comments) > 0 {
		for _, msg := range m.Comments {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintSilence(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *MeshSilence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MeshSilence) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Silence != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSilence(dAtA, i, uint64(m.Silence.Size()))
		n5, err := m.Silence.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintSilence(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiresAt)))
	n6, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.ExpiresAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	return i, nil
}

func encodeFixed64Silence(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Silence(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSilence(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Matcher) Size() (n int) {
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovSilence(uint64(m.Type))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	l = len(m.Pattern)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	return n
}

func (m *Comment) Size() (n int) {
	var l int
	_ = l
	l = len(m.Author)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	l = len(m.Comment)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovSilence(uint64(l))
	return n
}

func (m *Silence) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSilence(uint64(l))
	}
	if len(m.Matchers) > 0 {
		for _, e := range m.Matchers {
			l = e.Size()
			n += 1 + l + sovSilence(uint64(l))
		}
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartsAt)
	n += 1 + l + sovSilence(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndsAt)
	n += 1 + l + sovSilence(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovSilence(uint64(l))
	if len(m.Comments) > 0 {
		for _, e := range m.Comments {
			l = e.Size()
			n += 1 + l + sovSilence(uint64(l))
		}
	}
	return n
}

func (m *MeshSilence) Size() (n int) {
	var l int
	_ = l
	if m.Silence != nil {
		l = m.Silence.Size()
		n += 1 + l + sovSilence(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.ExpiresAt)
	n += 1 + l + sovSilence(uint64(l))
	return n
}

func sovSilence(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSilence(x uint64) (n int) {
	return sovSilence(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Matcher) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilence
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Matcher: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Matcher: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Matcher_Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pattern", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pattern = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSilence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilence
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
func (m *Comment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilence
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Comment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Comment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Author", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Author = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSilence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilence
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
func (m *Silence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilence
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Silence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Silence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Matchers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Matchers = append(m.Matchers, &Matcher{})
			if err := m.Matchers[len(m.Matchers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartsAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartsAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndsAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndsAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Comments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Comments = append(m.Comments, &Comment{})
			if err := m.Comments[len(m.Comments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSilence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilence
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
func (m *MeshSilence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSilence
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MeshSilence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MeshSilence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Silence", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Silence == nil {
				m.Silence = &Silence{}
			}
			if err := m.Silence.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSilence
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.ExpiresAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSilence(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSilence
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
func skipSilence(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSilence
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
					return 0, ErrIntOverflowSilence
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSilence
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSilence
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSilence
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSilence(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSilence = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSilence   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("silence.proto", fileDescriptorSilence) }

var fileDescriptorSilence = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x3b, 0xb9, 0xbd, 0x49, 0x73, 0x8a, 0x97, 0x32, 0x88, 0x86, 0x82, 0x69, 0xc9, 0xaa,
	0xa0, 0x4c, 0xa1, 0xae, 0x5d, 0xa4, 0x97, 0xe2, 0xc6, 0x0b, 0x3a, 0x5e, 0xc1, 0x9d, 0x4c, 0x9b,
	0x31, 0x0d, 0xdc, 0x64, 0x86, 0xe4, 0x14, 0xbc, 0x2b, 0x05, 0x5f, 0xc0, 0x47, 0x72, 0xd9, 0xa5,
	0x4f, 0xe0, 0x9f, 0x3e, 0x89, 0x64, 0x32, 0x89, 0x48, 0x57, 0xd9, 0x9d, 0x33, 0xf9, 0x7d, 0xe7,
	0x9c, 0xef, 0x0b, 0x3c, 0xa8, 0xb2, 0x3b, 0x59, 0xec, 0x24, 0xd3, 0xa5, 0x42, 0x45, 0x7d, 0xdb,
	0xea, 0xed, 0x74, 0x96, 0x2a, 0x95, 0xde, 0xc9, 0xa5, 0xf9, 0xb0, 0x3d, 0x7c, 0x5c, 0x62, 0x96,
	0xcb, 0x0a, 0x45, 0xae, 0x1b, 0x76, 0xfa, 0x30, 0x55, 0xa9, 0x32, 0xe5, 0xb2, 0xae, 0x9a, 0xd7,
	0xe8, 0x2b, 0x01, 0xef, 0x46, 0xe0, 0x6e, 0x2f, 0x4b, 0xfa, 0x14, 0x86, 0x78, 0xaf, 0x65, 0x40,
	0xe6, 0x64, 0x71, 0xb5, 0x7a, 0xcc, 0xba, 0xe1, 0xcc, 0x12, 0xec, 0xf6, 0x5e, 0x4b, 0x6e, 0x20,
	0x4a, 0x61, 0x58, 0x88, 0x5c, 0x06, 0xce, 0x9c, 0x2c, 0x7c, 0x6e, 0x6a, 0x1a, 0x80, 0xa7, 0x05,
	0xa2, 0x2c, 0x8b, 0xe0, 0xc2, 0x3c, 0xb7, 0x6d, 0xf4, 0x04, 0x86, 0xb5, 0x96, 0xfa, 0x70, 0xb9,
	0x79, 0xf3, 0x2e, 0x7e, 0x35, 0x19, 0x50, 0x00, 0x97, 0x6f, 0x5e, 0x6e, 0xde, 0xbf, 0x9e, 0x90,
	0xe8, 0x33, 0x78, 0xd7, 0x2a, 0xcf, 0x65, 0x81, 0xf4, 0x11, 0xb8, 0xe2, 0x80, 0x7b, 0x55, 0x9a,
	0x33, 0x7c, 0x6e, 0xbb, 0x7a, 0xf6, 0xae, 0x41, 0xec, 0xca, 0xb6, 0xa5, 0x6b, 0xf0, 0x3b, 0xaf,
	0x66, 0xef, 0x78, 0x35, 0x65, 0x4d, 0x1a, 0xac, 0x4d, 0x83, 0xdd, 0xb6, 0xc4, 0x7a, 0x74, 0xfc,
	0x39, 0x1b, 0x7c, 0xfb, 0x35, 0x23, 0xfc, 0x9f, 0x2c, 0xfa, 0xee, 0x80, 0xf7, 0xb6, 0xb1, 0x4b,
	0xaf, 0xc0, 0xc9, 0x12, 0xbb, 0xdd, 0xc9, 0x12, 0xca, 0x60, 0x94, 0x37, 0xfe, 0xab, 0xc0, 0x99,
	0x5f, 0x2c, 0xc6, 0x2b, 0x7a, 0x1e, 0x0d, 0xef, 0x18, 0x1a, 0x83, 0x5f, 0xa1, 0x28, 0xb1, 0xfa,
	0x20, 0xb0, 0xd7, 0x3d, 0xa3, 0x46, 0x16, 0x23, 0x7d, 0x01, 0x9e, 0x2c, 0x12, 0x33, 0x60, 0xd8,
	0x63, 0x80, 0x5b, 0x8b, 0x62, 0xa4, 0xd7, 0x00, 0x07, 0x9d, 0x08, 0x94, 0x49, 0x3d, 0xe1, 0xb2,
	0x4f, 0x24, 0x56, 0x17, 0x63, 0x6d, 0xdb, 0x26, 0x5c, 0x05, 0xde, 0x99, 0x6d, 0xfb, 0xbb, 0x78,
	0xc7, 0x44, 0x5f, 0x08, 0x8c, 0x6f, 0x64, 0xb5, 0x6f, 0x63, 0x7c, 0x06, 0x9e, 0xc5, 0x4d, 0x96,
	0xff, 0xcb, 0x2d, 0xc4, 0x5b, 0xa4, 0x3e, 0x59, 0x7e, 0xd2, 0x59, 0x29, 0x8d, 0x69, 0xa7, 0xcf,
	0xc9, 0x56, 0x17, 0xe3, 0x7a, 0x72, 0xfc, 0x13, 0x0e, 0x8e, 0xa7, 0x90, 0xfc, 0x38, 0x85, 0xe4,
	0xf7, 0x29, 0x24, 0x5b, 0xd7, 0x48, 0x9f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x2e, 0x01,
	0x5f, 0x38, 0x03, 0x00, 0x00,
}
