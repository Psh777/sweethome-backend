package lib

import (
	"time"
)

func TimeNow(timee time.Time) int64 {
	return timee.UnixNano() / int64(time.Millisecond)
}