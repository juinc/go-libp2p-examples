// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat.proto

package main

import (
	fmt "fmt"
	io "io"
	math "math"

	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
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

type Request_Type int32

const (
	Request_SEND_MESSAGE Request_Type = 0
	Request_UPDATE_PEER  Request_Type = 1
)

var Request_Type_name = map[int32]string{
	0: "SEND_MESSAGE",
	1: "UPDATE_PEER",
}

var Request_Type_value = map[string]int32{
	"SEND_MESSAGE": 0,
	"UPDATE_PEER":  1,
}

func (x Request_Type) Enum() *Request_Type {
	p := new(Request_Type)
	*p = x
	return p
}

func (x Request_Type) String() string {
	return proto.EnumName(Request_Type_name, int32(x))
}

func (x *Request_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Request_Type_value, data, "Request_Type")
	if err != nil {
		return err
	}
	*x = Request_Type(value)
	return nil
}

func (Request_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0, 0}
}

type Request struct {
	Type                 *Request_Type `protobuf:"varint,1,req,name=type,enum=main.Request_Type" json:"type,omitempty"`
	SendMessage          *SendMessage  `protobuf:"bytes,2,opt,name=sendMessage" json:"sendMessage,omitempty"`
	UpdatePeer           *UpdatePeer   `protobuf:"bytes,3,opt,name=updatePeer" json:"updatePeer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() Request_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Request_SEND_MESSAGE
}

func (m *Request) GetSendMessage() *SendMessage {
	if m != nil {
		return m.SendMessage
	}
	return nil
}

func (m *Request) GetUpdatePeer() *UpdatePeer {
	if m != nil {
		return m.UpdatePeer
	}
	return nil
}

type SendMessage struct {
	Data                 []byte   `protobuf:"bytes,1,req,name=data" json:"data,omitempty"`
	Created              *int64   `protobuf:"varint,2,req,name=created" json:"created,omitempty"`
	Id                   []byte   `protobuf:"bytes,3,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessage) Reset()         { *m = SendMessage{} }
func (m *SendMessage) String() string { return proto.CompactTextString(m) }
func (*SendMessage) ProtoMessage()    {}
func (*SendMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}
func (m *SendMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessage.Merge(m, src)
}
func (m *SendMessage) XXX_Size() int {
	return m.Size()
}
func (m *SendMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessage proto.InternalMessageInfo

func (m *SendMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SendMessage) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *SendMessage) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type UpdatePeer struct {
	UserHandle           []byte   `protobuf:"bytes,1,opt,name=userHandle" json:"userHandle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePeer) Reset()         { *m = UpdatePeer{} }
func (m *UpdatePeer) String() string { return proto.CompactTextString(m) }
func (*UpdatePeer) ProtoMessage()    {}
func (*UpdatePeer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}
func (m *UpdatePeer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdatePeer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdatePeer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdatePeer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePeer.Merge(m, src)
}
func (m *UpdatePeer) XXX_Size() int {
	return m.Size()
}
func (m *UpdatePeer) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePeer.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePeer proto.InternalMessageInfo

func (m *UpdatePeer) GetUserHandle() []byte {
	if m != nil {
		return m.UserHandle
	}
	return nil
}

func init() {
	proto.RegisterEnum("main.Request_Type", Request_Type_name, Request_Type_value)
	proto.RegisterType((*Request)(nil), "main.Request")
	proto.RegisterType((*SendMessage)(nil), "main.SendMessage")
	proto.RegisterType((*UpdatePeer)(nil), "main.UpdatePeer")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x3b, 0x93, 0x40, 0xe1, 0x26, 0xf4, 0xcf, 0x7f, 0x57, 0xb3, 0x0a, 0x21, 0x0b, 0x89,
	0x20, 0x41, 0xea, 0x13, 0x54, 0x3a, 0x28, 0x48, 0x25, 0x4c, 0xda, 0x75, 0x19, 0x3a, 0x17, 0x0d,
	0x68, 0x1a, 0x93, 0xc9, 0xa2, 0x6f, 0xe2, 0xfb, 0xb8, 0x71, 0xe9, 0x23, 0x48, 0x7c, 0x11, 0x49,
	0xaa, 0x36, 0xbb, 0x99, 0x73, 0xbe, 0xc3, 0xb9, 0xf7, 0x02, 0xec, 0x1e, 0xb5, 0x4d, 0xab, 0x7a,
	0x6f, 0xf7, 0xe8, 0x3e, 0xeb, 0xa2, 0x8c, 0xdf, 0x18, 0x4c, 0x15, 0xbd, 0xb4, 0xd4, 0x58, 0x3c,
	0x03, 0xd7, 0x1e, 0x2a, 0x12, 0x2c, 0xe2, 0xc9, 0x6c, 0x8e, 0x69, 0x0f, 0xa4, 0x3f, 0x66, 0xba,
	0x3e, 0x54, 0xa4, 0x06, 0x1f, 0xaf, 0xc0, 0x6b, 0xa8, 0x34, 0x2b, 0x6a, 0x1a, 0xfd, 0x40, 0x82,
	0x47, 0x2c, 0xf1, 0xe6, 0xff, 0x8f, 0x78, 0x7e, 0x32, 0xd4, 0x98, 0xc2, 0x4b, 0x80, 0xb6, 0x32,
	0xda, 0x52, 0x46, 0x54, 0x0b, 0x67, 0xc8, 0x04, 0xc7, 0xcc, 0xe6, 0x4f, 0x57, 0x23, 0x26, 0x3e,
	0x07, 0xb7, 0x2f, 0xc5, 0x00, 0xfc, 0x5c, 0xde, 0x2f, 0xb7, 0x2b, 0x99, 0xe7, 0x8b, 0x1b, 0x19,
	0x4c, 0xf0, 0x1f, 0x78, 0x9b, 0x6c, 0xb9, 0x58, 0xcb, 0x6d, 0x26, 0xa5, 0x0a, 0x58, 0x7c, 0x07,
	0xde, 0xa8, 0x18, 0x11, 0x5c, 0xa3, 0xad, 0x1e, 0x16, 0xf1, 0xd5, 0xf0, 0x46, 0x01, 0xd3, 0x5d,
	0x4d, 0xda, 0x92, 0x11, 0x3c, 0xe2, 0x89, 0xa3, 0x7e, 0xbf, 0x38, 0x03, 0x5e, 0x18, 0xe1, 0x0c,
	0x2c, 0x2f, 0x4c, 0x7c, 0x01, 0x70, 0x9a, 0x08, 0x43, 0x80, 0xb6, 0xa1, 0xfa, 0x56, 0x97, 0xe6,
	0xa9, 0x3f, 0x0d, 0x4b, 0x7c, 0x35, 0x52, 0xae, 0x83, 0xf7, 0x2e, 0x64, 0x1f, 0x5d, 0xc8, 0x3e,
	0xbb, 0x90, 0xbd, 0x7e, 0x85, 0x93, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0xa3, 0x90, 0xc9,
	0x65, 0x01, 0x00, 0x00,
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintChat(dAtA, i, uint64(*m.Type))
	}
	if m.SendMessage != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintChat(dAtA, i, uint64(m.SendMessage.Size()))
		n1, err := m.SendMessage.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.UpdatePeer != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintChat(dAtA, i, uint64(m.UpdatePeer.Size()))
		n2, err := m.UpdatePeer.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SendMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintChat(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.Created == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintChat(dAtA, i, uint64(*m.Created))
	}
	if m.Id == nil {
		return 0, new(github_com_golang_protobuf_proto.RequiredNotSetError)
	} else {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintChat(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *UpdatePeer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdatePeer) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.UserHandle != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintChat(dAtA, i, uint64(len(m.UserHandle)))
		i += copy(dAtA[i:], m.UserHandle)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintChat(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != nil {
		n += 1 + sovChat(uint64(*m.Type))
	}
	if m.SendMessage != nil {
		l = m.SendMessage.Size()
		n += 1 + l + sovChat(uint64(l))
	}
	if m.UpdatePeer != nil {
		l = m.UpdatePeer.Size()
		n += 1 + l + sovChat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SendMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovChat(uint64(l))
	}
	if m.Created != nil {
		n += 1 + sovChat(uint64(*m.Created))
	}
	if m.Id != nil {
		l = len(m.Id)
		n += 1 + l + sovChat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UpdatePeer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserHandle != nil {
		l = len(m.UserHandle)
		n += 1 + l + sovChat(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovChat(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozChat(x uint64) (n int) {
	return sovChat(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var v Request_Type
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= Request_Type(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Type = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendMessage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SendMessage == nil {
				m.SendMessage = &SendMessage{}
			}
			if err := m.SendMessage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatePeer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UpdatePeer == nil {
				m.UpdatePeer = &UpdatePeer{}
			}
			if err := m.UpdatePeer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SendMessage) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: SendMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Created = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(github_com_golang_protobuf_proto.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UpdatePeer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChat
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
			return fmt.Errorf("proto: UpdatePeer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdatePeer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserHandle", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowChat
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
				return ErrInvalidLengthChat
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthChat
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserHandle = append(m.UserHandle[:0], dAtA[iNdEx:postIndex]...)
			if m.UserHandle == nil {
				m.UserHandle = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipChat(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthChat
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipChat(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
					return 0, ErrIntOverflowChat
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
				return 0, ErrInvalidLengthChat
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthChat
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowChat
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
				next, err := skipChat(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthChat
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
	ErrInvalidLengthChat = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChat   = fmt.Errorf("proto: integer overflow")
)
