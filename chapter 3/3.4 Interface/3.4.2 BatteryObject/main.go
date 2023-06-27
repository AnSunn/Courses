/*
Реализовать объект, удовлетворяющий интерфейсу fmt.Stringer.
Назвать его "Батарейка".
Во-первых, вы должны объявить новый тип, удовлетворяющий интерфейсу fmt.Stringer.
Ваш тип должен предусматривать, что на печати он будет выглядеть так:
[      XXXX]: где пробелы - "опустошенная" емкость батареи, а X - "заряженная".
Во-вторых, на стандартный ввод вы получаете строку, состоящую ровно из 10 цифр:
0 или 1 (порядок 0/1 случайный). Ваша задача считать эту строку любым возможным
способом и создать на основе этой строки объект объявленного вами на первом этапе типа:
надеюсь, вы понимаете, что строка символизирует емкость батарейки: 0 - это "опустошенная"
часть, а 1 - "заряженная".
В-третьих, созданный вами объект должен называться batteryForTest
(использование этого имени обязательно).
*/
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
