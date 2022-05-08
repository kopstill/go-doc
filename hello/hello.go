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

	fmt.Println("------------multiple greetings below------------")
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err1 := greetings.Hellos(names)
	if err1 != nil {
		log.Fatal(err1)
	}
	for _, v := range messages {
		fmt.Println(v)
	}
}
