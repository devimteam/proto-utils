// Please, run go generate on root of this repository (github.com/devimteam/proto-utils)
package _go

import (
	"time"

	devim_time "github.com/devimteam/proto-utils/src/main/go/time"
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
	if t == nil {
		return time.Time{}, nil
	}
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

func PeriodToProto(p *Period) *devim_time.Period {
	if p == nil {
		return nil
	}
	start, _ := TimePtrToProto(p.Start)
	end, _ := TimePtrToProto(p.End)
	return &devim_time.Period{
		Start: start,
		End:   end,
	}
}

func ProtoToPeriod(period *devim_time.Period) *Period {
	if period == nil {
		return nil
	}
	start, _ := ProtoToTimePtr(period.Start)
	end, _ := ProtoToTimePtr(period.End)
	return &Period{
		Start: start,
		End:   end,
	}
}

func ManyTimesToProto(times ...interface{}) ([]*timestamp.Timestamp, error) {
	var list []*timestamp.Timestamp
	for _, t := range times {
		switch tt := t.(type) {
		case time.Time:
			ts, err := TimeToProto(tt)
			if err != nil {
				return nil, err
			}
			list = append(list, ts)
		case *time.Time:
			ts, err := TimePtrToProto(tt)
			if err != nil {
				return nil, err
			}
			list = append(list, ts)
		}
	}
	return list, nil
}
