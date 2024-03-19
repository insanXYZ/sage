package main

import (
	"fmt"
	"os"

	"github.com/insanXYZ/sage"
)

func main() {
	//open file in your pc with os.open
	open, _ := os.Open("giphy.gif")

	//initialize sage package
	s := sage.New()

	//validate image with Validate
	err := s.Validate(open, "gif") //not error
	if err != nil {
		fmt.Println(err.Error())
	}

	err = s.Validate(open, "png") //error
	if err != nil {
		fmt.Println(err.Error())
	}
}
