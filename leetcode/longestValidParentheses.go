package main

import "fmt"

func main() {
	d := longestValidParentheses(")()())")

	fmt.Println(d)

}

/**
 * ========== 最长可用括号计算问题 ==========
 * 给定一个只包含`(` 和 `)`的字符串，找出最长包含有效括号的子串
 * ========== 解题思路 ==========
 * 1. 有效括号即连续的`(`和`)`组成的字符串。`(` 个数和`)` 个数一致。
 * 2. 从左开始遍历，使用 l，r 分别记录 '(' 和 ')' 的个数，pl 记录子串长度。
 * 3. 当 l = r 时，则组成一个封闭的有效子串。记录子串的长度，如果子串长度增加则修改。
 * 4. 继续遍历当 r > l 时，表示多了 ')' 字符，重置 i, j 继续遍历。
 * 5. 由于左侧遍历只记录了 '(' 开头的子串，再从右侧重新遍历一次以 ')' 开头的子串。
 */

func longestValidParentheses(s string) int {
	length := len(s)

	if length < 2 {
		return 0
	}
	byteLeftParentheses := byte('(')
	byteRightParentheses := byte(')')
	l, r, pl := 0, 0, 0

	for i := 0; i < length; i++ {
		if s[i] == byteLeftParentheses {
			l ++
		} else {
			r ++
			if l == r {
				if r * 2 > pl {
					pl = r * 2
				}
			} else if r >= l {
				l = 0
				r = 0
			}
		}
	}

	l, r = 0, 0
	for i := length - 1; i > 0; i-- {
		if s[i] == byteRightParentheses {
			r ++
		} else {
			l ++
			if l == r {
				if l * 2 > pl {
					pl = l * 2
				}
			} else if r <= l{
				l = 0
				r = 0
			}
		}
	}

	return pl
}