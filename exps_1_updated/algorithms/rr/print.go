package rr

import (
	"exps_1_updated/algorithms"
	"fmt"
)

// PrintTS 打印指定的时间片序列，以及就绪队列
func PrintTS(pid, lastT, currT int, readyQueue algorithms.ReadyQueue) {
	fmt.Printf("%d\t%d\t%d\t", pid, lastT, currT)
	// readyQueue
	readyQueue.Print()
}
