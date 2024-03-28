package main

import (
	"fmt"

	"github.com/linkedinlearning/domina-go/package-layout/geometry"
)

func main() {
	c := geometry.NewCircle(1, 2, 3)
	fmt.Println(c)

	tr := geometry.NewTriangle(1, 2, 3, 4, 5, 6)
	fmt.Println(tr)

	p := geometry.NewParallellogram(1, 2, 3, 4)
	fmt.Println(p)

	point := geometry.point{}
	fmt.Println(point)
}
