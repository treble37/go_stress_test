package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	totalRequests := 200
	for i := 0; i < totalRequests; i++ {
		st := NewDefaultStressTest()
		go st.PostRequest(ch)
	}
	for i := 0; i < totalRequests; i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
