package tool

func Min(arr ...int) int {
	rtn := arr[0]

	for _, x := range arr {
		if x < rtn {
			rtn = x
		}
	}

	return rtn
}
