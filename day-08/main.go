package main

import "fmt"

type Animal interface {
	says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Fluffy",
		NumberOfTeeth: 40,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.says(), "and has", a.NumberOfLegs(), "legs")
}

func (d *Dog) says() string {
	return "Woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}
func (g *Gorilla) says() string {
	return "ugh"
}
func (g *Gorilla) NumberOfLegs() int {
	return 2
}
