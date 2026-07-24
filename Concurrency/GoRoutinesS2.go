package main

import (
	"fmt"
	"sync"
)

//Session 2 Practice
//
//Problem 1 — Basic WaitGroup
//Write a program that launches 3 goroutines, each printing "Worker <n> done" (where n is 1, 2, 3), using a properly matched Add/Done/Wait — no time.Sleep anywhere. Confirm all three print before "main finished" shows up.
//
//Problem 2 — Fix the loop-capture bug yourself
//Write a loop launching 5 goroutines that each print their index, deliberately using the buggy shared-variable version first (no parameter passed in) — run it and see what you actually get on your Go version. Then fix it using the parameter-passing pattern and confirm the difference (or lack of one, if you're on Go 1.22+).
//
//Problem 3 — Goroutines returning data via a shared slice + WaitGroup (no channels yet — that's Session 3)
//Write a function squareNumbers(nums []int, results []int, index int, wg *sync.WaitGroup) that squares nums[index] and stores it into results[index]. Launch one goroutine per number in nums (each writing to its own unique index in results — no two goroutines touch the same index, so this is safe without a mutex), wait for all to finish, then print results.

func Worker(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker ", n, " done")
}

func squareNumbers(nums []int, res []int, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	res[index] = nums[index] * nums[index]
}
func main() {
	fmt.Println("starting main")
	var wg sync.WaitGroup
	wg.Add(3)
	go Worker(1, &wg)
	go Worker(2, &wg)
	go Worker(3, &wg)
	wg.Wait()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n)
		}(i)

	}
	wg.Wait()

	fmt.Println("---------------------------------")
	nums := []int{2, 4, 6, 8, 10}
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		wg.Add(1)
		go squareNumbers(nums, result, i, &wg)
	}
	wg.Wait()
	fmt.Println(nums)
	fmt.Println(result)
	fmt.Println("---------------------------------")

	fmt.Println("main finished")

}
