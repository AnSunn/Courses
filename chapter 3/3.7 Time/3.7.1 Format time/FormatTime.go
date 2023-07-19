package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	const layout = time.RFC3339
	firstTime, err := time.Parse(layout, str)
	if err != nil {
		panic(err)
	}
	fmt.Println(firstTime.Format(time.UnixDate))
}
