package main

import (
	"fmt"
	"strings"
)

type battery string

// String functions is created to satisfy Stringer interface
func (b battery) String() string {
	return (fmt.Sprintf("[%10s]", strings.Repeat("X", strings.Count(string(b), "1"))))
}

func main() {
	var batteryForTest battery
	/* if it is required to take initial value using random, then use the code below
	var rnd int64
	for i := 0; i < 10; i++ {
		rnd = int64(rand.Intn(2))
		batteryForTest = batteryForTest + battery(strconv.FormatInt(rnd, 10))
		fmt.Print(rnd)
	}
	fmt.Println()
	fmt.Println(batteryForTest)
	*/
	//else
	fmt.Scan(&batteryForTest)
	fmt.Println(batteryForTest)
}
