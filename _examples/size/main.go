package size

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
	err := s.Validate(open, "minsize=100") //not error
	if err != nil {
		fmt.Println(err.Error())
	}

	err = s.Validate(open, "maxsize=100") //error
	if err != nil {
		fmt.Println(err.Error())
	}
}
