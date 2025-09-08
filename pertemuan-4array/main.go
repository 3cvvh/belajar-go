package main

import "fmt"

func main() {
	var angka [5]int = [5]int{10, 20, 30, 40, 50} //cara membuat array
	fmt.Println("array:", angka)                  //cara memanggil semua array
	angka[0] = 200                                //menimpa nilai array key 0
	fmt.Println("array key index 0:", angka[0])   //cara memanggil array dengan array key tertentu
	//function array
	array2 := [2]string{"andrew", "graves"}
	array := [2]string{"ashley", "graves"}    //array yang di tulis dengan cara modern
	fmt.Println("jumlah elemen:", len(array)) //mengecek ada berapa key atau elemen di dalam array
	for i, v := range array {                 //fungsi range mengambil jumlah key dan value dari array
		fmt.Println("index ke:", i, "yang value nya:", v) //me loping
	}
	fmt.Println("array 1 == arry 2?", array2 == array) //membanding kan array jika sama bakal true
	fmt.Println("array 1 != arry 2?", array != array2) //jika tidak sama bakal true
}
