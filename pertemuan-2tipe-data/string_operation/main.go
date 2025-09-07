package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `ohhh ashleyyy,
	kau lah wanita tercantik yang pernah kulihat dalam hidup ku`
	//megunbah menjadi besar
	fmt.Println("besar:\n", strings.ToUpper(text))
	//megubah menjadi kecil
	fmt.Println("kecil:\n", strings.ToLower(text))
	//mengecek apakah dimulai dengan ohhh
	fmt.Println("dimulah dari ohh?", strings.HasPrefix(text, "ohh"))
	//memisahkan string berdasarkan delimiter
	fmt.Println("pisah berdasarkan ,", strings.Split(text, ","))
	//mengecek apakah ada karakter tertentu dalam sebuah string
	fmt.Println("apakah mengandung kata tercantik?", strings.Contains(text, "tercantik"))
	//untuk mengubah karakter tertentu di variabel
	replace := strings.ReplaceAll(text, "tercantik", "ter imut")
	fmt.Println(replace)
}
