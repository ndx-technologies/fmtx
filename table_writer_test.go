package fmtx

import (
	"strings"
	"testing"
)

func TestDisplayWidth(t *testing.T) {
	tests := map[string]int{
		"":              0,
		"hello":         5,
		"12345":         5,
		"hello 123":     9,
		"こんにちは":         10, // 5 hiragana chars * 2
		"コンニチハ":         10, // 5 katakana chars * 2
		"漢字":            4,  // 2 kanji chars * 2
		"hello 漢字":      10, // 6 + 4
		RedS("hello"):   5,  // ANSI stripped
		GreenS("こんにちは"): 10, // ANSI stripped, CJK counted
		DimS("test"):    4,
		"💎":             2, // emoji is double-width
		"🔑":             2, // emoji is double-width
		"💎Store%":       8, // 2 + 6
		"🔑Perm%":        7, // 2 + 5
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

func TestTableWriter_WriteRow(t *testing.T) {
	tests := []struct {
		cols []TablCol
		data []string
		row  string // without trailing newline
	}{
		{
			cols: []TablCol{
				{Header: "Name", Alignment: AlignLeft, Width: 10},
				{Header: "Age", Alignment: AlignRight, Width: 3},
			},
			data: []string{"Alice", "25"},
			row:  "Alice       25",
		},
		{
			cols: []TablCol{
				{Header: "名前", Alignment: AlignLeft, Width: 8},
				{Header: "年齢", Alignment: AlignRight, Width: 4},
			},
			data: []string{"田中", "30"},
			row:  "田中       30",
		},
		{
			cols: []TablCol{
				{Header: "Status", Alignment: AlignLeft, Width: 10},
				{Header: "Count", Alignment: AlignRight, Width: 5},
			},
			data: []string{GreenS("OK"), RedS("123")},
			row:  GreenS("OK") + "        " + " " + "  " + RedS("123"),
		},
		{
			cols: []TablCol{
				{Header: "Title", Alignment: AlignLeft, Width: 10},
			},
			data: []string{"This is a very long title that should be truncated"},
			row:  "This is...",
		},
		{
			cols: []TablCol{
				{Header: "Title", Alignment: AlignLeft, Width: 10},
			},
			data: []string{BlueS("This is a very long title that should be truncated")},
			row:  BlueS("This is..."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.row, func(t *testing.T) {
			var output strings.Builder
			tw := TableWriter{Cols: tt.cols, Out: &output}

			tw.WriteRow(tt.data...)

			result := strings.TrimSuffix(output.String(), "\n") // Remove trailing newline
			if result != tt.row {
				t.Error(result, tt.row)
			}
		})
	}
}

func TestTableWriter_WriteHeader(t *testing.T) {
	tests := []struct {
		name   string
		cols   []TablCol
		header string // expected full header output (including separator)
	}{
		{
			name: "basic header",
			cols: []TablCol{
				{Header: "Name", Alignment: AlignLeft, Width: 10},
				{Header: "Age", Alignment: AlignRight, Width: 3},
			},
			header: DimS("Name") + "      " + " " + DimS("Age") + "\n" +
				DimS("──────────") + " " + DimS("───") + "\n",
		},
		{
			name: "cjk header",
			cols: []TablCol{
				{Header: "名前", Alignment: AlignLeft, Width: 8},
			},
			header: DimS("名前") + "    " + "\n" +
				DimS("────────") + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var output strings.Builder
			tw := TableWriter{Cols: tt.cols, Out: &output}
			tw.WriteHeader()
			tw.WriteHeaderLine()

			if header := output.String(); header != tt.header {
				t.Error(header, tt.header)
			}
		})
	}
}

func TestColorCleaner(t *testing.T) {
	tests := map[string]string{
		"hello":                         "hello",
		RedS("hello"):                   "hello",
		GreenS("world"):                 "world",
		DimS("test"):                    "test",
		RedS("hello") + GreenS("world"): "helloworld",
		"hello" + Reset + "world":       "helloworld",
	}
	for s, clean := range tests {
		t.Run(s, func(t *testing.T) {
			if v := colorCleaner.Replace(s); v != clean {
				t.Error(s, clean, v)
			}
		})
	}
}

func TestNoColor(t *testing.T) {
	original := EnableColor
	defer func() { EnableColor = original }()

	EnableColor = true
	if RedS("test") == "test" {
		t.Error("Expected colored output when colors enabled")
	}

	EnableColor = false
	if RedS("test") != "test" {
		t.Error("Expected plain text when colors disabled")
	}
	if GreenS("hello") != "hello" {
		t.Error("Expected plain text when colors disabled")
	}
	if BoldS("bold") != "bold" {
		t.Error("Expected plain text when colors disabled")
	}
}
