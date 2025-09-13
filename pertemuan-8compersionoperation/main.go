package main

import "fmt"

func main() {
	//operasi perbandingan digunakan untuk membanding kan dua nilai, hasil pebadingan biasa nya hasil nya boolean(true false),sering dipake di if
	//operator pebandingan
	var angka1, angka2 int
	fmt.Println("masukan angka ke 1:")
	fmt.Scanln(&angka1)
	fmt.Println("masukan angka ke 2:")
	fmt.Scanln(&angka2)
	fmt.Println("\n == hasil dari perbandingan 2 angka adalah == \n")
	fmt.Printf("\n %d == %v ? %v\n", angka1, angka2, angka1 == angka2)
	fmt.Printf("\n %d != %v ? %v\n", angka1, angka2, angka1 != angka2)
	fmt.Printf("\n %d >= %v ? %v\n", angka1, angka2, angka1 >= angka2)
	fmt.Printf("\n %d <= %v ? %v\n", angka1, angka2, angka1 <= angka2)
	fmt.Printf("\n %d > %v ? %v\n", angka1, angka2, angka1 > angka2)
	fmt.Printf("\n %d < %v ? %v\n", angka1, angka2, angka1 < angka2)
	//operator logika
	if angka1 > angka2 || angka1 == angka2 {
		fmt.Println(angka1, "dan", angka2, "nilai nya sama atau lebih besar angka 1")
	}
	if angka1 == 0 && angka1 < 1 {
		fmt.Println(angka1, "bernilai lebih kecil dari 1")
	}
}
