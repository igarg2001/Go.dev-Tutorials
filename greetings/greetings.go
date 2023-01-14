package greetings

import (
	"fmt"
)
//In Go, a function that begins with a capital letter is an exported name, i.e., it can be used in other packages
func Hello (name string) string {
	//var message string <-------------- only declaration
	message:= fmt.Sprintf("Hi, %v. Welcome!", name) //declaration and initalization
	return message
}