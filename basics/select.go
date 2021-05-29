package main

import (
	"fmt"
	"time"
)

func main() {

	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		one <- "hey"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		two <- "Hello"
	}()

	for x := 0; x < 2; x++ {
		// select is similar to the switch statement
		select {
			case rec1 := <-one:
				fmt.Println("I received from channel one", rec1)
			case rec2 := <-two:
				fmt.Println("I received from channel two", rec2)
		}
	}

}