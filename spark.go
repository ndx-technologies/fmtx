package fmtx

import "strings"

// SparkBlocks are Unicode block characters for drawing sparklines,
// from empty (space) to full block. Levels 0..8.
var SparkBlocks = []rune{' ', '▁', '▂', '▃', '▄', '▅', '▆', '▇', '█'}

// Spark renders a sparkline from a sparse bucket→count map.
// Each of the width positions maps to a SparkBlocks character
// whose level is proportional to its count relative to maxCount.
func Spark(counts map[int]int, width int) string {
	if width == 0 {
		return ""
	}
	maxCount := 0
	for _, c := range counts {
		if c > maxCount {
			maxCount = c
		}
	}
	if maxCount == 0 {
		return strings.Repeat(string(SparkBlocks[0]), width)
	}
	var b strings.Builder
	for i := range width {
		level := counts[i] * (len(SparkBlocks) - 1) / maxCount
		b.WriteRune(SparkBlocks[level])
	}
	return b.String()
}

// SparkLine renders a sparkline from a slice of values.
// Each value is mapped to a SparkBlocks character proportional to maxVal.
// The result is stretched to width characters.
func SparkLine(vals []int, maxVal, width int) string {
	if maxVal == 0 || len(vals) == 0 || width == 0 {
		return ""
	}
	var b strings.Builder
	for i := range width {
		idx := i * len(vals) / width
		v := vals[idx]
		level := max((v*(len(SparkBlocks)-1))/maxVal, 0)
		if level >= len(SparkBlocks) {
			level = len(SparkBlocks) - 1
		}
		b.WriteRune(SparkBlocks[level])
	}
	return b.String()
}
