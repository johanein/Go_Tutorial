package main

import (
	"fmt"

	"rsc.io/quote/v4"

	"github.com/johanein/goTutorialGreetings"
)

func main() {
	fmt.Println("Hello december")
	fmt.Println(quote.Go())
	// Get a greeting message and print it.
	message := goTutorialGreetings.Hello("Gladys")
	fmt.Println(message)
}
