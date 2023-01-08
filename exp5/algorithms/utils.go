package algorithms

import (
	"math"
)

// GetNearest 在arr的未访问过的节点中，返回与curr最接近的值的下标retIdx和值retVal
func GetNearest(arr []int, visited []bool, curr int) (int, int) {
	retIdx := 0
	retVal := math.MaxInt
	for idx, val := range arr {
		d := abs(val - curr)
		if d < retVal && !visited[idx] {
			retIdx, retVal = idx, d
		}
	}
	return retIdx, retVal
}

// reverse 反转原序列
func reverse(arr []int) {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-1-i] = arr[l-1-i], arr[i]
	}
}

// getFirst 获取第一个大于等于(direction==0)或小于等于(direction==1)curr的元素下标，若不存在则返回-1
func getFirst(arr []int, curr, direction int) int {
	for idx, val := range arr {
		if (direction == 0 && val >= curr) || (direction == 1 && val <= curr) {
			return idx
		}
	}
	return -1
}

// abs 获取n的绝对值
func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}
