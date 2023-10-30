package util

import "strconv"

func String2Int64(str string) int64 {
	str2int, _ := strconv.Atoi(str)
	return int64(str2int)
}
