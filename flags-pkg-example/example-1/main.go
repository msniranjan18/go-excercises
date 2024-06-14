package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This program is uses the flag package")

	filePathPtr := flag.String("filepath", "", "inputfile name with path")
	fileNamePtr := flag.String("filename", "", "inputfile name with path")

	flag.Parse()

	contentBS, _ := os.ReadFile(*filePathPtr + "/" + *fileNamePtr)

	content := string(contentBS)

	fmt.Println(content)
}
