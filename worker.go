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
}

// NewWorker creates a new worker with n processors
func NewWorker(n int, id string) Worker {
	w := &Worker{id: id, ch: make(chan string), maxWorkers: n}
	go func() {
		for {
			select {
			case c := <-ch:
				go func() {
					w.processInput(c)
				}()
			}
		}

	}()
	return w
}

func (w *Worker) Process(input string) {
	w.Buffer = append(w.Buffer, input)
}

func (w *Worker) processInput(input string) {
	t := rand.Intn(500) * time.Millisecond
	time.After(t)
	fmt.Printf("Worker %s: %s after %v\n", w.Id, input, t)

}
