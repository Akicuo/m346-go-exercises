package main

import (
	"fmt"
	"errors"
)

// TODO: implement the function computeGrade
func computeGrade(gotPoints, maxPoints float64) (float64, error) {
	// Check for invalid inputs
	if maxPoints <= 0 {
		return 0, errors.New("maxPoints must be greater than 0")
	}
	if gotPoints < 0 {
		return 0, errors.New("gotPoints cannot be negative")
	}
	if gotPoints > maxPoints {
		return 0, errors.New("gotPoints cannot exceed maxPoints")
	}
	
	// Calculate grade: N = 1 + 5 * (E / M)
	grade := 1.0 + 5.0 * (gotPoints / maxPoints)
	return grade, nil
}

func main() {
	// TODO: call the function computeGrade
	
	// Test case 1
	grade1, err1 := computeGrade(17.5, 28.0)
	if err1 != nil {
		fmt.Println("Error:", err1)
	} else {
		fmt.Println(grade1) // 4.125
	}
	
	// Test case 2
	grade2, err2 := computeGrade(28.0, 28.0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(grade2) // 6.0
	}
	
	// Test case 3
	grade3, err3 := computeGrade(14.0, 28.0)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println(grade3) // 3.5
	}
	
	// Test error handling: negative points
	_, err4 := computeGrade(-5.0, 28.0)
	if err4 != nil {
		fmt.Println("Error:", err4) // Error: gotPoints cannot be negative
	}
	
	// Test error handling: more points than possible
	_, err5 := computeGrade(30.0, 28.0)
	if err5 != nil {
		fmt.Println("Error:", err5) // Error: gotPoints cannot exceed maxPoints
	}
}
