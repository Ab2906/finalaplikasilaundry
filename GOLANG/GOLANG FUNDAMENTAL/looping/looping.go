package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Angka", i)
	}

	fmt.Println("==============================================================================")

	var i = 6
	for i <= 10 {
		fmt.Println("Angka", i)
		i++
	}

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	var index = 0
	for {
		fmt.Println("Enigma")
		index++
		if index == 5 {
			break
		}
	}


	fmt.Println("$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}


	fmt.Println("$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$%$")

	fmt.Println("Start")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, " ")
	}

	fmt.Println()
	fmt.Println("Done")


	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	fmt.Println("Start")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i, " ")
	}

	fmt.Println()
	fmt.Println("Done")

	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	fmt.Println("Start")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			if j == 4 {
				// break
				continue
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
		fmt.Println("Done")


}