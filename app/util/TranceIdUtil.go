package util

import (
	"fmt"
	"runtime"
)

//打印当前 协程 id

func printTranceId() int64 {
	// 获取当前协程的 ID
	buf := make([]byte, 64)
	n := runtime.Stack(buf, false)
	id := int64(0)
	fmt.Sscanf(string(buf[:n]), "goroutine %d ", &id)
	return id
}
