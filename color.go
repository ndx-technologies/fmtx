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
