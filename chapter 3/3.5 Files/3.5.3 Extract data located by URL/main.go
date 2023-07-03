/*
В тестовом файле, который вы можете скачать из нашего репозитория
https://github.com/semyon-dev/stepik-go/tree/master/work_with_files/task_sep_1 на github.com,
содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на какой позиции
находится число 0 и указать её в качестве ответа. Требуется вывести именно позицию числа,
а не индекс (то-есть порядковый номер, нумерация с 1).

Например:  12;234;6;0;78 :
Правильный ответ будет порядковый номер числа: 4
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Position_num_of_elem(path string) {
	var pos = 1
	data, err := http.Get(path)
	if err != nil {
		log.Fatalln(err)
	}
	r := bufio.NewReader(data.Body)
	defer data.Body.Close()
	for {
		str, err2 := r.ReadString(';')
		if err2 == io.EOF {
			break
		}
		if err2 != nil {
			log.Fatalln(err2)
		}
		if str == "0;" {
			fmt.Println(pos)
		}
		pos++
	}

}
func main() {
	const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_files/task_sep_1/task.data?raw=true"
	Position_num_of_elem(URL)
	//to see the result of exe file
	//the beginning
	var k int
	fmt.Scan(&k)
	//the end
}
