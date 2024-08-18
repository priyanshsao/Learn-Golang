package main

import (
	"log"
)

type Mystruct struct {
	Name    string
	Age     int
	Section string
}

func (m *Mystruct) StudentDetails() (string, int, string) {
	return m.Name, m.Age, m.Section
}

func main() {

	var Student Mystruct = Mystruct{
		Name:    "priyansh",
		Age:     19,
		Section: "A1",
	}

	studentName, studentAge, studentSection := Student.StudentDetails()

	log.Println("Name: ", studentName)
	log.Println("Age: ", studentAge)
	log.Println("Section: ", studentSection)
}
