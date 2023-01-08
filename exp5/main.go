package main

import (
	"exp5/algorithms"
	"exp5/model"
	"fmt"
)

const MaxRequestNum = 100

func main() {
	// 输入部分
	var trackNum int     // 磁道访问数量
	var initTrackNum int // 开始磁道号
	fmt.Print("input trackNum & initTrackNum: ")
	_, _ = fmt.Scan(&trackNum, &initTrackNum)
	for trackNum > MaxRequestNum {
		fmt.Printf("error: invalid track num")
		fmt.Print("input trackNum again: ")
		_, _ = fmt.Scan(&trackNum)
	}

	trackOrder := make([]int, trackNum) // 请求的磁道序列
	fmt.Print("input track order: ")
	for i := 0; i < trackNum; i++ {
		_, _ = fmt.Scan(&trackOrder[i])
	}

	var mode int // 算法
	fmt.Print("input mode(1-FCFS, 2-SSTF, 3-SCAN, 4-C-SCAN): ")
	_, _ = fmt.Scan(&mode)
	mode--

	var direction int // 移动方向，0-从外到里，1-从里到外
	if mode == model.ModeSCAN || mode == model.ModeCSCAN {
		fmt.Print("input direction: ")
		_, _ = fmt.Scan(&direction)
	}

	switch mode {
	case model.ModeFCFS: // 运行FCFS
		r := &algorithms.FCFSRunner{}
		r.Init(trackOrder, initTrackNum)
		if err := r.Run(); err != nil {
			fmt.Printf("r.Run() failed, err: %v\n", err)
		}
	case model.ModeSSTF: // 运行SSTF
		r := &algorithms.SSTFRunner{}
		r.Init(trackOrder, initTrackNum)
		if err := r.Run(); err != nil {
			fmt.Printf("r.Run() failed, err: %v\n", err)
		}
	case model.ModeSCAN: // 运行SCAN
		r := &algorithms.SCANRunner{}
		r.Init(trackOrder, initTrackNum, direction)
		if err := r.Run(); err != nil {
			fmt.Printf("r.Run() failed, err: %v\n", err)
		}
	case model.ModeCSCAN: // 运行C-SCAN
		r := &algorithms.CSCANRunner{}
		r.Init(trackOrder, initTrackNum, direction)
		if err := r.Run(); err != nil {
			fmt.Printf("r.Run() failed, err: %v\n", err)
		}
	default:
		fmt.Println("error: invalid mode")
	}
}
