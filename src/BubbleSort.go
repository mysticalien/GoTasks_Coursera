// Package main demonstrates a simple program for sorting integers using Bubble Sort.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PrintOutput prints the sorted array of integers.
func PrintOutput(numbers []int) {
	fmt.Println("Sorted array: ")
	for i, number := range numbers {
		fmt.Print(number)
		if i < len(numbers)-1 {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
	}
}

// Swap swaps the elements at index i and i+1 in the given slice of integers.
func Swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}

// BubbleSort performs the bubble sort algorithm on the given slice of integers.
func BubbleSort(numbers []int) {
	length := len(numbers)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(numbers, j)
			}
		}
	}
}

// ScanLine reads input from the user and returns it as a string.
func ScanLine() (string, error) {
	fmt.Println("Enter up to 10 integers: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input, nil
}

// ParseInput parses the input string containing space-separated integers and returns them as a slice of integers.
// If the input is invalid, it returns an error.
func ParseInput(input string) ([]int, error) {
	var numbers []int
	var err error
	for _, strNumber := range strings.Fields(input) {
		number, error := strconv.Atoi(strNumber)
		if error != nil {
			fmt.Println("Error, invalid input")
			err = error
			break
		}
		numbers = append(numbers, number)
	}
	return numbers, err
}

func main() {
	input, err := ScanLine()
	if err != nil {
		fmt.Println("Error occurs while reading input:", err)
		return
	}

	numbers, err := ParseInput(input)

	if err != nil {
		return
	}

	if len(numbers) > 10 {
		fmt.Println("Enter up to 10 integers only")
		return
	}

	BubbleSort(numbers)

	PrintOutput(numbers)
}
