package algorithms

// SSTFRunner SSTF执行器
type SSTFRunner struct {
	Runner
}

// Init 初始化SSTFRunner
func (sstf *SSTFRunner) Init(trackOrder []int, initTrackNum int) {
	sstf.Runner.Init(trackOrder, initTrackNum)
}

// Run 执行SSTF算法
func (sstf *SSTFRunner) Run() error {
	trackOrder := sstf.trackOrder
	moveDistance := make([]int, len(trackOrder))
	visited := make([]bool, len(trackOrder)) // 记录请求track的访问情况
	finalTrackOrder := make([]int, len(trackOrder))

	curr := sstf.initTrackNum // 当前track
	for i := 0; i < len(trackOrder); i++ {
		// 1.计算与当前最近的track
		idx, d := GetNearest(trackOrder, visited, curr)
		visited[idx] = true
		// 2.存至moveDistance
		moveDistance[i] = d
		finalTrackOrder[i] = trackOrder[idx]
		curr = trackOrder[idx]
	}

	sstf.moveDistance = moveDistance
	sstf.trackOrder = finalTrackOrder

	sstf.Print()

	return nil
}
