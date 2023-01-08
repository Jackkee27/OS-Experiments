package algorithms

import (
	"exps_1_updated/models"
	"exps_1_updated/pkg/queue"
	"fmt"
)

// ReadyQueue 就绪队列
type ReadyQueue struct {
	q *queue.Queue
}

// Remove 出队并返回进程
func (rq ReadyQueue) Remove() *models.Process {
	return rq.q.Remove().(*models.Process)
}

// Add 添加进程
func (rq ReadyQueue) Add(p *models.Process) {
	rq.q.Add(p)
}

// Length 获取长度
func (rq ReadyQueue) Length() int {
	return rq.q.Length()
}

// Get 获取位置为i的进程
func (rq ReadyQueue) Get(i int) *models.Process {
	return rq.q.Get(i).(*models.Process)
}

// Peek 获取队首进程
func (rq ReadyQueue) Peek() *models.Process {
	return rq.q.Peek().(*models.Process)
}

// Print 打印队列: [...->...]
func (rq ReadyQueue) Print() {
	fmt.Print("[")
	for i := 0; i < rq.Length(); i++ {
		fmt.Printf("%d", rq.Get(i).Pid)
		if i < rq.Length()-1 {
			fmt.Print("->")
		}
	}
	fmt.Print("]\n")
}
