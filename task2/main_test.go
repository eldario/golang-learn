package main

import (
	"testing"
)

type testCase struct {
	array  []int
	value  int
	result []int
}

func TestInsert(t *testing.T) {
	cases := []testCase{
		{[]int{1, 2}, -1, []int{2}},
		{[]int{0, 1}, -1, []int{0}},
		{[]int{}, -1, []int{}},
		{[]int{10, 4, 3, 9}, -1, []int{3, 4, 9, 10}},
		{[]int{10, 4, 3, 1}, -1, []int{3, 4, 10}},
	}
	for _, c := range cases {
		got := insert(c.array, c.value)
		if got == c.result {
			t.Fail()
		}
	}
}
