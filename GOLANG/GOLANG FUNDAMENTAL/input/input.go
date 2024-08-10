package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var firtsName string
	var lastsName string
	var age int

	fmt.Print("Enter Your Name : ")
	fmt.Scanln(&firtsName, &lastsName)
	fmt.Print("Your Name is", firtsName, lastsName)
	fmt.Print("Enter Your Age : ")
	fmt.Scanln(&age)
	birthYear := (2024 - age)

	fmt.Println("Your Name is", firtsName, lastsName, "You are born in", birthYear)

	fmt.Println("===============================================================================")

	var fullName string
	var address string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Data Diri Pegawai Enigma")
	fmt.Print("Masukan Nama Anda : ")
	scanner.Scan()
	fullName = scanner.Text()
	fmt.Print("Masukan Umur Anda : ")
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Masukan Alamat Anda : ")
	scanner.Scan()
	address = scanner.Text()

	fmt.Printf("%s | %d | %s", fullName, age, address)

}
