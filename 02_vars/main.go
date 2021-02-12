package main

import "fmt"

func main() {

	var age int = 44
	name, email := "istiak", "istinishat@gmail.com"
	size := 1.3

	fmt.Println(name, age, size, email)
	fmt.Printf("%T\n", size)
}
