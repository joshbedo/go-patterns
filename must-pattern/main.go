package main

import (
	"fmt"
	"regexp"
)

func Must[T any](x T, err error) T {
	if err != nil {
		panic(err)
	}
	return x
}

func main() {
	r := Must(regexp.Compile("123"))
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println(r)
}
