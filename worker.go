package fanout

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	Buffer     []string
	Id         string
	Ch         chan string
	maxWorkers int
	nActive    int
	Exit       <-chan bool
	Done       chan<- bool
}

// NewWorker creates a new worker with n processors
func NewWorker(n int, id string, exit <-chan bool) Worker {
	w := &Worker{id: id, ch: make(chan string), maxWorkers: n,
		exit: make(chan bool), done: make(chan bool)}
	go func() {
		for {
			select {
			case c := <-ch:
				go func() {
					w.processInput(c)
				}()
			case <-w.exit:
				// Perform any cleanup here
				w.done <- true
			}
		}

	}()
	return w
}

func (w *Worker) processInput(input string) {
	t := rand.Intn(500) * time.Millisecond
	time.After(t)
	fmt.Printf("Worker %s: %s after %v\n", w.Id, input, t)

}
