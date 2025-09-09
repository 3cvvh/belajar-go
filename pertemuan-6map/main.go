package main

import "fmt"

func main() {
	mahasiswa := map[string]string{ //cara membuat map(mirip mirip array assosiatif kalo di php) bedanya tipe data key sama value nya harus di setting dulu
		"nama":  "aqil",
		"kelas": "11 RPL 2",
	}
	fmt.Println(mahasiswa["nama"])
	fmt.Println(mahasiswa["kelas"])
	mahasiswa["sekolah"] = "smkn7"    //bisa menambahkan key baru
	fmt.Println(mahasiswa["sekolah"]) //mengambil value map
	//function map
	M := make(map[string]string) //membuat map kosong dan bisa diisi
	M["nama"] = "aqil"
	delete(M, "nama") //delete key dari map
	fmt.Println(M["nama"])
	fmt.Println(len(mahasiswa)) //mengecek panjang map

}
