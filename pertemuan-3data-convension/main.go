package main

import (
	"fmt"
	"strconv"
)

func main() {
	//ubah tipe data integer ke float
	// var a int = 2
	// var b float32 = float32(a) //diubah
	// fmt.Println("nilai a:", a)
	// fmt.Println("nilai b:", b)
	// var text string = "100b"
	// number, err := strconv.Atoi(text)
	// if err != nil {
	// 	fmt.Println("pesan Error :", err.Error())
	// } else {
	// 	fmt.Println("angka:", number)
	// }
	//boolean ke string
	// truth := true
	// val1 := strconv.FormatBool(truth)
	// fmt.Println(val1)
	//string ke boolean
	benar := "true"
	val, _ := strconv.ParseBool(benar)
	fmt.Println(val)
}
