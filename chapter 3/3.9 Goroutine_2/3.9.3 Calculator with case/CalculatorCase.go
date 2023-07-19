package main

import "fmt"

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	c := make(chan int)
	go func(c chan int) {
		defer close(c)
		select {
		case v := <-firstChan:
			c <- v * v
		case v := <-secondChan:
			c <- v * 3
		case <-stopChan:
			break
		}
	}(c)
	return c
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	stop := make(chan struct{})
	r := calculator(ch1, ch2, stop)
	ch1 <- 3
	//ch2 <- 3
	//close(stop)
	fmt.Println(<-r)
}
