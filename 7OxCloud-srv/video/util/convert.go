package util

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertTime2Timestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

func ConvertTimestamp2Time(ts *timestamppb.Timestamp) time.Time {
	return ts.AsTime()
}
