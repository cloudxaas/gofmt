package cxfmtreadable

import (
	"testing"
	"time"
)

func BenchmarkAppendBytes(b *testing.B) {
	var buf []byte
	sizes := []uint64{500, 1024, 1048576, 1073741824, 1099511627776} // Various sizes: 500B, 1KB, 1MB, 1GB, 1TB
	for _, size := range sizes {
		b.Run("Size "+strconv.FormatUint(size, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf = buf[:0]
				buf = AppendBytes(buf, size)
			}
			b.ReportAllocs()
		})
	}
}

func BenchmarkFormatDuration(b *testing.B) {
	var buf []byte
	durations := []time.Duration{
		500 * time.Nanosecond,
		1500 * time.Microsecond,
		2500 * time.Millisecond,
		120 * time.Second,
		90 * time.Minute,
		48 * time.Hour,
	} // Various durations
	for _, d := range durations {
		b.Run("Duration "+d.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				buf = buf[:0]
				buf = FormatDuration(buf, d)
			}
			b.ReportAllocs()
		})
	}
}

func TestAppendBytes(t *testing.T) {
	tests := []struct {
		input    uint64
		expected string
	}{
		{500, "500B"},
		{1024, "1.0KB"},
		{1048576, "1.0MB"},
		{1073741824, "1.0GB"},
		{1099511627776, "1.0TB"},
	}

	for _, test := range tests {
		buf := make([]byte, 0, 32)
		result := string(AppendBytes(buf, test.input))
		if result != test.expected {
			t.Errorf("AppendBytes(%d) = %s; want %s", test.input, result, test.expected)
		}
	}
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{500 * time.Nanosecond, "500ns"},
		{1500 * time.Microsecond, "1500µs"},
		{2500 * time.Millisecond, "2500ms"},
		{120 * time.Second, "120s"},
		{90 * time.Minute, "90m"},
		{48 * time.Hour, "48h"},
	}

	for _, test := range tests {
		buf := make([]byte, 0, 32)
		result := string(FormatDuration(buf, test.input))
		if result != test.expected {
			t.Errorf("FormatDuration(%s) = %s; want %s", test.input, result, test.expected)
		}
	}
}
