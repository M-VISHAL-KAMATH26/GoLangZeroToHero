package main

import "fmt"

func main() {
	fmt.Println("welcome to pointer to pointer implementation")
	var value int = 50
	p := &value
	pp := &p

	fmt.Println(pp)   //address of the variable pp
	fmt.Println(*pp)  //dereference once: gives p's value = the address of value
	fmt.Println(**pp) //dereference twice: gives value's actual value = 50

	**pp = 100
	fmt.Println(value)
}
