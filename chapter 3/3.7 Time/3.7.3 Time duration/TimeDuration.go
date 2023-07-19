package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	const dt = "02.01.2006 15:04:05"
	var substr time.Duration
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}
	spl := strings.Split(str, ",")
	strings.TrimSpace(spl[0])
	spl[1] = strings.ReplaceAll(spl[1], "\x0d\x0a", "")

	tm, errT := time.Parse(dt, spl[0])
	if errT != nil {
		log.Fatalln(errT)
	}
	tm1, errT1 := time.Parse(dt, spl[1])
	if errT1 != nil {
		log.Fatalln(errT1)
	}
	if tm.After(tm1) {
		substr = tm.Sub(tm1)
	} else {
		substr = tm1.Sub(tm)
	}
	fmt.Println(substr)
}
