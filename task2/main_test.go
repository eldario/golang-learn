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
		{[]int{1, 2}, 3, []int{1, 2, 3}},
		{[]int{0, 1}, 12, []int{0, 1, 12}},
		{[]int{}, 1, []int{1}},
		{[]int{10, 4, 3, 9}, 1, []int{10, 4, 3, 9, 1}},
		{[]int{10, 4, 3, 1}, 21, []int{10, 4, 3, 1, 21}},
	}
	for _, c := range cases {
		got := insert(&c.array, c.value)
		if !isSame(c.result, got) {
			t.Fail()
		}
	}
}

func TestRemove(t *testing.T) {
	cases := []testCase{
		{[]int{1, 2}, 1, []int{2}},
		{[]int{0, 1}, 5, []int{0, 1}},
		{[]int{}, 1, []int{}},
		{[]int{10, 4, 3, 9, 1}, 1, []int{10, 4, 3, 9}},
		{[]int{10, 4, 3, 1}, 2, []int{10, 4, 3, 1}},
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
	x := rand.Intn(10)
	i := 0
	for ; i < b.N; i++ {
		insert(&data, x*i)
	}
	// 10000 times for repeat
	// fmt.Println(data, i)
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
	data := make([]int, 0, 1000)
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
