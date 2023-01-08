package algorithms

// FIFORunner running with FIFO algorithms
type FIFORunner struct {
	runner
}

// Init FIFO runner
func (fifo *FIFORunner) Init(pageSeq []int, blockNum int) {
	fifo.runner.Init(pageSeq, blockNum)
}

// Run FIFO algorithm
func (fifo *FIFORunner) Run() error {
	ptr := 0 // virtual pointer
	n := len(fifo.PageSeq)
	m := len(fifo.Block)
	pageInBlock := 0 // page num in Block
	fifo.Cache = make([][]int, n)

	for i := 0; i < n; i++ {
		page := fifo.PageSeq[i]
		// 1. check whether the current page is in Block or not
		idx, _ := findPageById(fifo.Block, page)
		if idx == -1 { // not found, add or replace
			if pageInBlock < m { // not full, add it directly
				fifo.Block[pageInBlock] = page
				pageInBlock++
			} else { // full, replace
				// find a place to be replaced
				fifo.Block[ptr] = page
			}
			fifo.LackNum++
			ptr = (ptr + 1) % m
		} else { // found, hit
			fifo.HitNum++
			fifo.HitPage[i] = true
		}
		// 2. save the current status into Cache
		fifo.Cache[i] = make([]int, m)
		copy(fifo.Cache[i], fifo.Block)
	}

	// 3. print
	fifo.PrintPageSeq()

	return nil
}
