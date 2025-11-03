package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var eyes = rand.Intn(5) + 1
	var when = time.Now()

	eyesFile, err := os.Create("eyes.txt")
	if err != nil {
		panic(err)
	}
	defer eyesFile.Close()

	diceFile, err := os.Create("dice.log")
	if err != nil {
		panic(err)
	}
	defer diceFile.Close()

	fmt.Fprintln(eyesFile, "the dice shows", eyes, "eyes")

	fmt.Fprintln(diceFile, "the dice was rolled at", when)
}
