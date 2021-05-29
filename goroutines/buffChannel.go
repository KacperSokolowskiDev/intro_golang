package main

import (
	"fmt"
	"time"
)

func main() {

	information := make(chan string, 5) // buffer channels, 5 represents that we can store 5 elements max 

	information <- "Hello"
	information <- "Hi"
	information <- "Hey"
	information <- "Hello world"
	information <- "Bye World"
	// 6th element will cause a deadlock

	fmt.Println(<-information)
	information <- "ABC" // now that 1 element is out of the channel we can again store information 
	fmt.Println(<-information)
	fmt.Println(<-information)
	fmt.Println(<-information)
	fmt.Println(<-information)


	time.Sleep(5 * time.Second)

}