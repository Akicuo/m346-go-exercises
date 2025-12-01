package main

import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(c float64) float64 {
	// F = C × 9/5 + 32
	return c * 9.0 / 5.0 + 32.0
}

// TODO: implement the function convertFahrenheitToCelsius
func convertFahrenheitToCelsius(f float64) float64 {
	// C = (F - 32) × 5/9
	return (f - 32.0) * 5.0 / 9.0
}

// Bonus: Custom types with methods
type Celsius float64
type Fahrenheit float64

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(float64(c) * 9.0 / 5.0 + 32.0)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((float64(f) - 32.0) * 5.0 / 9.0)
}

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	
	fmt.Println("Testing convertCelsiusToFahrenheit:")
	temp1 := convertCelsiusToFahrenheit(0.0)
	fmt.Println(temp1) // 32.0
	
	temp2 := convertCelsiusToFahrenheit(23.0)
	fmt.Println(temp2) // 73.4
	
	temp3 := convertCelsiusToFahrenheit(100.0)
	fmt.Println(temp3) // 212.0
	
	// TODO: call the function convertFahrenheitToCelsius
	
	fmt.Println("\nTesting convertFahrenheitToCelsius (round-trip):")
	fmt.Println(convertFahrenheitToCelsius(temp1)) // 0.0
	fmt.Println(convertFahrenheitToCelsius(temp2)) // 23.0
	fmt.Println(convertFahrenheitToCelsius(temp3)) // 100.0
	
	// Bonus: Test custom types with methods
	fmt.Println("\nTesting custom types with methods:")
	
	var cozy Celsius = 23.0
	fmt.Println(cozy.ConvertToFahrenheit()) // 73.4
	
	var cold Fahrenheit = 15.3
	fmt.Println(cold.ConvertToCelsius()) // -9.277777777777779
	
	// Demonstrating method chaining
	fmt.Println("\nDemonstrating method chaining:")
	var c Celsius = 23.0
	fmt.Println(c.ConvertToFahrenheit().ConvertToCelsius()) // 23.0
	
	// Answer to bonus question: The method chaining notation (second version)
	// is more readable and follows object-oriented principles
}
