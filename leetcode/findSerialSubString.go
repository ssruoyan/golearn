package main

import "fmt"

func main() {
	fmt.Println(findSerialSubString("barfoothefoobarman", []string{"foo","bar"}))
}


/**
 * ========== 查找串联的子串问题 ===========
 * 给定一个字符串 s 和长度一直的单词 words，找出 s 中恰好由 words 所有单词串联成的子串起始位置。words 串联顺序不固定。
 *
 * ========== 问题解决思路 ===========
 * 1. 遍历字符串中的每个子串长度的字符，判断子串是否符合条件即 0 -> (sLen - subSLen) 中每个 subSLen 字符串的是否符合条件
 * 2. 判断子串是否符合条件。先用 maps 存储 words 的每个单词出现的次数比如 maps['foo'] => 1，记录每个单词长度 length
 * 3. 循环 words 的长度，依次在 maps 查询子串中的每个单词。仅当每个单词都出先且出现的次数与 maps 完全一致时，子串符合标准。
 */
func findSerialSubString(s string, words []string) []int {
	if words == nil {
		return []int{}
	}
	nums := len(words)

	if nums == 0 {
		return []int{}
	}
	length := len(words[0])

	if length == 0 {
		return []int{}
	}
	maps := make(map[string]int)
	substringLen := length * nums

	// 如果长度不够
	if len(s) < substringLen {
		return []int{}
	}

	for _, str := range words {
		maps[str] += 1
	}

	rtn := make([]int, 0, 30)

	for i := 0; i <= len(s) - substringLen; i++ {
		isSub(s[i:(i + substringLen)], maps, length, nums, i, &rtn)
	}

	return rtn
}

func isSub(s string, maps map[string]int, length int, nums int, pos int, rtn *[]int) {
	find := make(map[string]int)

	for i := 0; i < nums; i ++ {
		str := s[i * length:(i + 1) * length]
		record, ok := maps[str]
		fRecord, fOk := find[str]

		if !ok || (fOk && fRecord >= record) {
			return
		} else {
			find[str] ++
		}
	}

	*rtn = append(*rtn, pos)
}
