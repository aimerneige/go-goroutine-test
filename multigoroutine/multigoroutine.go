package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Hello"
			time.Sleep(time.Second * 1)
		}
	}()
	go func() {
		for {
			c2 <- "World"
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}
