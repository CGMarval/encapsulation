package main

import (
	"fmt"

	"github.com/CGMarval/encapsulation/course"
)

func main() {
	Go := &course.Course{
		Name:    "GO desde Cero",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.ChangePrice(24.99)
	fmt.Println(Go.Price)
}
