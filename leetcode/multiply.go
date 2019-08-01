package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiply("98765432", "23456789"))
}

/**
 * ========== 字符串相乘问题（大数乘法） ==========
 * 以字符串的形式表示的非负整数 num1 和 num2，计算两者的乘积用字符串表示。
 *
 * ========== 解题思路 ==========
 * 乘法即每一项相乘，然后相加。即： 455 * 34 = 34 * 400 + 34 * 50 + 34 * 5 = 15470
 * 详见：大数相乘 http://www2.lssh.tp.edu.tw/~hlf/class-1/lang-c/big_num3.htm
 */
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1 = Reverse(num1)
	num2 = Reverse(num2)

	len1 := len(num1)
	len2 := len(num2)

	var arr = make([]int, len1 + len2)

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {

			// ASCII 码计算差值
			n1 := num1[i] - '0'
			n2 := num2[j] - '0'

			arr[i + j] += int(n1) * int(n2)
		}
	}

	fmt.Println(arr)

	rtn := ""
	var curry int = 0

	for k := 0; k < len(arr) - 1; k ++ {

		value := arr[k] + curry

		if value  > 9 {
			curry = value / 10

			rtn += string(value - curry * 10 + '0')
		} else {
			curry = 0
			rtn += string(value + '0')
		}
	}
	if curry != 0 {
		rtn =  rtn + string(curry + '0')
	}

	return Reverse(rtn)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}