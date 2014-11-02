package main

import "fmt"

func main() {
	//s := make([]int, 3)
	s := []int{1, 3, 5}

	// append
	s = append(s, 8, 2, 10)

	t := make([]int, len(s))
	//var w[8]int
	//t := w[2:]
	n := copy(t, s)
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(n)
}
