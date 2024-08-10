package main

import (
	"fmt"
	
)
	//membuat map dengan keyword map
func main() {
	user := map[string]string{
		"username": "kenjif",
		"email": 	"kenji@gmail.com",
	}
	fmt.Println(user)

	fmt.Println("=================================================================")
	//membuat map dengan keyword make
	var scores = make(map[string]int)
	fmt.Println(scores)

	scores["java"] = 85
	scores["react"] = 87
	scores["kotlin"] = 90

	fmt.Println("Scores:", scores)
	fmt.Println("Nilai Java:", scores["java"])
	fmt.Println("Nilai Kotlin:", scores["kotlin"])
	fmt.Println("Nilai Golang:", scores["golang"])
//mengubah value yg ada di dalam map
	scores["java"] = 90

	fmt.Println("=================================================================")
	fmt.Println("Scores:", scores)

	fmt.Println("=================================================================")
//menghapus value yg ada didalam map
	delete(scores, "kotlin")
	fmt.Println("Scores:", scores)

//menggunakan for range untuk melakukan looping didalam map
fmt.Println("=================================================================")

names := map[int]string {
	1 : "jhon",
	2 : "jane",
	3 : "lily",
	4 : "ruby",
}

	for _, value := range names {
		fmt.Println("value", value)
	}

}
