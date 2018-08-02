package main

import (
	"fmt"

	"github.com/j-haj/worker-fanout-example/fanout"
)

func main() {
	master := NewMaster()
	master.Run(1000)
}
