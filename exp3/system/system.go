package system

import (
	"encoding/json"
	"fmt"
)

// System 模拟系统
type System struct {
	Available []int     // 系统剩余的资源
	Processes []Process // 系统进程
}

// Request 进程pid发送请求
func (s *System) Request(pid int, req []int) error {
	data, _ := json.Marshal(s)
	sysTmp := &System{} // sysTmp 临时变量，便于状态回滚
	if err := json.Unmarshal(data, &sysTmp); err != nil {
		return err
	}

	// 1.发出请求之前，判断系统是否处于安全状态
	safeOrder, err := sysTmp.GetSystemState()
	if err != nil {
		return err
	}
	fmt.Println("before request: [safe], safeOrder = ", safeOrder)

	// 2.TryAllocate
	if err := sysTmp.Allocate(pid, req); err != nil {
		return err
	}

	// 3.请求发出之后，系统是否处于安全状态
	safeOrder, err = sysTmp.GetSystemState()
	if err != nil {
		return err
	}
	fmt.Println("after  request: [safe], safeOrder = ", safeOrder)

	s.Available = sysTmp.Available
	s.Processes = sysTmp.Processes
	return nil
}

// GetSystemState 获取系统当前状态
func (s *System) GetSystemState() (safeOrder []int, err error) {
	data, _ := json.Marshal(s)
	sysTmp := &System{} // sysTmp 临时变量，便于状态回滚
	if err := json.Unmarshal(data, &sysTmp); err != nil {
		return []int{}, err
	}

	n := len(sysTmp.Processes)                     // 进程数量
	finish := 0                                    // 已完成的进程数量
	visited := make([]bool, len(sysTmp.Processes)) // 记录已经访问的进程

	for finish < n {
		find := false // 本次是否找到符合条件的进程
		idx := -1     // 符合条件的进程下标
		// 1.按顺序找到第一个符合条件的进程: sysTmp.Available[] >= p.Need[]
		for idxP, p := range sysTmp.Processes {
			if !visited[idxP] { // 在没有遍历过的进程中查找
				find, _ = compare(sysTmp.Available, p.Need)
				if find { // 符合要求
					idx = idxP
					break
				}
			}
		}
		// 2.找到：Available+Available, finish++
		if find {
			visited[idx] = true
			sysTmp.Processes[idx].Finished = true
			// 切片拷贝
			sysTmp.Processes[idx].work = make([]int, len(sysTmp.Available))
			sysTmp.Processes[idx].workAndAllocation = make([]int, len(sysTmp.Available))
			copy(sysTmp.Processes[idx].work, sysTmp.Available)

			for i := 0; i < len(sysTmp.Available); i++ {
				sysTmp.Available[i] += sysTmp.Processes[idx].Allocation[i]       // 更新Work
				sysTmp.Processes[idx].workAndAllocation[i] = sysTmp.Available[i] // 记录Work+Available
			}

			safeOrder = append(safeOrder, sysTmp.Processes[idx].Pid)
			finish++
		} else { // 否则，返回系统不安全的提示
			return []int{}, ErrSystemUnsafe
		}
	}

	// 按照指定顺序输出结果
	sysTmp.PrintSafeOrder(safeOrder)

	return
}

// Allocate 运行进程
func (s *System) Allocate(pid int, req []int) error {
	data, _ := json.Marshal(s)
	sysTmp := &System{} // sysTmp 临时变量，便于状态回滚
	if err := json.Unmarshal(data, &sysTmp); err != nil {
		return err
	}

	// 查找pid
	idx, err := findProcessByPid(sysTmp.Processes, pid)
	if err != nil {
		return err
	}

	for i := 0; i < len(sysTmp.Available); i++ {
		// 尝试分配资源
		sysTmp.Available[i] -= req[i]
		sysTmp.Processes[idx].Allocation[i] += req[i]
		sysTmp.Processes[idx].Need[i] -= req[i]
		if sysTmp.Available[i] < 0 { // 系统的某一类资源不够这次请求
			return ErrSystemInsufficientResources
		}
		if sysTmp.Processes[idx].Need[i] < 0 { // 请求的某一类资源数量，超过Need[i]，不予分配
			return ErrProcessRequestExceeded
		}
	}

	s.Available = sysTmp.Available
	s.Processes = sysTmp.Processes
	return nil
}

// PrintDetail 格式化输出当前系统状态
func (s *System) PrintDetail() {
	fmt.Printf("Available: %v\n", s.Available)
	fmt.Println("Processes:")
	fmt.Println("pid\tNeed\t\tAvailable\tFinished?")
	for _, p := range s.Processes {
		fmt.Printf("%v\t%v\t\t%v\t\t%v\n",
			p.Pid,
			p.Need,
			p.Allocation,
			p.Finished,
		)
	}
}

// PrintSafeOrder 按照给定的安全序列格式化输出
func (s *System) PrintSafeOrder(safeOrder []int) {
	fmt.Println("Processes:")
	fmt.Println("Pid\tWork\t\tNeed\t\tAllocation\tWork+Allocation\tFinished?")
	for _, i := range safeOrder {
		p := s.Processes[i]
		fmt.Printf("%2v\t%2v\t%2v\t%2v\t%2v\t%v\n",
			p.Pid,
			p.work,
			p.Need,
			p.Allocation,
			p.workAndAllocation,
			p.Finished,
		)
	}
}
