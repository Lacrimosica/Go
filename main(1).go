package main

import (
	"errors"
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Bird struct {
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

type Cow struct {
}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (cow Cow) Move() {
	fmt.Println("walk")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

type Snake struct {
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (snake Snake) Speak() {
	fmt.Println("hsss")
}

func newAnimal(mp map[string]Animal, name, animalType string) error {
	var animal Animal
	switch animalType {
	case "cow":
		animal = Cow{}
	case "snake":
		animal = Snake{}
	case "bird":
		animal = Bird{}
	default:
		return errors.New("no such animal type")
	}
	mp[name] = animal
	return nil
}

func queryAnimal(mp map[string]Animal, name, action string) error {
	if val, ok := mp[name]; ok {
		switch action {
		case "move":
			val.Move()
		case "eat":
			val.Eat()
		case "speak":
			val.Speak()
		default:
			return errors.New("no such action")
		}
		return nil
	} else {
		return errors.New("no such animal name")
	}

}
func main() {
	mp := make(map[string]Animal)

	for {
		var (
			query, p1, p2 string
		)
		fmt.Print("> ")
		fmt.Scan(&query)
		fmt.Scan(&p1)
		fmt.Scan(&p2)

		if query == "newanimal" {
			err := newAnimal(mp, p1, p2)
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Println("created!!")
			}
		} else if query == "query" {
			err := queryAnimal(mp, p1, p2)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			fmt.Print("No such query")
		}
	}
}
