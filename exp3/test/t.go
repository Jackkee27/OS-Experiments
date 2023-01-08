package test

import "fmt"

func ttt() {
	a := []int{1, 2, 4}
	b := make([]int, 3)

	copy(b, a)
	b[0] = 999
	fmt.Println("a[0] = ", a[0])
}
