package main

import (
	"fmt"
)

type SortedArray struct {
	array []int
}

func main() {
	var (
		data = new(SortedArray)
		value,
		number int
	)

	fmt.Printf("\r How many objects do we have? ")

	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		fmt.Scan(&value)
		if value > 0 {
			data.Insert(value)
		} else {
			data.Remove(-value)
		}
	}

	fmt.Println("Sorted data", data.Sort())
}

/**
 * Add a value in array.
 */
func (data *SortedArray) Insert(value int) {
	data.array = append(data.array, value)
}

/**
 * Remove a value from array.
 */
func (data *SortedArray) Remove(value int) {
	for key, v := range data.array {
		if v == value {
			copy(data.array[key:], data.array[key+1:])
			data.array = data.array[:len(data.array)-1]
		}
	}
}

/**
 * Sort an array.
 */
func (data *SortedArray) Sort() []int {
	for i := 0; i < len(data.array); i++ {
		for j := i; j < len(data.array); j++ {
			if data.array[i] > data.array[j] {
				data.array[i], data.array[j] = data.array[j], data.array[i]
			}
		}
	}

	return data.array
}
