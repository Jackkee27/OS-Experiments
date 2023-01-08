package algorithms

// OPTRunner running with OPT algorithm
type OPTRunner struct {
	LRURunner
}

// Init OPT runner
func (opt *OPTRunner) Init(pageSeq []int, blockNum int) {
	opt.LRURunner.Init(pageSeq, blockNum)
}

// Run FIFO algorithm
func (opt *OPTRunner) Run() error {
	n := len(opt.PageSeq)
	m := len(opt.Block)
	pageInBlock := 0 // page num in Block
	opt.Cache = make([][]int, n)

	for i := 0; i < n; i++ {
		page := opt.PageSeq[i]
		// 1. check whether the current page is in Block or not
		idx, _ := findPageById(opt.Block, page)
		if idx == -1 { // not found, add or replace
			if pageInBlock < m { // not full, add it directly
				opt.Block[pageInBlock] = page
				pageInBlock++
			} else { // full, replace
				// find a place to be replaced(opt)
				idx, err := opt.optProcess(i)
				if err != nil {
					return err
				}
				opt.Block[idx] = page
			}
			opt.LackNum++
		} else { // found, hit
			opt.HitNum++
			opt.HitPage[i] = true
		}
		// 2. set leastRecentlyUsed
		opt.leastRecentlyUsed(page)
		// 3. save the current status into Cache
		opt.Cache[i] = make([]int, m)
		copy(opt.Cache[i], opt.Block)
	}

	// 3. print
	opt.PrintPageSeq()

	return nil
}

// optProcess returns the position to be replaced, and an error
func (opt *OPTRunner) optProcess(currentI int) (int, error) {
	mark := make([]int, len(opt.Block)) // mark if visited
	cnt := 0
	for i := currentI + 1; i < len(opt.PageSeq); i++ {
		page := opt.PageSeq[i]
		idx, _ := findPageById(opt.Block, page)
		if idx != -1 && mark[idx] == 0 { // found in block & did not visit
			cnt++
			mark[idx] = cnt
		}
	}

	// exists 0?
	zeros := zeroNumInArr(mark) // idx in mark
	if len(zeros) > 0 {         // y - existed not visited
		arr := make([]int, len(zeros))
		for idx, val := range zeros {
			arr[idx] = opt.latelyVisit[val]
		}
		idx, _ := getMaxValIndex(arr) // idx in zeros

		return zeros[idx], nil
	}
	// n - all visited
	idx, err := getMaxValIndex(mark)
	if err != nil {
		return -1, err
	}
	return idx, nil
}
