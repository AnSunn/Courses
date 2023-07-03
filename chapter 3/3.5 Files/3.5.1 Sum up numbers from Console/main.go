package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	// Все указанные в тексте задачи пакеты уже импортированы (strconv, bufio, os, io)
	// Другие использовать нельзя.
	defer func() {
		if r := recover(); r != nil {
			_, err := io.WriteString(os.Stdout, "Recovered from panic")
			if err != nil {
				return
			}
		}
	}()

	str := bufio.NewScanner(os.Stdin)
	var sum int
	for str.Scan() {
		if len(str.Text()) != 0 {
			num, err := strconv.Atoi(str.Text())
			if err != nil {
				panic(err)
			}
			sum += num
		} else {
			break
		}
	}
	_, err := io.WriteString(os.Stdout, strconv.FormatInt(int64(sum), 10))
	if err != nil {
		return
	}
}
