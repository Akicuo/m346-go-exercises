package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a, b float64) float64 {
	// c = √(a² + b²)
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

// Bonus: ShortSides struct with Hypotenuse method
type ShortSides struct {
	a float64
	b float64
}

func (s ShortSides) Hypotenuse() float64 {
	return math.Sqrt(math.Pow(s.a, 2) + math.Pow(s.b, 2))
}

func main() {
	// TODO: call the function computeHypotenuse
	
	// Test case 1: 3-4-5 triangle
	fmt.Println(computeHypotenuse(3.0, 4.0)) // 5.0
	
	// Test case 2: 5-12-13 triangle
	fmt.Println(computeHypotenuse(5.0, 12.0)) // 13.0
	
	// Test case 3: equal sides
	fmt.Println(computeHypotenuse(7.0, 7.0)) // 9.899494936611665
	
	// Bonus: Test with ShortSides struct
	fmt.Println("\nUsing ShortSides struct:")
	
	sides1 := ShortSides{a: 3.0, b: 4.0}
	fmt.Println(sides1.Hypotenuse()) // 5.0
	
	sides2 := ShortSides{a: 5.0, b: 12.0}
	fmt.Println(sides2.Hypotenuse()) // 13.0
	
	sides3 := ShortSides{a: 7.0, b: 7.0}
	fmt.Println(sides3.Hypotenuse()) // 9.899494936611665
}
