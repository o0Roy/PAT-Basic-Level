package main

import (
	"fmt"
	"sort"
)
//0. 根据出现频次，构造数组，然后升序排序即可
//1. 包含0，则取最小值放在第一位
//2. 不包含0，则直接排序（升序）即可
func main() {
	var arr []int
	var isZero bool
	var min, red int
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &red)
		for j := 0;j < red;j++ {
			arr = append(arr, i) // 根据出现次数，构造数组
		}
		// 判断是否出现 0
		if i == 0 && red != 0 {
			isZero = true
		}
	}
	// 排序（升序）
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	min = arr[0]
	// 包含 0 则找到最小值
	if isZero {
		for i := 0; i < len(arr); i++ {
			if arr[i] > min {
				arr[i], arr[0] = arr[0], arr[i]
				break
			}
		}
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
	}
	//
	fmt.Println()
}