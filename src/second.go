package main

import "fmt"

func main() {
	fmt.Println("start!")
	a := 10
	b := 12.3
	c := "hoge"
	var d bool
	fmt.Printf("a: %d, b: %f, c: %s, d: %t \n", a, b, c, d)

	const (
		name = "ishikawa"

	)
	fmt.Println(name)

	const (
		sun = iota
		mon
		tue
	)
	fmt.Println(sun, mon, tue)
}
