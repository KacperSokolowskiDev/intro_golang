package main

import "fmt"

func main() {

	hello()
	fmt.Println(sum(2,4))
	fmt.Println(minus(20,3,10))
	fmt.Println(calcAdd(10,20,43,2220,1,902))
	arr := []int{1,2,3,4,5,6,7,8}
	fmt.Println(calcAdd(arr...))

	next_even_number := even_numbers()
	fmt.Println(next_even_number())
	fmt.Println(next_even_number())
	fmt.Println(next_even_number())
}

func hello() {
	fmt.Println("Hello world")
}

func sum(x int, y int) int {
	return x + y
}

func minus(a,b,c int) int {
	return a - b - c
}

func calcAdd(numbers ...int) int { //unspecified number of parameters
	mysum := 0
	for _, number := range numbers {
		mysum += number
	}
	return mysum
}

// Closure (generated sequence)
func even_numbers() func() int { // the returned value is a function
	i := 0
	return func() int {
		i += 2
		return i
	}
}