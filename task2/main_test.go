package main

import (
	"fmt"
	"math/rand"
	"testing"
)

type testCase struct {
	array  []int
	value  int
	result []int
}

type testSortCase struct {
	array  []int
	result []int
}

func TestInsert(t *testing.T) {
	cases := []testCase{
		{[]int{1, 2}, -1, []int{1, 2}},
		{[]int{0, 1}, -1, []int{0, 1}},
		{[]int{}, -1, []int{}},
		{[]int{10, 4, 3, 9}, -1, []int{10, 4, 3, 9}},
		{[]int{10, 4, 3, 1}, -1, []int{10, 4, 3, 1}},
	}
	for _, c := range cases {
		got := insert(c.array, c.value)
		if !isSame(c.result, got) {
			t.Fail()
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []testCase{
		{[]int{1, 2}, -1, []int{2}},
		{[]int{0, 1}, -5, []int{0, 1}},
		{[]int{}, 1, []int{}},
		{[]int{10, 4, 3, 9, 1}, -1, []int{10, 4, 3, 9}},
		{[]int{10, 4, 3, 1}, -2, []int{10, 4, 3, 1}},
	}
	for _, c := range cases {
		got := remove(c.array, c.value)
		if !isSame(c.result, got) {
			t.Fail()
		}
	}
}

func TestSort(t *testing.T) {
	cases := []testSortCase{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 0}, []int{0, 1}},
		{[]int{}, []int{}},
		{[]int{10, 4, 3, 9, 1}, []int{1, 3, 4, 9, 10}},
		{[]int{10, 4, 9, 2, 3, 1}, []int{1, 2, 3, 4, 9, 10}},
	}
	for _, c := range cases {
		got := sort(c.array)
		if !isSame(c.result, got) {
			t.Fail()
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	rand.Seed(1)
	var data []int
	x := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		insert(data, x)
	}
}
func BenchmarkRemove(b *testing.B) {
	rand.Seed(1)
	var data []int
	x := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		remove(data, x)
	}
}
func BenchmarkSort(b *testing.B) {
	rand.Seed(1)
	var data []int = []int{0, 123, 234, 11, 234, 98, 0, 1234, 123, 23, 234, 6, 2, 4, 5, 6, 2, 1}
	for i := 0; i < b.N; i++ {
		sort(data)
	}
}

func isSame(expected []int, actual []int) bool {
	if len(expected) != len(actual) {
		fmt.Print(expected, actual)
		return false
	}
	for key, value := range expected {
		if value != actual[key] {

			return false
		}
	}

	return true
}
