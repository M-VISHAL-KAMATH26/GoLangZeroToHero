package main

import (
	"fmt"
	"time"
)

//func printNumbers() {
//	for i := 1; i <= 5; i++ {
//		fmt.Println(i)
//	}
//}
//
//func main() {
//	go printNumbers()
//	fmt.Println("main function line")
//	time.Sleep(100 * time.Millisecond) // artificially "wait" a bit
//	fmt.Println("main finishing now")
//}

func printSomething(label string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(label, i)
	}
}

func main() {
	go printSomething("A")
	go printSomething("B")
	time.Sleep(100 * time.Millisecond)
}
