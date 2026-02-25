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
		out.WriteString(s)
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
	Indent     string
	Cols       []TablCol
	ColDefault map[string]TablCol
	Out        interface {
		WriteString(s string) (n int, err error)
	}
}

func (s TableWriter) withDefault(v TablCol) TablCol {
	if s.ColDefault == nil {
		return v
	}
	if d, ok := s.ColDefault[v.Header]; ok {
		if v.Width == 0 {
			v.Width = d.Width
		}
		if v.Alignment == AlignUndefined {
			v.Alignment = d.Alignment
		}
	}
	return v
}

func (s TableWriter) WriteHeader() {
	s.Out.WriteString(s.Indent)
	for i, col := range s.Cols {
		if i > 0 {
			s.Out.WriteString(" ")
		}
		col = s.withDefault(col)
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
		s.withDefault(c).Write(s.Out, DimS(vs[i]))
	}
	s.Out.WriteString("\n")
}

func (s TableWriter) WriteHeaderLine() {
	s.Out.WriteString(s.Indent)
	for i, col := range s.Cols {
		if i > 0 {
			s.Out.WriteString(" ")
		}
		s.Out.WriteString(Dim)
		for range s.withDefault(col).Width {
			s.Out.WriteString("─") // box-drawing
		}
		s.Out.WriteString(Reset)
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
		s.withDefault(c).Write(s.Out, vs[i])
	}
	s.Out.WriteString("\n")
}
