package main

import "fmt"

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int {
	c := make(chan int)
	s := 0
	go func(c chan int) <-chan int {
		defer close(c)
		for {
			select {
			case n := <-arguments:
				s += n
			case <-done:
				c <- s
				return c
			}
		}
	}(c)
	return c
}

func main() {
	ch1, r := make(chan int), make(<-chan int)
	stop := make(chan struct{})
	r = calculator(ch1, stop)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(stop)
	fmt.Println(<-r)
}
