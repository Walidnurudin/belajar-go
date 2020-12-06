package main

import "fmt"

func main() {
	// map[key]value

	// 1
	// var binatang map[string]int
	// binatang = map[string]int{}

	// binatang["kadal"] = 1
	// binatang["kucing"] = 2

	// 2
	binatang := map[string]int{
		"kadal":  1,
		"kucing": 2,
	}

	for index, item := range binatang {
		fmt.Println(index, item)
	}
}
