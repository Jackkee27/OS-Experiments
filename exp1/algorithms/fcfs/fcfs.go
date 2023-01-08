package fcfs

import (
	"exp1/models"
	"exp1/pkg/queue"
	"fmt"
	"sort"
)

func Run(processes models.Processes, readyQueue *queue.Queue) {
	t := 0 // 程序当前运行时间

	readyMap, keys := MakeReadyMap(processes) // map[arrivalTime]processes

	fmt.Printf("\npid\tst\tet\trq\n")
	for _, p := range processes {
		p.StartTime = t
		p.EndTime = p.StartTime + p.ServiceTime
		// update ready queue
		readyQueue.Remove()
		// 把 (t, t+p.ServiceTime] 到达的节点入队
		for _, key := range keys {
			if key <= t {
				continue
			} else if key > p.EndTime {
				break
			}
			q := readyMap[key]
			for q.Length() > 0 {
				readyQueue.Add(q.Remove().(*models.Process).Pid)
			}
		}
		Print(p, readyQueue)
		t += p.ServiceTime
	}

	// output
	processes.Print()
	processes.PrintDetail(models.StrFCFS)
}

// MakeReadyMap returns a map(type:map[arrivalTime]processesSJF),
// and a sorted slice 'keys' for arrivalTime
func MakeReadyMap(processes models.Processes) (readyMap map[int]*queue.Queue, keys []int) {
	readyMap = map[int]*queue.Queue{}
	for _, p := range processes {
		if readyMap[p.ArrivalTime] == nil {
			readyMap[p.ArrivalTime] = queue.New()
		}
		readyMap[p.ArrivalTime].Add(p)
		keys = append(keys, p.ArrivalTime)
	}
	sort.Ints(keys)
	return
}

// Print the pid, start time & end time of process
func Print(p *models.Process, readyQueue *queue.Queue) {
	fmt.Printf("%d\t%d\t%d\t", p.Pid, p.StartTime, p.EndTime)
	// readyQueue
	readyQueue.Print()
}
