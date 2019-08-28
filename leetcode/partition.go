package main

import (
	"fmt"
	"github.com/ssruoyan/golearn/leetcode/tool"
)

func main() {
	fmt.Println(partition("aab"))
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

func slice(s string, tmp []string, res *[][]string) {
	if s == "" {
		*res = append(*res, tmp)
	}
	for i := 0; i < len(s); i++ {
		sl := s[0:i]

		if sl == tool.Reverse(sl) {
			nt := append(tmp, s[:i])
			slice(s[i:], nt, res)
		}
	}
}