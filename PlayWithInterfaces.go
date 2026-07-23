package main

import "fmt"

//basic interface implementation
type Speaker interface {
	Speak() string
}
type Human struct {
	name string
}
type Robot struct {
	name string
}

func (h Human) Speak() string {
	return h.name + " says hello"
}

func (r Robot) Speak() string {
	return r.name + " says beep beep."
}

//interfaces having multiple methods and to satisfy it the type should have those all methods .lets check this
type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct {
	name string
}

func (c Car) Start() string {
	return c.name + " is starting"
}
func (c Car) Stop() string {
	return c.name + " is stopping"
}

type Bike struct {
	name string
}

func (b Bike) Start() string {
	return b.name + " is starting"
}
func (b Bike) Stop() string {
	return b.name + " is stopping"
}

//embedded interfaces
type Flyer interface {
	Fly() string
}
type Swimmer interface {
	Swim() string
}
type Amphibious interface {
	Flyer
	Swimmer
}
type Duck struct {
	name string
}

func (d Duck) Fly() string {
	return d.name + " is flying"
}
func (d Duck) Swim() string {
	return d.name + " is swimming"
}

func main() {
	fmt.Println("welcome to interfaces demo...! ")
	var h Speaker = Human{
		name: "vishal",
	}
	var r Speaker = Robot{
		name: "Alexa 8.009",
	}

	fmt.Println(h.Speak())
	fmt.Println(r.Speak())

	fmt.Println("----------------------------------")
	var vehicle1 Vehicle = Car{name: "Mercedes benz"}
	fmt.Println(vehicle1)
	fmt.Println(vehicle1.Start())
	fmt.Println(vehicle1.Stop())
	fmt.Println()
	var vehicle2 Vehicle = Bike{name: "Yamaha"}
	fmt.Println(vehicle2)
	fmt.Println(vehicle2.Start())
	fmt.Println(vehicle2.Stop())

	fmt.Println("----------------------------------")
	var ducky Amphibious = Duck{name: "golden duck"}
	fmt.Println(ducky)
	fmt.Println(ducky.Fly())
	fmt.Println(ducky.Swim())

}
