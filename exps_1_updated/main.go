package main

import (
	"exps_1_updated/models"
	"exps_1_updated/run"
	"fmt"
)

func main() {
	// input 输入部分
	var n int    // 进程数量
	var mode int // 调度模式, 0 - fcfs, 1 - sjf, 2 - rr
	fmt.Print("Input n: ")
	_, _ = fmt.Scan(&n)
	fmt.Print("Input mode: ")
	_, _ = fmt.Scan(&mode)

	var processes models.Processes // 进程序列
	fmt.Println("input pid, arrivalTime and serviceTime:")
	for i := 0; i < n; i++ {
		p := &models.Process{}
		_, _ = fmt.Scan(&p.Pid, &p.ArrivalTime, &p.ServiceTime)
		processes = append(processes, p)
	}

	// 运行调度算法程序
	if err := run.Run(processes, mode); err != nil {
		fmt.Printf("run.Run() failed, err: %v\n", err)
		return
	}
}
