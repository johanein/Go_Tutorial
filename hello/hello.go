package main

import (
	"fmt"
	"log"

	"rsc.io/quote/v4"

	"github.com/johanein/goTutorialGreetings"
)

func main() {
	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request goTutorialGreetings messages for the names.
	messages0, err0 := goTutorialGreetings.Hellos(names)
	if err0 != nil {
		log.Fatal(err0)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages0)

	// Get a goTutorialGreetings message and print it.
	message, err := goTutorialGreetings.Hello("Albert")
	// If no error was returned, print the returned message
	// to the console.
	if err != nil {
		log.Fatal(err)
	}
	// Get a goTutorialGreetings message and print it.
	fmt.Println(message)
	message1, err1 := goTutorialGreetings.Hello("")
	// If no error was returned, print the returned message
	// to the console.
	if err1 != nil {
		log.Fatal(err1)
	}
	// Get a goTutorialGreetings message and print it.
	fmt.Println(message1)
}
