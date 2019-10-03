package tool

func ReverseArray(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]

		start ++
		end --
	}
}