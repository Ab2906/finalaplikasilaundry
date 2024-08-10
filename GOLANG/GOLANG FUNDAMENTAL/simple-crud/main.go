package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

type User struct {
    ID   int
    Name string
    Age  int
}

var users = []User{}
var currentID = 1

func createUser(name string, age int) {
    user := User{ID: currentID, Name: name, Age: age}
    users = append(users, user)
    currentID++
    fmt.Println("User created successfully!")
}

func readUsers() {
    if len(users) == 0 {
        fmt.Println("No users found.")
        return
    }
    fmt.Println("List of Users:")
    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
    }
}

func updateUser(id int, name string, age int) {
    for i, user := range users {
        if user.ID == id {
            users[i].Name = name
            users[i].Age = age
            fmt.Println("User updated successfully!")
            return
        }
    }
    fmt.Println("User not found.")
}

func deleteUser(id int) {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            fmt.Println("User deleted successfully!")
            return
        }
    }
    fmt.Println("User not found.")
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("\nSimple CRUD Application")
        fmt.Println("1. Create User")
        fmt.Println("2. Read Users")
        fmt.Println("3. Update User")
        fmt.Println("4. Delete User")
        fmt.Println("5. Exit")
        fmt.Print("Choose an option: ")

        scanner.Scan()
        option, _ := strconv.Atoi(scanner.Text())

        switch option {
        case 1:
            fmt.Print("Enter name: ")
            scanner.Scan()
            name := scanner.Text()
            fmt.Print("Enter age: ")
            scanner.Scan()
            age, _ := strconv.Atoi(scanner.Text())
            createUser(name, age)
        case 2:
            readUsers()
        case 3:
            fmt.Print("Enter user ID to update: ")
            scanner.Scan()
            id, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Enter new name: ")
            scanner.Scan()
            name := scanner.Text()
            fmt.Print("Enter new age: ")
            scanner.Scan()
            age, _ := strconv.Atoi(scanner.Text())
            updateUser(id, name, age)
        case 4:
            fmt.Print("Enter user ID to delete: ")
            scanner.Scan()
            id, _ := strconv.Atoi(scanner.Text())
            deleteUser(id)
        case 5:
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid option! Please try again.")
        }
    }
}