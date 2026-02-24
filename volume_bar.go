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
