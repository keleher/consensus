// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ProposeRequest struct {
	Identity string `protobuf:"bytes,1,opt,name=identity" json:"identity,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value    []byte `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ProposeRequest) Reset()                    { *m = ProposeRequest{} }
func (m *ProposeRequest) String() string            { return proto.CompactTextString(m) }
func (*ProposeRequest) ProtoMessage()               {}
func (*ProposeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ProposeRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *ProposeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProposeRequest) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type ProposeReply struct {
	Success  bool   `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Redirect string `protobuf:"bytes,2,opt,name=redirect" json:"redirect,omitempty"`
	Error    string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	Seq      uint64 `protobuf:"varint,4,opt,name=seq" json:"seq,omitempty"`
	Leader   int32  `protobuf:"varint,5,opt,name=leader" json:"leader,omitempty"`
	Name     string `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	Value    []byte `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ProposeReply) Reset()                    { *m = ProposeReply{} }
func (m *ProposeReply) String() string            { return proto.CompactTextString(m) }
func (*ProposeReply) ProtoMessage()               {}
func (*ProposeReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ProposeReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ProposeReply) GetRedirect() string {
	if m != nil {
		return m.Redirect
	}
	return ""
}

func (m *ProposeReply) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ProposeReply) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ProposeReply) GetLeader() int32 {
	if m != nil {
		return m.Leader
	}
	return 0
}

func (m *ProposeReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProposeReply) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposeRequest)(nil), "pb.ProposeRequest")
	proto.RegisterType((*ProposeReply)(nil), "pb.ProposeReply")
}

func init() { proto.RegisterFile("client.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xb1, 0x6a, 0xc4, 0x30,
	0x0c, 0x86, 0xf1, 0x5d, 0x92, 0xbb, 0x8a, 0x50, 0x8a, 0x28, 0xc5, 0x74, 0x32, 0x37, 0x79, 0xea,
	0xd2, 0x17, 0x29, 0x1e, 0xba, 0xe7, 0x1c, 0x0d, 0x06, 0x37, 0x76, 0x64, 0xa7, 0x90, 0x97, 0xea,
	0x33, 0x96, 0x38, 0xa9, 0xd7, 0x6e, 0xff, 0x27, 0xc1, 0xa7, 0x1f, 0x41, 0x6f, 0xbd, 0xa3, 0x29,
	0xbf, 0x45, 0x0e, 0x39, 0xe0, 0x29, 0xde, 0x6f, 0x9f, 0xf0, 0xf8, 0xc1, 0x21, 0x86, 0x44, 0x86,
	0xe6, 0x85, 0x52, 0xc6, 0x57, 0xb8, 0xba, 0x91, 0xa6, 0xec, 0xf2, 0x2a, 0x85, 0x12, 0xfa, 0xc1,
	0x54, 0x46, 0x84, 0x66, 0x1a, 0xbe, 0x48, 0x9e, 0xca, 0xbc, 0x64, 0x7c, 0x86, 0xf6, 0x7b, 0xf0,
	0x0b, 0xc9, 0x8b, 0x12, 0xba, 0x37, 0x3b, 0xdc, 0x7e, 0x04, 0xf4, 0x55, 0x1c, 0xfd, 0x8a, 0x12,
	0x2e, 0x69, 0xb1, 0x96, 0x52, 0x2a, 0xd6, 0xab, 0xf9, 0xc3, 0xed, 0x20, 0xd3, 0xe8, 0x98, 0x6c,
	0x3e, 0xc4, 0x95, 0x37, 0x39, 0x31, 0x07, 0x96, 0xe7, 0xb2, 0xd8, 0x01, 0x9f, 0xe0, 0x9c, 0x68,
	0x96, 0x8d, 0x12, 0xba, 0x31, 0x5b, 0xc4, 0x17, 0xe8, 0x3c, 0x0d, 0x23, 0xb1, 0x6c, 0x95, 0xd0,
	0xad, 0x39, 0xa8, 0x16, 0xee, 0xfe, 0x2b, 0x7c, 0xef, 0xca, 0x4f, 0xde, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x3d, 0xa6, 0x9e, 0xdd, 0x23, 0x01, 0x00, 0x00,
}
