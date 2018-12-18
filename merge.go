package main

import (
	"log"

	"bitbucket.org/hameesys/merge-xml-report/report"
)

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
	}

	return WriteReport(mergedReport, outputPath)
}
