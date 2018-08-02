package fanout

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	buffer     []string
	processors [](chan string)
	id         string
}

// NewWorker creates a new worker with n processors
func NewWorker(n int, id string) Worker {
	// TODO: Need to come up with a  bette way to handle the processors
	w := &Worker{id: id}
	return w
}

// newProcessor creates a new processor from an input string channel and returns
// a bool channel that is used to communicate when  it is ready
func newProcessor(ch <-chan string, id string) chan bool {
	response := make(chan bool)
	select {
	case s := <-ch:
		t := rand.Intn(1000) * time.Millisecond
		time.After(t)
		fmt.Printf("Worker %s: %s after %v\n", id, s, t)
		response <- true
	}
	return response
}
