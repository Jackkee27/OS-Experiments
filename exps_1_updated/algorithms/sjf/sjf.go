package sjf

import (
	"exps_1_updated/algorithms"
	"exps_1_updated/models"
	"fmt"
	"sort"
)

type Runner struct {
	Processes  models.Processes      // 进程序列
	ReadyQueue algorithms.ReadyQueue // 就绪队列
}

func (sjf *Runner) Init(processes models.Processes) error {
	sjf.Processes = processes
	// InitReadyQueue 初始化预备队列
	sjf.ReadyQueue = algorithms.InitReadyQueue(processes, models.ModeSJF)

	return nil
}

// SortByServiceTime 将ReadyQueue中的进程，按照服务时间升序排序
func (sjf *Runner) SortByServiceTime() {
	tmpSJF := models.ProcessesSJF{}
	for sjf.ReadyQueue.Length() > 0 {
		tmpSJF = append(tmpSJF, sjf.ReadyQueue.Remove())
	}
	sort.Sort(tmpSJF)
	// 重新入队
	for _, s := range tmpSJF {
		sjf.ReadyQueue.Add(s)
	}
}

func (sjf *Runner) Run() (err error) {
	doneNum := 0 // 已完成的进程数量

	processes := sjf.Processes
	readyMap, _ := algorithms.MakeReadyMap(processes) // map[arrivalTime]processes

	fmt.Printf("\npid\tst\tet\trq\n") // 输出标题
	isFree := true
	lastT := 0
	var currentProcess *models.Process
	for t := 0; doneNum < len(processes); t++ {
		// 是否有该时刻到达的程序
		if _, ok := readyMap[t]; ok { // Y
			// 加入就绪队列
			sjf.ReadyQueue.Add(readyMap[t].Remove())
			// 按照服务时间排序
			sjf.SortByServiceTime()
		}

		if !isFree { // CPU BUSY
			if lastT+currentProcess.ServiceTime == t { // 当前进程完成
				doneNum++
				currentProcess.EndTime = t
				isFree = true
				algorithms.Print(currentProcess, sjf.ReadyQueue)
			} else {
				continue
			}
		}
		if isFree { // CPU FREE
			if sjf.ReadyQueue.Length() <= 0 { // 就绪队列为空
				continue
			} else { // 就绪队列 > 0
				// 取队首进程
				currentProcess = sjf.ReadyQueue.Remove()
				currentProcess.StartTime = t

				lastT = t
				isFree = false
			}
		}
	}

	sort.Sort(processes) // 按照到达时间排序
	sjf.Processes = processes

	// output
	sjf.Processes.Print()
	sjf.Processes.PrintDetail(models.StrSJF)

	return
}
