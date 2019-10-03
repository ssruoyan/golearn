package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1,2,2,1}, []int{2,2}))
}

/**
 * ========== 两数组交集 ==========
 * 给定两个数组，编写一个函数来计算它们的交集
 * 进阶思考：
 * 1. 如果两个有序数列如何优化
 * 2. 如果 num1 远远大于 num2 该如何调整算法
 * 3. 如果 num2 存储在磁盘上无法一次性读取该怎么办
 * ========== 解题思路 ==========
 * 使用 map 记录 nums1 数组中每个元素出现的次数。遍历 nums2 查询 map 中出现的数字并存储。
 */
func intersect(nums1 []int, nums2 []int) []int {
	var mp = make(map[int]int)
	var arr []int

	for _, v := range nums1 {
		mp[v] ++
	}

	for _, v := range nums2 {
		if mp[v] > 0 {
			arr = append(arr, v)
			mp[v] --
		}
	}

	return arr
}