package system

// Process 模拟进程
type Process struct {
	Finished   bool
	Pid        int   // 进程ID
	Max        []int // 所需资源的最大值
	Need       []int // 还需要的资源
	Allocation []int // 某时刻，进程分配到的 j 类资源

	work              []int // 用于输出
	workAndAllocation []int // 用于输出
}
