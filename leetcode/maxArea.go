package main

import (
	"fmt"
)

func main() {
	area := maxArea([]int{1,8,6,2,5,4,8,3,7,9,11,9,7,2,5,8,4,7,6,13,12})

	fmt.Println(area)
}
/**
 * =========== 最大面积问题 ===========
 * n 个非负整数 a1， a2 ....... an，每个数代表坐标轴上的点（i，ai）。在坐标轴内画 n 条垂直的线，线的两个端点分别是 (i, 0) 和 (i, ai)
 * 找出两个坐标点 ai 和 aj 使得这两条先围成的区域面积最大。
 *
 * ========== 解题思路 ==========
 * 1. 在高度一致的情况下，两个点横向距离越远，面积越大
 * 2. 使用双向指针从两边不断缩小横向距离，存储最大面积值。
 * 3. 判断高度较小的指针并移动，计算。
 * 4. 在移动指针时，判断下一个位置高度是否大于当前位置的高度，如果不大于则跳过计算和比较的继续移动指针
 */
func maxArea(height []int) int {
	if height == nil || len(height) < 2 {
		return 0
	}

	i, j, area := 0, len(height) - 1, 0

	for i < j {
		h := height[i]

		if height[i] > height[j] { h = height[j] }

		newArea := (j - i) * h

		if newArea > area {
			area = newArea
		}

		if height[i] < height[j] {
			for i < j && height[i] > height[i + 1] {
				i ++
			}
			i ++
		} else {
			for i < j && height[j] > height[j - 1] {
				j --
			}
			j --
		}
	}

	return area
}

