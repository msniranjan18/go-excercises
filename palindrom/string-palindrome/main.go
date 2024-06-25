package main

import (
	"fmt"
)

func main() {
	fmt.Println("String-Palindrome")
	var str string
	fmt.Println("Enter a string without space")
	fmt.Scan(&str)
	//fmt.Println(str)
	fmt.Println(isStringPalindrome(str))

}

func isStringPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("%v---%v\n", i, j)
		fmt.Printf("%v==%v\n", string(str[i]), string(str[j]))
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
