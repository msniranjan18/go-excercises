package main

import (
	"fmt"
	"reflect"
)

type User struct {
	FName string `color:"dark" default:"RAM"`
	LName string `foo:"bar" ram:"shyam"`
}

func main() {
	fmt.Println("Use-of-custom-struct-tags")
	u := User{}

	uType := reflect.TypeOf(u)

	fnameField := uType.Field(0)
	lnameField := uType.Field(1)

	fmt.Println(fnameField.Tag.Get("color"), fnameField.Tag.Get("default"))
	fmt.Println(lnameField.Tag.Get("foo"), lnameField.Tag.Get("ram"))

}

/*
go run main.go

Output:
Use-of-custom-struct-tags
dark RAM
bar shyam
*/

/*
Note: Please check below examples:
https://github.com/msniranjan18/go-gin-basic/blob/master/ex2/main.go
https://github.com/msniranjan18/go-excercises/blob/master/custom-struct-tags/main.go
*/
