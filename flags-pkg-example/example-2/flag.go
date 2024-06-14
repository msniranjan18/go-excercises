package main

import (
	"flag"
)

var (
	filePathPtr = flag.String("filepath", "", "inputfile name with path")
	fileNamePtr = flag.String("filename", "", "inputfile name with path")
)

func GetFilename() string {
	return *filePathPtr + "/" + *fileNamePtr
}
