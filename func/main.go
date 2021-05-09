package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	test()
	test()
	fmt.Println("main")
	wg.Wait()
}

func test() {
	go func() {
		defer wg.Done()
		for i := 0; i < 8; i++ {
			fmt.Println("Hello")
			time.Sleep(time.Millisecond * 1000)
			fmt.Println("World")
		}
	}()
	fmt.Println("done")
}
