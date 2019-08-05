package main

import "fmt"

func main() {
	fmt.Println(findSerialSubString("aaaa", []string{"a","a"}))
}


/**
 * ========== 查找串联的子串问题 ===========
 * 给定一个字符串 s 和长度一直的单词 words，找出 s 中恰好由 words 所有单词串联成的子串起始位置。words 串联顺序不固定。
 *
 * ========== 问题解决思路 ===========
 *
 */

type Word struct {
	num int
	pos []int
}
func findSerialSubString(s string, words []string) []int {
	if words == nil || len(words) == 0 {
		return nil
	}

	wordsMaps := make(map[string]int)
	initialMaps := make(map[string]int)
	wordLen := len(words[0])
	substringLen := wordLen * len(words)

	for i := 0; i < len(words); i++ {
		initialMaps[words[i]] ++
		wordsMaps[words[i]]++
	}

	i := 0
	pos := 0
	var rtn []int
	match := false
	for i <= len(s) - wordLen {
		if !match && i + substringLen > len(s) {
			return rtn
		}
		str := s[i:i + wordLen]
		if wordsMaps[str] > 0 {
			if !match {
				match = true
				pos = i
			}
			wordsMaps[str] --
			i = i + wordLen

			if i - pos == substringLen {
				rtn = append(rtn, pos)
				match = false
				i = pos + 1

				for k, v := range initialMaps {
					wordsMaps[k] = v
				}
			}
		} else {
			if match {
				match = false
				i = pos + 1
			} else {
				i++
			}

			pos = 0

			for k, v := range initialMaps {
				wordsMaps[k] = v
			}
		}
	}

	return rtn
}
