package main

import (
	"log"
)

var d string

func work() {
	d = "nil"
}
func main() {
	c := make(chan struct{})
	go func(c chan struct{}) {
		work()
		close(c)
	}(c)
	<-c
	if d != "nil" {
		log.Fatal("Error")
	}
	log.Print("All right")
}
