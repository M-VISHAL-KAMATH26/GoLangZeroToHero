package main

//Session 1 Practice
//
//Problem 1: Write a program with two functions, brewCoffee() and toastBread(), each printing a few messages with small labels (e.g. "Coffee: step 1", "Coffee: step 2", etc.). Run both as goroutines from main, using a time.Sleep at the end of main (hacky, but fine for now — proper fix comes in Session 2). Observe and paste the actual output you get.
//
//Problem 2: Modify Problem 1 so main does not sleep at all. Predict what you'll see before running, then confirm.
//
//Problem 3: Spawn 5 goroutines in a loop, each just printing its own number (1 through 5) — using the loop-variable-passing style (go func(n int) { ... }(i)) even though we haven't formally covered why that pattern's needed yet (that's Session 2) — just get comfortable with the syntax of passing a value into an anonymous goroutine function.

import (
	"fmt"
	"sync"
)

func brewCoffee(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("Coffee :step ", i)
	}
}
func toastBread(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println("Bread :step ", i)
	}
}
func main() {
	fmt.Println("welcome to the session 1 problem solving....!")
	var wg sync.WaitGroup

	wg.Add(1)
	go brewCoffee(&wg)

	wg.Add(1)
	go toastBread(&wg)

	wg.Wait()
	//time.Sleep(100 * time.Millisecond)
	fmt.Println("end of the main go routine")
}
