package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addStr("999", "998"))
}

/**
 *  字符串相乘。
 * 以字符串的形式表示的非负整数 num1 和 num2，计算两者的乘积用字符串表示。
 */
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	len1 := len(num1)
	len2 := len(num2)

	var nums string

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			pos1 := len1 - 1 - i
			pos2 := len2 - 1 - j
			n1, _ := strconv.Atoi(string(num1[i]))
			n2, _ := strconv.Atoi(string(num2[j]))
			st := strconv.Itoa(n1 * n2)
			if nums != "" {
				nums = addStr(nums, paddingZero(st, pos1 + pos2))
			} else {
				nums = paddingZero(st, pos1 + pos2)
			}
		}
	}

	return nums
}

func paddingZero(str string, num int) string {
	for i := 0; i < num; i ++ {
		str = str + "0"
	}

	return str
}
// 实现大数相加
func addStr(num1 string, num2 string) string {
	curry := 0
	len1 := len(num1)
	len2 := len(num2)
	digit := 1
	rtn := ""

	for digit <= len1 || digit <= len2 {
		var d1 int
		var d2 int
		if len1 < digit {
			d1 = 0
		} else {
			d1, _ = strconv.Atoi(string(num1[len1 - digit]))
		}
		if len2 < digit {
			d2 = 0
		} else {
			d2, _ = strconv.Atoi(string(num2[len2 - digit]))
		}
		multi := d1 + d2
		rtn = strconv.Itoa((multi + curry) % 10) + rtn

		curry = (multi + curry) / 10
		digit ++
	}
	if curry != 0 {
		rtn = strconv.Itoa(curry) + rtn
	}

	return rtn
}

