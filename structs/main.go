package main

import (
	"fmt"
	"strconv"
)

// There is no objects or classes in GO but structs
type computer struct { // defining a struct 
	manufacturer string
	ram int
	cpuBrand string
	laptop bool
}

// The best practice is to define a constructor, it's a function that returns the computer pointer
func newComputer(m string, r int, c string, l bool) *computer {
	cmp := computer{manufacturer: m, ram: r, cpuBrand: c, laptop: l}
	return &cmp // computer pointer is returned
}

func newComputerDefault() *computer {
	cmp := computer{manufacturer: "default", ram: 4, cpuBrand: "default", laptop: false}
	return &cmp
}

// Methods, "c" is the "object" for exmpl -> comp1 and it's pointing to computer struct then we specify the name of the method.
func (c *computer) increaseRam(amount int) {
	c.ram += amount
}

func (c *computer) printInfo() {
	fmt.Println("Computer info:")
	fmt.Printf("Manufacturer is %s\n", c.manufacturer)
	fmt.Printf("Ram is %d\n", c.ram)
	fmt.Printf("CPU brand is %s\n", c.cpuBrand)
	fmt.Printf("Laptop is %s\n", strconv.FormatBool(c.laptop))

}

func main() {

	comp1 := computer{manufacturer: "Lenovo", ram: 16, cpuBrand: "AMD", laptop: true} //assigning values to the struct
	//comp2 := newComputer("Dell", 32, "Intel", false)
	//comp3 := newComputerDefault()
	//fmt.Println(*comp2)
	//fmt.Println(*comp3)

	comp1.printInfo()
	comp1.increaseRam(16)
	comp1.printInfo()

}