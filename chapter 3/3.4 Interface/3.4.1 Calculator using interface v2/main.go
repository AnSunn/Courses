package main

import (
	"fmt"
)

// функция проверяет первые два параметра на нужный тип
func readTask() (interface{}, interface{}, interface{}) {
	return 6.0, 4.7, "-"
}

func action(a, b float64, symbol string) (res float64) {
	if symbol == "+" {
		res = a + b
	} else if symbol == "*" {
		res = a * b
	} else if symbol == "/" {
		res = a / b
	} else if symbol == "-" {
		if b == 0.0 {
			panic("divide by zero")
		} else {
			res = a - b
		}
	}
	return res
}

func contains(elems []string, v string) bool {
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	v1, ok1 := value1.(float64)
	v2, ok2 := value2.(float64)
	optn, ok3 := operation.(string)
	if contains([]string{"*", "/", "+", "-"}, optn) {
		if ok1 && ok2 && ok3 {
			fmt.Printf("%.4f", action(v1, v2, optn))
		} else if !ok1 {
			fmt.Printf("value=%v: %T", value1, value1)
		} else if !ok2 {
			fmt.Printf("value=%v: %T", value2, value2)
		}
	} else {
		fmt.Println("неизвестная операция")
	}
}
