package cxfmtreadable

import (
	"math"
	"strconv"
	"time"
)

// AppendBytes converts bytes to a human-readable string with units and appends to the provided buffer
func AppendBytes(buf []byte, b uint64) []byte {
	const unit = 1024
	if b < unit {
		buf = strconv.AppendUint(buf, b, 10)
		buf = append(buf, 'B')
		return buf
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	buf = strconv.AppendFloat(buf, float64(b)/float64(div), 'f', 1, 64)
	buf = append(buf, "KMGTPE"[exp])
	buf = append(buf, 'B')
	return buf
}

// FormatDuration formats a time.Duration into a human-readable string without heap allocations.
func FormatDuration(buf []byte, d time.Duration) []byte {
	switch {
	case d < time.Microsecond:
		buf = strconv.AppendInt(buf, d.Nanoseconds(), 10)
		buf = append(buf, "ns"...)
	case d < time.Millisecond:
		buf = strconv.AppendInt(buf, d.Microseconds(), 10)
		buf = append(buf, "µs"...)
	case d < time.Second:
		buf = strconv.AppendInt(buf, d.Milliseconds(), 10)
		buf = append(buf, "ms"...)
	case d < time.Minute:
		buf = strconv.AppendFloat(buf, d.Seconds(), 'f', 0, 64)
		buf = append(buf, "s"...)
	case d < time.Hour:
		buf = strconv.AppendFloat(buf, d.Minutes(), 'f', 0, 64)
		buf = append(buf, "m"...)
	default:
		buf = strconv.AppendFloat(buf, d.Hours(), 'f', 0, 64)
		buf = append(buf, "h"...)
	}
	return buf
}


// FormatNumberCompact formats numbers into a compact form using 'k' for thousand, 'M' for million, etc.,
// aiming for zero heap allocations. The buffer must be pre-allocated with enough space.
func FormatNumberCompact(num int64, buf []byte) []byte {
	if num == 0 {
		return append(buf, '0')
	}

	// Determine the scale and suffix
	suffix := ""
	value := float64(num)
	if num < 0 {
		value = -value // Make positive for logarithmic calculation
	}

	switch {
	case value < 1000:
		return strconv.AppendInt(buf, num, 10) // No suffix needed for < 1000
	case value < 1_000_000:
		value /= 1000
		suffix = "k"
	case value < 1_000_000_000:
		value /= 1_000_000
		suffix = "M"
	default:
		value /= 1_000_000_000
		suffix = "B"
	}

	// Adjust the number to one decimal place and format
	value = math.Round(value*10) / 10
	buf = strconv.AppendFloat(buf, value, 'f', 1, 64)

	// Remove unnecessary decimal and zero if integer value
	if buf[len(buf)-2] == '.' && buf[len(buf)-1] == '0' {
		buf = buf[:len(buf)-2] // Remove trailing ".0"
	}

	// Append the suffix
	buf = append(buf, suffix...)

	return buf
}
