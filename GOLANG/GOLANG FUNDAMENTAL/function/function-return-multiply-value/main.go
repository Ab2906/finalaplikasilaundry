package main

import "fmt"

func main() {
	num := []int {4, 1, 3, 7, 5}
	_, besar := minMax(num)
	// fmt.Println("Kecil :", kecil)
	fmt.Println("Besar :", besar)
}

func minMax(numbers []int) (min int, max int) {
	min = numbers[0]
	max = numbers[0]

	for _, n := range numbers{
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return

}