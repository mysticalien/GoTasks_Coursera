// Package main demonstrates a simple program to create and interact with animals.
package main

import (
	"fmt"
	"strings"
)

// Animal represents an animal with its characteristics.
type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

// scanInput prompts the user to input the name of the animal and the information requested about it.
func scanInput() (string, string) {
	var nameOfAnimal, info string
	fmt.Print(">")
	fmt.Scan(&nameOfAnimal, &info)
	return nameOfAnimal, info
}

func Usage() {
	fmt.Println("\033[0;30;102mUsage:\033[0m")
	fmt.Println("1. Enter the name of the animal followed by the information requested:")
	fmt.Println("   <animal_name> <info>")
	fmt.Println("   <animal_name>: Name of the animal (cow, bird, snake)")
	fmt.Println("   <info>: Information to query about the animal (eat, move, speak)")
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		nameOfAnimal, info := scanInput()

		nameOfAnimal = strings.ToLower(nameOfAnimal)
		info = strings.ToLower(info)

		animal, err := animals[nameOfAnimal]
		if !err {
			fmt.Println("No such animal")
			Usage()
			continue
		}

		switch info {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("The name of the information requested about the animal is not right")
			Usage()
		}
	}
}
