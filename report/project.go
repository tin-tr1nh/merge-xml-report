package report

// ProjectMetrics struct of coverage report's metrics tag in project tag
type ProjectMetrics struct {
	FileMetrics
	Files int32 `xml:"files,attr"`
}

// Project struct of coverage report's project tag
type Project struct {
	Name           string         `xml:"name,attr"`
	TimeStamp      int32          `xml:"timestamp,attr"`
	Files          []File         `xml:"file"`
	ProjectMetrics ProjectMetrics `xml:"metrics"`
}

func (p *Project) ParseMapOfFiles() map[string]*File {
	mapOfFiles := map[string]*File{}
	for i, file := range p.Files {
		mapOfFiles[file.Name] = &p.Files[i]
	}

	return mapOfFiles
}

// Merge with other Project
func (p *Project) Merge(other Project) {
	// get the files of other that has the same name
	// with files of p and merge them
	mapFiles := other.ParseMapOfFiles()
	for i := range p.Files {
		if pFile, ok := mapFiles[p.Files[i].Name]; ok {
			p.Files[i].Merge(*pFile)
		}
	}

	// if there are some files of other that not exist in p
	// add it to p
	mapFiles = p.ParseMapOfFiles()
	for _, file := range other.Files {
		if _, ok := mapFiles[file.Name]; !ok {
			p.Files = append(p.Files, file)
		}
	}
}
