package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10.0
	b := 0.33
	c := fmt.Sprintf("%.2f", a/b)
	fmt.Println("c:", c)
	d := math.Round(a/b*1000) / 1000
	fmt.Println("d:", d)
}
