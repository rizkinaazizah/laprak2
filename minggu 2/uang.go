package main

import "fmt"

func main() {
	var rupiah, dollar float64
	fmt.Print("Masukkan nilai rupiah:")
	fmt.Scan(&rupiah)
	dollar = rupiah / 15000
	fmt.Println("Jadi ", rupiah, "rupiah = ", dollar, "dollar")

}
