package greetings

import (
	"fmt"
	"errors"
)
//In Go, a function that begins with a capital letter is an exported name, i.e., it can be used in other packages
func Hello (name string) (string, error) {
	//var message string <-------------- only declaration
	if name == "" {
		return "", errors.New("empty name") //return both return values, in error case return empty string
	}
	message:= fmt.Sprintf("Hi, %v. Welcome!", name) //declaration and initalization
	return message, nil //return both return values, in correct case return nil error
}