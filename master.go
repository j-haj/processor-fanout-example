package fanout

import (
	"fmt"
	"math/rand"
)

// generatePrevies creates an array of prefix strings of the format A1, A2, ..
// B1, B2, ...
func generatePrefixes() []string {
	prefixes := []string{}
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, l := range letters {
		for i := 1; i < 10; i++ {
			prefixes = append(prefixes, fmt.Sprintf("%s%d", l, i))
		}
	}
	return prefixes
}

type Master struct {
	workers  map[string]Worker
	prefixes []string
}

// NewMaster creates a new master struct with generated prefixes.
func NewMaster() Master {
	p := generatePrefixes()
	m := &Master{}
	m.workers = make(map[string]Worker)
	m.prefixes = p
	return m
}

// Run runs the master main loop
func (m *Master) Run(n int) {
	for i := 0; i < n; i++ {
		go func() {
			prefix := m.prefixes[rand.Intn(len(prefixes))]
			// If we have encountered this prefix before, use the
			// previously made channel, otherwise create a new
			// channel and store the prefix
			var worker Worker
			if worker, ok := m.workers[prefix]; !ok {
				worker = NewWorker(2, fmt.Sprintf("%s%d", prefix, rand.Intn(10)))
				m.workers[prefix] = worker
			}
			worker.Process(randomString())
		}()
	}
}

func randomString() string {
	letters := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"
	buffer := make([]byte, 20)
	for i := 0; i < 20; i++ {
		buffer[i] = letters[rand.Intn(20)]
	}
	return string(buffer)
}
