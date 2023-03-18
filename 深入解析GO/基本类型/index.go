package main

import "fmt"

type Point struct {
	x, y int
}

type Rect struct {
	x, y Point
}

type pRect struct {
	x, y *Point
}

func main() {
	x := Point{2, 10}
	y := Point{1, 9}
	z := Rect{x, y}
	c := pRect{&x, &y}
	fmt.Printf("x, y, z, c =====> 🚀🚀🚀 %v\n %v\n%v\n%v\n", x, y, z, c)
}
