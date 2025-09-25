package main

import "fmt"

func main() {
	var angka int
	fmt.Scanln(&angka)
	switch angka {
	case 1:
		fmt.Println("angka 1")
	case 2:
		fmt.Println("angka 2")
	case 3:
		fmt.Println("angka 3")
	default:
		fmt.Println("lebih besar dari 2")
	}
}
