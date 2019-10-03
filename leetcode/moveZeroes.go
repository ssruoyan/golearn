package main

import "fmt"

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}

/**
 * ========== 移动零问题 ==========
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序
 * ========== 解题思路 ==========
 * 思路一： 结果推导
 * 1. 遍历数组
 * 2. 计算每个非零元素在结果数组中的位置。可得： 最终位置 = 当前位置 - 当前位置前 0 的个数
 * 3. 直接设置对应的值
 * 4. 末尾补零
 * 思路二：标记替换
 * 1. 遍历数组
 * 2. 记录最近的非零指针的位置 lastNoneZeroAt，即值为 0 但是应该被最近的数字替换的位置。
 * 3. 当查询到最近的非零数字后，替换掉 lastNoneZeroAt 位置的零并 + 1
 */
func moveZeroes(nums []int) {
	count := 0
	ln := len(nums)

	for i := 0; i < ln; i ++ {
		if nums[i] == 0 {
			count ++
		} else {
			if count > 0 {
				nums[i - count] = nums[i]
			}
		}
	}
	for j := ln - count; j < ln; j ++ {
		nums[j] = 0
	}
}

func moveZeroes2(nums []int) {
	l := 0
	for i := 0; i < ln; i ++ {
		if nums[i] != 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l ++
		}
	}
}