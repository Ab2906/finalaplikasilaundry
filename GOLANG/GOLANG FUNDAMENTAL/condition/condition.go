package main

import "fmt"

func main() {
	if false {
		fmt.Println("Kode dijalankan")
	}

	fmt.Println("done")

	fmt.Println("============================================================")

	if result := "lulus"; result == "lulus" {
		fmt.Println("Selamat anda,", result)
	} else {
		fmt.Println("Maaf anda,", result)
	}

	// fmt.Println(result)

	fmt.Println("============================================================")

	hour := 20

	if hour > 8 && hour < 17 {
		fmt.Println("Masih dalam rentan waktu yang diperbolehkan")
	} else {
		fmt.Println("Diluar rentan waktu")
	}

	fmt.Println("============================================================")

	GPA := 1.0
	var graduationStatus string

	if GPA >= 3.5 && GPA <= 4.0 {
		graduationStatus = "CUMLAUDE"
	} else if GPA >= 3.0 && GPA < 3.5 {
		graduationStatus = "MEMUASKAN"
	} else if GPA > 2.75 && GPA < 3.0 {
		graduationStatus = "CUKUP MEMUASKAN"
	} else {
		graduationStatus = "KURANG MEMUASKAN"
	}

	fmt.Printf("Selamat kamu lulus dengan predikat %s dengan IPK %.2f\n", graduationStatus, GPA)

	fmt.Println("============================================================")

	x := 3
	y := -1

	if x > 0 {
		if y > 0 {
			fmt.Println("x dan y lebih besar dari 0")
		} else {
			fmt.Println("x lebih besar dari 0 dan y kurang dari atau sama dari 0")
		}
	}

	fmt.Println("============================================================")

	var poin = 9
	// switch poin {
	// case 8:
	// 	{
	// 		fmt.Println("Bagus")
	// 	}
	// case 7, 6, 5:
	// 	{
	// 		fmt.Println("Cukup")
	// 	}
	// default:
	// 	{
	// 		fmt.Println("Kurang")
	// 	}

	// }

	switch {
	case poin >= 8:
		{
			fmt.Println("Bagus")
		}
		fallthrough
	case poin >= 6 && poin < 8:
		{
			fmt.Println("Cukup")
		}
	case poin >= 4 && poin < 6:
		{
			fmt.Println("kurang")
		}
	default:
		{
			fmt.Println("Gagal")
			fmt.Println("kamu perlu belajar lebih giat lagi")
		}
	}

}
