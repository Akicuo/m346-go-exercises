package main

import "fmt"

type FullName struct {
	// TODO: add fields
	firstName string
	lastName  string
}

// TODO: declare a structure for birth date
type BirthDate struct {
	day   byte
	month byte
	year  uint16
}
type Profile struct {
	// TODO: embed full name and birth date information
	NumberOfSiblings byte
	ZodiacSign       string
	FullName         FullName
	BirthDate        BirthDate
}

func main() {
	var me = Profile{
		FullName:  FullName{firstName: "Lyan", lastName: "Labaguis"},
		BirthDate: BirthDate{day: 4, month: 4, year: 2008},
		// TODO: set name and birth date information
		NumberOfSiblings: 1,       // TODO: adjust
		ZodiacSign:       "Aries", // TODO: adjust
	}
	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	// TODO: imagine, you get a little brother or sister
	me.NumberOfSiblings += 1
	fmt.Println("Siblings After:", me.NumberOfSiblings)
}
