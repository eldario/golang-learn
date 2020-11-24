package sorter

type SortedArray struct {
	array []int
}
/**
 * Constructor
 */
func New() *SortedArray {
	return new(SortedArray)
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
func (data *SortedArray) GetItems() []int {
	for i := 0; i < len(data.array); i++ {
		for j := i; j < len(data.array); j++ {
			if data.array[i] > data.array[j] {
				data.array[i], data.array[j] = data.array[j], data.array[i]
			}
		}
	}

	return data.array
}

/**
 * Get max value from array.
 */
func (data *SortedArray) GetMax() int {
	var count = len(data.array)
	if count == 0 {
		return 0
	}

	return data.array[count-1]
}

/**
 * Get minimal values from array.
 */
func (data *SortedArray) GetMin() int {
	var count = len(data.array)
	if count == 0 {
		return 0
	}

	return data.array[0]
}

/**
 * Check equals of actual slice.
 */
func (data *SortedArray) Equals(actual []int) bool {
	if len(data.array) != len(actual) {
		return false
	}
	for key, value := range data.array {
		if value != actual[key] {

			return false
		}
	}

	return true
}