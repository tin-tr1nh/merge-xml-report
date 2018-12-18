package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"bitbucket.org/hameesys/merge-xml-report/report"
)

// ListDir list all file in the dir
func ListDir(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	filePaths := []string{}
	for _, file := range files {
		if !file.IsDir() {
			filePaths = append(filePaths, dir+"/"+file.Name())
		}
	}

	return filePaths, nil
}

func printReportInfo(r *report.Report) error {
	fmt.Println("Report info")
	if r == nil {
		return fmt.Errorf("Report is nil")
	}

	fmt.Printf("Num of file: %v\n", len(r.Project.Files))

	totalLine := 0
	totalClass := 0

	for _, file := range r.Project.Files {
		totalLine += len(file.Lines)
		totalClass += len(file.Classes)
	}

	fmt.Printf("Total of line: %v\n", totalLine)
	fmt.Printf("Total of class: %v\n", totalClass)

	return nil
}

// MergePaths merge xml report in the dir
func MergePaths(inputPaths []string, outputPath string) error {
	var mergedReport *report.Report
	for _, inputPath := range inputPaths {
		report, err := ReadReport(inputPath)
		if err != nil {
			return err
		}

		log.Println("Merge report")
		mergedReport = mergedReport.Merge(report)
		printReportInfo(report)
	}

	printReportInfo(mergedReport)
	return WriteReport(mergedReport, outputPath)
}

// ReadReport read path file and return Report object
func ReadReport(path string) (*report.Report, error) {
	log.Printf("Read file %v to string \n", path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	log.Println("Unmarshal to object")
	report := &report.Report{}
	err = xml.Unmarshal(data, report)
	if err != nil {
		return nil, err
	}

	return report, nil
}

// WriteReport write Report to path file
func WriteReport(r *report.Report, path string) error {
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	outputPath := "files/result/coverage_dump.xml"
	dir := os.Args[1]

	filePaths, err := ListDir(dir)
	check(err)
	err = MergePaths(filePaths, outputPath)
	check(err)
}
