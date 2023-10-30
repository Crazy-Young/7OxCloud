package util

import "strconv"

func ParsePageAndPageSize(page, pageSize string) (p, pSize int) {
	p, _ = strconv.Atoi(page)
	pSize, _ = strconv.Atoi(pageSize)
	if p <= 0 {
		p = 1
	}
	switch {
	case pSize > 100:
		pSize = 100
	case pSize <= 0:
		pSize = 10
	}
	return
}
