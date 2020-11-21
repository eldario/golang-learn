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

type testMaxMinCase struct {
	array  []int
	result int
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
		data := SortedArray{c.array}
		data.Insert(c.value)
		if !isSame(c.result, data.array) {
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
		data := SortedArray{c.array}
		data.Remove(c.value)
		if !isSame(c.result, data.array) {
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
		data := SortedArray{c.array}
		if !isSame(c.result, data.Sort()) {
			t.Fail()
		}
	}
}

func TestGetMin(t *testing.T) {
	cases := []testMaxMinCase{
		{[]int{2, 3, 4, 5, 6, 7, 8, 9}, 2},
		{[]int{12, 2, 4, 5, 1, 4, 6, 9, 5}, 1},
		{[]int{9, 6, 5, 4, 8, 6, 5, 7, 6, 5}, 4},
		{[]int{3, 1, 2, 5, 6, 4, 8, 9}, 1},
		{[]int{6, 4, 9, 8, 9, 8, 5, 3, 7, 4}, 3},
		{[]int{}, 0},
	}
	for _, c := range cases {
		data := SortedArray{c.array}
		if c.result != data.GetMin() {
			t.Fail()
		}
	}
}

func TestGetMax(t *testing.T) {
	cases := []testMaxMinCase{
		{[]int{2, 3, 4, 5, 6, 7, 8, 9}, 9},
		{[]int{12, 2, 4, 5, 1, 4, 6, 9, 5}, 12},
		{[]int{9, 6, 5, 14, 8, 6, 5, 72, 6, 5}, 72},
		{[]int{32, 1, 2, 50, 6, 44, 8, 19}, 50},
		{[]int{6, 4, 49, 8, 9, 82, 5, 3, 72, 4}, 82},
		{[]int{}, 0},
	}
	for _, c := range cases {
		data := SortedArray{c.array}
		if c.result != data.GetMax() {
			t.Fail()
		}
	}
}

func BenchmarkInsert(b *testing.B) {
	rand.Seed(1)
	var data = new(SortedArray)
	x := rand.Intn(10)
	i := 0
	for ; i < b.N; i++ {
		data.Insert(x * i)
	}
	// 10000 times for repeat
	// fmt.Println(data, i)
}

func BenchmarkRemove(b *testing.B) {
	rand.Seed(1)
	var data = new(SortedArray)
	x := rand.Intn(100)
	for i := 0; i < b.N; i++ {
		data.Remove(x)
	}
}

func BenchmarkSort(b *testing.B) {
	rand.Seed(1)
	var data = SortedArray{make([]int, 0, 1000)}
	for i := 0; i < b.N; i++ {
		data.Sort()
	}
}

func BenchmarkGetMax(b *testing.B) {
	rand.Seed(1)
	var data = SortedArray{make([]int, 0, 1000)}
	for i := 0; i < b.N; i++ {
		data.GetMax()
	}
}
func BenchmarkGetMin(b *testing.B) {
	rand.Seed(1)
	var data = SortedArray{make([]int, 0, 1000)}
	for i := 0; i < b.N; i++ {
		data.GetMin()
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
