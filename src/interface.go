package main

import (
	"fmt"
)

func show(t interface {}) {
	// 型アサーション
	_, ok := t.(japanese)
	if ok {
		fmt.Println("I am japanese")
	} else {
		fmt.Println("I am not japanese")
	}
}

func show2(t interface {}) {
	// 型Switch
	switch t.(type) {
	case japanese:
		fmt.Println("私は日本人です")
	case american:
		fmt.Println("私は日本人ではないです")
	}
}

type greeter interface {
	greet()
}

type japanese struct {

}

type american struct {

}

func (j japanese) greet() {
	fmt.Println("こんにちわ！")
}

func (a american) greet() {
	fmt.Println("Hello!")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
		show2(greeter)
	}
}
