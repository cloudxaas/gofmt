package cxfmtreadable

import (
	"strconv"
	"time"
)

// AppendBytes converts bytes to a human-readable string with units and appends to the provided buffer
func AppendBytes(buf []byte, b uint64) []byte {
	const unit = 1024
	if b < unit {
		return strconv.AppendUint(buf, b, 10)
	}
	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	buf = strconv.AppendFloat(buf, float64(b)/float64(div), 'f', 1, 64)
	buf = append(buf, " KMGTPE"[exp])
	buf = append(buf, 'B')
	return buf
}

// FormatDuration formats a time.Duration into a human-readable string without heap allocations.
func FormatDuration(buf []byte, d time.Duration) []byte {
	switch {
	case d < time.Microsecond:
		buf = append(buf, strconv.Itoa(int(d.Nanoseconds()))...)
		buf = append(buf, "ns"...)
	case d < time.Millisecond:
		buf = append(buf, strconv.Itoa(int(d.Microseconds()))...)
		buf = append(buf, "µs"...)
	case d < time.Second:
		buf = append(buf, strconv.Itoa(int(d.Milliseconds()))...)
		buf = append(buf, "ms"...)
	case d < time.Minute:
		buf = append(buf, strconv.Itoa(int(d.Seconds()))...)
		buf = append(buf, "s"...)
	case d < time.Hour:
		buf = append(buf, strconv.Itoa(int(d.Minutes()))...)
		buf = append(buf, "m"...)
	default:
		buf = append(buf, strconv.Itoa(int(d.Hours()))...)
		buf = append(buf, "h"...)
	}
	return buf
}
