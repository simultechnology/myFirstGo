package main

import (
	"fmt"
)

func main() {
	//var msg string
	//msg = "hello world"
	//var msg string = "hoge world"
	msg := "hello world"
	fmt.Println(msg)

	//var a, b int
	//a, b = 10, 15
	a, b := 10, 15

	var (
		c int
		d string
	)
	c = 20
	d = "hoge"

	fmt.Printf("%d , %b\n", a, b)
	fmt.Printf("%d , %s\n", c, d)
}
