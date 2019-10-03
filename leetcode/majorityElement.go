package main

func main() {

}

/**
 * ========== 众数问题 ==========
 * 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素
 * ========== 问题解法 ==========
 * 使用 Map 记录每个数字的出现次数，出现一次 +1。如果次数超过半数则返回数字。
 */
func majorityElement(nums []int) int {
	var length = len(nums)
	var threshold = length / 2
	var mp = make(map[int]int)

	for i := 0; i < length; i ++ {
		num := nums[i]
		mp[num]++
		if mp[num] > threshold {
			return nums[i]
		}
	}

	return -1
}
