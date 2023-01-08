package algorithms

import (
	"sort"
)

// SCANRunner SCAN执行器
type SCANRunner struct {
	Runner
	direction int
}

// Init 初始化scan
func (scan *SCANRunner) Init(trackOrder []int, initTrackNum, direction int) {
	scan.Runner.Init(trackOrder, initTrackNum)
	scan.direction = direction
}

// Run 执行SCAN算法
func (scan *SCANRunner) Run() error {
	trackOrder := scan.trackOrder
	var moveDistance []int
	var finalTrackOrder []int // 最终访问顺序
	// 1. 将原来的trackOrder排序
	sort.Ints(trackOrder)    // 升序
	if scan.direction == 1 { // 1-从里到外，即磁道号从大到小
		reverse(trackOrder) // 变为降序
	}
	// 2. 先计算访问顺序
	curr := scan.initTrackNum
	pos := getFirst(trackOrder, curr, scan.direction) // 搜索分界点
	if pos == -1 {
		pos = len(trackOrder)
	}
	finalTrackOrder = trackOrder[pos:]
	//fmt.Printf("%v", finalTrackOrder)
	reverse(trackOrder[:pos]) // 逆转方向
	finalTrackOrder = append(finalTrackOrder, trackOrder[:pos]...)
	// 3. for循环遍历访问顺序
	for _, track := range finalTrackOrder {
		d := abs(track - curr)
		moveDistance = append(moveDistance, d)

		curr = track
	}

	scan.trackOrder = finalTrackOrder
	scan.moveDistance = moveDistance

	scan.Print()

	return nil
}
