package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["ishikawa"] = 78
	m["ayu"] = 32
	fmt.Println(m)

	m2 := map[string]int{"hoge":100, "ayu":23}
	fmt.Println(m2)
	fmt.Println(len(m2))
	delete(m, "ishikawa")
	fmt.Println(m)

	v, ok := m["ayu"]
	fmt.Println(v)
	fmt.Println(ok)
}
