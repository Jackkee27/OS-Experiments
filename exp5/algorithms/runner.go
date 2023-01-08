package algorithms

import (
	"fmt"
	"github.com/spf13/cast"
)

// Runner base runner
type Runner struct {
	initTrackNum int
	trackOrder   []int
	moveDistance []int
}

// Init 初始化runner
func (r *Runner) Init(trackOrder []int, initTrackNum int) {
	r.trackOrder = trackOrder
	r.initTrackNum = initTrackNum
}

// Print 输出部分
func (r *Runner) Print() {
	trackNum := len(r.trackOrder)
	// print order
	fmt.Printf("TrackOrder: \n")
	fmt.Printf("%d->\t", r.initTrackNum)
	for i := 0; i < trackNum; i++ {
		fmt.Printf("%d", r.trackOrder[i])
		if i < trackNum-1 {
			fmt.Printf("->\t")
		}
	}
	fmt.Println()
	// print moveDistance
	fmt.Printf("\t")
	for i := 0; i < trackNum; i++ {
		fmt.Printf("%d\t", r.moveDistance[i])
	}
	fmt.Println()
	// print average move distance
	fmt.Printf("Sum Move Distance    : %f\n", r.GetSumMoveDistance())
	fmt.Printf("Average Move Distance: %f\n", r.GetAverageMoveDistance())
}

// GetSumMoveDistance 返回总寻道长度
func (r *Runner) GetSumMoveDistance() float64 {
	var ret int
	for _, val := range r.moveDistance {
		ret += val
	}
	return cast.ToFloat64(ret)
}

// GetAverageMoveDistance 返回平均寻道长度
func (r *Runner) GetAverageMoveDistance() float64 {
	return r.GetSumMoveDistance() / cast.ToFloat64(len(r.moveDistance))
}
