package main

import (
	"math/rand"
	"testing"
)

func generateBenchSlice(b *testing.B, n int) []int {
	b.StopTimer()
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}
	b.StartTimer()
	return s
}

func BenchmarkQuickSortA1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateBenchSlice(b, 1000)
		quickSortA(s)
	}
}

func BenchmarkQuickSortW1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateBenchSlice(b, 1000)
		quickSortW(s)
	}
}

func BenchmarkQuickSortA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateBenchSlice(b, n)
		quickSortA(s)
	}
}

func BenchmarkQuickSortW(b *testing.B) {
	for n := 0; n < b.N; n++ {
		s := generateBenchSlice(b, n)
		quickSortW(s)
	}
}
