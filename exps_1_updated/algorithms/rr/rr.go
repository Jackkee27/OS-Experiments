package rr

import (
	"exps_1_updated/algorithms"
	"exps_1_updated/models"
	"fmt"
)

// Runner 执行器
type Runner struct {
	SizeOfTimeSlice int                   // 时间片大小
	Processes       models.Processes      // 进程序列
	TimeSlices      []*timeSlice          // 进程时间片序列
	ReadyQueue      algorithms.ReadyQueue // 就绪队列
}

// Init 初始化
func (rr *Runner) Init(processes models.Processes, sizeOfTimeSlice int) (err error) {
	rr.SizeOfTimeSlice = sizeOfTimeSlice
	rr.Processes = processes
	// InitTimeSlice 初始化时间片序列
	rr.TimeSlices = initTimeSlice(processes)
	// InitReadyQueue 初始化预备队列
	rr.ReadyQueue = algorithms.InitReadyQueue(processes, models.ModeRR)

	return
}

// Run 运行调度算法
func (rr *Runner) Run() (err error) {
	t := 0       // 程序当前运行时间
	doneNum := 0 // 已经完成的进程数量
	lastT := t   // 上一次执行运行的时间

	processes := rr.Processes
	readyMap, keys := algorithms.MakeReadyMap(processes) // map[arrivalTime]processes

	fmt.Printf("\npid\tlast\tcurr\trq\n")
	for doneNum < len(processes) {
		if rr.ReadyQueue.Length() <= 0 { // CPU FREE NOW
			// 把下一个到达的进程入队
			for _, key := range keys {
				if key > t {
					que := readyMap[key]
					// 把que中的所有进程加入到rr运行器的就绪队列中
					for que.Length() > 0 {
						rr.ReadyQueue.Add(que.Remove())
					}
					break
				}
			}
			t = rr.ReadyQueue.Peek().ArrivalTime
			lastT = t
		}
		// 找到当前就绪队首进程 p
		p := rr.ReadyQueue.Peek()
		// 找到 p 对应的 时间片 ts
		tIdx, err := findTimeSliceByPid(rr.TimeSlices, p.Pid)
		if err != nil {
			return err
		}
		ts := rr.TimeSlices[tIdx]

		// 设置p.StartTime
		if !ts.isVisited {
			ts.isVisited = true
			ts.StartTime = t
		}

		ts.Left -= rr.SizeOfTimeSlice
		t = t + rr.SizeOfTimeSlice
		if ts.Left <= 0 { // 如果当前进程已经全部完成
			doneNum++
			t += ts.Left // 还原多加的时间
			ts.EndTime = t
		}

		// arriveTime 在(lastT, t]之间到达的进程入列
		for _, key := range keys {
			if key <= lastT {
				continue
			} else if key > t {
				break
			}
			que := readyMap[key]
			for que.Length() > 0 {
				rr.ReadyQueue.Add(que.Remove())
			}
		}
		// 如果当前进程还未完成，再次入列
		if ts.Left > 0 {
			rr.ReadyQueue.Add(p)
		}

		rr.ReadyQueue.Remove()                   // 当前进程出队
		PrintTS(ts.Pid, lastT, t, rr.ReadyQueue) // 打印刚出队进程时间片信息
		lastT = t
	}

	rr.Processes = processes

	// output
	rr.Processes.PrintRR()
	rr.Processes.PrintDetail(models.StrRR)

	return
}
