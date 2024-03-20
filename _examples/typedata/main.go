package main

import (
	"fmt"
	"os"

	"github.com/insanXYZ/sage"
)

func main() {
	//open file in your pc with os.open
	open, err := os.Open("giphy.gif")
	if err != nil {
		fmt.Println(err.Error())
	}

	//validate image with Validate
	err = sage.Validate(open, "gif") //not error
	if err != nil {
		fmt.Println(err.Error())
	}

	err = sage.Validate(open, "png") //error
	if err != nil {
		fmt.Println(err.Error())
	}
}
