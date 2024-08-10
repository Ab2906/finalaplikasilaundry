package main

import "fmt"


var name = "Kenji"

func main() {
	helloWorld()
	fmt.Println("main:", name)
	greet("Kenji", 25)
	greet("Jhon", 19)
	greet("Doe", 20)
	//greet(21)
	// total := add(2, 6)
	add(2, 6)
}
func helloWorld() {
	var name = "Ferdian"
	fmt.Println("Hello World!")
	fmt.Println("helloWorld:", name)
	
}

func greet (name string, age int)  {
	fmt.Println("Hello my name is", name, "I'm", age, "years old")
}

func add(a int, b int)  {
	fmt.Println("Result Add", a + b)
}