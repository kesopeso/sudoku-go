package util

import (
	"sort"
)

func Assert(condition bool, msg string) {
	if !condition {
		panic("assertion failed: " + msg)
	}
}

func ArraysIntersection(arrays ...[]int) []int {
	occurances := map[int]int{}

	for _, arr := range arrays {
		repeatingArrayValues := make(map[int]bool, len(arr))
		for _, v := range arr {
			if repeatingArrayValues[v] {
				continue
			}
			repeatingArrayValues[v] = true
			occurances[v]++
		}
	}

	intersection := []int{}
	for k, v := range occurances {
		if v == len(arrays) {
			intersection = append(intersection, k)
		}
	}

	sort.Ints(intersection)

	return intersection
}
