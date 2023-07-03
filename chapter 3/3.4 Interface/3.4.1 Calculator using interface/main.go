package main

import (
	"fmt"
)

// функция проверяет первые два параметра на нужный тип
func getFloat(itrf interface{}) float64 {
	if res, ok := itrf.(float64); !ok {
		panic(fmt.Sprintf("value=%v: %T", itrf, itrf))
	} else {
		return res
	}
}

// функция проверяет третий параметр на нужный тип
func getString(itrf interface{}) string {
	if res, ok := itrf.(string); !ok {
		panic(fmt.Sprintf("value=%v: %T", itrf, itrf))
	} else {
		return res
	}
}

// данная функция написана в условии задачи, служит для проверки кода
func readTask() (interface{}, interface{}, interface{}) {
	return 6.0, 6, "-"
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	// все полученные значения имеют тип пустого интерфейса
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	v1 := getFloat(value1)
	v2 := getFloat(value2)
	op := getString(operation)
	m := map[string]func() float64{
		"+": func() float64 { return v1 + v2 },
		"-": func() float64 { return v1 - v2 },
		"*": func() float64 { return v1 * v2 },
		"/": func() float64 {
			if v2 == 0.0 {
				panic("divide by zero")
			} else {
				return v1 / v2
			}
		},
	}
	if res, ok := m[op]; ok {
		fmt.Printf("%.4f", res())
	} else {
		fmt.Println("неизвестная операция")
	}
}
