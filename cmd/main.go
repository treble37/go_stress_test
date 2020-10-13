package main

import (
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	totalRequests = 200
	for i := 0; i < totalRequests; i++ {
		st := NewDefaultStressTest()
	}
}
