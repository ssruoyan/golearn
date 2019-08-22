package main

import (
	"fmt"
	"github.com/ssruoyan/golearn/leetcode/tool"
)

func main() {
	fmt.Println(isPalindrome("0P"))
}

/**
 * ========== 鸡蛋掉落问题 ==========
 * K 个鸡蛋，N 层楼。假设每个鸡蛋都是一样的，存在 F 层满足 0 < F <= N，当超过 F 时鸡蛋从楼层落下都会摔碎，当小于 F 时鸡蛋都不会摔碎。
 * 每次移动，你可以取一个鸡蛋（如果你有完整的鸡蛋）并把它从任一楼层 X 扔下（满足 1 <= X <= N）
 * 需要确切的知道 F 的层数，最小移动的次数
 *
 * ========== 解题思路 ==========
 * 定义 移动次数 T；楼层数 N；鸡蛋个数 K
 * 假设移动的次数为 T，移动次数 dp[K][N] = dp[K - 1][N - 1]
 * 上一次移动为
 */

//func superEggDrop(K int, N int) int {
//	if N == 0 || K == 0 {
//		return 0
//	}
//
//	// t1 表示 T1(k - 1, i) 状态
//	t1 := make([]int, N + 1)
//
//	for i := 0; i <= N; i ++ {
//		t1[i] = i
//	}
//
//	// k = 1 时，层数与次数相等
//	for k := 2; k <= K; k++ {
//		// t2 表示 T2(k, i) 状态
//		t2 := make([] int, N+1)
//		x := 1
//		for i := 1; i <= N; i++ {
//			// 此处 t1[x - 1] => T1(k-1, x - 1)
//			// t2[i - x] => T2(K, i - x)
//			for x < i && tool.Max(t1[x-1], t2[i-x]) > tool.Max(t1[x], t2[i-x-1]) {
//				x++
//			}
//			t2[i] = tool.Max(t2[i-x], t2[x-1]) + 1
//		}
//
//		t1 = t2
//	}
//
//	return t1[N]
//}

func superEggDrop(K int, N int) int {
	if K == 0 || N == 0 {
		return 0
	}

	D := tool.Make2DArr(K + 1, N + 1)

	for n := 1; n <= N; n ++ {
		for k := 1; k <= K; k ++ {
			if K == 1 {
				D[k][n] = n
			} else if N == 1 {
				D[k][n] = 1
			} else {
				D[k][n] = D[k - 1][n - 1] + D[k][n - 1] + 1
			}

			if D[k][n] >= N {
				return n
			}
		}
	}
	return N
}
