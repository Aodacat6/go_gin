package conf

import (
	"fmt"
	"testing"
)

// 测试文件规范
// 1、这个问题可能是因为你的测试文件没有被识别为测试文件。请确保你的测试文件名以 _test.go 结尾，例如 conf_test.go。
// 2、此外，你需要在测试方法名前面加上 Test 前缀，例如 TestSyncTable。这样 Go 语言才能识别这个方法是一个测试方法。
func Test_SyncTable(t *testing.T) {
	result := SyncTable()
	if result == 0 {
		fmt.Println("执行成功")
	} else {
		fmt.Println("执行失败")
	}
}
