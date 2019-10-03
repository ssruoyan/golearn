package main

import "fmt"

func main() {
	fmt.Println(isAnagram("ab", "a"))
}

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