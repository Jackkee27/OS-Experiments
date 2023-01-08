package Mutex

import (
	"context"
	"exp2/models"
	"golang.org/x/sync/semaphore"
)

// Process simulates a process with semaphore mutex
type Process struct {
	models.ProcessBasic

	mutex *semaphore.Weighted
	ctx   context.Context
}

// Init process
func (p *Process) Init(maxBufferSize int) {
	p.InitBasic(maxBufferSize)

	p.ctx = context.TODO()
	p.mutex = semaphore.NewWeighted(1)
}

// Produce producer thread
func (p *Process) Produce() {
	for {
		for p.Cnt >= p.MaxBufferSize {
		}
		_ = p.mutex.Acquire(p.ctx, 1)

		p.Write()

		p.mutex.Release(1)
	}
}

// Consume consumer thread
func (p *Process) Consume() {
	for {
		for p.Cnt <= 0 {
		}
		_ = p.mutex.Acquire(p.ctx, 1)

		p.Read()

		p.mutex.Release(1)
	}
}
