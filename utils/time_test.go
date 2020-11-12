package utils

import (
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestTimeToDateString(t *testing.T) {
	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	assert.Equal(t, TimeToDateString(t1), uint64(20121101))

	layout := "2006-01-02T15:04:05.000Z"
	str := "2014-11-12T11:45:26.371Z"
	t2, _ := time.Parse(layout, str)
	assert.Equal(t, TimeToDateString(t2), uint64(20141112))
}

func TestDateStringToTime(t *testing.T) {
	t1 := DateStringToTime(20121101)

	assert.Equal(t, t1.Year(), 2012)
	assert.Equal(t, t1.Month(), time.November)
	assert.Equal(t, t1.Day(), 1)
}
