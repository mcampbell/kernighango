package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, a := range(os.Args) {
		s += sep + a
		sep = " "
	}
	fmt.Println(s)
}
