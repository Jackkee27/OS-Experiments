package ts

import (
	"testing"
	"time"
)

// TestTS 测试TS方法
// cmd: go test -v ts_test.go ts.go
func TestTS(t *testing.T) {
	p := &Process{}
	p.Init(16)

	p.PrintTitle()
	go p.Produce()
	go p.Consume()
	time.Sleep(1 * time.Second)
}
