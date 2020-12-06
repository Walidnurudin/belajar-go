package main

import (
	"fmt"
	"strconv"
)

func penjumlahan(angka1 int, angka2 int) int {
	return angka1 + angka2
}

func bio(age int, name string, work string) string {
	// convert to string
	newAge := strconv.Itoa(age)
	return name + " adalah seorang " + work + " yang berumur " + newAge
}

// return 2 nilai
func klub(name string, born int) (string, int) {
	return "nama klub " + name, born
}

// return 2 nilai, predifined return value
func rumus(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}

// func main langsung jalan
func main() {
	fmt.Println(penjumlahan(10, 10))
	fmt.Println(bio(19, "walid", "programmer"))

	name, born := klub("persib bandung", 1933)
	fmt.Println(name, born)

	luas, keliling := rumus(10, 10)
	fmt.Println(luas, keliling)
}
