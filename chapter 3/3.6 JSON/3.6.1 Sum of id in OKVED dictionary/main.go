package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type dict struct {
	Id int64 `json:"global_id"`
}

func sumId(path string) int64 {
	var d []dict
	var sum int64
	resp, err := http.Get(path)
	if err != nil {
		log.Fatalln(err)
	}
	rb, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if err := json.Unmarshal(rb, &d); err != nil {
		log.Fatalln(err)
	}
	for i, _ := range d {
		sum += d[i].Id
	}
	return sum
}
func main() {
	const URL = "https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json"
	var sum = sumId(URL)
	fmt.Println(sum)
}
