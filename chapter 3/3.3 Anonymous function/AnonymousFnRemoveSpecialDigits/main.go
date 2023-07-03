package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, num_n uint
	fmt.Scan(&num)
	fn := func(num uint) uint {
		str := strconv.FormatUint(uint64(num), 10)
		for i := range str {
			if str[i]%2 == 0 && str[i] != '0' {
				num_n = num_n*10 + uint(str[i]-'0')
			}
		}
		if num_n == 0 {
			num_n = 100
		}
		return num_n
	}
	fmt.Println(fn(num))
}
