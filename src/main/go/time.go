// Please, run go generate on root of this repository (github.com/devimteam/proto-utils)
package _go

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/lib/pq"
)

type Period struct {
	Start *time.Time
	End   *time.Time
}

func TimeToProto(t time.Time) (*timestamp.Timestamp, error) {
	return ptypes.TimestampProto(t)
}

func ProtoToTime(t *timestamp.Timestamp) (time.Time, error) {
	return ptypes.Timestamp(t)
}

func TimePtrToProto(t *time.Time) (*timestamp.Timestamp, error) {
	if t == nil {
		return nil, nil
	}
	return TimeToProto(*t)
}

func ProtoToTimePtr(t *timestamp.Timestamp) (*time.Time, error) {
	if t == nil {
		return nil, nil
	}
	tm, err := ProtoToTime(t)
	return &tm, err
}

func TimestampToPqNullTime(t *timestamp.Timestamp) pq.NullTime {
	if t == nil {
		return pq.NullTime{
			Valid: false,
		}
	}
	tm, err := ProtoToTime(t)
	return pq.NullTime{
		Time:  tm,
		Valid: err != nil,
	}
}

func PqNullTimeToTimestamp(nt pq.NullTime) *timestamp.Timestamp {
	if nt.Valid == false {
		return nil
	}
	t, _ := TimeToProto(nt.Time)
	return t
}
