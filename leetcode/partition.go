package main

import (
	"fmt"
	"github.com/ssruoyan/golearn/leetcode/tool"
)

func main() {
	fmt.Println(partition("ababbbabbaba"))
}
/**
 * ========== 分割回文字符串 ===========
 * 给定一个字符串 s，将字符串分割成多个子串，使得每个子串都是回文串。返回 s 所有可能的分割方案。
 * ========== 解题思路 ==========
 * 标准递归解题。所有字符串都是子串即一定存在子串[n:m] 使得 [0:n]和[m:len(s)]都可以完全分割成回文子串。
 * 再切割 [0:n]和[m:len(s)]
 */
func partition(s string) [][]string {
	var res [][]string
	slice(s, []string{}, &res)

	return res
}

func slice(s string, stack []string, res *[][]string) {
	// 如果已经解决完最后一个字符，则表示分割结束，存储分割结果
	if s == "" {
		*res = append(*res, stack)
	}
	for i := 1; i <= len(s); i++ {
		ss := s[:i]
		if ss == tool.Reverse(ss) {
			// 如果是回文串则存下来，slice 需要拷贝
			tmp := make([]string, len(stack))
			copy(tmp, stack)
			tmp = append(tmp, ss)
			// 继续分割剩余子串的回文串
			slice(s[i:], tmp, res)
		}
	}
}