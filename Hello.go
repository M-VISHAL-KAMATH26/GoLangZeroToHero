package main

import "fmt"

func main() {
	fmt.Println("welcome to the ticket booking App")

	var trainName string = "rajdhani express"

	const trainTicketsNo int = 50
	var remainingTickets int = 50
	//fmt.Println("the train from delhi is ", trainName, " at9:00...! ", remainingTickets, " tickets are remaining")
	fmt.Printf("the train rom delhi is %s and the tickets remaining is %d ...! \n", trainName, remainingTickets)
	fmt.Println("Number of tickets", trainTicketsNo)
	//fmt.Printf("%T type of tickets\n", trainTicketsNo)

	var passengerName string
	fmt.Println("please enter the passenger name..!")
	fmt.Scan(&passengerName)
	fmt.Println(passengerName)
}
