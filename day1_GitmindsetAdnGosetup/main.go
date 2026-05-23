package main

import "fmt"

func main() {
	createAccount("Thong", 26, "nam")
}

func createAccount(name string, age int, gender string) {
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Gender: %s\n", gender)
}
