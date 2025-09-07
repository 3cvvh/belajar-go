package main

import "fmt"

func main() {
	// cara klasik: sebutkan tipe data dan nama nya
	var nama string = "agus"
	var umur int = 20
	//membuat variabel tanpa mendefinisikan tipe data nya(otomatis di definisikan oleh golang)
	city := "bendung"
	waifu := "ashley graves"
	year := 2007
	//constanta tidak bisa di timpa oleh nilai terbaru
	const suami = "aqil"
	//tipe data boolean
	benar := true
	salah := false
	//raw string
	paragaraf := `ashley kau lah segalanya bagiku
	hanya kau`
	//bisa juga pake cara klasik yang di didefinisikan dulu tipe data nya
	fmt.Println("nama:", nama)
	fmt.Println("umur:", umur)
	fmt.Println("kota:", city)
	fmt.Println("waifu:", waifu)
	fmt.Println("tahun lahir:", year)
	fmt.Println("nama suami:", suami)
	fmt.Println("benar", benar)
	fmt.Println("salah", salah)
	fmt.Println("kutipan\n", paragaraf)

}
