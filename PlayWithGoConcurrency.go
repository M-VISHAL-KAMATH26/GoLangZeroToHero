package main

import (
	"fmt"
	"sync"
)

func printNum(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go printNum(&wg)
	wg.Wait()
	fmt.Println("hello world......!")
}
