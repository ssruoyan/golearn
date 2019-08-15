package main

import (
	"fmt"
	"github.com/ssruoyan/golearn/leetcode/tool"
	"golang.org/x/tools/container/intsets"
)

func main() {
	fmt.Println(maxProduct([]int{-2}))
}

/**
 * ========== 最大乘积问题 ==========
 */
func maxProduct(nums []int) int {
	imax, imin, max := 1, 1, intsets.MinInt
	for _, v := range nums {
		if v < 0 {
			imax, imin = imin, imax
		}
		imax = tool.Max(imax * v, v)
		imin = tool.Min(imin * v, v)

		max = tool.Max(imax, max)
	}

	return max
}