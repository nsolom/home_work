package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

func main() {
	greeting := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(greeting))
}
