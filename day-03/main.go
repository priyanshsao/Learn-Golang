package main

import (
	"time"
    "log"
)

type Data struct {

	Name        string
	Sirname     string
	PhoneNumber int
	Date        time.Time

}

func main() {
	
	user1 := Data{

		Name: "Priyansh",
		Sirname: "Sao",
		PhoneNumber: 1000000001,
		Date: time.Now(),
	}

	log.Println("Your name is: ",user1.Name, user1.Sirname)
	log.Println("Your Phone no. is: ",user1.PhoneNumber)
    log.Println("Date of registration is: ",user1.Date)
}