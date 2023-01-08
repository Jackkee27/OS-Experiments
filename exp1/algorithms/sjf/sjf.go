package sjf

import (
	"fmt"
	"sort"

	"exp1/models"
	"exp1/pkg/queue"
)

// Run SJF algorithm, including outputs
func Run(processes models.Processes, readyQueue *queue.Queue) {
	t := 0 // 程序当前运行时间

	readyMap, keys := MakeReadyMap(processes) // map[arrivalTime]processesSJF

	fmt.Printf("\npid\tst\tet\trq\n") // print 1st row
	for readyQueue.Length() > 0 {
		// 找到当前就绪队列的开头(pID)对应的进程 p
		pID := readyQueue.Peek().(int)
		pIdx := findProcess(processes, pID)
		if pIdx == -1 {
			fmt.Println("findProcess(processes, pID) failed: not found")
			return
		}
		p := processes[pIdx]
		p.StartTime = t
		p.EndTime = p.StartTime + p.ServiceTime
		// update ready queue
		readyQueue.Remove()
		// 选择一个优先级最高的进程
		// 将 (t, t+p.ServiceTime] 到达的所有进程按照 serviceTime 升序排序加入 readyQueue 队列
		sjfs := models.ProcessesSJF{}
		// readyQueue中原来的也参与排序
		for readyQueue.Length() > 0 {
			sjfs = append(sjfs, processes[findProcess(processes, readyQueue.Remove().(int))])
		}
		for _, key := range keys {
			if key <= t {
				continue
			} else if key > p.EndTime {
				break
			}
			q := readyMap[key]
			for q.Length() > 0 {
				sjfs = append(sjfs, q.Remove().(*models.Process))
			}
		}
		// sorted by serviceTime first
		sort.Sort(sjfs)
		for _, sjf := range sjfs {
			readyQueue.Add(sjf.Pid)
		}

		Print(p, readyQueue)

		t += p.ServiceTime
	}

	// output
	processes.Print()
	processes.PrintDetail(models.StrSJF)
}

// Print the pid, start time & end time of process
func Print(p *models.Process, readyQueue *queue.Queue) {
	fmt.Printf("%d\t%d\t%d\t", p.Pid, p.StartTime, p.EndTime)
	// readyQueue
	readyQueue.Print()
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

// findProcess returns index of the goal process, and returns -1 if not exists
func findProcess(processes models.Processes, pid int) int {
	for idx, p := range processes {
		if p.Pid == pid {
			return idx
		}
	}
	return -1
}
