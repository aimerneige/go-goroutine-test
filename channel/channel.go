package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go count(5, "A", c)
	for message := range c {
		fmt.Println(message)
	}
}

func count(n int, animal string, c chan string) {
	for i := 0; i < n; i++ {
		c <- animal
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
