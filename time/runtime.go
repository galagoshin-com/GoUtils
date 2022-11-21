package time

import (
	"time"
)

func MeasureExecution(runtime func()) float64 {
	startTime := time.Now().UnixNano()
	runtime()
	endTime := time.Now().UnixNano()
	seconds := (float64(endTime) - float64(startTime)) / float64(time.Second)
	return seconds
}
