package algorithms

import (
	"testing"
)

func TestGetCloset(t *testing.T) {
	arr := []int{86, 1470, 913, 1774, 948, 1509, 1022, 1750, 130}

	//arr := []int{144, 145, 154, 143}

	r := &CSCANRunner{}
	r.Init(arr, 143, 1)
	_ = r.Run()
}
