package main

import "fmt"

func swap2(a, b int) (int, int) {
	return b, a
}



func main() {
	fmt.Println(swap2(5, 2))

	f := func(a, b int) (int, int) {
		return b, a
	}
	fmt.Println(f(2, 3))

	//　即時関数的
	func(msg string) {
		fmt.Println(msg)
	}("ayu")
}
