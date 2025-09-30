// package main
package main

import "fmt"

// import "fmt"

// //dpr adalah singkatan dari defer, panic, recover
// func main() {
// 	x := 1
// 	defer fmt.Println("ini harus nya 1 =", x) //defer akan menjalankan perintah menjadi urutan terakhir
// 	x = 2
// 	defer fmt.Println("ini harus nya 2 =", x) //defer megunakan konsep last in,first out jadi yg terakhir masuk yang pertama keluar
// 	fmt.Println("void:", x)
// 	defer fmt.Println("paniccc")
// 	fmt.Println("sebelum panic")
// 	panic("ada sesuatu yg fatal") //panic akan membuat code yang dibawah nya tidak di eksekusi dan semua defer akan dipanggil
// 	fmt.Println("setelah panic")
// }
func clenup() {
	fmt.Println("program dibersihkan...")
}
func bacaconfig(namafile string) string {
	defer clenup()
	defer func() {
		if r := recover(); r != nil { //func anonymus recover jika panic maka akan menjalankan recver ini
			fmt.Println("terjadi error:", r, "...")
		}
	}()
	if namafile == "" {
		panic("nama file tidak boleh kosong")
	}
	return namafile
}
func main() {

	fmt.Println(bacaconfig(""))
	fmt.Println("anjay")
}
