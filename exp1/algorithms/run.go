package algorithms

import (
	"errors"
	"fmt"
	"sort"

	"exp1/algorithms/fcfs"
	"exp1/algorithms/rr"
	"exp1/algorithms/sjf"
	"exp1/models"
	"exp1/pkg/queue"
)

var (
	ErrNOutOfRange = errors.New("n: out of range")
	ErrPOutOfRange = errors.New("p: out of range")
	ErrInvalidMode = errors.New("invalid mode")
)

// Run 运行程序
func Run(processes models.Processes, mode int) error {
	// 输入验证
	if len(processes) <= 0 || len(processes) > models.MaxProcessNumber {
		return ErrNOutOfRange
	}

	// sorted by arrivalTime first
	sort.Sort(processes)
	// init readyQueue 初始化预备队列
	readyQueue := queue.New()
	// t == 0, 把 arrivalTime == 0 时刻的入队
	for _, p := range processes {
		if p.ArrivalTime == 0 {
			readyQueue.Add(p.Pid)
		} else if p.ArrivalTime > 0 {
			break
		}
	}
	switch mode {
	case models.ModeFCFS:
		// Run FCFS
		fcfs.Run(processes, readyQueue)
	case models.ModeSJF:
		// Run SJF
		sjf.Run(processes, readyQueue)
	case models.ModeRR:
		// input q
		var q int
		fmt.Print("input q: ")
		_, _ = fmt.Scan(&q)
		if q <= 0 {
			return ErrPOutOfRange
		}
		// Run RR
		rr.Run(processes, readyQueue, q)
	default:
		return ErrInvalidMode
	}

	return nil
}
