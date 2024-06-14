package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This program is uses the flag package")

	flag.Parse()
	filename := GetFilename()
	fmt.Println(filename)
	contentBS, _ := os.ReadFile(filename)

	content := string(contentBS)

	fmt.Println(content)
}
