package algorithms

import (
	"sort"
)

// CSCANRunner C-SCAN执行器
type CSCANRunner struct {
	Runner
	direction int
}

// Init 初始化C-SCAN
func (cs *CSCANRunner) Init(trackOrder []int, initTrackNum, direction int) {
	cs.Runner.Init(trackOrder, initTrackNum)
	cs.direction = direction
}

// Run 执行C-SCAN算法
func (cs *CSCANRunner) Run() error {
	trackOrder := cs.trackOrder
	var moveDistance []int
	var finalTrackOrder []int // 最终访问顺序
	// 1. 将原来的trackOrder排序
	sort.Ints(trackOrder)  // 升序
	if cs.direction == 1 { // 1-从里到外，即磁道号从大到小
		reverse(trackOrder) // 变为降序
	}
	// 2. 先计算访问顺序
	curr := cs.initTrackNum
	pos := getFirst(trackOrder, curr, cs.direction) // 搜索分界点
	if pos == -1 {
		pos = len(trackOrder)
	}
	finalTrackOrder = trackOrder[pos:]
	finalTrackOrder = append(finalTrackOrder, trackOrder[:pos]...) // C-SCAN不用逆转
	// 3. for循环遍历访问顺序
	for _, track := range finalTrackOrder {
		d := abs(track - curr)
		moveDistance = append(moveDistance, d)

		curr = track
	}

	cs.trackOrder = finalTrackOrder
	cs.moveDistance = moveDistance

	cs.Print()

	return nil
}
