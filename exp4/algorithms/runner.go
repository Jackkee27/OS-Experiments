package algorithms

import (
	"fmt"
	"github.com/spf13/cast"
)

// runner basic runner
type runner struct {
	LackNum int     // lack page num
	HitNum  int     // hit page num
	HitPage []bool  // record the pages hit
	PageSeq []int   // page sequence
	Block   []int   // physical blocks
	Cache   [][]int // save block data every step
}

// Init basic runner
func (r *runner) Init(pageSeq []int, blockNum int) {
	r.PageSeq = pageSeq
	r.HitPage = make([]bool, len(pageSeq))
	r.Block = make([]int, blockNum)
}

// GetPageFaultRate returns page fault rate
func (r *runner) GetPageFaultRate() float64 {
	return cast.ToFloat64(r.LackNum) / cast.ToFloat64(len(r.PageSeq))
}

// PrintPageSeq prints the page sequence
func (r *runner) PrintPageSeq() {
	// 1. print page seq
	for _, p := range r.PageSeq {
		fmt.Printf("%d\t", p)
	}
	fmt.Println()
	for i := 0; i < len(r.PageSeq); i++ {
		print("--------")
	}
	fmt.Println()
	// 2. print rows
	for i := 0; i < len(r.Block); i++ {
		for c := 0; c < len(r.PageSeq); c++ {
			if r.HitPage[c] == true { // hit
				if r.PageSeq[c] == r.Cache[c][i] {
					fmt.Printf("hit")
				}
				fmt.Print("\t")
			} else { // do not hit
				if r.Cache[c][i] == 0 {
					fmt.Printf("\t")
				} else {
					fmt.Printf("%d\t", r.Cache[c][i])
				}
			}
		}
		fmt.Println()
	}
	// 3.print page fault rate
	fmt.Printf("HitNum  = %d\n", r.HitNum)
	fmt.Printf("LackNum = %d\n", r.LackNum)
	fmt.Printf("PageFaultRate = %f%%\n", 100*r.GetPageFaultRate())
}
