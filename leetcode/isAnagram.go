package main

import "fmt"

func main() {
	fmt.Println(isAnagram("ab", "a"))
}

/**
 * =========== 字母异位词 ==========
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 * =========== 解题思路 ==========
 * 使用 Map 记录 s 字符串中每个词的数量。
 * 遍历 t 反射删除 Map 中每个词的数量。如果最终 Map 中所有字母数量都为 0 则两个词由相同字母组成。
 */
func isAnagram(s string, t string) bool {
	iMap := make(map[int32]int)

	for _, v := range s {
		iMap[v - 'a']++
	}

	for _, v := range t {
		iMap[v - 'a'] --
	}

	for _, v := range iMap {
		if v != 0 {
			return false
		}
	}
	
	return true
}