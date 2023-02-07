package main

import (
	"fmt"
	"os"
	"strconv"
)


func main() {
	s, sep, nl := "", "", "\n"
	for i, arg := range os.Args[:] {
		s += arg + sep + strconv.FormatInt(int64(i), 10) + nl
		sep = " "
	}
	fmt.Println(s)
}
