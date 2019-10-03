package main

import (
	"fmt"
	trie2 "github.com/ssruoyan/golearn/basic/data-structure/trie"
)

func main() {
	word := []string{"oath","pea","eat","rain"}
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}

	fmt.Println(findWord2(board, word))
}

/**
 * =========== 单词查找问题 ===========
 * 给定一个二维网格 board 和一个字典中的单词列表 words，找出所有同时在二维网格和字典中出现的单词。
 * =========== 解题思路 ==========
 */
func findWord2(board [][]byte, word []string) []string {
	trie := trie2.Trie{}

	for _, v := range word {
		trie.Insert(v)
	}
	var res []string
	var visited = make([][]bool, len(board), len(board[0]))

	for i := 0; i < len(board); i ++ {
		for j := 0; j < len(board[i]); j ++ {
			dfsWord(board, visited, i, j, trie, "", &res)
		}
	}

	return res
}

func dfsWord(board [][]byte, visited [][]bool, i int, j int, trie trie2.Trie, tmp string, res *[]string) {
	fmt.Println(i, j)
	if (i < 0 || i == len(board)) || (j < 0 || j == len(board[0])) {
		return
	}

	str := board[i][j]
	tmp = tmp + string(str)
	visited[i][j] = true

	if trie.Search(tmp) {
		*res = append(*res, tmp)
	}

	dfsWord(board, visited, i + 1, j, trie, tmp, res)
	dfsWord(board, visited, i, j + 1, trie, tmp, res)
	dfsWord(board, visited, i, j - 1, trie, tmp, res)
	dfsWord(board, visited, i - 1, j, trie, tmp, res)

	visited[i][j] = false
}
