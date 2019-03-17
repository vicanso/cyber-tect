package util

import (
	"time"
)

// Now get the now time
// 可扩展此函数针对开发环境mock时间
func Now() time.Time {
	return time.Now()
}

// NowString get the now time string of time RFC3339
func NowString() string {
	return Now().Format(time.RFC3339)
}

// UTCNow get the utc time
func UTCNow() time.Time {
	return Now().UTC()
}

// UTCNowString get the utc time string of time RFC3339
func UTCNowString() string {
	return UTCNow().Format(time.RFC3339)
}

// ParseTime parse time
func ParseTime(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}

// FormatTime format time
func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
