package main

import "fmt"

func main() {

	var greeting string //creating a variable
	greeting = "Hello, world"

	fmt.Println(greeting)

	var num int
	num = 32

	fmt.Println(num)

	vehicle_1, vehicle_2 := program() // The two strings which are returned by program function is stored in vehicle_1 and vehicle-2
	fmt.Println("vehicle no 1 is", vehicle_1, "vehicle no 2 is", vehicle_2)
}

func program() (string, string) {
	return "CAR", "BUS" // returns two strings which are car and bus
}
