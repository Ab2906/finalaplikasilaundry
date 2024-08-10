package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Numerik")
	var positiveNumber uint8 = 90
	var negativeNumber int = -90
	fmt.Printf("Bilangan Positif : %d\n", positiveNumber)
	fmt.Printf("Bilangan Negative : %d\n", negativeNumber)

	var decimalNumber = 3.90
	fmt.Printf("Bilangan Pecahan : %f\n", decimalNumber)
	fmt.Printf("Bilangan Pecahan : %.2f\n", decimalNumber)

	fmt.Println("==========================================")

	fmt.Println("Boolean")
	var exist = true
	fmt.Printf("exist ? %t\n", exist)

	fmt.Println("==========================================")

	fmt.Println("String")
	var message string = ("Halo")
	fmt.Printf("message : %s\n", message)
	var otherMessage string = `Nama Saya "Kenji",
Salam Kenal,
Mari Belajar di "Enigma Camp"`
	fmt.Println(otherMessage)

	fmt.Println("Kenji" + "Ferdian")

	fmt.Println(len(message))

}
