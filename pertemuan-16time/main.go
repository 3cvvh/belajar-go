package main

import (
	"fmt"
	"time"
)

func mengubah_waktu(waktu time.Time, zona string) time.Time {
	zone, err := time.LoadLocation(zona)
	if err != nil {
		fmt.Println("err buat lokasi")
		return waktu
	}
	return waktu.In(zone)
}
func main() {
	nowUtc := time.Now().UTC()
	fmt.Println("utc time:", nowUtc)
	jakartaTime := mengubah_waktu(nowUtc, "Asia/Jakarta")
	fmt.Println("jakarta waktu", jakartaTime)
	new_YorkTime := mengubah_waktu(nowUtc, "America/New_York")
	fmt.Println("newyork time:", new_YorkTime)
}
