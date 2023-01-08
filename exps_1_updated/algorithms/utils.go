package algorithms

import (
	"exps_1_updated/models"
	"exps_1_updated/pkg/queue"
	"fmt"
	"sort"
)

// InitReadyQueue 初始化就绪队列
func InitReadyQueue(processes models.Processes, mode int) ReadyQueue {
	q := queue.New()
	if mode == models.ModeRR {
		for _, p := range processes {
			if p.ArrivalTime == 0 {
				q.Add(p)
			} else if p.ArrivalTime > 0 {
				break
			}
		}
	}

	return ReadyQueue{q: q}
}

// MakeReadyMap returns a map(type:map[arrivalTime]processesSJF),
// and a sorted slice 'keys' for arrivalTime
func MakeReadyMap(processes models.Processes) (readyMap map[int]ReadyQueue, keys []int) {
	readyMap = map[int]ReadyQueue{}
	for _, p := range processes {
		if readyMap[p.ArrivalTime].q == nil {
			readyMap[p.ArrivalTime] = ReadyQueue{q: queue.New()}
		}
		readyMap[p.ArrivalTime].Add(p)
		keys = append(keys, p.ArrivalTime)
	}
	sort.Ints(keys)
	return
}

// Print the pid, start time & end time of process
func Print(p *models.Process, rq ReadyQueue) {
	fmt.Printf("%d\t%d\t%d\t", p.Pid, p.StartTime, p.EndTime)
	// readyQueue
	rq.Print()
}
