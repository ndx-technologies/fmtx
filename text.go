package fmtx

import (
	"strings"
	"unicode"
)

func displayWidth(s string) int {
	w := 0
	for _, r := range colorCleaner.Replace(s) {
		// Combining/nonspacing marks are zero-width (e.g. Thai tone marks, diacritics)
		if unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) {
			continue
		}
		// Format characters are zero-width (e.g. Arabic RTL mark U+200F, directional embeddings)
		if unicode.Is(unicode.Cf, r) {
			continue
		}
		// CJK characters and fullwidth forms are double-width
		// Box drawing and geometric shapes are single-width
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hangul, r) ||
			(r >= 0x3040 && r <= 0x309F) || // Hiragana
			(r >= 0x30A0 && r <= 0x30FF) || // Katakana
			(r >= 0x3000 && r <= 0x303F) || // CJK Symbols and Punctuation
			(r >= 0xFF00 && r <= 0xFFEF) || // Fullwidth Forms
			(r >= 0x1F300 && r <= 0x1FAFF) { // Emoji
			w += 2
		} else {
			w++
		}
	}
	return w
}

// hasRTL reports whether s contains RTL characters (Arabic, Hebrew) that would
// cause bidi reordering in a terminal, making left-aligned padding appear on the wrong side.
func hasRTL(s string) bool {
	for _, r := range colorCleaner.Replace(s) {
		if unicode.Is(unicode.Arabic, r) || unicode.Is(unicode.Hebrew, r) {
			return true
		}
	}
	return false
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
		r := runes[i]
		// Combining/nonspacing marks are zero-width
		if unicode.Is(unicode.Mn, r) || unicode.Is(unicode.Me, r) {
			continue
		}
		// Format characters are zero-width (e.g. Arabic RTL mark U+200F, directional embeddings)
		if unicode.Is(unicode.Cf, r) {
			continue
		}
		cw := 1
		// CJK characters, fullwidth forms, and emoji are double-width
		if unicode.Is(unicode.Han, r) || unicode.Is(unicode.Hangul, r) ||
			(r >= 0x3040 && r <= 0x309F) || // Hiragana
			(r >= 0x30A0 && r <= 0x30FF) || // Katakana
			(r >= 0x3000 && r <= 0x303F) || // CJK Symbols and Punctuation
			(r >= 0xFF00 && r <= 0xFFEF) || // Fullwidth Forms
			(r >= 0x1F300 && r <= 0x1FAFF) { // Emoji
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
