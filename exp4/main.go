package main

import (
	"exp4/algorithms"
	"exp4/models"
	"fmt"
)

const MaxNum = 100

func main() {
	// 1.input part
	// input blockNum and pageNum
	var blockNum, pageNum int
	fmt.Print("input blockNum & pageNum: ")
	_, _ = fmt.Scan(&blockNum, &pageNum)
	if pageNum > MaxNum {
		fmt.Printf("input page num error")
		return
	}
	// input page sequences
	fmt.Print("input page sequences: ")
	pageSeq := make([]int, pageNum)
	for i := 0; i < pageNum; i++ {
		_, _ = fmt.Scan(&pageSeq[i])
	}
	for {
		// input mode
		var mode int
		fmt.Print("input mode(1-FIFO, 2-OPT, 3-LRU): ")
		_, _ = fmt.Scan(&mode)
		mode--

		// 2.run algorithms

		switch mode {
		case models.ModeFIFO:
			r := &algorithms.FIFORunner{}
			r.Init(pageSeq, blockNum)
			if err := r.Run(); err != nil {
				fmt.Printf("r.Run() failed, err: %v\n", err)
				return
			}
		case models.ModeOPT:
			r := &algorithms.OPTRunner{}
			r.Init(pageSeq, blockNum)
			if err := r.Run(); err != nil {
				fmt.Printf("r.Run() failed, err: %v\n", err)
				return
			}
		case models.ModeLRU:
			r := &algorithms.LRURunner{}
			r.Init(pageSeq, blockNum)
			if err := r.Run(); err != nil {
				fmt.Printf("r.Run() failed, err: %v\n", err)
				return
			}
		default:
			fmt.Println("input mode failed")
		}

		var flag string
		fmt.Print("Input [q] to quit, others to continue: ")
		_, _ = fmt.Scan(&flag)
		if flag == "q" {
			break
		}
	}

}
