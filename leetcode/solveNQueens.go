package main

import (
	"fmt"
)

func main() {
	fmt.Println(solveNQueens(5))
}

/**
 * ========= N 皇后问题 ==========
 * 将 N 个皇后摆在 N x N 的方格上，保证彼此皇后之间不能相互攻击
 * ========= 解题思路 ==========
 * 经典回溯法解题。
 * 1）对于所有每一行、每一列、每一对角线上都不存在两个皇后。
	即对于任意两个皇后 queen1[a1,b1], queen2[a2,b2]
	a1 != a2
	b1 != b2
	b1 - a1 != b2 - a2
    b1 + a1 != b2 + a2
 * 2）
 */
func solveNQueens(n int) [][]string {
	// 有皇后的行
	rowUsed := make(map[int]bool)
	// 有皇后的列
	colUsed := make(map[int]bool)
	// 有皇后的反对角
	oppoAngleUsed := make(map[int]bool)
	// 有皇后的正对角
	posAngleUsed := make(map[int]bool)
	// 存储结果
	result := [][]string{}
	list := []string{}

	searchQueens(rowUsed, colUsed, oppoAngleUsed, posAngleUsed, &result, &list, n)

	return result
}

func searchQueens(rowUsed map[int]bool, colUsed map[int]bool, oppoAngleUsed map[int]bool, posAngleUsed map[int]bool, result *[][]string, list *[]string, n int) {
	for i := 0; i < n; i ++ {
		if i == len(*list) {
			for j := 0; j < n; j++ {
				// 判断当前位置是否可用
				canUse := !(rowUsed[i] == true || colUsed[j] == true || posAngleUsed[j-i] == true || oppoAngleUsed[i+j] == true)
				if canUse {
					str := combineString(j, n)
					*list = append(*list, str)
					rowUsed[i] = true
					colUsed[j] = true
					posAngleUsed[j-i] = true
					oppoAngleUsed[i+j] = true

					if len(*list) == n {
						arr := copySlice(*list)

						*result = append(*result, arr)
					}

					searchQueens(rowUsed, colUsed, oppoAngleUsed, posAngleUsed, result, list, n)

					*list = (*list)[0:len(*list) - 1]
					rowUsed[i] = false
					colUsed[j] = false
					posAngleUsed[j-i] = false
					oppoAngleUsed[i+j] = false

				}
			}
		}
	}
}

func copySlice(slice []T) []T {
	arr := make([]string, len(*list))

	for k := 0; k < len(slice); k ++ {
		arr[k] = slice[k]
	}

	return arr
}

func combineString(pos int, n int) string {
	str := ""
	for i := 0; i < n; i ++ {
		if i == pos {
			str += "Q"
		} else {
			str += "."
		}
	}

	return str
}