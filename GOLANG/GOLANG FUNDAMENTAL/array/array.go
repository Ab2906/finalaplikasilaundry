package main

import "fmt"

func main() {
	var fruits = [4] string{"Apel", "Pisang", "Anggur", "Semangka"}
	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fruits[2] = "Jeruk"
	fmt.Println(fruits)

	fmt.Println("===============================================================================")

	var scores [3] int
	scores[0] = 87
	scores[1] = 78
	scores[2] = 92
	// scores[3] = 70 error out of bounds

	fmt.Println(scores)

	fmt.Println("===============================================================================")

	var days = [...]string {"Senin", "Selasa", "Rabu","kamis", "Jumat", "Sabtu", "Minggu"}

	// days[7] = "Libur" error out of bounds

	fmt.Println(days)
	fmt.Println("Jumlah elemen :", len(days))

	fmt.Println("===============================================================================")

	for i := 0; i < len(days); i++ {
		fmt.Printf("Elemen %d: %s\n", i, days[i])
	}

	fmt.Println("===============================================================================")
	
	for i, day := range days {
		fmt.Printf("Elemen %d: %s\n", i, day)
	}

	fmt.Println("===============================================================================")
	
	for _, day := range days {
		fmt.Printf("Nama Hari : %s\n", day )
	}

	fmt.Println("===============================================================================")

	for i := range days {
		fmt.Printf("Index hari ke : %d\n", i)
	}

	fmt.Println("===============================================================================")

	var numbers = [5]int {3, 8, 2, 7, 4}

	for _, number := range numbers{
		if number % 2 == 0 {
			fmt.Println(number)
		}
	}

	fmt.Println("===============================================================================")

	fmt.Println("Sebelum :", numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] *=2
	}
	fmt.Println("Setelah :", numbers)

	fmt.Println("===============================================================================")

	var matrix = [2] [3]int {
		{3, 2, 3},
		{3, 4, 5},
	}
	fmt.Println(matrix)
	fmt.Println(matrix[1][1])
	fmt.Println(matrix[1][2])
}