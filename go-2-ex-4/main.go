package main

import "fmt"

type Student struct {
	FirstName string
	LastName  string
}

type Class struct {
	Students []Student
}

type Module struct {
	Number  int
	Classes []Class
}

func main() {
	// 1. Klasse
	class1 := Class{
		Students: []Student{
			{FirstName: "Anna", LastName: "Müller"},
			{FirstName: "Ben", LastName: "Schmidt"},
			{FirstName: "Clara", LastName: "Weber"},
		},
	}

	// 2. Klasse
	class2 := Class{
		Students: []Student{
			{FirstName: "David", LastName: "Fischer"},
			{FirstName: "Emma", LastName: "Wagner"},
			{FirstName: "Felix", LastName: "Becker"},
		},
	}

	// 3. Klasse
	modules := map[int]Module{
		346: {
			Number:  346,
			Classes: []Class{class1, class2},
		},
		104: {
			Number:  104,
			Classes: []Class{class1},
		},
		117: {
			Number:  117,
			Classes: []Class{class2},
		},
	}

	// Output all data
	for moduleNumber, module := range modules {
		fmt.Printf("Modul %d:\n", moduleNumber)
		for i, class := range module.Classes {
			fmt.Printf("  Klasse %d:\n", i+1)
			for _, student := range class.Students {
				fmt.Printf("    - %s %s\n", student.FirstName, student.LastName)
			}
		}
		fmt.Println()
	}

	// TODO: declare a type for Student (with first and last name)
	// TODO: declare a type for Class (	consisting of multiple students)
	// TODO: declare a map of modules being attended by multiple classes
	// TODO: output everything using fmt.Println()
}
