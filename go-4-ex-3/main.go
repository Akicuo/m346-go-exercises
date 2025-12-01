package main

import (
	"fmt"
	"math"
)

// TODO: implement the function computeQuadraticFormula

// Bonus: Separate function for computing discriminant
func computeDiscriminant(a, b, c float64) float64 {
	// D = b² - 4ac
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	// Calculate discriminant using the separate function
	discriminant := computeDiscriminant(a, b, c)
	
	if discriminant > 0 {
		// Two solutions
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return []float64{x1, x2}
	} else if discriminant == 0 {
		// One solution
		x := -b / (2 * a)
		return []float64{x}
	} else {
		// No solution (D < 0)
		return []float64{}
	}
}

func main() {
	// TODO: call the function computeQuadraticFormula
	
	// Test discriminant function separately
	fmt.Println("Testing computeDiscriminant:")
	fmt.Println(computeDiscriminant(3, 4, 1)) // 4.0 (positive)
	fmt.Println(computeDiscriminant(2, 4, 2)) // 0.0 (zero)
	fmt.Println(computeDiscriminant(3, 4, 2)) // -8.0 (negative)
	
	fmt.Println("\nTesting computeQuadraticFormula:")
	
	// Test case 1: D > 0, two solutions (a=3, b=4, c=1)
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // [-0.3333333333333333 -1]
	
	// Test case 2: D = 0, one solution (a=2, b=4, c=2)
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // [-1]
	
	// Test case 3: D < 0, no solution (a=3, b=4, c=2)
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // []
}
