package main

import "fmt"

func main() {
	// 1
	// var languages [3]string

	// languages[0] = "JavaScript"
	// languages[1] = "Go"
	// languages[2] = "Python"

	// 2
	languages := [...]string{
		"Go",
		"JavaScript",
		"Python",
		"Ruby",
	}

	for i := 0; i < len(languages); i++ {
		fmt.Println(i, languages[i])
	}

}
