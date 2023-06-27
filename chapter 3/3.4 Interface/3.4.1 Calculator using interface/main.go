/*
/*Обязательные условия выполнения: данные со стандартного ввода читаются функцией readTask(),
которая возвращает 3 значения типа пустой интерфейс. Эта функция использует пакеты encoding/json, fmt,
и os - не удаляйте их из импорта. Скорее всего, вам понадобится пакет "fmt", но вы можете использовать
любой другой пакет для записи в стандартный вывод не удаляя fmt.
Итак, вы получаете 3 значения типа пустой интерфейс: если все удачно, то первые 2 значения будут float64.
Третье значение в идеальном случае будет строкой, которая может иметь значения:
"+", "-", "*", "/" (определенная математическая операция). Но такие идеальные случаи будут не всегда,
вы можете получить значения других типов, а также строка в третьем значении может не относится к одной из
указанных математических операций.
Результат выполнения программы должен быть такой:
в штатной ситуации вы должны напечатать в стандартный вывод результат выполнения математической операции
с точностью до 4 цифры после запятой (fmt.Printf(%.4f)); если первое или второе значение не является типом float64,
вы должны напечатать сообщение об ошибке вида: value=полученное_значение: тип_значения (например: value=true: bool)
если третье значение имеет неверный тип или передан знак, не относящийся к указанным выше математическим операциям,
сообщение об ошибке должно иметь вид: неизвестная операция
Гарантируется, что ошибка в аргументах может быть только одна, поэтому если вы при проверке первого значения
увидели, что оно содержит ошибку - печатайте сообщение об ошибке и завершайте работу программы, проверка остальных
аргументов значения уже не имеет, а проверяющая система воспримет 2 сообщения об ошибке как нарушение условия
выполнения задания.
*/
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