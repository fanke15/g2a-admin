package basic

import (
	"crypto/rand"
	"math/big"
)

// 生成数字范围内的随机数
func GetRandNum(max int) int {
	num, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(num.Int64() + One)
}

// 获取未重复的随机数
func GetUniqueNum(max int, data []int) int {
	var val int
	for i := Zero; i < Hundred; i++ {
		val = GetRandNum(max)
		status := true
		for _, v := range data {
			if v == val {
				status = false
				break
			}
		}
		if status {
			break
		}
	}
	return val
}

// 从一组数据中获取大于给定值得随机数
func GetGreaterNum(num int, data []int) int {
	var (
		filter = make([]int, Zero)
		res    int
	)
	for _, v := range data {
		if v > num {
			filter = append(filter, v)
		}
	}
	if len(filter) > Zero {
		return filter[GetRandNum(len(filter))-One]
	}
	return res
}

// 判断数字是否在集合内
func InIntSlice(num int, data []int) bool {
	for _, v := range data {
		if num == v {
			return true
		}
	}
	return false
}
