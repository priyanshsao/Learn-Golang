package main

import (
	"log"

	"github.com/Learn-Golang/modules/day-09/helpers"
)

const numpool = 100

/*
	                               ---2---
    Creating a function which takes intChan as an argument(a "channel"
    which is of int datatype). After this go to helpers.go and read
	section 3.

    Now we will store the value of the generated no. inside a variable
    called "randomvariable". Then  we will  pass that random number
	that we just generated into  intChan. Then go to section 4.
*/
func CalculateValue(intChan chan int) {
	randomnumber := helpers.RandomNumber(numpool)
	intChan <- randomnumber
}

// ---1---
func main() {
	// Create a channel
	intChan := make(chan int)
	defer close(intChan) //Ensures the channel is closed after all operations are done.

	/*                 ---4---
	Start CalculateValue in a new goroutine  */
	go CalculateValue(intChan)
	// Receive the value from the channel
	num := <-intChan
	// Log the received value
	log.Println(num)
}
