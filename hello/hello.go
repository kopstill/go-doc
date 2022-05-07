package main

import (
	"fmt"
	"kopever.com/greetings"
	"kopever.com/quotes"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("kopever")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	fmt.Println("------------quotes below------------")

	quotes.Print()
}
