package _go

import "github.com/golang/protobuf/ptypes/wrappers"

func StringToProto(str string) *wrappers.StringValue {
	if str == "" {
		return nil
	}
	return &wrappers.StringValue{Value: str}
}

func ProtoToString(value *wrappers.StringValue) string {
	if value == nil {
		return ""
	}
	t := value.Value
	return t
}

func NilStringToProto(str *string) *wrappers.StringValue {
	if str == nil {
		return nil
	}
	return StringToProto(*str)
}

func ProtoToNilString(value *wrappers.StringValue) *string {
	if value == nil {
		return nil
	}
	ret := ProtoToString(value)
	return &ret
}

func NilUInt32ToProto(d *uint32) *wrappers.UInt32Value {
	if d == nil {
		return nil
	}
	return &wrappers.UInt32Value{
		Value: *d,
	}
}

func ProtoToNilUInt32(d *wrappers.UInt32Value) *uint32 {
	if d == nil {
		return nil
	}
	t := d.Value
	return &t
}

func NilUInt64ToProto(d *uint64) *wrappers.UInt64Value {
	if d == nil {
		return nil
	}
	return &wrappers.UInt64Value{
		Value: *d,
	}
}

func ProtoToNilUInt64(d *wrappers.UInt64Value) *uint64 {
	if d == nil {
		return nil
	}
	t := d.Value
	return &t
}

func NilInt64ToProto(d *int64) *wrappers.Int64Value {
	if d == nil {
		return nil
	}
	return &wrappers.Int64Value{
		Value: *d,
	}
}

func ProtoToNilInt64(d *wrappers.Int64Value) *int64 {
	if d == nil {
		return nil
	}
	t := d.Value
	return &t
}

func NilInt32ToProto(d *int32) *wrappers.Int32Value {
	if d == nil {
		return nil
	}
	return &wrappers.Int32Value{
		Value: *d,
	}
}

func ProtoToNilInt32(d *wrappers.Int32Value) *int32 {
	if d == nil {
		return nil
	}
	t := d.Value
	return &t
}