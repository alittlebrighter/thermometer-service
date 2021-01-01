// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: messages/temperature.proto

package messages

import (
	encoding_binary "encoding/binary"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TemperatureUnits int32

const (
	celsius    TemperatureUnits = 0
	fahrenheit TemperatureUnits = 1
)

var TemperatureUnits_name = map[int32]string{
	0: "celsius",
	1: "fahrenheit",
}

var TemperatureUnits_value = map[string]int32{
	"celsius":    0,
	"fahrenheit": 1,
}

func (TemperatureUnits) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f681de90b539975a, []int{0}
}

type TemperatureRead struct {
	// actor.PID Sender = 1; // this is the PID the remote actor should reply to
	//
	// Types that are valid to be assigned to Result:
	//	*TemperatureRead_Value
	//	*TemperatureRead_Error
	Result   isTemperatureRead_Result `protobuf_oneof:"result"`
	Units    TemperatureUnits         `protobuf:"varint,3,opt,name=units,proto3,enum=messages.TemperatureUnits" json:"units,omitempty"`
	Location string                   `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
}

func (m *TemperatureRead) Reset()      { *m = TemperatureRead{} }
func (*TemperatureRead) ProtoMessage() {}
func (*TemperatureRead) Descriptor() ([]byte, []int) {
	return fileDescriptor_f681de90b539975a, []int{0}
}
func (m *TemperatureRead) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TemperatureRead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TemperatureRead.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TemperatureRead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemperatureRead.Merge(m, src)
}
func (m *TemperatureRead) XXX_Size() int {
	return m.Size()
}
func (m *TemperatureRead) XXX_DiscardUnknown() {
	xxx_messageInfo_TemperatureRead.DiscardUnknown(m)
}

var xxx_messageInfo_TemperatureRead proto.InternalMessageInfo

type isTemperatureRead_Result interface {
	isTemperatureRead_Result()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type TemperatureRead_Value struct {
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3,oneof"`
}
type TemperatureRead_Error struct {
	Error string `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

func (*TemperatureRead_Value) isTemperatureRead_Result() {}
func (*TemperatureRead_Error) isTemperatureRead_Result() {}

func (m *TemperatureRead) GetResult() isTemperatureRead_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *TemperatureRead) GetValue() float64 {
	if x, ok := m.GetResult().(*TemperatureRead_Value); ok {
		return x.Value
	}
	return 0
}

func (m *TemperatureRead) GetError() string {
	if x, ok := m.GetResult().(*TemperatureRead_Error); ok {
		return x.Error
	}
	return ""
}

func (m *TemperatureRead) GetUnits() TemperatureUnits {
	if m != nil {
		return m.Units
	}
	return celsius
}

func (m *TemperatureRead) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TemperatureRead) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TemperatureRead_OneofMarshaler, _TemperatureRead_OneofUnmarshaler, _TemperatureRead_OneofSizer, []interface{}{
		(*TemperatureRead_Value)(nil),
		(*TemperatureRead_Error)(nil),
	}
}

func _TemperatureRead_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TemperatureRead)
	// result
	switch x := m.Result.(type) {
	case *TemperatureRead_Value:
		_ = b.EncodeVarint(1<<3 | proto.WireFixed64)
		_ = b.EncodeFixed64(math.Float64bits(x.Value))
	case *TemperatureRead_Error:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Error)
	case nil:
	default:
		return fmt.Errorf("TemperatureRead.Result has unexpected type %T", x)
	}
	return nil
}

func _TemperatureRead_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TemperatureRead)
	switch tag {
	case 1: // result.value
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.Result = &TemperatureRead_Value{math.Float64frombits(x)}
		return true, err
	case 2: // result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Result = &TemperatureRead_Error{x}
		return true, err
	default:
		return false, nil
	}
}

func _TemperatureRead_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TemperatureRead)
	// result
	switch x := m.Result.(type) {
	case *TemperatureRead_Value:
		n += 1 // tag and wire
		n += 8
	case *TemperatureRead_Error:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Error)))
		n += len(x.Error)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterEnum("messages.TemperatureUnits", TemperatureUnits_name, TemperatureUnits_value)
	proto.RegisterType((*TemperatureRead)(nil), "messages.TemperatureRead")
}

func init() { proto.RegisterFile("messages/temperature.proto", fileDescriptor_f681de90b539975a) }

var fileDescriptor_f681de90b539975a = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3d, 0x4e, 0xc3, 0x40,
	0x10, 0x85, 0x77, 0x80, 0x84, 0x30, 0x48, 0xc1, 0xda, 0x02, 0x59, 0x2e, 0x46, 0x16, 0x95, 0x45,
	0xe1, 0x20, 0xe0, 0x04, 0xa9, 0xa8, 0x2d, 0x38, 0xc0, 0x12, 0x06, 0x62, 0xc9, 0xf1, 0x46, 0xfb,
	0x43, 0xcd, 0x11, 0x68, 0xb9, 0x01, 0x47, 0xa1, 0x74, 0x99, 0x12, 0xaf, 0x1b, 0xca, 0x1c, 0x01,
	0x39, 0x51, 0x00, 0x51, 0x7e, 0xef, 0x3d, 0x7d, 0xc5, 0xc3, 0x64, 0xc1, 0xd6, 0xaa, 0x27, 0xb6,
	0x13, 0xc7, 0x8b, 0x25, 0x1b, 0xe5, 0xbc, 0xe1, 0x7c, 0x69, 0xb4, 0xd3, 0x72, 0xb4, 0xeb, 0xce,
	0xde, 0x00, 0x4f, 0x6e, 0x7f, 0xfb, 0x82, 0xd5, 0x83, 0x3c, 0xc5, 0xc1, 0xb3, 0xaa, 0x3c, 0xc7,
	0x90, 0x42, 0x06, 0x37, 0xa2, 0xd8, 0x62, 0x9f, 0xb3, 0x31, 0xda, 0xc4, 0x7b, 0x29, 0x64, 0x47,
	0x7d, 0xbe, 0x41, 0x79, 0x81, 0x03, 0x5f, 0x97, 0xce, 0xc6, 0xfb, 0x29, 0x64, 0xe3, 0xcb, 0x24,
	0xdf, 0xd9, 0xf3, 0x3f, 0xe6, 0xbb, 0x7e, 0x51, 0x6c, 0x87, 0x32, 0xc1, 0x51, 0xa5, 0x67, 0xca,
	0x95, 0xba, 0x8e, 0x0f, 0x7a, 0x59, 0xf1, 0xc3, 0xd3, 0x11, 0x0e, 0x0d, 0x5b, 0x5f, 0xb9, 0xf3,
	0x09, 0x46, 0xff, 0x05, 0xf2, 0x18, 0x0f, 0x67, 0x5c, 0xd9, 0xd2, 0xdb, 0x48, 0xc8, 0x31, 0xe2,
	0xa3, 0x9a, 0x1b, 0xae, 0xe7, 0x5c, 0xba, 0x08, 0xa6, 0xd7, 0x4d, 0x4b, 0x62, 0xd5, 0x92, 0x58,
	0xb7, 0x04, 0x2f, 0x81, 0xe0, 0x3d, 0x10, 0x7c, 0x04, 0x82, 0x26, 0x10, 0x7c, 0x06, 0x82, 0xaf,
	0x40, 0x62, 0x1d, 0x08, 0x5e, 0x3b, 0x12, 0x4d, 0x47, 0x62, 0xd5, 0x91, 0xb8, 0x1f, 0x6e, 0x3e,
	0xb9, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x03, 0x17, 0x4b, 0x31, 0x01, 0x00, 0x00,
}

func (x TemperatureUnits) String() string {
	s, ok := TemperatureUnits_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *TemperatureRead) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TemperatureRead)
	if !ok {
		that2, ok := that.(TemperatureRead)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Result == nil {
		if this.Result != nil {
			return false
		}
	} else if this.Result == nil {
		return false
	} else if !this.Result.Equal(that1.Result) {
		return false
	}
	if this.Units != that1.Units {
		return false
	}
	if this.Location != that1.Location {
		return false
	}
	return true
}
func (this *TemperatureRead_Value) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TemperatureRead_Value)
	if !ok {
		that2, ok := that.(TemperatureRead_Value)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (this *TemperatureRead_Error) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TemperatureRead_Error)
	if !ok {
		that2, ok := that.(TemperatureRead_Error)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Error != that1.Error {
		return false
	}
	return true
}
func (this *TemperatureRead) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&messages.TemperatureRead{")
	if this.Result != nil {
		s = append(s, "Result: "+fmt.Sprintf("%#v", this.Result)+",\n")
	}
	s = append(s, "Units: "+fmt.Sprintf("%#v", this.Units)+",\n")
	s = append(s, "Location: "+fmt.Sprintf("%#v", this.Location)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *TemperatureRead_Value) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&messages.TemperatureRead_Value{` +
		`Value:` + fmt.Sprintf("%#v", this.Value) + `}`}, ", ")
	return s
}
func (this *TemperatureRead_Error) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&messages.TemperatureRead_Error{` +
		`Error:` + fmt.Sprintf("%#v", this.Error) + `}`}, ", ")
	return s
}
func valueToGoStringTemperature(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *TemperatureRead) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TemperatureRead) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		nn1, err := m.Result.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	if m.Units != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTemperature(dAtA, i, uint64(m.Units))
	}
	if len(m.Location) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintTemperature(dAtA, i, uint64(len(m.Location)))
		i += copy(dAtA[i:], m.Location)
	}
	return i, nil
}

func (m *TemperatureRead_Value) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x9
	i++
	encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Value))))
	i += 8
	return i, nil
}
func (m *TemperatureRead_Error) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintTemperature(dAtA, i, uint64(len(m.Error)))
	i += copy(dAtA[i:], m.Error)
	return i, nil
}
func encodeVarintTemperature(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TemperatureRead) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		n += m.Result.Size()
	}
	if m.Units != 0 {
		n += 1 + sovTemperature(uint64(m.Units))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovTemperature(uint64(l))
	}
	return n
}

func (m *TemperatureRead_Value) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 9
	return n
}
func (m *TemperatureRead_Error) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Error)
	n += 1 + l + sovTemperature(uint64(l))
	return n
}

func sovTemperature(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTemperature(x uint64) (n int) {
	return sovTemperature(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *TemperatureRead) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TemperatureRead{`,
		`Result:` + fmt.Sprintf("%v", this.Result) + `,`,
		`Units:` + fmt.Sprintf("%v", this.Units) + `,`,
		`Location:` + fmt.Sprintf("%v", this.Location) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TemperatureRead_Value) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TemperatureRead_Value{`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TemperatureRead_Error) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TemperatureRead_Error{`,
		`Error:` + fmt.Sprintf("%v", this.Error) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTemperature(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *TemperatureRead) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTemperature
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
			return fmt.Errorf("proto: TemperatureRead: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TemperatureRead: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Result = &TemperatureRead_Value{float64(math.Float64frombits(v))}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemperature
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
				return ErrInvalidLengthTemperature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemperature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = &TemperatureRead_Error{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			m.Units = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemperature
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Units |= TemperatureUnits(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTemperature
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
				return ErrInvalidLengthTemperature
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTemperature
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTemperature(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTemperature
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTemperature
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
func skipTemperature(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTemperature
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
					return 0, ErrIntOverflowTemperature
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
					return 0, ErrIntOverflowTemperature
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
				return 0, ErrInvalidLengthTemperature
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthTemperature
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTemperature
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
				next, err := skipTemperature(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthTemperature
				}
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
	ErrInvalidLengthTemperature = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTemperature   = fmt.Errorf("proto: integer overflow")
)