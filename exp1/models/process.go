package models

import (
	"fmt"
	"github.com/spf13/cast"
)

// Process 进程
type Process struct {
	ArrivalTime int // 到达时间
	Pid         int // 进程ID
	StartTime   int // 开始执行时间
	EndTime     int // 执行时间
	ServiceTime int // 服务时间, CPU burst
}

// GetTurnaroundTime 获取周转时间
func (p Process) GetTurnaroundTime() int {
	return p.EndTime - p.ArrivalTime
}

// GetWeightedTurnaroundTime 获取加权周转时间
func (p Process) GetWeightedTurnaroundTime() float64 {
	return cast.ToFloat64(p.GetTurnaroundTime()) / cast.ToFloat64(p.ServiceTime)
}

// GetWaitTime 获取等待时间（不包括RR）
func (p Process) GetWaitTime() int {
	return p.StartTime - p.ArrivalTime
}

// GetWaitTimeRR 获取等待时间（RR）
func (p Process) GetWaitTimeRR() int {
	return p.GetTurnaroundTime() - p.ServiceTime
}

// Processes 进程切片
type Processes []*Process

func (ps Processes) Len() int {
	return len(ps)
}

func (ps Processes) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// Less 排列优先级：到达时间 > 服务时间
func (ps Processes) Less(i, j int) bool {
	if ps[i].ArrivalTime < ps[j].ArrivalTime {
		return true
	} else if ps[i].ArrivalTime == ps[j].ArrivalTime {
		if ps[i].ServiceTime < ps[j].ServiceTime {
			return true
		}
		return false
	}
	return false
}

// Print 格式化输出进程（不包括RR）
func (ps Processes) Print() {
	fmt.Printf("pid\ttt\twtt\twt\n")
	for _, p := range ps {
		fmt.Printf("%d\t%d\t%.2f\t%d\n",
			p.Pid,
			p.GetTurnaroundTime(),
			p.GetWeightedTurnaroundTime(),
			p.GetWaitTime())
	}
}

// PrintRR 格式化输出（RR）
func (ps Processes) PrintRR() {
	fmt.Printf("pid\ttt\twtt\twt\n")
	for _, p := range ps {
		fmt.Printf("%d\t%d\t%.2f\t%d\n",
			p.Pid,
			p.GetTurnaroundTime(),
			p.GetWeightedTurnaroundTime(),
			p.GetWaitTimeRR())
	}
}

// GetAverageTurnaroundTime 获取平均周转时间
func (ps Processes) GetAverageTurnaroundTime() float64 {
	tmpSum := 0
	for _, p := range ps {
		tmpSum += p.GetTurnaroundTime()
	}
	return cast.ToFloat64(tmpSum) / cast.ToFloat64(len(ps))
}

// GetAverageWeightedTurnaroundTime 获取平均加权周转时间
func (ps Processes) GetAverageWeightedTurnaroundTime() float64 {
	tmpSum := 0.0
	for _, p := range ps {
		tmpSum += p.GetWeightedTurnaroundTime()
	}
	return tmpSum / cast.ToFloat64(len(ps))
}

// GetAverageWaitTime 获取平均等待时间
func (ps Processes) GetAverageWaitTime() float64 {
	tmpSum := 0
	for _, p := range ps {
		tmpSum += p.GetWaitTime()
	}
	return cast.ToFloat64(tmpSum) / cast.ToFloat64(len(ps))
}

// GetAverageWaitTimeRR 获取平均等待时间（RR）
func (ps Processes) GetAverageWaitTimeRR() float64 {
	tmpSum := 0
	for _, p := range ps {
		tmpSum += p.GetWaitTimeRR()
	}
	return cast.ToFloat64(tmpSum) / cast.ToFloat64(len(ps))
}

// PrintDetail 打印平均类数据
func (ps Processes) PrintDetail(algorithm string) {
	fmt.Printf("Algorithm:   %s\n", algorithm)
	fmt.Printf("AverageTAT:  %f\n", ps.GetAverageTurnaroundTime())
	fmt.Printf("AverageWTAT: %f\n", ps.GetAverageWeightedTurnaroundTime())
	if algorithm == StrRR {
		fmt.Printf("AverageWT:   %f\n", ps.GetAverageWaitTimeRR())
	} else {
		fmt.Printf("AverageWT:   %f\n", ps.GetAverageWaitTime())
	}
}

/*==================================================*/

// ProcessesSJF SJF临时使用的进程序列，用于排序
type ProcessesSJF []*Process

func (ps ProcessesSJF) Len() int {
	return len(ps)
}

func (ps ProcessesSJF) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// Less 排列优先级：服务时间 > 到达时间
func (ps ProcessesSJF) Less(i, j int) bool {
	if ps[i].ServiceTime < ps[j].ServiceTime {
		return true
	} else if ps[i].ServiceTime == ps[j].ServiceTime {
		return ps[i].ArrivalTime < ps[j].ArrivalTime
	}
	return false
}

// Print 格式化输出
//func (ps ProcessesSJF) Print() {
//	fmt.Printf("pid\ttt\twtt\twt\n")
//	for _, p := range ps {
//		fmt.Printf("%d\t%d\t%.2f\t%d\n",
//			p.Pid,
//			p.EndTime-p.ArrivalTime,
//			cast.ToFloat64(p.EndTime-p.ArrivalTime)/cast.ToFloat64(p.ServiceTime),
//			p.StartTime-p.ArrivalTime)
//	}
//}
