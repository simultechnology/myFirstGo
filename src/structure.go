package main

import (
	"fmt"
)

type user struct {
	name string
	score int
}

func main() {
	u := new(user)
	u.name = "ishikawa"
	(*u).score = 500
	fmt.Println(u)
	fmt.Println(*u)

	u2 := user{"ishi", 666}
	fmt.Println(u2)
	u3 := user{score: 777, name: "ta"}
	fmt.Println(u3)
}
