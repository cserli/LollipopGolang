package Log_Eio

import (
	"fmt"
)

const (
	DEBUG_FMT = 2 // debug
	// RELEASE_FMT = 2 // release
)

// 打印函数，发布版本打日志，否则打输出
func Fmt(data ...interface{}) {

	if DEBUG_FMT == 1 {
		fmt.Println(data)
		return
	}
}
