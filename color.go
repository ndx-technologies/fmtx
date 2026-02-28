package fmtx

import "strings"

const (
	Reset    = "\033[0m"
	Red      = "\033[91m"
	Green    = "\033[92m"
	Yellow   = "\033[93m"
	Blue     = "\033[94m"
	Magenta  = "\033[95m"
	Cyan     = "\033[96m"
	Dim      = "\033[2m"
	Bold     = "\033[1m"
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
)

var EnableColor = true

var colorCleaner = strings.NewReplacer(
	Red, "", Green, "", Yellow, "", Blue, "", Magenta, "", Cyan, "",
	Dim, "", Bold, "", BgRed, "", BgGreen, "", BgYellow, "", Reset, "",
)

func RedS(s string) string {
	if !EnableColor {
		return s
	}
	return Red + s + Reset
}
func GreenS(s string) string {
	if !EnableColor {
		return s
	}
	return Green + s + Reset
}
func YellowS(s string) string {
	if !EnableColor {
		return s
	}
	return Yellow + s + Reset
}
func BlueS(s string) string {
	if !EnableColor {
		return s
	}
	return Blue + s + Reset
}
func DimS(s string) string {
	if !EnableColor {
		return s
	}
	return Dim + s + Reset
}
func BoldS(s string) string {
	if !EnableColor {
		return s
	}
	return Bold + s + Reset
}
func BgRedS(s string) string {
	if !EnableColor {
		return s
	}
	return BgRed + s + Reset
}
func BgGreenS(s string) string {
	if !EnableColor {
		return s
	}
	return BgGreen + s + Reset
}

func ColorizeMinMax[T int | float32 | float64](s string, v, min, max T, minC, maxC string) string {
	if !EnableColor {
		return s
	}
	if v == min {
		return minC + s + Reset
	}
	if v == max {
		return maxC + s + Reset
	}
	return s
}

// ColorizeDist colorizes s based on which interval v falls into.
// points must be sorted ascending; colors must have len(points)+1 elements.
// colors[0]: v < points[0]
// colors[i]: points[i-1] <= v < points[i]
// colors[len(points)]: v >= points[len(points)-1]
func ColorizeDist[T int | float32 | float64](s string, v T, points []T, colors []string) string {
	if !EnableColor {
		return s
	}
	for i, p := range points {
		if v < p {
			if i < len(colors) {
				return colors[i] + s + Reset
			}
			return s
		}
	}
	if len(colors) > len(points) {
		return colors[len(points)] + s + Reset
	}
	return s
}
