# The fundamental concepts within Go
The purpose of this section is to equip our listeners with all the knowledge they require for our session, but we will not teach them Go from scratch.

## Strong typing
- The compiler is your best friend. 
- At all points, we will know the type of our variables, and what behaviour they expose.
- Dynamic typing is slower and the compiler avoids runtime errors and fatal error cases due to undefined behaviours.
- The fmt package is part of the standard library and allows us to format and print strings.
- The Go toolchain builds and runs our programs.

[Playground](https://go.dev/play/p/DCh2szuTTbm)
```go
package main

import "fmt"

func main() {
    var greeting string
	hosts := []string{"Adelina", "Liam"}
	greeting = fmt.Sprintf("Hello, friends! We're %s and %s.", hosts[0], hosts[1])
	fmt.Println(greeting)
}
```

## Functions
- Go functions are natively supported and can be passed as variables, return types and parameters for later invocation. 
- Anonymous functions are also allowed.
- Function composition is easy to do in Go. 
- Deferred functions are useful for guaranteed clean up tasks. 

[Playground](https://go.dev/play/p/fDUKmkVvezv)

```go
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

```

## Custom types
- Go is not an object oriented language, as it does not support type hierarchy. 
- Structs allow us to build custom types and behaviours.
- They are a collection of fields, which can be partially initialised.
- Custom types can also define methods by using a special receiver argument that is the implicit first argument of the method.

[Playground](https://go.dev/play/p/DDT86lbYhhz)
```go
type host struct {
	name string
}

func (h host) sayHi() string {
	return fmt.Sprintf("Hello, friends! I'm %s.", h.name)
}

func main() {
	hosts := []host{
		host{name: "Adelina"},
		host{name: "Liam"},
	}
	for _, h := range hosts {
		fmt.Println(h.sayHi())
	}

}
```

## Visibility
- Go code is organised in packages, which control the visibility of the variables, types and functions they contain.
- A folder may only contain a single package, but the package does not need to be named after the directory.
- Names can only be used once inside the same package.
- Runnable programs have a main function defined in a main package.
- We can export fields outside their package by capitalising the first letter of their name. 

[Playground](https://go.dev/play/p/m0U35Q6fWMJ)
```go
type host struct {
	name string
}

func (h host) sayHi() string {
	return fmt.Sprintf("Hello, friends! I'm %s.", h.name)
}

func (h host) name() {
	return h.name
}
```

## Interfaces
- Interfaces are collections of method signatures. 
- They are automatically implemented by the compiler on types which satisfy the entire collection of methods. 
- They are the primary way of implementing polymorphism in Go.
- Interfaces are often exported, while the structs remain visible only inside the package. 

[Playground](https://go.dev/play/p/GHiwVPPVdPF)

```go
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
```

## Goroutines
- Goroutines are known as lightweight threads. They are used to run functions concurrently inside our Go programs.
- We instruct the Go runtime to run a function in a new goroutine by using the `go` keyword.
- Starting a goroutine is non-blocking by design, otherwise we'd be running things sequentially.
- The program runs in its own goroutine, known as the main goroutine. 
- The main goroutine has a parent child relationship with the goroutines it starts up.

```go
func main() {
	hosts := []Host{
		newHost{name: "Adelina"},
		experiencedHost{name: "Liam"},
	}
	for _, h := range hosts {
		go fmt.Println(h.SayHi())
	}
	fmt.Println("Goodbye, friends!")
}
```

## Channels
- It's discouraged to pass information between goroutines using shared memory variables.
- Channels are pipes which allow passing information in a threadsafe way.
- The type of variable that the channel supports is part of its initialisation.
- The send operation writes information through a channel, while the receive operation reads information from the channel.
- Sends and receives on a channel are blocking operations. They can be used for the synchronisation of goroutines.
- Messages are only read once.
- Once operations are completed, channels can be closed to signal to others that no more values will be sent through it.

[Playground](https://go.dev/play/p/S7x-Om6Qu4Z)

```go
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
    ch := make(chan string)
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
```

## Buffered channels
- By default, channels are unbuffered. 
- They require both the sender and receiver to be available for the operation to be completed. These operations are synchronous.
- If one side is available without the other, then it will be blocked until the corresponding opposite operation is possible.
- Channels can be buffered with a pre-determined capacity to hold senders' values until receivers arrive.
- If there is space in the channel's queue, then the operation completes immediately.

[Playground](https://go.dev/play/p/ZirmrNjv1zo)
```go
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
```

## Unit testing
- Go's testing package allows us to write tests, verifications and benchmarks.
- Coming from other languages, it might seem that Go's standard testing package is barebones.
- We can supplement it with other third-party libraries, but it's good to start with understanding how to write tests first.
- Testing concurrent code cannot prove the absence of bugs, but it can give us a statistical confidence of our code's behaviour under certain conditions.

```go
func TestSayHi(t *testing.T) {
	testCases := map[string]struct {
		name string
		host Host
		want string
	}{
		"new host": {
			host: newHost{name: "New Host"},
			want: "Hello, friends! I'm New Host and I'm new here.",
		},
		"experienced host": {
			host: experiencedHost{name: "Exp Host"},
			want: "Hello again, friends! I'm Exp Host.",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ch := make(chan string)
			go tc.host.SayHi(ch)
			greeting := <-ch
			if greeting != tc.want {
				t.Errorf("got: %s, want %s", greeting, tc.want)
			}
		})

	}
}
```

## The net/http package
- Go's HTTP package is easy to use and one of the reasons that Go is so widely used in web application development.
- Handlers respond to HTTP requests. Functions serving as handlers take in two parameters: `http.ResponseWriter` and `http.Request`.
- Handler functions are registered to a particular HTTP route using `http.HandleFunc`.
- The `net/http` package is responsible for passing the headers and requests to our custom registered handler functions.

```go
package main

import (
	"fmt"
	"net/http"
)

// welcome is a handler function as it satisfies the Handler signature.
func welcomeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello, friends! Welcome to the Microsoft Tech Days with Liam & Adelina!\n")
}

func main() {
	// Register the welcomeHandler function to serve the root endpoint.
	http.HandleFunc("/", welcomeHandler)

	// Start the default router on port 4321 and block the main goroutine from terminating.
	fmt.Println("Listening on 4321...")
	http.ListenAndServe(":4321", nil)
}
```