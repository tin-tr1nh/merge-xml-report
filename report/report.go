package report

import (
	"encoding/xml"
)

// Report struct of coverage report
type Report struct {
	XMLName       xml.Name `xml:"coverage"`
	GeneratedTime int32    `xml:"generated,attr"`
	Project       Project  `xml:"project"`
}

// Merge merge two Report into one
func (r *Report) Merge(other *Report) *Report {
	if other == nil {
		return r
	}

	if r == nil {
		return other
	}

	r.Project.Merge(other.Project)

	return r
}
