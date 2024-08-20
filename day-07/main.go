package main

import "log"

func main() {
    // Creating slice 
	Student := []string{"Alex", "Mary", "Josh", "Max", "Louis"}
    // Looping using for 
	// The index number is stored in i and the names are stored in Names
	for i, Names := range Student {  

		log.Println(i, ":", Names)
	}
    // Creating a Map
	/* A map in Go is a built-in data structure that stores key-value
	   pairs. The key acts like an index, and each key is associated 
	   with a specific value.*/
	StudentByMap := make(map[int]string)//"[int]" shows that keys are of type int.
                                        //"string" shows that  values are of type string.
	StudentByMap[1] = "harry"
	StudentByMap[2] = "Joe"
	StudentByMap[3] = "Gary"
    // Here the keys are stored in i and values in NamesByMap
	for i, NamesByMap := range StudentByMap {
		log.Println(i, ":", NamesByMap)
	}
    /* The underscore(_) is used as a blank identifier and 
	   and by using it we can ignore a value that we don't need. */
	for _, NamesByMap := range StudentByMap { // here keys are ignored
		log.Println(NamesByMap)
	}
    // Creating a variable of type string 
	var str string = "HELLO"
    /*  Here the variable "RuneValue" stores The Unicode code 
	    point (rune) of the character at that index.*/
	for i, RuneValue := range str {
		log.Println(i, ":", RuneValue)
	}
    // Creating a struct
	type User struct {
		Name    string
		Age     int
		Section string
	}
    // Now creating a slice of User type
	var UserData []User
    // Adding data into the Slice
	UserData = append(UserData, User{
		Name:    "Alex",
		Age:     34,
		Section: "A",
	})
	UserData = append(UserData, User{
		Name:    "Joe",
		Age:     28,
		Section: "B",
	})

	for _, UserViewData := range UserData {
		log.Println(UserViewData.Name, "|", UserViewData.Age, "|", UserViewData.Section)
	}
}
