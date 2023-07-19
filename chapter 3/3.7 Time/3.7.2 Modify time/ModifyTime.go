package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	const t = "13:00:00"
	const to = "15:04:05"
	const dt = "2006-01-02 15:04:05"
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	str = strings.TrimSpace(str)
	inpttime, err := time.Parse(dt, str)
	switch {
	case inpttime.Format(to) < t:
		fmt.Println(inpttime.Format(dt))
	case inpttime.Format(to) >= t:
		fmt.Println(inpttime.AddDate(0, 0, 1).Format(dt))
	}
}
