package main

import (
	"flag"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var outputPath = flag.String("output", "/files/result/merged_coverage.xml", "Output path of merged report file")
	var inputDirPath = flag.String("input", "/files/reports", "Input path of reports dir")
	flag.Parse()

	filePaths, err := ListDir(*inputDirPath)
	check(err)
	err = MergePaths(filePaths, *outputPath)
	check(err)
}
