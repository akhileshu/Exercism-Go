package cards


// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetSliceItemAtIndex[T any](slice []T, index int) (T, bool) {
	if index < 0 || index >= len(slice) {
		var zero T
		return zero, false
	}
	return slice[index], true
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	value, ok := GetSliceItemAtIndex(slice, index)
	if !ok {
		return -1
	}
	return value
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	_, ok := GetSliceItemAtIndex(slice, index)
	if !ok {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    _, ok := GetSliceItemAtIndex(slice, index)
    if !ok {
		return slice
	}
    updatedSlice :=append(slice[:index], slice[index+1:]...)
	return updatedSlice
}
