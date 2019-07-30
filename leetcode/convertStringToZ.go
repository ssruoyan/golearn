package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := convertStringToZ("LEETCODEISHIRING", 4)

	fmt.Println("result: ", str)
}

func convertStringToZ(s string, num int) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 || num == 1 {
		return s
	}
	arr := make([][]string, num)

	for index, x := range s {
		str := string(x)

		remain := index % (2 * num - 2) + 1

		if remain <= num {
			arr[remain - 1] = append(arr[remain - 1], str)
		} else {
			arr[2 * num - remain - 1] = append(arr[2 * num - remain - 1], str)
		}
	}

	var rtn []string

	for i := 0; i < num; i ++ {
		rtn = append(rtn, arr[i]...)
	}

	return strings.Join(rtn, "")
}
