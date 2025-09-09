package main

import "fmt"

func main() {
	example := [3]int{10, 20, 30}
	s0 := example[:] //mengambilsemua elemen di dalam array
	fmt.Println(s0)
	s1 := example[:2] //mengambil elemen sampai inputan key(key yang di input gabakal ada)
	fmt.Println(s1)
	s2 := example[1:] // kebalikan nya. cuma bedanya ini di awal bukan di akhir
	fmt.Println(s2)
	fmt.Println(example[0:2]) //bisa diatur mulai sampai mana akhir sampai mana
	//function slice
	s := make([]int, 3, 5)
	fmt.Println(s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	s = append(s, 10, 20, 30, 40, 50)
	fmt.Println("setelah di append:", s)
	slice := make([]int, 3)
	slice3 := copy(slice, s)
	fmt.Println("setelah di copy:", slice)
	fmt.Println("jumlah elemen:", slice3)
}
