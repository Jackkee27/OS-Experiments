package fcfs

import (
	"exps_1_updated/algorithms"
	"exps_1_updated/models"
	"fmt"
	"sort"
)

// Runner 运行器
type Runner struct {
	Processes  models.Processes      // 进程序列
	ReadyQueue algorithms.ReadyQueue // 就绪队列
}

// Init 初始化
func (fcfs *Runner) Init(processes models.Processes) (err error) {
	fcfs.Processes = processes
	// InitReadyQueue 初始化预备队列
	fcfs.ReadyQueue = algorithms.InitReadyQueue(processes, models.ModeFCFS)

	return
}

// Run 运行调度程序
func (fcfs *Runner) Run() (err error) {
	doneNum := 0 // 已完成的进程数量
	processes := fcfs.Processes
	readyMap, _ := algorithms.MakeReadyMap(processes) // map[arrivalTime]processes

	fmt.Printf("\npid\tst\tet\trq\n")
	isFree := true // CPU状态
	lastT := 0
	var currentProcess *models.Process
	for t := 0; doneNum < len(processes); t++ {
		// 是否有该时刻到达的程序
		if _, ok := readyMap[t]; ok { // Y
			// 加入就绪队列
			fcfs.ReadyQueue.Add(readyMap[t].Remove())
		}

		// 处理当前进程，并判断是否完成
		if !isFree { // CPU BUSY
			if lastT+currentProcess.ServiceTime == t { // 当前进程完成
				doneNum++
				currentProcess.EndTime = t
				isFree = true
				algorithms.Print(currentProcess, fcfs.ReadyQueue)
			} else { // 未完成，t++
				continue
			}
		}
		// CPU空闲，查看是否有下一个进程
		if isFree { // CPU FREE
			if fcfs.ReadyQueue.Length() <= 0 { // 就绪队列为空
				continue
			} else { // 就绪队列 > 0
				// 取队首进程
				currentProcess = fcfs.ReadyQueue.Remove()
				currentProcess.StartTime = t

				lastT = t
				isFree = false
			}
		}
	}

	sort.Sort(processes) // 按照到达时间排序
	fcfs.Processes = processes

	// output
	fcfs.Processes.Print()
	fcfs.Processes.PrintDetail(models.StrFCFS)

	return
}
