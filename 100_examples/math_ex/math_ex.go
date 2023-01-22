package main

import (
	"fmt"
	"math"
)

func main() {
	var rad float64 = 1234.32145
	grad1 := (math.Pi / 180) * rad
	grad2 := (math.Pi * rad) / 180

	fmt.Println("grad", grad1)
	fmt.Println("grad", grad2)

}
