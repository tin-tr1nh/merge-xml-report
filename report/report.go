package report

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
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

// ReadReport read path file and return Report object
func ReadReport(path string) (*Report, error) {
	log.Printf("Read file %v to string \n", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	log.Println("Unmarshal to object")
	report := &Report{}
	err = xml.Unmarshal(data, report)
	if err != nil {
		return nil, err
	}

	return report, nil
}

// WriteReport write Report to path file
func WriteReport(r *Report, path string) error {
	if r == nil {
		return fmt.Errorf("Report object is nil")
	}

	log.Printf("Marshal to file %v \n", path)
	output, err := xml.MarshalIndent(*r, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, output, 0644)
	if err != nil {
		return err
	}

	log.Println("Finished")
	return nil
}
