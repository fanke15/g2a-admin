package basic

import "strconv"

// 过滤重复
func RemoveRepeat(filter, origin [][]int) [][]int {
	var (
		filterMap = map[string][]int{}
		res       = make([][]int, Zero)
	)
	for _, v := range filter {
		filterMap[string(Marshal(v))] = v
	}

	for _, v := range origin {
		if _, ok := filterMap[string(Marshal(v))]; !ok {
			res = append(res, v)
		}
	}
	return res
}

// string->int
func SliceStrToInt(data []string) []int {
	var res = make([]int, Zero)
	for _, v := range data {
		n, _ := strconv.Atoi(v)
		res = append(res, n)
	}
	return res
}

// int->string
func SliceIntToStr(data []int) []string {
	var res = make([]string, Zero)
	for _, v := range data {
		res = append(res, strconv.Itoa(v))
	}
	return res
}

//  去重
func RemoveRepeatStr(filter, origin []string) []string {
	var (
		filterMap = make(map[string]string)
		res       = make([]string, Zero)
	)
	for _, v := range filter {
		filterMap[v] = v
	}

	for _, v := range origin {
		if _, ok := filterMap[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}

//  去重
func RemoveRepeatInt(filter, origin []int) []int {
	var (
		filterMap = make(map[int]int)
		res       = make([]int, Zero)
	)
	for _, v := range filter {
		filterMap[v] = v
	}

	for _, v := range origin {
		if _, ok := filterMap[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
