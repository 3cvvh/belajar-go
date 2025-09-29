package main

import "fmt"

// func tambah(a, b int //tipe data value harus di deklarasikan terlebih dahulu) int { //kalo mau return harus deklarasi tipe data value yang di return
// 	hasil := a + b
// 	return hasil
// }
func variadick(angka ...int) int { //function variadic adalah function yang bisa memiliki banyak argumen tanpa harus di tambah 1 1 di param, biasanya di pake kalo gatau di argument bakal ngasih berapa value
	hasil := 0
	for _, v := range angka {
		hasil += v
	}
	return hasil
}
func tambah(a, b int) (hasil int) { //return bisa di beri nama dan tinggal menjadi wadah untuk value
	hasil = a + b
	return
}
func bagi(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "tidak bisa dibagi dengan 0 brengsek"
	}
	return a / b, "BERHASIL"
}

// func main() {
// 	// anjay := tambah(1, 2)
// 	anjay, pesan := bagi(2, 0) //kalo return value nya lebih dari 1 harus ada variabel tambahan buat nampung
// 	fmt.Println(anjay, pesan)
// }
func main() {
	var angka1, angka2, angka3 int
	fmt.Println("masukan angka 1")
	fmt.Scan(&angka1)
	fmt.Scan(&angka2)
	fmt.Scan(&angka3)
	fmt.Println(variadick(angka1, angka2, angka3))
}
