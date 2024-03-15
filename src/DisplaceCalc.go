// Package main demonstrates a simple program to calculate displacement
// using displacement formula in physics.
package main

import "fmt"

// GetUserInput prompts the user for input and returns the entered value as a float64.
func GetUserInput(phrase string) float64 {
	var input float64
	fmt.Print(phrase)
	fmt.Scan(&input)
	return input
}

// GenDisplaceFn generates a displacement function based on provided acceleration, initial velocity, and initial displacement.
// It returns a function that takes time as input and returns displacement at that time.
func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}
}

func main() {
	a := GetUserInput("Enter acceleration (a): ")
	v0 := GetUserInput("Enter initial velocity (v0): ")
	s0 := GetUserInput("Enter initial displacement (s0): ")

	fn := GenDisplaceFn(a, v0, s0)

	t := GetUserInput("Enter time (t): ")

	displacement := fn(t)
	fmt.Println("The displacement after", t, "seconds is", displacement)
}
