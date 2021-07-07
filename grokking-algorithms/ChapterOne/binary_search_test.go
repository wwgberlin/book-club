package main

import (
	"testing"
)

func TestBinarySearchString(t *testing.T){

	testCases := []struct{
		name string
		input []string
		element string
		expected int
	}{
		{
			"Happy Path",
			[]string{"a","aa","b"},
			"aa",
			1,
		},
		{
			"Happy Path: even length",
			[]string{"a","aa","b", "zz"},
			"a",
			0,
		},
		{
			"Not found",
			[]string{"a","aa","b", "zz"},
			"c",
			-1,
		},
		{
			"Empty input",
			[]string{},
			"aa",
			-1,
		},

	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T){
			got := binarySearchString(tt.input, tt.element)
			if got != tt.expected {
				t.Errorf("got %v want %v", got, tt.expected)
			}
		})
	}
}
func TestBinarySearchInteger(t *testing.T){

	testCases := []struct{
		name string
		input []int
		element int
		expected int
	}{
		{
			"Happy Path",
			[]int{1,2,3},
			2,
			1,
		},
		{
			"Happy Path: even length",
			[]int{1,2,3,4},
			2,
			1,
		},
		{
			"Not found",
			[]int{1,2,3},
			7,
			-1,
		},
		{
			"Empty input",
			[]int{},
			7,
			-1,
		},

	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T){
			got := binarySearchInteger(tt.input, tt.element)
			if got != tt.expected {
				t.Errorf("got %v want %v", got, tt.expected)
			}
		})
	}
}