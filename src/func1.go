package main

import "fmt"

func main() {
	fmt.Println(hi("isihkawa san"))
}

func hi(name string) (msg string) {
	msg = "hi! " + name
	fmt.Println(msg)
	return
}
