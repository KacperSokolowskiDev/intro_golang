package main

import "fmt"

func change_value(val int) { // the value won't change outside of the scope of this function by calling with the value
	val = 200
	fmt.Println(val)
}

// "*int" -> integer pointer, the value will 
//change even outside the scope of this func because we change it from it's memory address not the value
func change_ref(val *int) { 
	*val = 500
	fmt.Println(*val)
} 

func main() {

	var a int
	a = 10

	ptr := &a // pointing at the memory address of the element a 

	fmt.Println(ptr)
	fmt.Println(*ptr) // the sign "*" is used to depoint and print the value of element a

	fmt.Println("---Value---")

	fmt.Println("Before value change ->", a)
	change_value(a)
	fmt.Println("After value change ->", a)

	fmt.Println("---Reference---")

	fmt.Println("Before ref change ->", a)
	change_ref(ptr)
	fmt.Println("After ref change ->", a)

	// Calling pointers 
	var ptr2 *int

	b := 1218
	ptr2 = &b

	var pptr **int
	pptr = &ptr2 //pointer pointing to a different pointer

	fmt.Println(ptr2) // memory address of b
	fmt.Println(pptr) // memory address of ptr2
	fmt.Println(*pptr) // depointed -> memory address of b
	fmt.Println(**pptr) // double depointed -> value of b
}