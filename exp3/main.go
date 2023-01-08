package main

import (
	"exp3/system"
	"fmt"
)

func main() {
	// 1.输入进程信息，系统信息
	var n, m int
	fmt.Print("input n (num of processes) & m (num of system resource): ")
	_, _ = fmt.Scan(&n, &m)

	sys := &system.System{}
	fmt.Print("input system Max: ")
	for i := 0; i < m; i++ {
		var num int
		_, _ = fmt.Scan(&num)
		sys.Available = append(sys.Available, num)
	}

	fmt.Print("input the process info(pid, resources & need):")
	var processes []system.Process
	for i := 0; i < n; i++ {
		var (
			pid        int   // 进程ID
			allocation []int // 已经分配的资源
			need       []int // 还需要的资源
			max_       []int // 总共需要的资源
		)
		fmt.Print("input Pid: ")
		// pid
		_, _ = fmt.Scan(&pid)
		fmt.Print("input Allocation: ")
		// allocation
		for j := 0; j < m; j++ {
			var t int
			_, _ = fmt.Scan(&t)
			sys.Available[j] -= t
			allocation = append(allocation, t)
		}
		fmt.Print("input Need: ")
		// need
		for j := 0; j < m; j++ {
			var t int
			_, _ = fmt.Scan(&t)
			need = append(need, t)
		}
		// max
		for j := 0; j < m; j++ {
			max_ = append(max_, allocation[j]+need[j])
		}
		processes = append(processes, system.Process{
			Pid:        pid,
			Max:        max_,
			Need:       need,
			Allocation: allocation,
		})
	}
	sys.Processes = processes
	// 2.Request
	for {
		var (
			pid  int
			req  = make([]int, m)
			flag string
		)
		// 输入q退出，其他则继续发送request
		fmt.Print("quit?(input q to quit): ")
		_, _ = fmt.Scan(&flag)
		if flag == "q" {
			break
		}
		// 输入请求信息
		fmt.Print("input pid of reqs: ")
		_, _ = fmt.Scan(&pid)
		fmt.Print("input reqs: ")
		for i := 0; i < m; i++ {
			_, _ = fmt.Scan(&req[i])
		}
		// 处理请求
		if err := sys.Request(pid, req); err != nil {
			fmt.Printf("sys.Request() failed, err: %v \n", err)
		}
	}
}
