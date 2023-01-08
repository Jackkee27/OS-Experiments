package run

import (
	"exps_1_updated/algorithms/fcfs"
	"exps_1_updated/algorithms/rr"
	"exps_1_updated/algorithms/sjf"
	"exps_1_updated/models"
	"fmt"
)

func Run(processes models.Processes, mode int) error {
	// 输入验证
	if len(processes) <= 0 || len(processes) > models.MaxProcessNumber {
		return ErrNOutOfRange
	}

	switch mode {
	case models.ModeFCFS:
		r := fcfs.Runner{}
		if err := r.Init(processes); err != nil {
			return err
		}
		if err := r.Run(); err != nil {
			return err
		}
	case models.ModeSJF:
		r := sjf.Runner{}
		if err := r.Init(processes); err != nil {
			return err
		}
		if err := r.Run(); err != nil {
			return err
		}
	case models.ModeRR:
		// input the size of time slice
		var q int // 时间片大小
		fmt.Print("input the size of time slice: ")
		_, _ = fmt.Scan(&q)
		if q <= 0 { // 验证q的范围
			return ErrPOutOfRange
		}
		r := rr.Runner{}
		if err := r.Init(processes, q); err != nil {
			return err
		}
		if err := r.Run(); err != nil {
			return err
		}
	default: // 返回InvalidMode错误
		return ErrInvalidMode
	}
	return nil
}
