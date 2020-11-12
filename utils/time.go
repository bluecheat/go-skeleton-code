package utils

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"strconv"
	"time"
)

func TimeToShortDateNumber(t time.Time) uint64 {
	date, _ := strconv.ParseUint(fmt.Sprintf("%d%02d%02d",
		t.Year(), t.Month(), t.Day()), 10, 64)
	return date
}

func TimeToDateString(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func DateStringToTime(date uint64) time.Time {
	dateString := strconv.FormatUint(date, 10)
	yyyy := dateString[:4]
	mm := dateString[4:6]
	dd := dateString[6:]
	t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s-%02s-%02sT%02s:%02s:%02s-00:00", yyyy, mm, dd, "00", "00", "00"))
	return t
}

func UintToString(i uint64) string {
	return strconv.FormatUint(i, 10)
}

func StringToUint(s string) uint64 {
	i, _ := strconv.ParseUint(s, 10, 64)
	return i
}

func TimeToTimestamp(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func TimestampToTime(t timestamp.Timestamp) time.Time {
	return time.Unix(t.Seconds, int64(t.Nanos))
}
