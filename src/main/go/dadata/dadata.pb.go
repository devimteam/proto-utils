// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devim/protobuf/dadata.proto

package dadata // import "github.com/devimteam/proto-utils/src/main/go/dadata"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Address struct {
	FiasId               *wrappers.StringValue `protobuf:"bytes,1,opt,name=fias_id,json=fiasId,proto3" json:"fias_id,omitempty"`
	Country              *wrappers.StringValue `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	RegionType           *wrappers.StringValue `protobuf:"bytes,3,opt,name=region_type,json=regionType,proto3" json:"region_type,omitempty"`
	RegionTypeFull       *wrappers.StringValue `protobuf:"bytes,4,opt,name=region_type_full,json=regionTypeFull,proto3" json:"region_type_full,omitempty"`
	Region               *wrappers.StringValue `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	AreaType             *wrappers.StringValue `protobuf:"bytes,6,opt,name=area_type,json=areaType,proto3" json:"area_type,omitempty"`
	AreaTypeFull         *wrappers.StringValue `protobuf:"bytes,7,opt,name=area_type_full,json=areaTypeFull,proto3" json:"area_type_full,omitempty"`
	Area                 *wrappers.StringValue `protobuf:"bytes,8,opt,name=area,proto3" json:"area,omitempty"`
	CityType             *wrappers.StringValue `protobuf:"bytes,9,opt,name=city_type,json=cityType,proto3" json:"city_type,omitempty"`
	CityTypeFull         *wrappers.StringValue `protobuf:"bytes,10,opt,name=city_type_full,json=cityTypeFull,proto3" json:"city_type_full,omitempty"`
	City                 *wrappers.StringValue `protobuf:"bytes,11,opt,name=city,proto3" json:"city,omitempty"`
	CityDistrict         *wrappers.StringValue `protobuf:"bytes,12,opt,name=city_district,json=cityDistrict,proto3" json:"city_district,omitempty"`
	SettlementType       *wrappers.StringValue `protobuf:"bytes,13,opt,name=settlement_type,json=settlementType,proto3" json:"settlement_type,omitempty"`
	SettlementTypeFull   *wrappers.StringValue `protobuf:"bytes,14,opt,name=settlement_type_full,json=settlementTypeFull,proto3" json:"settlement_type_full,omitempty"`
	Settlement           *wrappers.StringValue `protobuf:"bytes,15,opt,name=settlement,proto3" json:"settlement,omitempty"`
	StreetType           *wrappers.StringValue `protobuf:"bytes,16,opt,name=street_type,json=streetType,proto3" json:"street_type,omitempty"`
	StreetTypeFull       *wrappers.StringValue `protobuf:"bytes,17,opt,name=street_type_full,json=streetTypeFull,proto3" json:"street_type_full,omitempty"`
	Street               *wrappers.StringValue `protobuf:"bytes,18,opt,name=street,proto3" json:"street,omitempty"`
	HouseType            *wrappers.StringValue `protobuf:"bytes,19,opt,name=house_type,json=houseType,proto3" json:"house_type,omitempty"`
	HouseTypeFull        *wrappers.StringValue `protobuf:"bytes,20,opt,name=house_type_full,json=houseTypeFull,proto3" json:"house_type_full,omitempty"`
	House                *wrappers.StringValue `protobuf:"bytes,21,opt,name=house,proto3" json:"house,omitempty"`
	BlockType            *wrappers.StringValue `protobuf:"bytes,22,opt,name=block_type,json=blockType,proto3" json:"block_type,omitempty"`
	BlockTypeFull        *wrappers.StringValue `protobuf:"bytes,23,opt,name=block_type_full,json=blockTypeFull,proto3" json:"block_type_full,omitempty"`
	Block                *wrappers.StringValue `protobuf:"bytes,24,opt,name=block,proto3" json:"block,omitempty"`
	FlatType             *wrappers.StringValue `protobuf:"bytes,25,opt,name=flat_type,json=flatType,proto3" json:"flat_type,omitempty"`
	FlatTypeFull         *wrappers.StringValue `protobuf:"bytes,26,opt,name=flat_type_full,json=flatTypeFull,proto3" json:"flat_type_full,omitempty"`
	Flat                 *wrappers.StringValue `protobuf:"bytes,27,opt,name=flat,proto3" json:"flat,omitempty"`
	PostalCode           *wrappers.StringValue `protobuf:"bytes,28,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_dadata_38bd2606d89fdaa4, []int{0}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetFiasId() *wrappers.StringValue {
	if m != nil {
		return m.FiasId
	}
	return nil
}

func (m *Address) GetCountry() *wrappers.StringValue {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Address) GetRegionType() *wrappers.StringValue {
	if m != nil {
		return m.RegionType
	}
	return nil
}

func (m *Address) GetRegionTypeFull() *wrappers.StringValue {
	if m != nil {
		return m.RegionTypeFull
	}
	return nil
}

func (m *Address) GetRegion() *wrappers.StringValue {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *Address) GetAreaType() *wrappers.StringValue {
	if m != nil {
		return m.AreaType
	}
	return nil
}

func (m *Address) GetAreaTypeFull() *wrappers.StringValue {
	if m != nil {
		return m.AreaTypeFull
	}
	return nil
}

func (m *Address) GetArea() *wrappers.StringValue {
	if m != nil {
		return m.Area
	}
	return nil
}

func (m *Address) GetCityType() *wrappers.StringValue {
	if m != nil {
		return m.CityType
	}
	return nil
}

func (m *Address) GetCityTypeFull() *wrappers.StringValue {
	if m != nil {
		return m.CityTypeFull
	}
	return nil
}

func (m *Address) GetCity() *wrappers.StringValue {
	if m != nil {
		return m.City
	}
	return nil
}

func (m *Address) GetCityDistrict() *wrappers.StringValue {
	if m != nil {
		return m.CityDistrict
	}
	return nil
}

func (m *Address) GetSettlementType() *wrappers.StringValue {
	if m != nil {
		return m.SettlementType
	}
	return nil
}

func (m *Address) GetSettlementTypeFull() *wrappers.StringValue {
	if m != nil {
		return m.SettlementTypeFull
	}
	return nil
}

func (m *Address) GetSettlement() *wrappers.StringValue {
	if m != nil {
		return m.Settlement
	}
	return nil
}

func (m *Address) GetStreetType() *wrappers.StringValue {
	if m != nil {
		return m.StreetType
	}
	return nil
}

func (m *Address) GetStreetTypeFull() *wrappers.StringValue {
	if m != nil {
		return m.StreetTypeFull
	}
	return nil
}

func (m *Address) GetStreet() *wrappers.StringValue {
	if m != nil {
		return m.Street
	}
	return nil
}

func (m *Address) GetHouseType() *wrappers.StringValue {
	if m != nil {
		return m.HouseType
	}
	return nil
}

func (m *Address) GetHouseTypeFull() *wrappers.StringValue {
	if m != nil {
		return m.HouseTypeFull
	}
	return nil
}

func (m *Address) GetHouse() *wrappers.StringValue {
	if m != nil {
		return m.House
	}
	return nil
}

func (m *Address) GetBlockType() *wrappers.StringValue {
	if m != nil {
		return m.BlockType
	}
	return nil
}

func (m *Address) GetBlockTypeFull() *wrappers.StringValue {
	if m != nil {
		return m.BlockTypeFull
	}
	return nil
}

func (m *Address) GetBlock() *wrappers.StringValue {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Address) GetFlatType() *wrappers.StringValue {
	if m != nil {
		return m.FlatType
	}
	return nil
}

func (m *Address) GetFlatTypeFull() *wrappers.StringValue {
	if m != nil {
		return m.FlatTypeFull
	}
	return nil
}

func (m *Address) GetFlat() *wrappers.StringValue {
	if m != nil {
		return m.Flat
	}
	return nil
}

func (m *Address) GetPostalCode() *wrappers.StringValue {
	if m != nil {
		return m.PostalCode
	}
	return nil
}

func init() {
	proto.RegisterType((*Address)(nil), "devim.protobuf.Address")
}

func init() { proto.RegisterFile("devim/protobuf/dadata.proto", fileDescriptor_dadata_38bd2606d89fdaa4) }

var fileDescriptor_dadata_38bd2606d89fdaa4 = []byte{
	// 592 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x95, 0xfd, 0x68, 0xd7, 0xd7, 0xb5, 0x1d, 0x66, 0x80, 0xd9, 0x06, 0x42, 0x9c, 0xb8,
	0x90, 0xa0, 0x0d, 0x90, 0x10, 0x70, 0x58, 0x29, 0x43, 0x08, 0x09, 0x55, 0x80, 0x40, 0xe2, 0x52,
	0xb9, 0x89, 0x9b, 0x45, 0xb8, 0x71, 0x64, 0x3b, 0x4c, 0xbd, 0xf3, 0x87, 0x20, 0x8e, 0xfc, 0x29,
	0xfc, 0x45, 0x1c, 0x91, 0xed, 0x3a, 0xd9, 0x38, 0xbd, 0xa3, 0x9d, 0xef, 0xe7, 0xe5, 0x63, 0xcb,
	0x4f, 0x0f, 0x0e, 0x33, 0xfe, 0xbd, 0x58, 0x26, 0x95, 0x92, 0x46, 0xce, 0xeb, 0x45, 0x92, 0xb1,
	0x8c, 0x19, 0x16, 0xbb, 0x35, 0x19, 0xba, 0x8f, 0x71, 0xf8, 0x78, 0x70, 0x37, 0x97, 0x32, 0x17,
	0xbc, 0x4d, 0x5f, 0x28, 0x56, 0x55, 0x5c, 0x69, 0x1f, 0xb9, 0xff, 0x67, 0x00, 0xdd, 0xd3, 0x2c,
	0x53, 0x5c, 0x6b, 0xf2, 0x04, 0xba, 0x8b, 0x82, 0xe9, 0x59, 0x91, 0xd1, 0xe8, 0x5e, 0xf4, 0xa0,
	0x7f, 0x7c, 0x14, 0x7b, 0xba, 0x29, 0x17, 0x7f, 0x34, 0xaa, 0x28, 0xf3, 0xcf, 0x4c, 0xd4, 0xfc,
	0x43, 0xc7, 0x86, 0xdf, 0x66, 0xe4, 0x29, 0x74, 0x53, 0x59, 0x97, 0x46, 0xad, 0xe8, 0x06, 0x02,
	0x0b, 0x61, 0xf2, 0x12, 0xfa, 0x8a, 0xe7, 0x85, 0x2c, 0x67, 0x66, 0x55, 0x71, 0xba, 0x89, 0x60,
	0xc1, 0x03, 0x9f, 0x56, 0x15, 0x27, 0x67, 0xb0, 0x77, 0x09, 0x9f, 0x2d, 0x6a, 0x21, 0xe8, 0x16,
	0xa2, 0xc6, 0xb0, 0xad, 0x71, 0x56, 0x0b, 0x41, 0x1e, 0x43, 0xc7, 0xef, 0xd0, 0x6d, 0xcc, 0xa1,
	0x7d, 0x96, 0x3c, 0x83, 0x1e, 0x53, 0x9c, 0x79, 0xf5, 0x0e, 0x02, 0xdc, 0xb1, 0x71, 0x27, 0x3e,
	0x86, 0x61, 0x83, 0x7a, 0xed, 0x2e, 0x82, 0xdf, 0x0d, 0xbc, 0x93, 0x7e, 0x04, 0x5b, 0x76, 0x4d,
	0x77, 0x10, 0xa4, 0x4b, 0x5a, 0xe1, 0xb4, 0x30, 0x2b, 0x2f, 0xdc, 0xc3, 0x08, 0xdb, 0x78, 0x10,
	0x6e, 0x50, 0x2f, 0x0c, 0x18, 0xe1, 0xc0, 0x07, 0x61, 0xbb, 0xa6, 0x7d, 0x8c, 0xb0, 0x4d, 0x92,
	0x53, 0x18, 0xb8, 0xbf, 0x66, 0x85, 0x36, 0xaa, 0x48, 0x0d, 0xdd, 0xc5, 0xfe, 0x74, 0xb2, 0x26,
	0xc8, 0x6b, 0x18, 0x69, 0x6e, 0x8c, 0xe0, 0x4b, 0x5e, 0x1a, 0x7f, 0xf2, 0x01, 0xe6, 0x85, 0xb4,
	0x90, 0x3b, 0xff, 0x7b, 0xd8, 0xff, 0xaf, 0x8c, 0xbf, 0x85, 0x21, 0xa2, 0x16, 0xb9, 0x5a, 0xcb,
	0xdd, 0xc5, 0x0b, 0x80, 0x76, 0x97, 0x8e, 0x30, 0xef, 0xbe, 0xcd, 0xdb, 0xb6, 0xd1, 0x46, 0x71,
	0xbe, 0x3e, 0xd0, 0x1e, 0x0a, 0x77, 0x40, 0x68, 0x9b, 0x4b, 0xb8, 0x3f, 0xc8, 0x35, 0xd4, 0xa5,
	0x34, 0x35, 0x42, 0xdb, 0xf8, 0x1d, 0x4a, 0x30, 0x6d, 0xe3, 0xb3, 0xe4, 0x39, 0xc0, 0xb9, 0xac,
	0x35, 0xf7, 0xee, 0xd7, 0x11, 0x64, 0xcf, 0xe5, 0x9d, 0xfa, 0x04, 0x46, 0x2d, 0xec, 0xcd, 0xf7,
	0x11, 0x15, 0x06, 0x4d, 0x05, 0x27, 0x7e, 0x0c, 0xdb, 0x6e, 0x83, 0xde, 0x40, 0xb0, 0x3e, 0x6a,
	0xb5, 0xe7, 0x42, 0xa6, 0xdf, 0xbc, 0xf6, 0x4d, 0x8c, 0xb6, 0xcb, 0x07, 0xed, 0x16, 0xf6, 0xda,
	0xb7, 0x30, 0xda, 0x4d, 0x85, 0xa0, 0xed, 0x36, 0x28, 0xc5, 0x68, 0xbb, 0xa8, 0xed, 0xf9, 0x85,
	0x60, 0xeb, 0x87, 0x72, 0x1b, 0xd3, 0xf3, 0x36, 0x1e, 0x7a, 0xbe, 0x41, 0xbd, 0xf3, 0x01, 0xa6,
	0xfd, 0x02, 0x1f, 0x7a, 0xde, 0xae, 0xe9, 0x21, 0xa6, 0xe7, 0x6d, 0xd2, 0xbe, 0xed, 0x4a, 0x6a,
	0xc3, 0xc4, 0x2c, 0x95, 0x19, 0xa7, 0x47, 0x98, 0xb7, 0xed, 0x81, 0x57, 0x32, 0xe3, 0xe3, 0x1f,
	0x11, 0x90, 0x54, 0x2e, 0xe3, 0xab, 0x33, 0x70, 0xdc, 0x9f, 0xb8, 0x09, 0x39, 0xb5, 0xeb, 0x69,
	0xf4, 0xf5, 0x24, 0x2f, 0xcc, 0x79, 0x3d, 0x8f, 0x53, 0xb9, 0x4c, 0x5c, 0xd2, 0x70, 0xb6, 0x1e,
	0xa7, 0x0f, 0x6b, 0x53, 0x08, 0x9d, 0x68, 0x95, 0x26, 0x4b, 0x56, 0x94, 0x49, 0x2e, 0xd7, 0xd3,
	0xf5, 0x67, 0x14, 0xfd, 0x8d, 0xa2, 0x5f, 0x1b, 0x9b, 0x6f, 0xa6, 0xe3, 0xdf, 0x1b, 0x77, 0x26,
	0xee, 0x07, 0xd3, 0xe0, 0xf3, 0x85, 0x0b, 0xf1, 0xae, 0x94, 0x17, 0x6e, 0xa8, 0xe8, 0x79, 0xc7,
	0xd5, 0x3a, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xb2, 0xdf, 0xc7, 0xa9, 0x07, 0x00, 0x00,
}