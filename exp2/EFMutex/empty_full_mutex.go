package EFMutex

import (
	"context"
	"exp2/models"
	"github.com/spf13/cast"
	"golang.org/x/sync/semaphore"
)

// Process simulates a process with semaphores empty, full & mutex
type Process struct {
	models.ProcessBasic

	mutex *semaphore.Weighted
	empty *semaphore.Weighted
	full  *semaphore.Weighted
	ctx   context.Context
}

// Init init process
func (p *Process) Init(maxBufferSize int) {
	p.ctx = context.TODO()     // create a context
	p.InitBasic(maxBufferSize) // init models.ProcessBasic
	// init semaphore empty, mutex & full
	p.empty = semaphore.NewWeighted(cast.ToInt64(maxBufferSize))
	p.full = semaphore.NewWeighted(cast.ToInt64(maxBufferSize))
	p.mutex = semaphore.NewWeighted(1)
	// set p.full.cur = 0
	p.full.TryAcquire(cast.ToInt64(maxBufferSize))
}

// Produce producer thread
func (p *Process) Produce() {
	for {
		for p.Cnt >= p.MaxBufferSize {
		}
		// wait
		_ = p.empty.Acquire(p.ctx, 1)
		_ = p.mutex.Acquire(p.ctx, 1)

		p.Write()

		// signal
		p.mutex.Release(1)
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
		_ = p.mutex.Acquire(p.ctx, 1)

		p.Read()

		// signal
		p.mutex.Release(1)
		p.empty.Release(1)
	}

}
