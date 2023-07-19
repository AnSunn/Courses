package main

import (
	"fmt"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var s string
	defer close(outputStream)
	for v := range inputStream {
		if s != v {
			outputStream <- v
		}
		s = v
	}
}
func main() {
	var str string
	fmt.Print("Enter string: ")
	fmt.Scan(&str)
	in := make(chan string)
	out := make(chan string)
	go removeDuplicates(in, out)
	go func() {
		defer close(in)
		for _, r := range str {
			in <- string(r)
		}
	}()
	for x := range out {
		fmt.Print(x)
	}
}
