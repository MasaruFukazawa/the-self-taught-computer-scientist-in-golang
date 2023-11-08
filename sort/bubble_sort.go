package bubble_sort

// BubbleSort sorts a slice of integers using the bubble sort algorithm.
func BubbleSort(nums []int) {

	for i := len(nums) - 1; i > 0; i-- {

		for j := 1; j <= i; j++ {

			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}

		}
	}

}
