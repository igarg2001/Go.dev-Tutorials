package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

// import "rsc.io/quote"

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names:=[]string{"John", "Alex", "Samantha"}

	// message, err := greetings.Hello("Ishan")
	// if err!=nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)

	messages, err := greetings.Hellos(names)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}