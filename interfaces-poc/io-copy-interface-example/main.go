package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello", os.Args)
	cmdArgs := os.Args[1:]
	filename := ""
	if len(cmdArgs) >= 1 {
		filename = cmdArgs[0]
	}
	filePtr, _ := os.Open(filename)
	io.Copy(os.Stdout, filePtr)
}
