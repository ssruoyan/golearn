package main

import (
	"fmt"
	"github.com/ssruoyan/golearn/leetcode/tool"
)

func main() {
	d := minDistance("memberofthisclass", "mes")

	fmt.Println("distance: ", d)
}

/**
 * ========== 最短距离计算问题 ==========
 * 问题： 给定两个单词 word1 和 word2，计算出 word1 转换成 word2 所使用的最少操作次数。可以进行一下操作：
 * 1. 删除一个字符  2. 插入一个字符 3. 替换一个字符
 *
 * 示例：horse => ros   结果：3
 * 1. h 替换成 r  2. 删除 r 3. 删除 e
 */

/**
 * ========== 解题思路 ==========
 * 参考由少到多的递归解题思路
 *
 * 1. 假设 D[i][j] 表示 word1 的前 i 个字符到 word2 的前 j 个字符的操作距离。则可表示为 D[len(word1)][len(word2)]
 * 2. 从上一个状态到最后一步 D[i][j] 只有三种可能。插入，删除，替换（如果相等则不需要替换操作）:
 *   1）word1 插入一个与 word2 最后字母相同的字母，这时 word1 和 word2 相等。
 *      那么word1 => word2去掉最后一个字符的距离 D[i][j-1] + 插入一个字符的操作 1 即可得D[i][j]。即D[i][j] = D[i][j-1] + 1
 *   2）删除 word1 最后一个字符到 word2 的编辑距离加上删除一个字符的操作。 即 D[i][j] = D[i - 1][j] + 1
 *   3）把 word1 的最后一个字符替换成 word2 的最后一个字符，两者相等。即 D[i][j] = D[i-1][j-1] + 1。
 *      这时如果最后一个字符相同则D[i][j] = D[i-1][j-1]，则不需要替换。
 * 3. 然后取三种情况的最小操作距离
 */

func minDistance(s1 string, s2 string) int {

	s1Len := len(s1)
	s2Len := len(s2)

	D := tool.Make2DArr(s1Len + 1, s2Len + 1)

	for i := 0; i <= s1Len; i++ {
		for j := 0; j <= s2Len; j ++ {
			if i == 0 {
				D[i][j] = j
			} else if j == 0 {
				D[i][j] = i
			} else {
				if string(s1[i - 1]) == string(s2[j - 1]) {
					D[i][j] = D[i - 1][j - 1]
				} else {
					D[i][j] = tool.Min(D[i][j-1], D[i-1][j-1], D[i - 1][j]) + 1
				}
			}
		}
	}

	return D[s1Len][s2Len]
}