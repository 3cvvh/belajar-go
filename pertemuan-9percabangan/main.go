package main

import "fmt"

func main() {
	fmt.Println("masukan angka")
	var angka int
	fmt.Scanln(&angka)
	if angka >= 10 { //tanpa ()
		fmt.Println(angka, "lebih besar atau sama dengan 10")
	} else if angka == 2 { //harus ada spasi antara else if
		fmt.Println(angka, "sama dengan 2")
	} else { //harus ada jarak antar pembuka kurung
		fmt.Println(angka, "lebih kecil dari 10")
	}
	if v := angka * 2; v > 20 {
		fmt.Println(v, "lebih besar dari 10")
	}
}
