package main

import "fmt"

func addOne(n *int) {
	*n = *n + 1
}

func main() {

	fmt.Println("Hello welcome to the pointers in Go..!")

	x := 10
	fmt.Println(x)  //10
	fmt.Println(&x) //some address

	var pointer *int
	pointer = &x
	fmt.Println(*pointer)

	//alternate way using the new keyword
	pointerNew := new(int)
	fmt.Println(*pointerNew)
	*pointerNew = 20
	fmt.Println(*pointerNew)
	fmt.Println("----------------------------------------------")
	var currentValue int
	fmt.Println("please enter the value")
	fmt.Scanln(&currentValue)

	pointerX := &currentValue
	fmt.Println(*pointerX, "is the current value.")
	addOne(pointerX)
	addOne(pointerX)
	addOne(pointerX)

	fmt.Println("expected value is 3 times increase the current value")
	fmt.Println(*pointerX)

}
