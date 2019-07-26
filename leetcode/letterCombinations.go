package main

import (
	"fmt"
)

func main() {
	m := letterCombinations("23")

	fmt.Println("result: ", m)
}

func letterCombinations(s string) []string {
	var rtn []string
	if len(s) <= 0 {
		return rtn
	}
	maps := map[string][]string{
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}

	dfs(s, "", &rtn, maps)

	return rtn
}

func dfs(s string, combination string, result *[]string, maps map[string][]string) {
	// 如果 combination 长度和 s 一样，表示已经拼接完成一个有效的结果，直接放入 result
	if len(combination) == len(s) {
		*result = append(*result, combination)
		return
	}

	// 获取combination 当前位置对应的 s 的值对应的 maps 中的数据
	str := maps[string(s[len(combination)])]

	for _, char := range str {
		combination = combination + string(char)

		dfs(s, combination, result, maps)

		combination = combination[0:len(combination) - 1]
	}
}