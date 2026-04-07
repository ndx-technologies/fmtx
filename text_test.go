package fmtx

import (
	"testing"
)

func TestDisplayWidth(t *testing.T) {
	tests := map[string]int{
		"":                     0,
		"hello":                5,
		"12345":                5,
		"hello 123":            9,
		"こんにちは":                10, // 5 hiragana chars * 2
		"コンニチハ":                10, // 5 katakana chars * 2
		"漢字":                   4,  // 2 kanji chars * 2
		"hello 漢字":             10, // 6 + 4
		RedS("hello"):          5,  // ANSI stripped
		GreenS("こんにちは"):        10, // ANSI stripped, CJK counted
		DimS("test"):           4,
		"⭐":                    2, // U+2B50 wide emoji star
		"⚠️":                   2, // U+26A0 + VS-16: text symbol upgraded to emoji
		"👇🏻":                   2, // base emoji + skin-tone modifier = 1 unit
		"👇🏻💬":                  4, // skin-tone sequence (2) + chat bubble (2)
		"💎":                    2, // emoji is double-width
		"🔑":                    2, // emoji is double-width
		"💎Store%":              8, // 2 + 6
		"🔑Perm%":               7, // 2 + 5
		"ช้อปปี้":              4, // 4 base glyphs * 1 + 3 combining marks * 0
		"كلنا امن":             8, // Arabic: 7 letters + 1 space, each 1 wide
		"\u200Fكلنا امن":       8, // Arabic with U+200F RIGHT-TO-LEFT MARK (zero-width Cf)
		"\u200Eكلنا امن":       8, // Arabic with U+200E LEFT-TO-RIGHT MARK (zero-width Cf)
		"\u202Bكلنا امن\u202C": 8, // Arabic with directional embedding marks (zero-width Cf)
	}

	for s, w := range tests {
		t.Run(s, func(t *testing.T) {
			if displayWidth(s) != w {
				t.Error(s, w, displayWidth(s))
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		s         string
		width     int
		truncated string
	}{
		{"hello", 10, "hello"},
		{"hello world", 8, "hello..."},
		{"こんにちは世界", 10, "こんに..."}, // 5 chars * 2 = 10, should truncate
		{RedS("hello world"), 8, RedS("hello...")},
		{GreenS("こんにちは世界"), 10, GreenS("こんに...")},
		{"hello", 5, "hello"},
		{"漢字", 4, "漢字"}, // 2 chars * 2 = 4
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			result := truncate(tt.s, tt.width)
			if result != tt.truncated {
				t.Error(tt.s, tt.width, result, tt.truncated)
			}
			if displayWidth(colorCleaner.Replace(result)) > tt.width {
				t.Error(tt.s, tt.width, result, tt.truncated)
			}
		})
	}
}
