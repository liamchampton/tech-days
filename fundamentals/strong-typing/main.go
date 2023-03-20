package main

import "fmt"

func main() {
	var greeting string
	hosts := []string{"Adelina", "Liam"}
	greeting = fmt.Sprintf("Hello, friends! We're %s and %s.", hosts[0], hosts[1])
	fmt.Println(greeting)
}
