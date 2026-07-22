package main

import (
	"fmt"
)

//the function definition of swapping 2 numbers
func swappingOfNumbersUsingPointers(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp

}
func main() {

	fmt.Println("Welcome to the swapping logic implemented using the pointers")

	//let us take two variables here from the user
	var a, b int
	fmt.Println("enter the number a")
	fmt.Scanln(&a)
	fmt.Println("enter the number b")
	fmt.Scanln(&b)
	fmt.Println("the value of a is ", a)
	fmt.Println("the value of b is ", b)

	//now lets make a pointer variable and point to these two variables
	fmt.Println("-------------------------------")
	pointerofA := &a
	var pointerofB *int
	pointerofB = &b

	//now call the swapping function
	fmt.Println("-------------------------------")

	swappingOfNumbersUsingPointers(pointerofA, pointerofB)
	fmt.Println("after swapping the values of a and b")

	fmt.Println("the value of a is ", a)
	fmt.Println("the value of b is ", b)

}
