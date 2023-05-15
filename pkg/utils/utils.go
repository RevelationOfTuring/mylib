package utils

import (
	"fmt"
	"time"
)

func GetUnixTimestamp() int64 {
	fmt.Println("调用mylib - GetUnixTimestamp()")
	return time.Now().Unix()
}
