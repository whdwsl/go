package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	if slice == nil {
		return 0, false
	}
	if len(slice) <= index || index < 0 {
		return 0, false
	}

	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice) <= index || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}

	return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length <= 0 {
		return make([]int, 0)
	}
	ret := make([]int, length)
	for b := 0; b < length; b++ {
		ret[b] = value
	}
	return ret
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if len(slice) <= index || index < 0 {
		return slice
	}

	if index == len(slice)-1 {
		return append(slice[:index])
	}

	ret := append(slice[:index], slice[index+1:]...)
	return ret
}
