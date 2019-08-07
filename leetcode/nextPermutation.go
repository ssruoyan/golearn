package main

import "fmt"

func main() {
	nums := []int{4,3,2,1}

	nextPermutation(nums)

	fmt.Println(nums)
}

/**
 * =========== 下一个排列问题 ==========
 * 实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
 * =========== 解题思路 ==========
 * 1. 需要获获取一个比目前排列数大的排列，需要查找一个需要变动的元素。即满足以下条件：
 *    1）元素靠右。即需要替换的元素尽可能保持小进位。即从右侧查询
 *    2）元素右侧存在比该元素大的数字，即满足低位向高位替换。如：54321不满足；54123满足3>2
 * 2. 查找到该元素后，用右侧大于该元素的最小元素，替换该元素。
 * 3. 由于右侧查询满足左 > 右，因此替换完成之后需要翻转元素右侧的剩余数据，得到最终结果。
 */

func nextPermutation(nums []int)  {
	if nums == nil {
		return
	}

	length := len(nums)

	if length == 0 {
		return
	}

	j := length - 2


	for j > 0 && nums[j] >= nums[j + 1] {
		j --
	}

	if j >= 0 {
		for i := length - 1; i > j; i-- {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				reverse(&nums, j + 1, length - 1)

				return
			}
		}
	}

	reverse(&nums, 0, length - 1)
}
func reverse(nums *[]int, start int, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
	}
}