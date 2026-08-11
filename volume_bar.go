package fmtx

import (
	"strings"
)

func VolumeBar[T int | int32 | int64 | float32 | float64](v, max T, width int) string {
	if max <= 0 {
		return ""
	}
	filled := min(width, int(float64(v)*float64(width)/float64(max)))
	return BlueS(strings.Repeat("█", int(filled)) + strings.Repeat("░", int(width-filled)))
}

func VolumeBar2[T int | int32 | int64](a, b, max T, width int) string {
	if max <= 0 {
		return ""
	}
	if a < 0 {
		a = 0
	}
	if b < 0 {
		b = 0
	}
	barA := min(width, int(float64(a)*float64(width)/float64(max)))
	barB := min(width-barA, int(float64(b)*float64(width)/float64(max)))
	return BlueS(strings.Repeat("█", barA) + strings.Repeat("▒", barB) + strings.Repeat("░", width-barA-barB))
}
