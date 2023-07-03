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
