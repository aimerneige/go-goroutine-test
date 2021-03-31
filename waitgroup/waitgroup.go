package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		count(5, "Hello")
		wg.Done()
	}()

	go func() {
		go count(5, "World")
		wg.Done()
	}()

	wg.Wait()
}

func count(n int, s string) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, s)
		time.Sleep(time.Millisecond * 200)
	}
}
