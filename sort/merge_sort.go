package sort

func MergeSort(nums []int) {

	if len(nums) > 1 {

		var mid int = len(nums) / 2

		var left []int = nums[:mid]
		var right []int = nums[mid:]

		MergeSort(left)
		MergeSort(right)

		var left_i int = 0
		var right_i int = 0
		var nums_i int = 0

		for (left_i < len(left)) && (right_i < len(right)) {

			if left[left_i] <= right[right_i] {
				nums[nums_i] = left[left_i]
				left_i++
			} else {
				nums[nums_i] = right[right_i]
				right_i++
			}

			nums_i++

		}

		for left_i < len(left) {
			nums[nums_i] = left[left_i]
			left_i++
			nums_i++
		}

		for right_i < len(right) {
			nums[nums_i] = right[right_i]
			right_i++
			nums_i++
		}

	}

}
