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

func RangeVolume[T int | int32 | int64 | float32 | float64](from, to, minV, maxV T, width int) string {
	if width <= 0 {
		return ""
	}
	if maxV <= minV {
		// single point: no range to show, render all empty
		return strings.Repeat("░", width)
	}
	norm := func(v T) int {
		n := int((float64(v) - float64(minV)) * float64(width) / (float64(maxV) - float64(minV)))
		return min(width, max(0, n))
	}
	start, end := norm(from), norm(to)
	if end < start {
		start, end = end, start
	}
	return strings.Repeat("░", start) + strings.Repeat("█", end-start) + strings.Repeat("░", width-end)
}

func RangeVolume2[T int | int32 | int64 | float32 | float64](from1, to1, from2, to2, minV, maxV T, width int) string {
	if width <= 0 {
		return ""
	}
	if maxV <= minV {
		return strings.Repeat("░", width)
	}
	norm := func(v T) int {
		n := int((float64(v) - float64(minV)) * float64(width) / (float64(maxV) - float64(minV)))
		return min(width, max(0, n))
	}
	p1, q1 := norm(from1), norm(to1)
	p2, q2 := norm(from2), norm(to2)
	if q1 < p1 {
		p1, q1 = q1, p1
	}
	if q2 < p2 {
		p2, q2 = q2, p2
	}
	if q1 > p2 {
		q1 = p2 // overlapping segments: let the second one win
	}
	if q1 < p1 {
		q1 = p1
	}
	if q2 < p2 {
		q2 = p2
	}
	return strings.Repeat("░", min(width, max(0, p1))) +
		strings.Repeat("█", max(0, q1-p1)) +
		strings.Repeat("░", max(0, p2-q1)) +
		strings.Repeat("▒", max(0, q2-p2)) +
		strings.Repeat("░", max(0, width-q2))
}
