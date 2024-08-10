package main

import "fmt"

func main() {
	var firstName string = "Keji"

	var lastName string = "Ferdian"

	fmt.Println(firstName, lastName)
	fmt.Println("Halo, Kenji Ferdian")

	fmt.Println("Halo, saya Kenji Ferdian berumur 25 dan saya tinggal di Jakarta")

	var (
		bootcampName, bootcampAddress = "Enigma Camp", "Jakarta Selatan"
	)
	fmt.Println(bootcampName, bootcampAddress)

	day := "Monday"
	date := "31"
	month := "October"

	fmt.Println(day, date, month+" 2024")

	var number = 21
	number = 20

	const phi = 3.14

	fmt.Println(number, phi)

	bootcamp, _ := "Enigma Camp", "Aktif7"
	fmt.Println(bootcamp)

}
