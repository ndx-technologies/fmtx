package fmtx

import (
	"strings"
	"unicode"
)

func displayWidth(s string) int {
	w := 0
	for _, r := range colorCleaner.Replace(s) {
		// CJK characters and fullwidth forms are double-width
		// Box drawing and geometric shapes are single-width
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hangul, r) ||
			(r >= 0x3040 && r <= 0x309F) || // Hiragana
			(r >= 0x30A0 && r <= 0x30FF) || // Katakana
			(r >= 0x3000 && r <= 0x303F) || // CJK Symbols and Punctuation
			(r >= 0xFF00 && r <= 0xFFEF) { // Fullwidth Forms
			w += 2
		} else {
			w++
		}
	}
	return w
}

func truncate(s string, width int) string {
	visible := colorCleaner.Replace(s)
	if displayWidth(visible) <= width {
		return s
	}
	// Truncate the visible string
	runes := []rune(visible)
	w := 0
	for i := range runes {
		cw := 1
		r := runes[i]
		// CJK characters and fullwidth forms are double-width
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hangul, r) ||
			(r >= 0x3040 && r <= 0x309F) || // Hiragana
			(r >= 0x30A0 && r <= 0x30FF) || // Katakana
			(r >= 0x3000 && r <= 0x303F) || // CJK Symbols and Punctuation
			(r >= 0xFF00 && r <= 0xFFEF) { // Fullwidth Forms
			cw = 2
		}
		if w+cw > width-3 {
			truncatedVisible := string(runes[:i]) + "..."
			// Preserve ANSI codes from the start
			if strings.HasPrefix(s, "\033[") {
				end := strings.Index(s, "m")
				if end != -1 {
					return s[:end+1] + truncatedVisible + "\033[0m"
				}
			}
			return truncatedVisible
		}
		w += cw
	}
	return s
}
