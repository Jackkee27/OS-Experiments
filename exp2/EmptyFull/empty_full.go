package EmptyFull

import (
	"context"
	"exp2/models"
	"github.com/spf13/cast"
	"golang.org/x/sync/semaphore"
)

// Process simulates a process with semaphores empty & full
type Process struct {
	models.ProcessBasic

	ctx   context.Context
	empty *semaphore.Weighted
	full  *semaphore.Weighted
}

// Init process
func (p *Process) Init(maxBufferSize int) {
	p.InitBasic(maxBufferSize) // init models.ProcessBasic

	p.ctx = context.TODO() // create a context
	// init semaphore empty & full
	p.empty = semaphore.NewWeighted(cast.ToInt64(maxBufferSize))
	p.full = semaphore.NewWeighted(cast.ToInt64(maxBufferSize))

	p.full.TryAcquire(cast.ToInt64(maxBufferSize))
}

// Produce producer thread
func (p *Process) Produce() {
	for {
		for p.Cnt >= p.MaxBufferSize {
		}
		// wait
		_ = p.empty.Acquire(p.ctx, 1)
		// produce
		p.Write()
		// signal
		p.full.Release(1)
	}
}

// Consume consumer thread
func (p *Process) Consume() {
	for {
		for p.Cnt <= 0 {
		}
		// wait
		_ = p.full.Acquire(p.ctx, 1)
		// consume
		p.Read()
		// signal
		p.empty.Release(1)
	}

}
