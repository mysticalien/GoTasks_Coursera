// Package main demonstrates a simple command-line program to manage animals.
// It allows users to create animals of different types (cow, bird, snake)
// and query information about them (eat, move, speak).
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal represents an interface for different types of animals.
type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct {
	name string
}

type Bird struct {
	name string
}

type Snake struct {
	name string
}

func (c Cow) Eat() string {
	return "grass"
}

func (c Cow) Move() string {
	return "walk"
}

func (c Cow) Speak() string {
	return "moo"
}

func (b Bird) Eat() string {
	return "worms"
}

func (b Bird) Move() string {
	return "fly"
}

func (b Bird) Speak() string {
	return "peep"
}

func (s Snake) Eat() string {
	return "mice"
}

func (s Snake) Move() string {
	return "slither"
}

func (s Snake) Speak() string {
	return "hsss"
}

// CreateAnimal creates a new animal based on the animal type and name.
func CreateAnimal(animalType, name string) Animal {
	switch animalType {
	case "cow":
		return Cow{name}
	case "bird":
		return Bird{name}
	case "snake":
		return Snake{name}
	default:
		return nil
	}
}

func Usage() {
	fmt.Println("\033[0;30;102mUsage:\033[0m")
	fmt.Println("1. For creating a new animal:")
	fmt.Println("   \u001B[0;30;102mnewanimal\u001B[0m <name> <type>")
	fmt.Println("   <name>: Name of the new animal")
	fmt.Println("   <type>: Type of the new animal (cow, bird, snake)")
	fmt.Println("2. For querying information about an animal:")
	fmt.Println("   \u001B[0;30;102mquery\u001B[0m <name> <info>")
	fmt.Println("   <name>: Name of the animal to query")
	fmt.Println("   <info>: Information to query about the animal (eat, move, speak)")
}

func main() {
	var input string
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">")
		scanner.Scan()
		input = scanner.Text()
		words := strings.Fields(input)
		if len(words) == 3 {
			var query = words[0]
			var name = words[1]
			var info = words[2]
			switch query {
			case "newanimal":
				if info != "cow" && info != "bird" && info != "snake" {
					fmt.Println("Animal not found!")
					Usage()
					continue
				}
				animals[name] = CreateAnimal(info, name)
				fmt.Println("Animal created!")
			case "query":
				animal, ok := animals[name]
				if !ok {
					fmt.Println("Animal not found!")
					Usage()
					continue
				}
				switch info {
				case "eat":
					fmt.Println(animal.Eat())
				case "move":
					fmt.Println(animal.Move())
				case "speak":
					fmt.Println(animal.Speak())
				default:
					fmt.Println("Invalid query!")
					Usage()
				}
			default:
				fmt.Println("Invalid command!")
			}
		} else {
			fmt.Println("Invalid number of arguments!")
			Usage()
		}
	}
}
