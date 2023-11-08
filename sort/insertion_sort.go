package bubble_sort

func InsertionSort(nums []int) {

	var nums_len int = len(nums)

	for i := 1; i < nums_len; i++ {

		target_num := nums[i]

		j := i

		for ; j > 0; j-- {

			if nums[j-1] > target_num {
				nums[j] = nums[j-1]
			} else {
				break
			}

		}

		nums[j] = target_num

	}

}
