package main

import "fmt"

func main() {
	fmt.Println("Operator Logika")

	a := true
	b := false
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)
	fmt.Println(false || true && false)

}
