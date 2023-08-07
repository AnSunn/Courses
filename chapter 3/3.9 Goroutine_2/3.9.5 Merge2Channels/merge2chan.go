package main

import (
	"fmt"
	"sync"
	"time"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	st := make(chan struct{})
	go func(st chan struct{}) {
		wg := new(sync.WaitGroup)
		defer close(out)
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2
			wg.Add(2)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				x1 = fn(x1)
			}(wg)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				x2 = fn(x2)
			}(wg)
			wg.Wait()
			out <- x1 + x2

		}
	}(st)

}

func main() {
	in1 := make(chan int, 10)
	in2 := make(chan int, 10)
	out := make(chan int, 10)
	n := 10
	for i := 0; i < 10; i++ {
		in1 <- i
		in2 <- i
	}
	go merge2Channels(fn, in1, in2, out, n)
	fmt.Println("not blocked if printed first")
	time.Sleep(time.Second * 3)
	for i := 0; i < n; i++ {
		fmt.Println(<-out)

	}

}

func fn(x int) int {
	return x
}
