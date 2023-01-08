package EmptyFull

import (
	"testing"
	"time"
)

func TestEmptyFull(t *testing.T) {
	p := &Process{}
	p.Init(16)
	p.PrintTitle()
	go p.Produce()
	go p.Consume()

	time.Sleep(1 * time.Second)
}
