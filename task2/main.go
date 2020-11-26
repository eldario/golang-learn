package main

import (
	"fmt"
	"tasks/task2/pkg/list"
	"tasks/task2/pkg/sorter"
)

func main() {
	var choice string

	fmt.Printf("\rWhich package u want to use type `sorter` or `list`?")
	var _, err = fmt.Scan(&choice)

	if err != nil {
		fmt.Println("Wasted")
	}

	if choice == "list" {
		witList()
	} else {
		withSorter()
	}
}

func withSorter() {
	var (
		data = sorter.New()
		value,
		number int
	)

	fmt.Printf("\r How many objects do we have? ")

	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		fmt.Scan(&value)
		if value > 0 {
			data.Insert(value)
			continue
		}

		data.Remove(-value)
	}

	fmt.Println("Sorted data", data.GetItems())
	fmt.Printf("Max value is %d and min value is %d", data.GetMax(), data.GetMin())
}

func witList() {
	var (
		data = list.New()
		value,
		number int
	)

	fmt.Printf("\r How many objects do we have? ")

	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		fmt.Scan(&value)
		if value > 0 {
			data.Insert(value)
			continue
		}

		data.Remove(-value)
	}

	fmt.Println("Sorted data", data.GetItems())
	fmt.Printf("Max value is %d and min value is %d", data.GetMax(), data.GetMin())
}
