package util

import (
	"strconv"
)

func String2Int64(str string) int64 {
	result, _ := strconv.ParseInt(str, 10, 64)
	return result
}
