package main

import "fmt"

func main() {

	// 1
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("Baris ke", i)
	// }

	// 2
	// number := 0
	// for number <= 100 {
	// 	fmt.Println("Baris ke", number)
	// 	number++
	// }

	// 3
	// name := "Walid nurudin"
	// for index, huruf := range name {
	// 	fmt.Println("index ke", index, "huruf = ", string(huruf))
	// }

	// case
	name := "Walid nurudin"
	for index, huruf := range name {
		if index%2 == 0 {
			fmt.Println(string(huruf))
		}
	}
}
