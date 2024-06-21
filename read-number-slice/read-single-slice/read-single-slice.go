package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func main() {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Space seperated numbers: ")
	scanner.Scan()
	nums = numbers(scanner.Text())
	fmt.Println("You have entered: ", nums)
}
