package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ZipReader(path string, err error) {
	resp, err := http.Get(path) //get request
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body) //read body answer
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	reader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		fmt.Println("Ошибка", err)
	}
	for _, elem := range reader.File {
		if !elem.FileInfo().IsDir() { // if there is a file
			el, err1 := elem.Open() // then open it
			if err1 != nil {
				panic(err1)
			}
			defer func(el io.ReadCloser) {
				err := el.Close()
				if err != nil {
					panic(err)
				}
			}(el)
			r1 := csv.NewReader(el)    //create reader of csv data
			data, err2 := r1.ReadAll() //read csv data
			if err2 != nil {
				panic(err2)
			}
			if len(data) != 10 || len(data[0]) != 10 {
				continue
			} else if len(data) == 10 {
				fmt.Println(data[4][2])
			}
			err3 := el.Close()
			if err3 != nil {
				panic(err3)
			}
		}
	}
	return
}

// If you want to read data from the file, downloaded at your local machine, then use the code below
//version 2
/*
	func Walkfunc(path string, info os.FileInfo, err2 error) error {
		reader, _ := zip.OpenReader(path) //open archive
		defer func(reader *zip.ReadCloser) {
			err := reader.Close()
			if err != nil {

			}
		}(reader)
		for _, elem := range reader.File { //look through every file
			if !elem.FileInfo().IsDir() { // if there is a file
				el, err := elem.Open() // then open it
				if err != nil {
					panic(err)
				}
				r1 := csv.NewReader(el)    //create reader of csv data
				data, err1 := r1.ReadAll() //read csv data
				if err1 != nil {
					panic(err1)
				}
				if len(data) != 10 || len(data[0]) != 10 {
					continue
				} else if len(data) == 10 {
					fmt.Println(data[4][2])
				}
				err4 := el.Close()
				if err4 != nil {
					return err4
				}
			}
			if err2 != nil {
				return err2
			}
		}
		return nil

}
*/
// The end of version 2

func main() {
	const URL = "https://github.com/semyon-dev/stepik-go/blob/master/work_with_files/task_csv_1/task.zip?raw=true"
	var err error
	ZipReader(URL, err)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from func ", r)
		}
	}()

	// If you want to read data from the file, downloaded at your local machine, then use the code below
	//version 2
	/*	path := "C:\\Users\\Nastya\\Downloads\\task.zip"
		if err := filepath.Walk(path, Walkfunc); err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from func ", r)
			}
		}()*/
	// The end of version 2
}
