// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devim/protobuf/person.proto

package person // import "github.com/devimteam/proto-utils/src/main/go/person"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Gender int32

const (
	Gender_UNSET  Gender = 0
	Gender_MALE   Gender = 1
	Gender_FEMALE Gender = 2
)

var Gender_name = map[int32]string{
	0: "UNSET",
	1: "MALE",
	2: "FEMALE",
}
var Gender_value = map[string]int32{
	"UNSET":  0,
	"MALE":   1,
	"FEMALE": 2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}
func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_person_0dccffbc00f0ff40, []int{0}
}

func init() {
	proto.RegisterEnum("devim.protobuf.Gender", Gender_name, Gender_value)
}

func init() { proto.RegisterFile("devim/protobuf/person.proto", fileDescriptor_person_0dccffbc00f0ff40) }

var fileDescriptor_person_0dccffbc00f0ff40 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x49, 0x2d, 0xcb,
	0xcc, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x48, 0x2d, 0x2a, 0xce,
	0xcf, 0xd3, 0x03, 0xf3, 0x85, 0xf8, 0xc0, 0x92, 0x7a, 0x30, 0x49, 0x2d, 0x4d, 0x2e, 0x36, 0xf7,
	0xd4, 0xbc, 0x94, 0xd4, 0x22, 0x21, 0x4e, 0x2e, 0xd6, 0x50, 0xbf, 0x60, 0xd7, 0x10, 0x01, 0x06,
	0x21, 0x0e, 0x2e, 0x16, 0x5f, 0x47, 0x1f, 0x57, 0x01, 0x46, 0x21, 0x2e, 0x2e, 0x36, 0x37, 0x57,
	0x30, 0x9b, 0xc9, 0xa9, 0x85, 0x91, 0x4b, 0x28, 0x39, 0x3f, 0x57, 0x0f, 0xd5, 0x04, 0x27, 0xee,
	0x00, 0xb0, 0xf9, 0x01, 0x20, 0x7e, 0x00, 0x63, 0x94, 0x71, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92,
	0x5e, 0x72, 0x7e, 0xae, 0x3e, 0x58, 0x65, 0x49, 0x6a, 0x22, 0xd4, 0x31, 0xba, 0xa5, 0x25, 0x99,
	0x39, 0xc5, 0xfa, 0xc5, 0x45, 0xc9, 0xfa, 0xb9, 0x89, 0x99, 0x79, 0xfa, 0xe9, 0xf9, 0x50, 0xb7,
	0x2d, 0x60, 0x64, 0xfc, 0xc1, 0xc8, 0xb8, 0x88, 0x89, 0xd9, 0x3d, 0xc0, 0x69, 0x15, 0x93, 0xac,
	0x0b, 0xd8, 0x82, 0x00, 0xa8, 0x05, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x79, 0xf9, 0xe5, 0x79,
	0x21, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xed, 0xfa, 0x09, 0xbe, 0xe7, 0x00, 0x00, 0x00,
}