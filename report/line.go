package report

// Line struct of coverage report's line tag
type Line struct {
	Num   int32  `xml:"num,attr"`
	Type  string `xml:"type,attr"`
	Count int32  `xml:"count,attr"`
	Name  string `xml:"name,attr,omitempty"`
}

// Merge with other Line
func (l *Line) Merge(other Line) {
	l.Count = l.Count + other.Count
}
