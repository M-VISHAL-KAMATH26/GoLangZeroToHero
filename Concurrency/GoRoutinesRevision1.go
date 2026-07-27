package main

import (
	"fmt"
	"sync"
)

func squareInt(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var res int = n * n
	ch <- res

}
func main() {
	fmt.Println("welcome to the go routines revision1")
	var wg sync.WaitGroup
	ch := make(chan int)
	res := []int{11, 22, 13}
	for i := 0; i < len(res); i++ {
		wg.Add(1)
		go squareInt(res[i], ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := 0; i < len(res); i++ {
		fmt.Println(<-ch)
	}
}
