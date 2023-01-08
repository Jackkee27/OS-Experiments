package algorithms

// LRURunner running with LRU algorithm
type LRURunner struct {
	runner
	latelyVisit []int // count lately visit, set current page = 0, and others++
}

// Init LRU runner
func (lru *LRURunner) Init(pageSeq []int, blockNum int) {
	lru.runner.Init(pageSeq, blockNum)
	lru.latelyVisit = make([]int, blockNum)
}

// Run LRU algorithm
func (lru *LRURunner) Run() error {
	n := len(lru.PageSeq)
	m := len(lru.Block)
	pageInBlock := 0 // page num in Block
	lru.Cache = make([][]int, n)

	for i := 0; i < n; i++ {
		page := lru.PageSeq[i]
		// 1. check whether the current page is in Block or not
		idx, _ := findPageById(lru.Block, page)
		if idx == -1 { // not found, add or replace
			if pageInBlock < m { // not full, add it directly
				lru.Block[pageInBlock] = page
				pageInBlock++
			} else { // full, replace
				// find a place to be replaced(least recently use)
				idx, _ := getMaxValIndex(lru.latelyVisit)
				lru.Block[idx] = page
			}
			lru.LackNum++
		} else { // found, hit
			lru.HitNum++
			lru.HitPage[i] = true
		}
		// 2. set leastRecentlyUsed
		lru.leastRecentlyUsed(page)
		// 3. save the current status into Cache
		lru.Cache[i] = make([]int, m)
		copy(lru.Cache[i], lru.Block)
	}

	// 3. print
	lru.PrintPageSeq()

	return nil
}

// leastRecentlyUsed sets current latelyVisit[page] = 0, and others++
func (lru *LRURunner) leastRecentlyUsed(page int) {
	for i := 0; i < len(lru.latelyVisit); i++ {
		if lru.Block[i] == page {
			lru.latelyVisit[i] = 0
		} else {
			lru.latelyVisit[i]++
		}
	}
}
