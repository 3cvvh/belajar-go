package main

import "fmt"

type human struct { //mendefinisikan struct
	name   string
	age    int
	school string
}

func (h human) perkenalan() { //method struct
	fmt.Println("nama saya", h.name, "umur saya", h.age, "sekolah saya di", h.school)
}

func main() {
	user := human{
		name:   "aqil",
		age:    18,
		school: "baleendah", //memberikan value pada struct yg telah di definisikan
	}
	user.perkenalan() //memmanggil method struct dengan data yang sudah diisi
}
