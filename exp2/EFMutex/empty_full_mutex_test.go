package EFMutex

import (
	"testing"
	"time"
)

func TestEFMutex(t *testing.T) {
	p := &Process{}
	p.Init(16)
	p.PrintTitle()

	go p.Consume()
	time.Sleep(3 * time.Second)
	go p.Produce()

	time.Sleep(5 * time.Second)
}
