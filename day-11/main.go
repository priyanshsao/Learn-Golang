package main

import (
	"log"
	"errors"
)

func main() {
	result, err := divide(2.0, 0.0)

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("the result is: ", result)

}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("error in dividing")
	}

	result = x/y
	return result, nil
}