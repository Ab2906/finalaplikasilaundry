package main

import "fmt"

func main() {
	numbersSlice := []int{2, 7, 9, 4}
	fmt.Println(numbersSlice)

	fmt.Println("Panjang Slice:", len(numbersSlice))
	fmt.Println("Kapasitas Slice:", cap(numbersSlice))

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	scores := make([]int, 4, 7)
	scores[0] = 70
	scores[1] = 80
	scores[2] = 95
	scores[3] = 65

	fmt.Println(scores)
	fmt.Println("Panjang scores :", len(scores))
	fmt.Println("Kapasitas scores :", cap(scores))

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	scores2 := make([]int, 4)
	scores2[0] = 70
	scores2[1] = 80
	scores2[2] = 95
	scores2[3] = 65

	fmt.Println(scores2)
	fmt.Println("Panjang scores 2 :", len(scores2))
	fmt.Println("Kapasitas scores 2 :", cap(scores2))

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	heroes := [4]string{"Superman", "Batman", "Spiderman", "Ironman"}
	fmt.Println("Heroes:", heroes)
	//heroes[5] = "Catwomen" // error saat penambahan

	villain := make([]string, 3, 5)
	villain[0] = "Thanos"
	villain[1] = "Joker"
	villain[2] = "HomeLander"

	fmt.Println("Villain:", villain)
	fmt.Println("Panjang villain:", len(villain))
	fmt.Println("Kapasitas villain:", cap(villain))
	villain = append(villain, "Ultron")
	villain = append(villain, "Joko", "Sumanto")
	fmt.Println("Villain:", villain)
	fmt.Println("Panjang villain:", len(villain))
	fmt.Println("Kapasitas villain:", cap(villain))

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	//versi array fase by value
	var numbers = [4]int{2, 1, 6, 8}
	var anotherNumbers = numbers
	fmt.Println("Numbers:", numbers)
	fmt.Println("Another Numbers:", anotherNumbers)
	anotherNumbers[1] = 11
	fmt.Println("Numbers:", numbers)
	fmt.Println("Another Numbers:", anotherNumbers)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	//versi slice fase by reference
	var prices = []int{20000, 13000, 10000}
	var anotherPrices = prices
	fmt.Println("Prices:", prices)
	fmt.Println("Another Prices:", anotherPrices)
	anotherPrices[1] = 15000
	fmt.Println("Prices:", prices)
	fmt.Println("Another Prices:", anotherPrices)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	//membuat slice dari array yg sudah ada // fase by reference
	ages := [4]int{20, 22, 25, 19}
	sliceAges := ages [0:3]
	fmt.Println("Ages:", ages)
	fmt.Println("slice Ages:", sliceAges)
	sliceAges[2] = 33
	fmt.Println("Ages:", ages)
	fmt.Println("slice Ages:", sliceAges)
}