package rr

import (
	"exp1/models"
	"exp1/pkg/queue"
	"fmt"
	"sort"
)

// timeSlice 时间片
type timeSlice struct {
	Left      int  // 剩下的时间
	IsVisited bool // 是否已经访问过
	*models.Process
}

func Run(processes models.Processes, readyQueue *queue.Queue, q int) {
	t := 0     // 程序当前运行时间
	lastT := t // 上一次执行运行的时间

	tSlice := initTimeSlice(processes)        // 初始化时间片序列
	readyMap, keys := MakeReadyMap(processes) // map[arrivalTime]processes

	fmt.Printf("\npid\tlast\tcurr\trq\n")
	for readyQueue.Length() > 0 {
		// 找到当前就绪队列的开头(tID)对应的进程 ts
		tID := readyQueue.Peek().(int)
		tIdx := findTimeSlice(tSlice, tID)
		if tIdx == -1 {
			fmt.Println("findProcess(processes, pID) failed: not found")
			return
		}
		ts := tSlice[tIdx]

		if ts.IsVisited == false { // 还未访问过
			ts.StartTime = t
			ts.IsVisited = true
		}

		ts.Left -= q
		if ts.Left <= 0 {
			t = t + q + ts.Left
			// arriveTime 在(lastT, t]之间的进程入列
			for _, key := range keys {
				if key <= lastT {
					continue
				} else if key > t {
					break
				}
				que := readyMap[key]
				for que.Length() > 0 {
					readyQueue.Add(que.Remove().(*models.Process).Pid)
				}
			}
			ts.EndTime = t
		} else { // ts.Left > 0
			t = t + q
			// arriveTime 在(lastT, t]之间的进程入列
			for _, key := range keys {
				if key <= lastT {
					continue
				} else if key > t {
					break
				}
				que := readyMap[key]
				for que.Length() > 0 {
					readyQueue.Add(que.Remove().(*models.Process).Pid)
				}
			}

			// 再次入列
			readyQueue.Add(ts.Pid)
		}

		readyQueue.Remove()
		Print(ts.Pid, lastT, t, readyQueue)
		lastT = t
	}

	// output
	processes.PrintRR()
	processes.PrintDetail(models.StrRR)
}

// Print pid, lastT, currentT & readyQueue
func Print(pid, lastT, currT int, readyQueue *queue.Queue) {
	fmt.Printf("%d\t%d\t%d\t", pid, lastT, currT)
	// readyQueue
	readyQueue.Print()
}

// findTimeSlice returns index of the goal process, and returns -1 if not exists
func findTimeSlice(tSlice []*timeSlice, tid int) int {
	for idx, t := range tSlice {
		if t.Pid == tid {
			return idx
		}
	}
	return -1
}

// initTimeSlice initialize tSlice
func initTimeSlice(processes models.Processes) (tSlice []*timeSlice) {
	for _, p := range processes {
		ts := &timeSlice{
			Left:      p.ServiceTime,
			IsVisited: false,
			Process:   p,
		}
		tSlice = append(tSlice, ts)
	}
	return
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
