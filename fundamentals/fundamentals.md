# The fundamental concepts within Go
The purpose of this section is to equip our listeners with all the knowledge they require for our session, but we will not teach them Go from scratch.

## Strong typing
- The compiler is your best friend. 
- At all points, we will know the type of our variables, and what behaviour they expose.
- Dynamic typing is slower and the compiler avoids runtime errors and fatal error cases due to undefined behaviours.
- The fmt package is part of the standard library and allows us to format and print strings.
- The Go toolchain builds and runs our programs.

```bash
$ go run fundamentals/strong-typing/main.go
```

## Functions
- Go functions are natively supported and can be passed as variables, return types and parameters for later invocation. 
- Anonymous functions are also allowed.
- Function composition is easy to do in Go. 
- Deferred functions are useful for guaranteed clean up tasks. 

```bash
$ go run fundamentals/functions/main.go
```

## Error handling
- Go functions can return multiple values.
- By convention, the error is the last returned value using the built-in error type. 
- The zero value of the error type is nil.
- Errors should be handled first, keeping code minimally indented.


```bash
$ go run fundamentals/error-handling/main.go
```

## Custom types
- Go is not an object oriented language, as it does not support type hierarchy. 
- Structs allow us to build custom types and behaviours.
- They are a collection of fields, which can be partially initialised.
- Custom types can also define methods by using a special receiver argument that is the implicit first argument of the method.

```bash
$ go run fundamentals/structs/main.go
```

## Visibility
- Go code is organised in packages, which control the visibility of the variables, types and functions they contain.
- A folder may only contain a single package, but the package does not need to be named after the directory.
- Names can only be used once inside the same package.
- Runnable programs have a main function defined in a main package.
- We can export fields outside their package by capitalising the first letter of their name. 

```bash
$ go run fundamentals/visibility/main.go
```

## Interfaces
- Interfaces are collections of method signatures. 
- They are automatically implemented by the compiler on types which satisfy the entire collection of methods. 
- They are the primary way of implementing polymorphism in Go.
- Interfaces are often exported, while the structs remain visible only inside the package. 

```bash
$ go run fundamentals/visibility/main.go
```

## Goroutines
- Goroutines are known as lightweight threads. They are used to run functions concurrently inside our Go programs.
- We instruct the Go runtime to run a function in a new goroutine by using the `go` keyword.
- Starting a goroutine is non-blocking by design, otherwise we'd be running things sequentially.
- The program runs in its own goroutine, known as the main goroutine. 
- The main goroutine has a parent child relationship with the goroutines it starts up.

```bash
$ go run fundamentals/goroutines/main.go
```

## Channels
- It's discouraged to pass information between goroutines using shared memory variables.
- Channels are pipes which allow passing information in a threadsafe way.
- The type of variable that the channel supports is part of its initialisation.
- The send operation writes information through a channel, while the receive operation reads information from the channel.
- Sends and receives on a channel are blocking operations. They can be used for the synchronisation of goroutines.
- Messages are only read once.
- Once operations are completed, channels can be closed to signal to others that no more values will be sent through it.

```bash
$ go run fundamentals/channels/main.go
```

## Buffered channels
- By default, channels are unbuffered. 
- They require both the sender and receiver to be available for the operation to be completed. These operations are synchronous.
- If one side is available without the other, then it will be blocked until the corresponding opposite operation is possible.
- Channels can be buffered with a pre-determined capacity to hold senders' values until receivers arrive.
- If there is space in the channel's queue, then the operation completes immediately.

```bash
$ go run fundamentals/buffered-channels/main.go
```

## Unit testing
- Go's testing package allows us to write tests, verifications and benchmarks.
- Coming from other languages, it might seem that Go's standard testing package is barebones.
- We can supplement it with other third-party libraries, but it's good to start with understanding how to write tests first.
- Testing concurrent code cannot prove the absence of bugs, but it can give us a statistical confidence of our code's behaviour under certain conditions.

```bash
$ go test -run=TestSayHi ./fundamentals/unit-test
```

## The net/http package
- Go's HTTP package is easy to use and one of the reasons that Go is so widely used in web application development.
- Handlers respond to HTTP requests. Functions serving as handlers take in two parameters: `http.ResponseWriter` and `http.Request`.
- Handler functions are registered to a particular HTTP route using `http.HandleFunc`.
- The `net/http` package is responsible for passing the headers and requests to our custom registered handler functions.

```bash
$ go run fundamentals/server/main.go
```