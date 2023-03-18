package main

import (
	"fmt"
)

func f(i int) func() int {
	return func() int {
		i++
		return i
	}
}

func adder() func(int2 int) int {
	a := 0
	return func(x int) int {
		a += x
		return a
	}
}

func sum(num int) func() int {
	a := 0
	return func() int {
		a += num
		return a
	}
}

func main() {
	c1 := f(0)
	c2 := f(0)
	c1() // reference to i, i = 0, return 1
	fmt.Printf("c1() =====> 🚀🚀🚀 %v\n", c1())
	c2() // reference to another i, i = 0, return 1
	fmt.Printf("c2() =====> 🚀🚀🚀 %v\n", c2())
	a := adder()
	fmt.Printf("a(2) =====> 🚀🚀🚀 %v\n", a(2))
	b := adder()
	fmt.Printf("b(1) =====> 🚀🚀🚀 %v\n", b(1))
	z := sum(1)
	fmt.Printf("z() =====> 🚀🚀🚀 %v\n", z())
	fmt.Printf("z() =====> 🚀🚀🚀 %v\n", z())
	fmt.Printf("z() =====> 🚀🚀🚀 %v\n", z())
	fmt.Printf("z() =====> 🚀🚀🚀 %v\n", z())
}
