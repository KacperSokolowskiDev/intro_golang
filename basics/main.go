package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello world")

	var word = "Test"
	var b int8 = 10
	fmt.Println((reflect.TypeOf(word)))
	fmt.Println((reflect.TypeOf(b)))

	x:= 10 //initial declaration
	x = 20
	fmt.Println(x)

	const PI = 3.14159
	fmt.Println(PI)

	a := 0
	a++ // but ++a is not working
	fmt.Println(a)

	//input 
	var myinput int
	fmt.Println("Please enter a number: ")
	fmt.Scan(&myinput)
	fmt.Printf("Your number was %d \n", myinput) //formated

	switch(myinput) {
		case 10:
			fmt.Println("10 was your input")
		case 20: 
			fmt.Println("20 was your input")
		default:
			fmt.Println("Something else")
	}

	//loops

	y := 0
	for y < 10 {
		y++
		fmt.Println(y)
	}

	for z := 0; z < 3; z++ {
		fmt.Println(z)
	}
}