package algorithms

import (
	"errors"
)

// findPageById returns the index of page exist in block, and an error if the page does not exist
func findPageById(block []int, pageId int) (int, error) {
	for idx, p := range block {
		if p == pageId {
			return idx, nil
		}
	}
	return -1, errors.New("not found the page in Block")
}

// getMaxValIndex returns the index of maxVal in arr
func getMaxValIndex(arr []int) (int, error) {
	if len(arr) < 0 {
		return -1, errors.New("arr is nil")
	}
	maxIdx := 0
	for idx, val := range arr {
		if val > arr[maxIdx] {
			maxIdx = idx
		}
	}

	return maxIdx, nil
}

// returns indices of 0 in arr
func zeroNumInArr(arr []int) []int {
	var ret []int
	for idx, v := range arr {
		if v == 0 {
			ret = append(ret, idx)
		}
	}

	return ret
}
