package main

import "fmt"

type host struct {
	name string
}

func (h host) sayHi() string {
	return fmt.Sprintf("Hello, friends! I'm %s.", h.name)
}

func main() {
	hosts := []host{
		{name: "Adelina"},
		{name: "Liam"},
	}
	for _, h := range hosts {
		fmt.Println(h.sayHi())
	}

}
