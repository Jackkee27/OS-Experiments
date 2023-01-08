package system

// compare 如果切片a的每一个元素都大于b，就返回true；否则返回false.
// 若两个数组的大小不相等，则返回一个error
func compare(a, b []int) (bool, error) {
	if len(a) != len(b) {
		return false, ErrInvalidComparison
	}

	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return false, nil
		}
	}

	return true, nil
}

// findProcessByPid 根据pid查找进程在process中的下标，没找到则返回-1和一个error
func findProcessByPid(processes []Process, pid int) (idx int, err error) {
	for idx, p := range processes {
		if p.Pid == pid {
			return idx, nil
		}
	}
	return -1, ErrProcessNotFound
}
