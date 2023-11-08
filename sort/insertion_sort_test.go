package bubble_sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {

	nums := []int{4, 1, 3, 5, 2}

	InsertionSort(nums)

	// Check if the array is sorted
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Error("Array is not sorted.")
		}
	}
}
