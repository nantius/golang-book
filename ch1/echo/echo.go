package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var s, sep string
	start := time.Now()
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i] + " " + strconv.Itoa(i)
		sep = " \n"
	}
	fmt.Printf("%.10fs elapsed\n",time.Since(start).Seconds())
	fmt.Println(s)
}
