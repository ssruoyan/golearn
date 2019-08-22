package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("asddsa"))
}

func isPalindrome(s string) bool {
	length := len(s)
	i, j := 0, length - 1
	s = strings.ToLower(s)

	for i < j {
		for !isValidChar()(s[i]) {
			i ++
			if i == length {
				return true
			}
		}
		for !isValidChar()(s[j]) {
			j --
		}

		if s[i] != s[j] {
			return false
		}
		i ++
		j --
	}

	return true
}


func isValidChar() func(str uint8) bool {

	return func(str uint8) bool {
		return (str >= 'a' && str <= 'z') || (str >= 'A' && str <= 'Z') || (str >= '0' && str <= '9')
	}
}