package main

import (
	"fmt"
)

func main() {
	//anonymous function
	func(){
		fmt.Println("Hello World!")
	}()

	//anonymous function dalam variabel
	halo := func() {
		fmt.Println("Halo Dunia!")
	}
	halo()

	// passing argument kedalam anonymous function
	func(word string){
		fmt.Println(word)
	}("EnigmaCamp")

	// passing argument kedalam variabel function

	hello := func(name string) {
		fmt.Println("Hello, ", name)
	}
	hello("Jhonny")

// passing anonymous function sebagai argument
greetEnglish := func (name string) string {
	return "Hello " + name
}
greetRussian := func (name string) string {
	return "Привет " + name
}
greetKorean := func (name string) string {
	return "안녕하세요 " + name
}
greet("Joko", greetEnglish)
greet("Joko", greetRussian)
greet("Joko", greetKorean)

	add := func(num1 int, num2 int) int {
	return num1 + num2	}
	multiply := func(num1 int, num2 int) int {
	return num1 * num2
	}

	calculate(3,2, add)
	calculate(4,5, multiply)
}


func greet(name string, f func(name string) string) {
	fmt.Println(f(name))
} 

func calculate(a int, b int, operator func(x int, y int) int) {
	fmt.Println(operator(a,b))
}
	
	