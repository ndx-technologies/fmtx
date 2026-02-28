package fmtx

import (
	"io"
	"strings"
)

func HeaderTo(w io.StringWriter, title string) {
	w.WriteString(strings.Repeat("═", 100))
	w.WriteString("\n")
	w.WriteString(title)
	w.WriteString("\n")
	w.WriteString(strings.Repeat("═", 100))
	w.WriteString("\n")
}
