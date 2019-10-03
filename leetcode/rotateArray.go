package main

import (
	"fmt"
	"github.com/ssruoyan/golearn/leetcode/tool"
)

func main() {
	var nums = []int{1,2,3,4,5,6,7}
	rotateArray2(nums, 2)

	fmt.Println(nums)
}

/**
 * =========== 旋转数组问题 ==========
 * 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。要求三种以上解法
 * =========== 解题思路 ==========
 * 解法一：移动的本质即数据交换。每次移动都是将最后一个元素移动到第一个位置，总尾部遍历依次替换元素。
 * 解法二：使用 整体旋转 + 局部旋转 实现尾部子数组移动到前面。即
 * 1. 先翻转 数组 n，实现 n.....(n-k).....1
 * 2. 翻转 n....(n-k) => (n-k).....n
 * 3. 翻转 (n-k)...1 => 1...(n-k)
 * 解法三：迭代计算。即计算每一个元素的最终位置，然后放入到对应的位置上。需要记录被替换位置的元素。
 * 1. 从第一个位置开始 nums[0]，查询第一个元素的最终位置 k，并放入 nums[k]，同时临时存储num[k] 旧值。
 * 2. 从 k 个位置开始，继续查找最终位置循环计算。
 */
func rotateArray(nums []int, k int) {
	for i := 0; i < k; i ++ {
		for j := len(nums) - 1; j > 0; j -- {
			nums[j], nums[j - 1] = nums[j - 1], nums[j]
		}
	}
}

func rotateArray1(nums []int, k int) {
	n := len(nums)
	k = k % n

	tool.ReverseArray(nums, 0, n - 1)
	tool.ReverseArray(nums, 0, k - 1)
	tool.ReverseArray(nums, k, n - 1)
}

func rotateArray2(nums []int, k int) {
	var cur = 0 // 当前位置
	var tmp = nums[cur] // 临时存储
	var prev = -1 // 前一个位置
	var length = len(nums)
	var count = 0

	for count < length {
		// 获取当前位置需要填充的值和位置
		next := cur - k

		// 如果小于 0 则从后往前取值
		if next < 0 {
			next = next + length
		}

		// 如果上一个位置跟当前位置下一个位置一样
		if prev == next {
			nums[cur] = tmp
			cur = next + 1
		} else {
			tmp = nums[cur]
			nums[cur] = nums[next]
			cur = next
		}
		cur = next
		count ++
	}
}