package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	x += 2 // sama saja dengan x = x + 2
	fmt.Println("hasil dari x di tambah 2:", x)
	x -= 2 //sama saja dengan x = x - 2
	fmt.Println("hasil dari x di kurang 2:", x)
	x /= 2
	fmt.Println("hasil dari x di bagi 2:", x)
	x %= 2
	fmt.Println("hasil dari x di modulus 2:", x)
}
