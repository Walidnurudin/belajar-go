package main

import "fmt"

type Person struct {
	Nama   string
	Umur   int
	Alamat string
}

// method
func (person Person) show() string {
	return fmt.Sprintf("Nama saya %s, umur %d", person.Nama, person.Umur)
}

func main() {
	walid := Person{"Walid nurudin", 20, "Indramayu"}
	fmt.Println(walid.show())
}
