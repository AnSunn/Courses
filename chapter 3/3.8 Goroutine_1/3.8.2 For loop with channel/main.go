package main

import (
	"fmt"
)

func task2(c chan string, str string) {
	for i := 0; i < 5; i++ {
		c <- str + " "
	}
	defer close(c)
}

func main() {
	var str string
	const N = 5
	c := make(chan string, N)
	fmt.Print("Enter string: ")
	_, err := fmt.Scan(&str)
	if err != nil {
		return
	}
	go task2(c, str)
	for num := range c {
		fmt.Print(num)
	}
}
