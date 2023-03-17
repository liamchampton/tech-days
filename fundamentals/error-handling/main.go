package main

import (
	"errors"
	"fmt"
	"log"
)

func sayHi(hosts []string) error {
	if len(hosts) == 0 {
		return errors.New("empty hosts")
	}
	greeting := fmt.Sprintf("Hello, friends! We're %s and %s.", hosts[0], hosts[1])
	log.Println(greeting)
	return nil
}

func main() {
	defer func() {
		log.Println("Goodbye, friends!")
	}()
	if err := sayHi([]string{}); err != nil {
		log.Fatalf("error saying hi:%v\n", err)
	}
}
