package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func reverse(s string) string {
	return stringutil.Reverse(s)
}

func main() {
	srcString := "Hello, OTUS!"

	retString := reverse(srcString)

	fmt.Println(retString)
}
