package main

import (
	"fmt"
	"time"
)

func main() {
	var min, sec time.Duration
	const c int64 = 1589570165
	fmt.Scanf("%d мин. %d", &min, &sec)
	te := time.Unix(c, 0).UTC()
	k := time.Minute*min + time.Second*sec
	fmt.Println(te.Add(k).Format(time.UnixDate))
}
