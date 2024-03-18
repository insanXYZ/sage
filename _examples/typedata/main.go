package typedata

import (
	"fmt"
	"github.com/insanXYZ/sage/sage"
	"os"
)

func main() {
	//open file in your pc with os.open
	open, _ := os.Open("giphy.gif")

	//initialize sage package
	s := sage.New()

	//validate image with Validate
	err := s.Validate(open, "image", "gif") //not error
	if err != nil {
		fmt.Println(err.Error())
	}

	err = s.Validate(open, "image", "png") //error
	if err != nil {
		fmt.Println(err.Error())
	}
}
