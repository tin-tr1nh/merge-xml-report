package main

import (
	"flag"
	"io/ioutil"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

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

func main() {
	var outputPath = flag.String("output", "/files/result/merged_coverage.xml", "Output path of merged report file")
	var inputDirPath = flag.String("input", "/files/reports", "Input path of reports dir")
	flag.Parse()

	filePaths, err := ListDir(*inputDirPath)
	check(err)
	err = MergePaths(filePaths, *outputPath)
	check(err)
}
