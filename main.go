package main

import "os"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	outputPath := "files/result/merged_coverage.xml"
	dir := os.Args[1]

	filePaths, err := ListDir(dir)
	check(err)
	err = MergePaths(filePaths, outputPath)
	check(err)
}
