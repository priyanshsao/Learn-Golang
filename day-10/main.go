package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	/*
	     Part-1(Converting JSON string into a slice of Go structs)
	   Creating a JSON string that represents an array of objects,
	   where each object contains information about a person,
	   including their first name, last name,hair color, and whether
	   they have a dog.
	*/
	myJson := `
[
	{
		"first_name":"Tony",
		"last_name":"stark",
		"hair_color":"brown",
		"has_dog": true
	},
	{
		"first_name":"Steve",
		"last_name":"Rogers",
		"hair_color":"brown",
		"has_dog": false
	}
	
]`
	// Creating a variable that stores a slice of 'person' structs
	var unmarshalled []person
	/*
	   The "[]byte(myJson)" converts the JSON string into a slice of bytes,
	   which is the format required for unmarshalling. For example,
	   the string "hello" would be converted to the byte slice
	   {'h','e','l','l','o'}. The json.Unmarshal function then decodes
	   this byte slice into the 'unmarshalled' variable, which should be
	   a pointer to a slice or struct where the JSON data will be stored
	*/
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("error unmarshalling json", err)
	}

	fmt.Printf("unmarshalled:%v\n", unmarshalled)
	/*
	   Part-2(converting a slice of go struct into json string)
	*/
	var students_info []person

	var student_1 person = person{
		FirstName: "Stefan",
		LastName:  "Salvatore",
		HairColor: "Brown",
		HasDog:    false,
	}
	students_info = append(students_info, student_1)

	var student_2 person = person{
		FirstName: "Daemon",
		LastName:  "Salvatore",
		HairColor: "Black",
		HasDog:    false,
	}
	students_info = append(students_info, student_2)

	/*
	   Here we are using MarshalIndent to marshal the slice of structs
	   into a JSON format with indentation for better readability.
	   You might wonder why there are two variables, resultJson and err,
	   storing the output of a single function. This is because the function
	   actually returns two values: the marshalled JSON data and an error.
	   In a successful case, the error will be nil.
	*/
	resultJson, err := json.MarshalIndent(students_info, "", "      ")

	if err != nil {
		log.Println("error marshalling:", err)
	}

	fmt.Println("Marshalled: \n",string(resultJson))
}
