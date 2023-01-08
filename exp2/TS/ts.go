package ts

import (
	"exp2/models"
)

// TS simulates Test-and-Set instruction
func TS(lock *bool) bool {
	old := *lock
	*lock = true
	return old
}

// Process simulate a process with TS
type Process struct {
	models.ProcessBasic

	lock bool
}

// Init process
func (p *Process) Init(maxBufferSize int) {
	p.InitBasic(maxBufferSize)

	p.lock = false
}

// Produce producer thread
func (p *Process) Produce() {
	for {
		for TS(&p.lock) == false || p.Cnt >= p.MaxBufferSize {
		}

		p.Write()
		p.lock = false
	}
}

// Consume consumer thread
func (p *Process) Consume() {
	for {
		for TS(&p.lock) == false || p.Cnt <= 0 {
		}
		p.Read()
		p.lock = false
	}
}
