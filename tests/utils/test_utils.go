package test_utils

import (
	"fmt"
	"migoro/utils"
	"time"
)

func MockUtilsTimestamp(t int64) {
	utils.TimeUnix = func() int64 {
		fmt.Println(t)
		return t
	}
}

func RestoreUtilsTimestamp() {
	utils.TimeUnix = time.Now().UnixNano
}
