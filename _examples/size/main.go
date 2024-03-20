package size

import (
	"fmt"
	"os"

	"github.com/insanXYZ/sage"
)

func main() {
	//open file in your pc with os.open
	//this gif image have size 167kb
	open, _ := os.Open("giphy.gif")

	//validate image with Validate
	err := sage.Validate(open, "minsize=100") //not error
	if err != nil {
		fmt.Println(err.Error())
	}

	err = sage.Validate(open, "maxsize=100") //error
	if err != nil {
		fmt.Println(err.Error())
	}
}
