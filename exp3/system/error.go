package system

import "errors"

// in system.go
var (
	ErrSystemUnsafe                = errors.New("system unsafe")          // 系统不安全
	ErrSystemInsufficientResources = errors.New("insufficient resources") // 系统资源不够

	ErrProcessRequestExceeded = errors.New("exceeded the required(max) value") // 进程请求超过其所需
)

// in util.go
var (
	ErrInvalidComparison = errors.New("invalid compare between int slices") // 无效的比较（切片长度不同）
	ErrProcessNotFound   = errors.New("process not found")                  // 没找到进程
)
