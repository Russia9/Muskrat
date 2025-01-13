package utils

import (
	"math"
	"strconv"
	"time"
)

func FormatDuration(d time.Duration) string {
	days := int(math.Floor(d.Hours() / 24))
	hours := int(math.Floor(d.Hours() - float64(days*24)))
	minutes := int(math.Floor(d.Minutes() - float64(days*24*60) - float64(hours*60)))

	if days > 0 {
		return strconv.Itoa(days) + "d"
	}
	if hours >= 10 {
		return strconv.Itoa(hours) + "h"
	}
	if hours > 0 {
		return strconv.Itoa(hours) + "h" + strconv.Itoa(minutes) + "m"
	}
	return strconv.Itoa(minutes) + "m"
}

func ShortString(s string, n int) string {
	if len(s) > n {
		return s[:n-3] + "..."
	}

	return s
}
