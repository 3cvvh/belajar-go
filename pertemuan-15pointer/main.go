package main

import "fmt"

func ubah(nama *string) {
	*nama = "budi" //mengambil lokasi varibel disimpan dan menganti nya dengan budi
}
func main() {
	nama := "aqil"
	ubah(&nama)
	fmt.Println(nama)
}
