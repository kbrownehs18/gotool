package randutils

import (
	"math/rand"
	"time"
)

// New return *rand.Rand
func New() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Random rand string slice
func Random(arr []string) string {
	return arr[New().Intn(len(arr))]
}

// RangeInt return min<=x<max
func RangeInt(min, max int) int {
	if max == min {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return min + New().Intn(max-min)
}

// RangeInt32 return min<=x<max
func RangeInt32(min, max int32) int32 {
	if max == min {
		return min
	}
	if max < min {
		min, max = max, min
	}
	return min + New().Int31n(max-min)
}
