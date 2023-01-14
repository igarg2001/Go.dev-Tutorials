package main

import (
	"fmt"
	"example.com/greetings"
)

// import "rsc.io/quote"

func main() {
	message:= greetings.Hello("Ishan")
	fmt.Println(message)
}