package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))

	taken := time.Since(start).Seconds()
	fmt.Printf("%.8f Seconds\n", taken)
}
