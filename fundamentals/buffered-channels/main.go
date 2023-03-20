package main
import "fmt"

type Host interface {
	SayHi(chan string)
}

type experiencedHost struct {
	name string
}

func (eh experiencedHost) SayHi(ch chan string) {
	ch <- fmt.Sprintf("Hello again, friends! I'm %s.", eh.name)
}

type newHost struct {
	name string
}

func (nh newHost) SayHi(ch chan string) {
	ch <- fmt.Sprintf("Hello, friends! I'm %s and I'm new here.", nh.name)
}

func main() {
	ch := make(chan string, 2)
	hosts := []Host{
		newHost{name: "Adelina"},
		experiencedHost{name: "Liam"},
	}
	for _, h := range hosts {
		go h.SayHi(ch)
	}
	for i := 0; i < len(hosts); i++ {
		fmt.Println(<-ch)
	}
	close(ch)

	fmt.Println("Goodbye, friends!")
}