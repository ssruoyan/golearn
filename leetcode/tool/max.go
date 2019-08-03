package tool

func Max(arr ...int) int {
	rtn := arr[0]

	for _, x := range arr {
		if x > rtn {
			rtn = x
		}
	}

	return rtn
}