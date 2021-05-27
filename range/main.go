package main

import "fmt"

func main() {

	numbers := []int{10,728,22,1,90}

	for idx, number := range numbers {
		fmt.Printf("Value at index %d is %d\n", idx, number)
	}

	for _, number := range numbers { // underscore is a placeholder used when we want to avoid the index but we can't leave it blank
		fmt.Printf("Value is %d\n", number)
	}

	myMap := map[string]int{"A": 1, "B": 19, "C": 4}

	for key, value := range myMap {
		fmt.Printf("Key is %s and Value is %d\n", key, value)
	}

}