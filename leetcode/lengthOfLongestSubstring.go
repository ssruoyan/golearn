package main

import (
	"fmt"
	"github.com/ssruoyan/golearn/leetcode/tool"
)

func main() {
	m := lengthOfLongestSubstring("abasdaser4rswdfadbfs")

	fmt.Printf("result:  %d", m)
}

/**
 * ========== 最长字符子串问题 ==========
 * 计算字符串中不重复字符的最长子串
 *
 * ========== 解题思路 ===========
 * 1. 计算不重复的子串，即从存在字符 si 和 sj，si 到 sj 中没有重复字符，且长度最长。
 * 2. 遍历整个字符串，使用 map 记录字符串出现的位置。（因为 map 的 key 不会重复，即可以用来判断上次是否出现过字符以及出现的位置）
 * 3. 使用 i 作为起始位置（初始值为0），j 作为结束位置。j - i 即为子串长度
 * 4. 遍历字符串，如果没有重复字符串出现则 j 一直累加。
 * 5. 如果出现重复字符，并且重复字符上次出现位置为 k， 则初始位置变成重复字符的下一位 i = k + 1；继续遍历直到结束。
 */
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// maps 记录字符的出现位置
	// rtn 查询结果
	// start 记录子串的开始位置
	maps, rtn, start := make(map[string]int), 0, 0

	for i := 0; i < len(s); i++ {
		str := string(s[i])

		if _, ok := maps[str]; ok {
			start = tool.Max(start, maps[str] + 1)
			rtn = tool.Max(rtn, i - start + 1)

			// 子串的长度已经大于剩余计算最大值，则直接返回
			if rtn > (len(s) - start) {
				return rtn
			}

			maps[str] = i
		} else {
			maps[str] = i

			rtn = tool.Max(rtn, i - start + 1)
		}
	}

	return rtn
}