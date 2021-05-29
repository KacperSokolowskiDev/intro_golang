package main

import (
	"fmt"
	"time"
)

// function sending the information, specifying the parameter, type (chan) and if it receives or sends (the arrow) + the content (info)
func sendInfo(channel chan<- string, info string) {
	channel <- info
}

// function receiving the information and printing it 
func printInfo(channel <-chan string) {
	fmt.Println(<-channel)
}

func main() {

	//channels -> sending & receiving info (this time strings)
	information := make(chan string) // defining a channel

	go sendInfo(information, "Hello") // goroutines is similar to async, it makes functions work independently 
	go printInfo(information)

	time.Sleep(2 * time.Second)

}