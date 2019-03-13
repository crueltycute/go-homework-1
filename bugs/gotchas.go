package main

import (
	"sort"
	"strconv"
)

func ReturnInt() int {
	return 1
}

func ReturnFloat() float32 {
	return 1.1
}

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(values []int ) (out string) {
	for i := range values {
		text := strconv.Itoa(values[i])
		out += text
	}
	return out
}

func MergeSlices(first []float32, second []int32) []int {
	var temp []int
	for i := range first {
		temp = append(temp, int(first[i]))
	}
	for i := range second {
		temp = append(temp, int(second[i]))
	}
	return temp
}

func GetMapValuesSortedByKey(item map[int]string) []string {
	var keys []int
	var values []string

	for i := range item {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for i := range keys {
		values = append(values, item[keys[i]])
	}
	return values
}