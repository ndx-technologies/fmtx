package fmtx

type Alignment uint8

const (
	AlignUndefined Alignment = iota
	AlignLeft
	AlignRight
)

func (a Alignment) Write(out interface{ WriteString(string) (int, error) }, s string, width int) {
	switch a {
	case AlignRight:
		for i := displayWidth(s); i < width; i++ {
			out.WriteString(" ")
		}
		out.WriteString(s)
	case AlignLeft, AlignUndefined:
		// Wrap RTL text in Left-to-Right Isolate so the terminal bidi algorithm
		// keeps the trailing padding spaces visually to the right of the text.
		if hasRTL(s) {
			out.WriteString("\u2066") // LEFT-TO-RIGHT ISOLATE
			out.WriteString(s)
			out.WriteString("\u2069") // POP DIRECTIONAL ISOLATE
		} else {
			out.WriteString(s)
		}
		for i := displayWidth(s); i < width; i++ {
			out.WriteString(" ")
		}
	}
}

type TablCol struct {
	Header    string
	Alignment Alignment
	Width     int
}

func (s TablCol) Write(out interface{ WriteString(string) (int, error) }, v string) {
	if s.Width <= 0 {
		out.WriteString(v)
		return
	}
	s.Alignment.Write(out, truncate(v, s.Width), s.Width)
}

// TableWriter provides simple table formatting with ANSI color support.
// Values that are too long will get truncated into column width.
type TableWriter struct {
	Indent string
	Cols   []TablCol
	Out    interface {
		WriteString(s string) (n int, err error)
	}
}

func (s TableWriter) WriteHeader() {
	s.Out.WriteString(s.Indent)
	for i, col := range s.Cols {
		if i > 0 {
			s.Out.WriteString(" ")
		}
		col.Alignment.Write(s.Out, DimS(col.Header), col.Width)
	}
	s.Out.WriteString("\n")
}

func (s TableWriter) WriteSubHeader(vs ...string) {
	s.Out.WriteString(s.Indent)
	for i, c := range s.Cols {
		if i >= len(vs) {
			break
		}
		if i > 0 {
			s.Out.WriteString(" ")
		}
		c.Write(s.Out, DimS(vs[i]))
	}
	s.Out.WriteString("\n")
}

func (s TableWriter) WriteHeaderLine() {
	s.Out.WriteString(s.Indent)
	for i, col := range s.Cols {
		if i > 0 {
			s.Out.WriteString(" ")
		}
		if EnableColor {
			s.Out.WriteString(Dim)
		}
		for range col.Width {
			s.Out.WriteString("─") // box-drawing
		}
		if EnableColor {
			s.Out.WriteString(Reset)
		}
	}
	s.Out.WriteString("\n")
}

func (s TableWriter) WriteRow(vs ...string) {
	s.Out.WriteString(s.Indent)
	for i, c := range s.Cols {
		if i >= len(vs) {
			break
		}
		if i > 0 {
			s.Out.WriteString(" ")
		}
		c.Write(s.Out, vs[i])
	}
	s.Out.WriteString("\n")
}
