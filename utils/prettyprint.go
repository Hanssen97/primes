package utils

import (
	"fmt"
	"time"
)

func PrettyDuration(duration time.Duration) string {
	if duration.Minutes() > 1 {
		return fmt.Sprintf("%.3fm", duration.Minutes())
	}
	if duration.Seconds() > 1 {
		return fmt.Sprintf("%.3fs", duration.Seconds())
	}
	if duration.Milliseconds() != 0 {
		return fmt.Sprintf("%dms", duration.Milliseconds())
	}
	return fmt.Sprintf("%dµs", duration.Microseconds())
}