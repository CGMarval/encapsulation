package main

import (
	"curso-poo-go/course"
	"fmt"
)

func main() {
	Go := &course.Course{
		"GO desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.ChangePrice(24.99)
	fmt.Println(Go.Price)
}
