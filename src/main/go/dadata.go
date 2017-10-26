// Please, run go generate on root of this repository (github.com/devimteam/proto-utils)
package _go

import "github.com/devimteam/proto-utils/src/main/go/dadata"

type Address struct {
	FiasID             *string
	Country            *string
	RegionType         *string
	RegionTypeFull     *string
	Region             *string
	AreaType           *string
	AreaTypeFull       *string
	Area               *string
	CityType           *string
	CityTypeFull       *string
	City               *string
	CityDistrict       *string
	SettlementType     *string
	SettlementTypeFull *string
	Settlement         *string
	StreetType         *string
	StreetTypeFull     *string
	Street             *string
	HouseType          *string
	HouseTypeFull      *string
	House              *string
	BlockType          *string
	BlockTypeFull      *string
	Block              *string
	FlatType           *string
	FlatTypeFull       *string
	Flat               *string
	PostalCode         *string
}

func AddressPtrToProto(address *Address) (*dadata.Address, error) {
	if address == nil {
		return nil, nil
	}

	return &dadata.Address{
		FiasId:             NilStringToProto(address.FiasID),
		Country:            NilStringToProto(address.Country),
		RegionType:         NilStringToProto(address.RegionType),
		RegionTypeFull:     NilStringToProto(address.RegionTypeFull),
		Region:             NilStringToProto(address.Region),
		AreaType:           NilStringToProto(address.AreaType),
		AreaTypeFull:       NilStringToProto(address.AreaTypeFull),
		Area:               NilStringToProto(address.Area),
		CityType:           NilStringToProto(address.CityType),
		CityTypeFull:       NilStringToProto(address.CityTypeFull),
		City:               NilStringToProto(address.City),
		CityDistrict:       NilStringToProto(address.CityDistrict),
		SettlementType:     NilStringToProto(address.SettlementType),
		SettlementTypeFull: NilStringToProto(address.SettlementTypeFull),
		Settlement:         NilStringToProto(address.Settlement),
		StreetType:         NilStringToProto(address.StreetType),
		StreetTypeFull:     NilStringToProto(address.StreetTypeFull),
		Street:             NilStringToProto(address.Street),
		HouseType:          NilStringToProto(address.HouseType),
		HouseTypeFull:      NilStringToProto(address.HouseTypeFull),
		House:              NilStringToProto(address.House),
		BlockType:          NilStringToProto(address.BlockType),
		BlockTypeFull:      NilStringToProto(address.BlockTypeFull),
		Block:              NilStringToProto(address.Block),
		FlatType:           NilStringToProto(address.FlatType),
		FlatTypeFull:       NilStringToProto(address.FlatTypeFull),
		Flat:               NilStringToProto(address.Flat),
		PostalCode:         NilStringToProto(address.PostalCode),
	}, nil
}

func ProtoToAddressPtr(protoAddress *dadata.Address) (*Address, error) {
	if protoAddress == nil {
		return nil, nil
	}
	addr, err := ProtoToAddress(protoAddress)
	if err != nil {
		return nil, err
	}

	return &addr, nil
}

func ProtoToAddress(protoAddress *dadata.Address) (Address, error) {
	if protoAddress == nil {
		return Address{}, nil
	}

	return Address{
		FiasID:             ProtoToNilString(protoAddress.FiasId),
		Country:            ProtoToNilString(protoAddress.Country),
		RegionType:         ProtoToNilString(protoAddress.RegionType),
		RegionTypeFull:     ProtoToNilString(protoAddress.RegionTypeFull),
		Region:             ProtoToNilString(protoAddress.Region),
		AreaType:           ProtoToNilString(protoAddress.AreaType),
		AreaTypeFull:       ProtoToNilString(protoAddress.AreaTypeFull),
		Area:               ProtoToNilString(protoAddress.Area),
		CityType:           ProtoToNilString(protoAddress.CityType),
		CityTypeFull:       ProtoToNilString(protoAddress.CityTypeFull),
		City:               ProtoToNilString(protoAddress.City),
		CityDistrict:       ProtoToNilString(protoAddress.CityDistrict),
		SettlementType:     ProtoToNilString(protoAddress.SettlementType),
		SettlementTypeFull: ProtoToNilString(protoAddress.SettlementTypeFull),
		Settlement:         ProtoToNilString(protoAddress.Settlement),
		StreetType:         ProtoToNilString(protoAddress.StreetType),
		StreetTypeFull:     ProtoToNilString(protoAddress.StreetTypeFull),
		Street:             ProtoToNilString(protoAddress.Street),
		HouseType:          ProtoToNilString(protoAddress.HouseType),
		HouseTypeFull:      ProtoToNilString(protoAddress.HouseTypeFull),
		House:              ProtoToNilString(protoAddress.House),
		BlockType:          ProtoToNilString(protoAddress.BlockType),
		BlockTypeFull:      ProtoToNilString(protoAddress.BlockTypeFull),
		Block:              ProtoToNilString(protoAddress.Block),
		FlatType:           ProtoToNilString(protoAddress.FlatType),
		FlatTypeFull:       ProtoToNilString(protoAddress.FlatTypeFull),
		Flat:               ProtoToNilString(protoAddress.Flat),
		PostalCode:         ProtoToNilString(protoAddress.PostalCode),
	}, nil
}
