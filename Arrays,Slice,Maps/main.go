package main

import "fmt"

func main() {

	// Arrays
	fmt.Println("------Arrays------")
	var arr1[5] int // defining a size and type of elements
	arr1[1] = 20
	arr1[4] = 10
	//arr1[5] = 2 -> the last element is out of bounds (5)

	fmt.Println(arr1)
	fmt.Println(arr1[4])

	arr2 := [5]int{1,2,3,4,5} // declaring everything at once
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	var stringArr[5] string
	fmt.Println(stringArr) // empty array

	var boolArr[5] bool
	fmt.Println(boolArr) // all values are false by default

	var arrTwoD[4][5] int // 2D arrays
	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			arrTwoD[i][j] = i*j
		}
	}
	fmt.Println(arrTwoD)

	// Slices -> Lists in Python
	fmt.Println("------Slices------")

	s := make([]int, 5) // make function is needed to use slices
	s = append(s, 60, 22, 31) // increase the size of the slice at the end
	cp1 := make([]int, len(s)) // making new slice with length of s
	copy(cp1, s) // copying elements from slice s to cp1

	cp1[2] = 18 // changes cp1 only

	fmt.Println("Original", s)
	fmt.Println("Copy", cp1)

	fmt.Println(cp1[1:4]) // print 1,2,3, slicing works only with positive numbers 


	// Maps -> Dictionnaries in Python
	fmt.Println("------MAPS------")

	m1 := make(map[string]int) // strings referencing integers

	m1["Key1"] = 20
	m1["Key2"] = 30
	m1["Key3"] = 18
	m1["None"] = 0

	value, present := m1["None"]
	fmt.Println(value) // verifying the volue of specified key
	fmt.Println(present) // verifying if the key is present in the map

	delete(m1, "Key1") // deleting a key from the map

	fmt.Println(m1)
}