package main

import "fmt"

func main() {
	var x, fx float64
	fmt.Scan(&fx)
	x = 2/(fx+5) + 5
	fmt.Print(x)
}
