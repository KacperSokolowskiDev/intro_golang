package main

import (
	"errors"
	"fmt"
)

func someFunction(param int) (int, error) {
	if param == 60 {
		return -1, errors.New("i don't like 60") // errors are type in lowercase
	}else {
		return param+10, nil
	}
}


func main() {

	var number int 

	_, err1 := fmt.Scan(&number) // must be a pointer
	// nil is null
	if err1 != nil {
		fmt.Println("We got an error!")	
		fmt.Println(err1)
	}else {
		fmt.Println(number)
	}

	ret, err2 := someFunction(60)
	if err2 != nil {
		fmt.Println("error encountered")
		fmt.Println(err2)
	}else {
		fmt.Println(ret)
	}

}