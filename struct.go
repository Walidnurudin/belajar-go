package main

import "fmt"

type User struct {
	ID      int
	Name    string
	Address string
	Email   string
}

// struct sebagai parameter func
func perkenalan(user User) string {
	return fmt.Sprintf("Hallo, nama saya %s, saya asal %s.", user.Name, user.Address)
}

func main() {
	// 1
	// walid := User{}
	// walid.ID = 1
	// walid.Name = "Walid nurudin"
	// walid.Address = "Indramayu"
	// walid.Email = "walid@gmail.com"

	// 2
	// walid := User{
	// 	ID:      1,
	// 	Name:    "Walid nurudin",
	// 	Address: "Indramayu",
	// 	Email:   "walid@gmail.com",
	// }

	// 3
	walid := User{1, "Walid nurudin", "Indramayu", "walid@gmail.com"}

	fmt.Println(perkenalan(walid))
}
