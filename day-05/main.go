package main

import "log"

func main() {

	Firstmap := make(map[int]string)

	Firstmap[1]= "Sunday"
	Firstmap[2]= "Monday"
	Firstmap[3]= "Tuesday"
	Firstmap[4]= "Wednesday"
	Firstmap[5]= "Thursday"
	Firstmap[6]= "Friday"
	Firstmap[7]= "Saturday"

	var Week []string 

	Week = append(Week, Firstmap[1])
	Week = append(Week, Firstmap[2])

    log.Println("The first two days are: ",Week)
}