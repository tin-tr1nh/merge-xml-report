package report

// FileMetrics struct of coverage report's metrics tag in file tag
type FileMetrics struct {
	ClassMetrics
	Loc   int32 `xml:"loc,attr"`
	NcLoc int32 `xml:"ncloc,attr"`
}

// File struct of coverage report's file tag
type File struct {
	Name    string      `xml:"name,attr"`
	Classes []Class     `xml:"class"`
	Lines   []Line      `xml:"line"`
	Metrics FileMetrics `xml:"metrics"`
}

func (f *File) parseMapOfLines() map[int32]*Line {
	mapOfLines := map[int32]*Line{}
	for i, line := range f.Lines {
		mapOfLines[line.Num] = &f.Lines[i]
	}

	return mapOfLines
}

// Merge with other File
func (f *File) Merge(other File) {
	mapOfLines := other.parseMapOfLines()
	for i := range f.Lines {
		if pLine, ok := mapOfLines[f.Lines[i].Num]; ok {
			f.Lines[i].Merge(*pLine)
		}
	}

	mapOfLines = f.parseMapOfLines()
	for _, line := range other.Lines {
		if _, ok := mapOfLines[line.Num]; !ok {
			f.Lines = append(f.Lines, line)
		}
	}
}
