package main

import "fmt"

func sayHi(hosts []string) {
	greeting := fmt.Sprintf("Hello, friends! We're %s and %s.", hosts[0], hosts[1])
	fmt.Println(greeting)

}

func main() {
	defer func() {
		fmt.Println("Goodbye, friends!")
	}()
	hosts := []string{"Adelina", "Liam"}
	sayHi(hosts)
}
