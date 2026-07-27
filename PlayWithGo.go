package main

import (
	"fmt"
	"sync"
)

func setUsernameWithCharacter(name *string, character string) string {
	*name = *name + " " + character
	return *name
}

func AddGoldAndElixer(p *Player, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Added the resources")
	p.stats.gold += 25000
	p.stats.elixer += 25000
	p.stats.darkElixer += 2500
}

type Player struct {
	name   string
	age    int
	gender string
	stats  Stats
}
type Stats struct {
	gold       int64
	elixer     int64
	darkElixer int64
}

type Engine interface {
	Start() string
	End() string
}

type game struct {
	gameName string
}

func (g game) Start() string {
	return "game engine has been started"
}

func (g game) End() string {
	return "game engine has been ended"
}

//creating the method to displaay the gold elixer and dar elixer status
func (p Player) DisplayStats() {
	fmt.Println(p.stats.gold, " is the gold earned ", p.stats.elixer, "is the elixer earned  ", p.stats.darkElixer, " is your current dark elixer status")
}
func main() {
	fmt.Println("Welcome to the GoLang Game...!")
	var gaming Engine = game{gameName: "Clash of Titans"}
	fmt.Println(gaming.Start())
	fmt.Println("-------------------------------")
	var userName string
	var palyerage int
	fmt.Println("Enter your name: ")
	fmt.Scanln(&userName)
	fmt.Println("enter the age")
	fmt.Scanln(&palyerage)
	fmt.Println("welcone onboard: ", userName)
	fmt.Println("-------------------------------")
	fmt.Println("click 1: for male click 2: for female")
	var gender int
	fmt.Scanln(&gender)
	var character string
	if gender == 1 {
		character = "The King"
	} else {
		character = "The Queen"
	}

	gameChar := setUsernameWithCharacter(&userName, character)
	fmt.Println("you are now  ", gameChar)
	fmt.Println("-------------------------------")

	fmt.Println("u have entered the game so u have by default  some gold and elixer ")
	fmt.Println("here is your game profile")
	player1 := Player{
		name:   userName,
		age:    palyerage,
		gender: character,
		stats: Stats{
			gold:   100000,
			elixer: 100000,
		},
	}
	fmt.Println(player1)
	player1.DisplayStats()
	fmt.Println("---------------------------")

	fmt.Println("Attacking in the battle...!")
	var winOrLoose bool
	fmt.Scan(&winOrLoose)

	var wg sync.WaitGroup

	if winOrLoose {
		wg.Add(1)
		go AddGoldAndElixer(&player1, &wg)
	}
	wg.Wait()

	player1.DisplayStats()
	fmt.Println("---------------------------")

	fmt.Println(gaming.End())

}
