package main

import (
	"exp1/algorithms"
	"exp1/models"
	"fmt"
)

func main() {
	var n int    // 进程数量
	var mode int // 调度模式
	fmt.Print("Input n: ")
	_, _ = fmt.Scan(&n)
	fmt.Print("Input mode: ")
	_, _ = fmt.Scan(&mode)

	// input 输入部分
	var processes models.Processes // 进程序列
	fmt.Println("input pid, arrivalTime and serviceTime:")
	for i := 0; i < n; i++ {
		p := &models.Process{}
		_, _ = fmt.Scan(&p.Pid, &p.ArrivalTime, &p.ServiceTime)
		processes = append(processes, p)
	}

	if err := algorithms.Run(processes, mode); err != nil {
		fmt.Println("algorithms.Run() failed, err:", err)
		return
	}
}
