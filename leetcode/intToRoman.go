package main

import "fmt"

func main() {
	roman := intToRoman(2933)

	fmt.Println("result: ", roman)
}
/**
 * 整数转罗马数字。规则如下：
 * 1. 罗马数包含以下基本数字：1 => I  5 => V  10 => X  L => 50  C => 100  D => 500 M => 1000
 * 2. 小数字在大数字左边表示减，但是只存在以下几种情况：IV(4) IX(9) XL(40) XC(90) CD(400) CM(900)。
 *
 * 解题思路：先枚举出可以表示的简单数字的表示法[0, 1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000]
 * 从高位往低位依次计算，累加即可。
 */
func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1, 0}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I", ""}
	digit := 0
	var rtn string

	for num > 0 {
		times := num / nums[digit]
		num = num - nums[digit] * times

		for i := 0; i < times; i++ {
			rtn = rtn + romans[digit]
		}
		digit ++
	}

	return rtn
}