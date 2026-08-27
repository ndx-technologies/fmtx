package fmtx

import (
	"fmt"
	"strings"
)

// TableWriter draws aligned, padded and truncated columns; headers are dimmed.
func ExampleTableWriter() {
	orig := EnableColor
	EnableColor = false // keep the docs readable; color is on by default
	defer func() { EnableColor = orig }()

	var b strings.Builder
	tw := TableWriter{
		Indent: "  ",
		Cols: []TablCol{
			{Header: "Fruit", Alignment: AlignLeft, Width: 10},
			{Header: "Count", Alignment: AlignRight, Width: 6},
			{Header: "Share", Alignment: AlignRight, Width: 6},
		},
		Out: &b,
	}
	tw.WriteHeader()
	tw.WriteHeaderLine()
	tw.WriteRow("apples", "42", "67%")
	tw.WriteRow("bananas", "7", "11%")
	tw.WriteRow("cherries", "14", "22%")
	fmt.Print(b.String())
	// Output:
	//   Fruit       Count  Share
	//   ────────── ────── ──────
	//   apples         42    67%
	//   bananas         7    11%
	//   cherries       14    22%
}

// WriteSubHeader writes a secondary, dimmed line (e.g. units) under the header.
func ExampleTableWriter_WriteSubHeader() {
	orig := EnableColor
	EnableColor = false
	defer func() { EnableColor = orig }()

	var b strings.Builder
	tw := TableWriter{
		Indent: "  ",
		Cols: []TablCol{
			{Header: "Item", Alignment: AlignLeft, Width: 10},
			{Header: "Price", Alignment: AlignRight, Width: 6},
		},
		Out: &b,
	}
	tw.WriteHeader()
	tw.WriteSubHeader("(kg)", "(USD)") // units under the headers
	tw.WriteRow("apples", "1.5")
	fmt.Print(b.String())
	// Output:
	//   Item        Price
	//   (kg)        (USD)
	//   apples        1.5
}

// TablCol writes a single cell, padding and truncating to its width.
func ExampleTablCol_Write() {
	orig := EnableColor
	EnableColor = false
	defer func() { EnableColor = orig }()

	var b strings.Builder
	c := TablCol{Header: "Title", Alignment: AlignLeft, Width: 10}
	c.Write(&b, "This is a very long title") // truncated to the width
	fmt.Printf("%q\n", b.String())

	b.Reset()
	c = TablCol{Header: "Count", Alignment: AlignRight, Width: 5}
	c.Write(&b, "42") // right-aligned
	fmt.Printf("%q\n", b.String())
	// Output:
	// "This is..."
	// "   42"
}

// Alignment.Write pads text in a given direction.
func ExampleAlignment_Write() {
	orig := EnableColor
	EnableColor = false
	defer func() { EnableColor = orig }()

	var b strings.Builder
	AlignLeft.Write(&b, "abc", 5)
	fmt.Printf("%q\n", b.String())

	b.Reset()
	AlignRight.Write(&b, "abc", 5)
	fmt.Printf("%q\n", b.String())
	// Output:
	// "abc  "
	// "  abc"
}
