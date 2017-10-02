// Please, run go generate on root of this repository (github.com/devimteam/proto-utils)
package _go

import "github.com/devimteam/proto-utils/src/main/go/collect"

type Uint32Range struct {
	Start *uint32
	End   *uint32
}

type Uint64Range struct {
	Start *uint64
	End   *uint64
}

func Uint32RangeToProto(r *Uint32Range) *collect.UInt32Range {
	return &collect.UInt32Range{
		Start: NilUInt32ToProto(r.Start),
		End:   NilUInt32ToProto(r.End),
	}
}

func ProtoToUint32Range(r *collect.UInt32Range) *Uint32Range {
	return &Uint32Range{
		Start: ProtoToNilUInt32(r.Start),
		End:   ProtoToNilUInt32(r.End),
	}
}

func Uint64RangeToProto(r *Uint64Range) *collect.UInt64Range {
	return &collect.UInt64Range{
		Start: NilUInt64ToProto(r.Start),
		End:   NilUInt64ToProto(r.End),
	}
}

func ProtoToUint64Range(r *collect.UInt64Range) *Uint64Range {
	return &Uint64Range{
		Start: ProtoToNilUInt64(r.Start),
		End:   ProtoToNilUInt64(r.End),
	}
}
