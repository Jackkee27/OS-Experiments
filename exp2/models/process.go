package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ProcessBasic basic process info
type ProcessBasic struct {
	In            int
	Out           int
	Cnt           int
	MaxBufferSize int
	Buf           []int
	cntMutex      sync.Mutex // counter lock
}

func (p *ProcessBasic) InitBasic(maxBufferSize int) {
	p.MaxBufferSize = maxBufferSize    // set max buffer size
	p.Buf = make([]int, maxBufferSize) // set buffer
	p.In, p.Out, p.Cnt = 0, 0, 0       // init in, out, cnt
}

func (p *ProcessBasic) PrintTitle() {
	fmt.Println("write/read\tcontent\tindex\tproduction num")
}

// Write writes a content
func (p *ProcessBasic) Write() {
	time.Sleep(100)
	p.Buf[p.In] = rand.Int() % 100
	p.cntMutex.Lock() // 临界区begin
	p.Cnt++
	fmt.Printf("write\t\t%d\t\t%d\t\t%d\n", p.Buf[p.In], p.In, p.Cnt)
	p.cntMutex.Unlock() // 临界区end
	p.In = (p.In + 1) % p.MaxBufferSize
}

// Read reads a content
func (p *ProcessBasic) Read() {
	time.Sleep(100)
	p.cntMutex.Lock() // 临界区begin
	p.Cnt--
	fmt.Printf("read \t\t%d\t\t%d\t\t%d\n", p.Buf[p.Out], p.Out, p.Cnt)
	p.cntMutex.Unlock() // 临界区end
	p.Out = (p.Out + 1) % p.MaxBufferSize
}
