package main

import (
	"fmt"
	"math"
)

// interface
type hitung interface {
	luas() int
	keliling() int
}

// 1
type Persegi struct {
	Sisi    int
	Panjang int
	Lebar   int
}

func (item Persegi) luas() int {
	return item.Sisi * item.Sisi
}

func (item Persegi) keliling() int {
	return 2 * (item.Panjang + item.Lebar)
}

// 2
type Lingkaran struct {
	Jari_jari int
}

func (item Lingkaran) luas() int {
	phi := float64(math.Pi)
	return int(phi) * item.Jari_jari * item.Jari_jari
}

func (item Lingkaran) keliling() int {
	phi := float64(math.Pi)
	return int(phi) * item.Jari_jari
}

func main() {
	kotak1 := Persegi{Sisi: 10, Panjang: 10, Lebar: 10}
	lingkaran1 := Lingkaran{Jari_jari: 5}

	fmt.Println("kotak 1 = ", hitungLuas(kotak1), hitungKeliling(kotak1))
	fmt.Println("lingkaran 1 = ", hitungLuas(lingkaran1), hitungKeliling(lingkaran1))
}

func hitungLuas(item hitung) int {
	return item.luas()
}

func hitungKeliling(item hitung) int {
	return item.keliling()
}
