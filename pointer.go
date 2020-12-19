package main

import "fmt"

func main() {
	number1 := 1
	number2 := &number1 // number2 menyimpan alamat nilai number1

	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(*number2) // mencetak nilainya

	*number2 = 10 // merubah nilai number2

	fmt.Println(*number2)
	fmt.Println(number1) // number1 ikut berubah, karena number1 & number2 menyimpan nilai di titik yang sama

	// case
	var angka1 int = 5
	rubah(&angka1, 100)
	fmt.Println(&angka1)
	fmt.Println(angka1)
}

func rubah(old *int, new int) {
	*old = new
	fmt.Println(old)
	fmt.Println(*old)
}
