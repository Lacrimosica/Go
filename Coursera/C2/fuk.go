package main

import "fmt"
import "bufio"
import "os"
import "strings"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	// cow specific stuff goes here
}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	// bird specific stuff goes here
}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	// snake specific stuff goes here
}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			os.Exit(1);
		}
		input := scanner.Text()
		sv := strings.Fields(input)
		if len(sv) != 3 {
			os.Exit(0);
		}

		switch sv[0] {
		case "newanimal":
			switch sv[2] {
			case "cow":
				animals[sv[1]] = Cow{}
			case "bird":
				animals[sv[1]] = Bird{}
			case "snake":
				animals[sv[1]] = Snake{}
			default:
				fmt.Println("Unknown animal", sv[2])
				continue
			}
			fmt.Println("Created it!")

		case "query":
			ai, known := animals[sv[1]]
			if !known {
				fmt.Println("Don't recognize", sv[1])
				continue
			}
			switch sv[2] {
			case "eat":
				ai.Eat()
			case "move":
				ai.Move()
			case "speak":
				ai.Speak()
			default:
				fmt.Println("Don't grok", sv[2])
			}

		default:
			fmt.Println("Don't grok", sv[0])
		}
	}
}
