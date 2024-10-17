package main

import (
	"fmt"

	"github.com/CGMarval/encapsulation/course"
)

func main() {
	Go := course.New("Go desde cero", 0, false)

	Go.UserIDs = []uint{12, 56, 89}
	Go.Classes = map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	}

	Go.ChangePrice(24.99)
	fmt.Println(Go.Price)
}
