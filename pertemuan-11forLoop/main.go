package main

import "fmt"

func main() {
	//for normal
	for i := 0; i < 10; i++ {
		if i == 4 { //nge skip angka 4(jadi gabakal ada di looping)
			continue
		}
		fmt.Printf("\n iterasi ke: %d", i)
	}
	fmt.Println()
	//for yang mirip seperti while di bahasa lain
	a := 0       //nilai awal
	for a <= 9 { //kondisi
		fmt.Printf("\niterasi ke: %d", a)
		a = a + 1 //incerment a
	}
	wife := map[string]string{"istri1": "ashley",
		"istri2": "ashley2",
		"istri3": "ashley3"}
	for index, value := range wife {
		fmt.Printf("\n istri ke: %d = %s", index, value)
	}
}
