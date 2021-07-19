package main

import (
	"fmt"
	"math"
	"unified_log/test"
)

func genlog() {
	var log test.Log
	err := log.GenerateLogs(int(math.Round(1024 * 6.66 * 1024)))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	genlog()
}
