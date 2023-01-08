package rr

import (
	"errors"
	"exps_1_updated/models"
)

// timeSlice 时间片
type timeSlice struct {
	isVisited bool // 是否已被访问
	Left      int  // 剩下的时间

	*models.Process // 对应进程
}

// initTimeSlice 初始化时间片序列
func initTimeSlice(processes models.Processes) (tSlice []*timeSlice) {
	for _, p := range processes {
		ts := &timeSlice{
			isVisited: false,
			Left:      p.ServiceTime,
			Process:   p,
		}
		tSlice = append(tSlice, ts)
	}
	return
}

// findTimeSliceByPid 通过PID来查找时间片
func findTimeSliceByPid(timeSlices []*timeSlice, pid int) (int, error) {
	for idx, ts := range timeSlices {
		if ts.Pid == pid {
			return idx, nil
		}
	}
	return -1, errors.New("time slice: not found")
}
