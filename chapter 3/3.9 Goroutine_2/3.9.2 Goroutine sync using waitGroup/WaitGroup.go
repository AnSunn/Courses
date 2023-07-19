package main

import (
	"fmt"
	"log"
	"sync"
)

var d string

func work() {
	fmt.Println("Goroutine start")
	d = "nil"
	fmt.Println("Goroutine stop")
}
func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, id int) {
			fmt.Printf("Goroutine %d is in process \n", id)
			work()
			defer wg.Done()
		}(wg, i)
	}
	wg.Wait()
	if d != "nil" {
		log.Fatal("Error")
	}
	log.Print("All right")
}
