package main

import (
	"fmt"
	"strings"
)


type Animal interface{
	Eat()
	Move()
	Speak()
}
type Cow struct{
	food string
	locomotion string
	noise string
}
func (x Cow) Eat(){fmt.Println(x.food)}
func (x Cow) Move(){fmt.Println(x.locomotion)}
func (x Cow) Speak(){fmt.Println(x.noise)}
type Bird struct{
	food string
	locomotion string
	noise string
}
func (x Bird) Eat(){fmt.Println(x.food)}
func (x Bird) Move(){fmt.Println(x.locomotion)}
func (x Bird) Speak(){fmt.Println(x.noise)}
type Snake struct{
	food string
	locomotion string
	noise string
}

func (x Snake) Eat(){fmt.Println(x.food)}
func (x Snake) Move(){fmt.Println(x.locomotion)}
func (x Snake) Speak(){fmt.Println(x.noise)}
func main(){
	var commandType, animalOrName, infoOrType string
	m := make(map[string]Animal)
	for {
		fmt.Printf(">")
		fmt.Scanf("%s %s %s\n",&commandType,&animalOrName,&infoOrType)
		switch strings.ToLower(commandType){
		case "newanimal":
			switch strings.ToLower(infoOrType){
			case "cow":
				m[animalOrName] = Cow{"grass","walk","moo"}
			case "bird":
				m[animalOrName] = Bird{"worms","fly","peep"}
			case "snake":
				m[animalOrName] = Snake{"mice","slither","hsss"}
			}
			fmt.Println("Created it")
		case "query":
			requestedAnimal := m[strings.ToLower(animalOrName)]
			switch strings.ToLower(infoOrType) {
			case "eat":
				requestedAnimal.Eat()
			case "move":
				requestedAnimal.Move()
			case "speak":
				requestedAnimal.Speak()
			}
		}
	}
}