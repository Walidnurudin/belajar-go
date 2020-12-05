package main

import (
	"fmt"
)

func main() {
	nilai := 70
	var grade string

	if nilai <= 50 {
		if nilai <= 55 {
			grade = "C-"
		} else {
			grade = "C+"
		}
	} else if nilai <= 60 {
		if nilai <= 65 {
			grade = "B-"
		} else {
			grade = "B+"
		}
	} else if nilai >= 70 {
		if nilai < 75 {
			grade = "A-"
		} else {
			grade = "A+"
		}
	} else {
		grade = "unknown"
	}

	fmt.Println(grade)
}
