package main

import (
	"fmt"
	"math"
)

func main() {
	s := make([]int, 0)
	oldCap := cap(s)
	fmt.Printf("len: %8d\tcap: %8d\n", len(s), cap(s))
	for i := 0; i < 10000000; i++ {
		s = append(s, i)
		cap := cap(s)
		if cap != oldCap {
			fmt.Printf("len: %8d\tcap: %8d\tfactor: %6.3f\n", len(s), cap, factor(cap, oldCap))
		}
		oldCap = cap
	}
}

func factor(n int, oldN int) float64 {
	if oldN == 0 {
		return math.NaN()
	} else {
		return float64(n) / float64(oldN)
	}
}
