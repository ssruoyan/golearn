package main

import (
	"fmt"
)
func main() {
	arr := threeSum([]int{1, 0, -1, 3, 2, 9, -2, -4, 6}, 0)
	fmt.Println(arr)
}


/**
 * ========== 三数之和问题 ==========
 * 给定一个包含 n 个整数的数组，判断是否存在三个元素 a，b，c 使得 a + b + c = m ？找出所有满足的三元组
 *
 * ========== 解题思路 ==========
 * 1. 取出一个数 a 作为基准，则只需要计算两个数 b + c = m - a
 * 2. 先对数组进行排序，方便循环过程中累加的值与基准值的大小，如果小于基准值则往升序方向偏移，大于则往降序方向偏移。
 * 3. 循环遍历整个数组，只需要计算基准数后面是否存在满足条件的数字。
 * 4. 每次循环判断当前值是否与上一个值相同，如果相同则跳过进入下一步
 */
func threeSum(nums []int, m int) [][]int {
	if nums == nil || len(nums) < 3 {
		return nil
	}

	sort(nums)

	var rtn = make([][]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		// 保证不重复出现
		if i == 0 || nums[i] != nums[i - 1] {
			a := nums[i]
			l := i + 1
			r := len(nums) - 1

			for l < r {
				if a+nums[l]+nums[r] == m {
					rtn = append(rtn, []int{a, nums[l], nums[r]})
					for l < r && nums[l] == nums[l + 1] {
						l ++
					}
					for r > l && nums[r] == nums[r - 1] {
						r --
					}
					l++
					r--
				} else if a+nums[l]+nums[r] > m {
					r --
				} else {
					l ++
				}
			}
		}
	}

	return rtn
}

func sort(nums []int) {
	mid, i := nums[0], 1
	head, tail := 0, len(nums) - 1

	for head < tail {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail --
		} else {
			nums[head], nums[i] = nums[i], nums[head]
			head ++
			i ++
		}
	}
	nums[head] =  mid

	if head > 0 {
		sort(nums[:head])
	}
	if head < len(nums) - 1 {
		sort(nums[head+1:])
	}
}
