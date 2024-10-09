package main

import "fmt"

func main() {
	var celcius, reamur, fahrenheit, kelvin int
	fmt.Print("Temperatur celcius :")
	fmt.Scan(&celcius)
	reamur = celcius * 4 / 5
	fahrenheit = (celcius * 9 / 5) + 32
	kelvin = (fahrenheit + 460) * 5 / 9
	fmt.Println("Derajat reamur:", reamur)
	fmt.Println("Derajat fahrenheit:", fahrenheit)
	fmt.Print("Derajat kelvin:", kelvin)
}
