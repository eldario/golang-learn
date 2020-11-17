package main

import (
	"fmt"
)

func main() {
	var (
		data []int
		value,
		number int
	)

	fmt.Printf("\r How many objects do we have? ")

	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		fmt.Scan(&value)
		if value > 0 {
			insert(&data, value)
		} else {
			data = remove(data, -value)
		}
	}

	fmt.Println("Sorted data", data)
}

/**
 * Add a value in array.
 */
func insert(data *[]int, value int) []int {
	*data = append(*data, value)

	return *data
}

/**
 * Remove a value from array.
 */
func remove(data []int, value int) []int {
	for key, v := range data {
		if v == value {
			copy(data[key:], data[key+1:])
			data = data[:len(data)-1]
		}
	}

	return data
}

/**
 * Sort an array.
 */
func sort(data []int) []int {
	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}

	return data
}
