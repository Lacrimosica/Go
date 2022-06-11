package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.

	if name == "" {
		return "", errors.New("empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	//message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

//init sets initial values for variable used in the function.
//go executes init functions automatically at program startup
//after global variables have been initialized
func init() {
	rand.Seed(time.Now().UnixNano())
}

//random Format returns one of a set of greeting message. The returned
//message ius selected at random.

func randomFormat() string {
	//A slice of message formats

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see  you, %v",
		"Hail, %v! Well met!",
	}

	//return a randomly selected message format by specifying
	//a random index fo the slice of formats
	return formats[rand.Intn(len(formats))]
}
