package random

import (
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	for i := 0; i < size; i++ {
		if is_all {
			ikind = r.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

func KrandNum(size int) string {
	return Krand(size, 0)
}

func KrandLowerChar(size int) string {
	return Krand(size, 1)
}

func KrandUpperChar(size int) string {
	return Krand(size, 2)
}

func KrandAll(size int) string {
	return Krand(size, 3)
}
