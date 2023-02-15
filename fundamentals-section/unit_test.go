package unit_test

import (
	"fmt"
	"testing"
)

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
