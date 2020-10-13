package main

import (
	"fmt"
	"time"
	"github.com/treble37/stress_test/pkg/stress"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	totalRequests := 200
	for i := 0; i < totalRequests; i++ {
		st := stress.NewDefaultStressTest()
		go st.PostRequest(ch)
	}
	for i := 0; i < totalRequests; i++ {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
