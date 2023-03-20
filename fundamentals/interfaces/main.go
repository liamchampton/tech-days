package main

import "fmt"

type Host interface {
	SayHi() string
}

type experiencedHost struct {
	name string
}

func (eh experiencedHost) SayHi() string {
	return fmt.Sprintf("Hello again, friends! I'm %s.", eh.name)
}

type newHost struct {
	name string
}

func (nh newHost) SayHi() string {
	return fmt.Sprintf("Hello, friends! I'm %s and I'm new here.", nh.name)
}

func main() {
	hosts := []Host{
		newHost{name: "Adelina"},
		experiencedHost{name: "Liam"},
	}
	for _, h := range hosts {
		fmt.Println(h.SayHi())
	}
}
