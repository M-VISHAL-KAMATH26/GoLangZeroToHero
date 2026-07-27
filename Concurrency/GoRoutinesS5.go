package main

//Session 5 Practice
//
//Problem 1 — Basic select between two channels
//Create two channels. Launch two goroutines, each sleeping a different short duration (e.g. 500ms and 1s) before sending a message into their respective channel. Use select to print whichever arrives first.
//
//Problem 2 — select in a loop (basic fan-in)
//Using the same two channels/goroutines as Problem 1, but have each goroutine send 3 messages (in a loop, with a small sleep between each) instead of just one. In main, use select inside a for loop to receive and print 6 total messages (3 from each), regardless of which order they arrive in.
//
//Problem 3 — Non-blocking check with default
//Write a small loop (say, 5 iterations) that each time tries to receive from a channel using select + default — printing "got: X" if something's there, or "nothing yet" if not. Have a goroutine send just one value into that channel after a short delay, so you can observe most iterations printing "nothing yet" until the one iteration where the value has actually arrived.
//
//Problem 4 — Timeout pattern
//Write a function slowOperation(ch chan string) that sleeps for 3 seconds, then sends "done". In main, use select with time.After(1 * time.Second) to timeout before the operation completes, printing an appropriate timeout message.

func main() {

}
