package algorithms

import "exps_1_updated/models"

type Runner interface {
	Init(processes models.Processes) error // 初始化
	Run() error                            // 执行方法
}
