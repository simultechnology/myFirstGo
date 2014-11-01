package main

import "fmt"

func main() {
	var x int
	//x = 10 % 3
	x += 3
	x++
	fmt.Println(x)

	var s string
	s = "hello " + "world!!!!"
	fmt.Println(s)

	a := true
	b := false
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)
}
