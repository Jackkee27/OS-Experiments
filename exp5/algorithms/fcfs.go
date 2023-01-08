package algorithms

// FCFSRunner FCFS执行器
type FCFSRunner struct {
	Runner
}

// Init 初始化FCFSRunner
func (fcfs *FCFSRunner) Init(trackOrder []int, initTrackNum int) {
	fcfs.Runner.Init(trackOrder, initTrackNum)
}

// Run 执行FCFS算法
func (fcfs *FCFSRunner) Run() error {
	trackOrder := fcfs.trackOrder
	moveDistance := make([]int, len(trackOrder))
	curr := fcfs.initTrackNum // 当前track

	for i := 0; i < len(trackOrder); i++ {
		// 1.计算移动的距离
		d := abs(trackOrder[i] - curr)
		moveDistance[i] = d

		curr = trackOrder[i]
	}

	fcfs.moveDistance = moveDistance

	fcfs.Print()

	return nil
}
