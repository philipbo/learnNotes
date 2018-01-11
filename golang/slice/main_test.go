package main

import (
	"testing"
)

const capacity = 1024

func array_init() [capacity]int {
	var a [capacity]int
	for i := 0; i < len(a); i++ {
		a[i] = 1
	}
	return a
}

func slice_init() []int {
	sl := make([]int, capacity)
	for i := 0; i < len(sl); i++ {
		sl[i] = 1
	}
	return sl
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = array_init()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slice_init()
	}
}
