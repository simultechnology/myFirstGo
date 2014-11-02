package main

import "fmt"

func main() {
	s := []int{2, 3, 8}
	for i, v := range s {
		fmt.Println(i, v)
	}

	for _, v := range s {
		fmt.Println(v)
	}

	m := map[string]int{"ishi":200, "taka":3000}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
