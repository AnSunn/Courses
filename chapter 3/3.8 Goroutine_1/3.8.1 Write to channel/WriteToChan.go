package main

import "fmt"

func task(c chan int, N int) {
	c <- N + 1
}
func main() {
	var N int
	c := make(chan int)
	fmt.Print("Enter N: ")
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println("Incorrect value")
		return
	}
	go task(c, N)
	fmt.Println("The result of N+1 is:", <-c)
	defer close(c)
}
