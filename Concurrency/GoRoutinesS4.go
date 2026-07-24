package main

import (
	"fmt"
	"sync"
)

//Session 4 Practice
//
//Problem 1 — Buffered channel basics
//Create a buffered channel with capacity 3. Send 3 values into it without any goroutine involved at all (should NOT block, since nothing's consuming yet — prove to yourself this works with a plain sequential program). Then receive and print all three.
//
//Problem 2 — Producer/consumer with close and range
//Write a function generateSquares(n int, ch chan int) that sends the squares of 1 through n into ch, then closes it. In main, launch it as a goroutine and use range to receive and print every value — no manual count-tracking needed.
//
//Problem 3 — Sending on a closed channel (see the panic yourself)
//Deliberately write code that closes a channel and then tries to send to it afterward. Run it, observe the actual panic message, and paste it.
//
//Problem 4 — Two producers, one consumer, using close correctly
//This one's trickier — write two producer functions (e.g., produceEvens(ch chan int) sending 2,4,6 and produceOdds(ch chan int) sending 1,3,5) both sending into the same channel. Think carefully: if both try to close(ch) themselves, what goes wrong? (Hint: closing an already-closed channel also panics.) Figure out a way to safely close the channel only once after both producers are done — this is a great excuse to combine what you learned about WaitGroup in Session 2 with channels.

func generateSquare(n int, ch chan int) {
	for i := 1; i <= n; i++ {
		ch <- i * i
	}
	close(ch)
}

func produceEven(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			ch <- i
		}
	}
}
func produceOdd(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < n; i++ {
		if i%2 != 0 {
			ch <- i
		}
	}
}

func main() {
	fmt.Println("welcome to the buffered channel demo....!")

	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("not blocking...!")
	fmt.Println(<-ch, <-ch, <-ch)

	fmt.Println("------------------------------")
	ch2 := make(chan int, 3)
	go generateSquare(5, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}

	fmt.Println("------------------------------")
	ch3 := make(chan int, 3)
	ch3 <- 100
	ch3 <- 200
	close(ch3)
	//causes panic
	//ch3 <- 300
	fmt.Println(<-ch3, <-ch3)
	fmt.Println("------------------------------")

	var wg sync.WaitGroup
	ch4 := make(chan int)
	wg.Add(2)
	go produceEven(20, ch4, &wg)
	go produceOdd(20, ch4, &wg)

	go func() {
		wg.Wait()
		close(ch4)
	}()

	for val := range ch4 {
		fmt.Println(val)
	}

	fmt.Println("end of the world...!")

}
