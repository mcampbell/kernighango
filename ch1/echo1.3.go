package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	taken := time.Since(start).Seconds()

	fmt.Printf("%.8f Seconds\n", taken)
}
