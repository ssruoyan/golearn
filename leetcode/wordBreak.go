package main

import "fmt"

func main() {
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "an", "cat"}))
}

/**
 * ========== 单词分割问题 ==========
 * 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。
 * PS： 字典中的单词可以重复出现
 * ========== 解题思路 ===========
 * 动态规划子问题解题。使用 DP[i] 表示 s 中第 i 个位置是否可以被字典拆分。
 * 对于任意子串 s[:i]，如果可以被拆分，肯定存在 DP[j] = true，s[j:i] 包含在字典中 或者 j = 0
 */
func wordBreak(s string, wordDict []string ) bool {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}

	length := len(s)
	DP := make([]bool, length + 1)
	DP[0] = true

	for i := 1; i <= length; i ++ {
		for j := 0; j < i; j ++ {
			if DP[j] && dict[s[j:i]] == true {
				DP[i] = true
			}
		}
	}

	return DP[length]
}