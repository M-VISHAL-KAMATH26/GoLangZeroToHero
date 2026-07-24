package main

//Problem 1 — Basic handoff
//Write a function calculateSquare(n int, ch chan int) that computes n * n and sends the result into ch. In main, create the channel, launch the goroutine, receive the result, and print it.
//
//Problem 2 — Multiple sends, received in a loop
//Write a function generateNumbers(ch chan int) that sends the numbers 1 through 5 into ch, one at a time (a loop with ch <- i each iteration — no close() yet, that's Session 4). In main, receive exactly 5 values in a loop and print each.
//
//Problem 3 — Two goroutines, one channel each, main collects both
//Write two functions, sendGreeting(ch chan string) (sends "hello") and sendFarewell(ch chan string) (sends "goodbye") — but have them share the same channel. Launch both as goroutines, then receive two values in main and print both. Think about what order you might get them in, and why.
//
//Problem 4 — Direction-restricted channel practice
//Rewrite Problem 1's calculateSquare function so its channel parameter is explicitly send-only (chan<- int), and write a small separate function that takes the same channel as receive-only (<-chan int) to print the result. Wire them together in main.

import "fmt"

func calculateSquare(n int, ch chan<- int) {
	fmt.Println("Calculating square")
	ch <- n * n

}
func printResultReceiveOnly(ch <-chan int) {
	result := <-ch
	fmt.Println("received result:", result)
}

func generateNumbers(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i
	}
}

func sendGreetings(ch chan string) {
	ch <- "Hello"
}

func sendFarewell(ch chan string) {
	ch <- "Goodbye"
}

func main() {
	fmt.Println("welcome to the go channels demo....!")
	ch := make(chan int)
	go calculateSquare(90, ch)
	res := <-ch
	fmt.Println(res)

	fmt.Println("----------------------------")
	ch2 := make(chan int)
	go generateNumbers(ch2)
	for i := 1; i <= 5; i++ {
		val := <-ch2
		fmt.Println(val)
	}

	fmt.Println("----------------------------")

	ch3 := make(chan string)
	go sendGreetings(ch3)

	go sendFarewell(ch3)

	res1 := <-ch3
	res2 := <-ch3
	fmt.Println(res1)
	fmt.Println(res2)

	fmt.Println("-----------------------------")
	go calculateSquare(8, ch)
	printResultReceiveOnly(ch)
	fmt.Println("done with main.....!")
}
