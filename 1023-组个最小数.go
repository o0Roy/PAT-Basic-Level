package main

import (
	"fmt"
	"sort"
)

//1.包含0，则取最小值放在第一位
//2.不包含0，则直接排序（升序）即可
func main() {
	var arr [10]int
	var isZero bool
	var min int
	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &arr[i])
		if arr[i] == 0 {
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
		for i := 0; i < 10; i++ {
			if arr[i] > min {
				arr[i], arr[0] = arr[0], arr[i]
				break
			}
		}
	}
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", arr[i])

	}
	//
	fmt.Println()
}
