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

func (f *File) ParseMapOfLines() map[int32]*Line {
	mapOfLines := map[int32]*Line{}
	for i, line := range f.Lines {
		mapOfLines[line.Num] = &f.Lines[i]
	}

	return mapOfLines
}

// Merge with other File
func (f *File) Merge(other File) {
	mapOfLines := other.ParseMapOfLines()
	for i := range f.Lines {
		if pLine, ok := mapOfLines[f.Lines[i].Num]; ok {
			f.Lines[i].Merge(*pLine)
		}
	}

	mapOfLines = f.ParseMapOfLines()
	for _, line := range other.Lines {
		if _, ok := mapOfLines[line.Num]; !ok && line.Count > 0 {
			f.Lines = append(f.Lines, line)
		}
	}

	f.updateClassMetric()
	f.updateFileMetric()
}

// Pull only the count of line that already existed
func (f *File) Pull(other File) {
	mapOfLines := other.ParseMapOfLines()
	for i := range f.Lines {
		if pLine, ok := mapOfLines[f.Lines[i].Num]; ok {
			f.Lines[i].Merge(*pLine)
		}
	}

	f.updateClassMetric()
	f.updateFileMetric()
}

func (f *File) countClassMetric() ClassMetrics {
	metrics := ClassMetrics{}
	for _, line := range f.Lines {
		if line.Type == "stmt" {
			metrics.Statements++
		}
		if line.Type == "method" {
			metrics.Methods++
		}
		if line.Type == "stmt" && line.Count > 0 {
			metrics.CoveredStatements++
		}

		if line.Type == "method" && line.Count > 0 {
			metrics.CoveredMethods++
		}
	}
	metrics.Elements = metrics.Methods + metrics.Statements
	metrics.CoveredElements = metrics.CoveredMethods + metrics.CoveredStatements

	return metrics
}

func (f *File) updateClassMetric() {
	metrics := f.countClassMetric()
	if len(f.Classes) > 0 {
		f.Classes[0].Metrics.Elements = metrics.Elements
		f.Classes[0].Metrics.CoveredElements = metrics.CoveredElements
		f.Classes[0].Metrics.Statements = metrics.Statements
		f.Classes[0].Metrics.CoveredStatements = metrics.CoveredStatements
		f.Classes[0].Metrics.Methods = metrics.Methods
		f.Classes[0].Metrics.CoveredMethods = metrics.CoveredMethods
	}
}

func (f *File) updateFileMetric() {
	metrics := f.countClassMetric()
	if len(f.Classes) > 0 {
		f.Metrics.Elements = metrics.Elements
		f.Metrics.CoveredElements = metrics.CoveredElements
		f.Metrics.Statements = metrics.Statements
		f.Metrics.CoveredStatements = metrics.CoveredStatements
		f.Metrics.Methods = metrics.Methods
		f.Metrics.CoveredMethods = metrics.CoveredMethods
	}
}
