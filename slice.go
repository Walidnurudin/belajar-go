package main

import "fmt"

func main() {
	// 1
	// var nama []string

	// nama = append(nama, "Walid")
	// nama = append(nama, "nurudin")

	// 2
	// nama := []string{"walid", "nurudin"}

	// fmt.Println(nama)

	// slice of map
	students := []map[int]string{
		{1: "walid", 2: "indramayu"},
		{1: "wawewow", 2: "pilang payung"},
	}

	for _, student := range students {
		fmt.Println(student[1] + " - " + student[2])
	}
}
