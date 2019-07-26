package main

import (
	"fmt"
)

func main() {
	m := lengthOfLongestSubstring("abasdaser4rswdfas")

	fmt.Printf("result:  %d", m)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var maps = make(map[string]int) // 记录字符的出现位置
	var rtn = 0  // 查询结果
	var start = 0  // 记录子串的开始位置

	for i := 0; i < len(s); i++ {
		str := string(s[i])

		if _, ok := maps[str]; ok {
			start = max(start, maps[str] + 1)
			rtn = max(rtn, i - start + 1)
			maps[str] = i
		} else {
			maps[str] = i

			rtn = max(rtn, i - start + 1)
		}
	}

	return rtn
}

func max(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
