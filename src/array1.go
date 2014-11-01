package main

import "fmt"

func main() {
	var a[5]int
	a[2] = 3
	a[4] = 10
	fmt.Println(a)
	fmt.Println(a[2])

	b := []int{1, 3, 5, 10}

	fmt.Println(b)
	fmt.Println(len(b))

	fmt.Println("-------------------------")

	myA := [5]int{2, 10, 8, 15, 4}
	fmt.Println(myA)

	// スライス
	s := myA[2:4]
	fmt.Println(s)
	fmt.Println(myA[:4])
	fmt.Println(myA[2:])

	s[1] = 12
	fmt.Println(myA)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
