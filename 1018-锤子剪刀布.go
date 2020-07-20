package main

import (
	"fmt"
	"strings"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/20 12:15 上午
* @description：最后一个测试点运行时间过不了，主要看 C++ 版本
 */
// 判断剪刀石头
func CJB(a, b uint8) int {

	//B = 66  C= 67 J = 74
	//1 赢 0 平 -1 负
	switch a {
	case 67: // 锤子
		if b == 74 {
			return 1
		} else if b == 67 {
			return 0
		} else {
			return -1
		}
	case 74: // 剪刀
		if b == 66 {
			return 1
		} else if b == 74 {
			return 0
		} else {
			return -1
		}
	case 66: // 布
		if b == 67 {
			return 1
		} else if b == 66 {
			return 0
		} else {
			return -1
		}
	default:
		return 0
	}
}

func main() {
	var N int
	var a, b string
	str := "BCJ"
	//  赢  平  负  最多
	var A3, A2, A1 int
	var flag [2][3]int
	fmt.Scanf("%d", &N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%s %s", &a, &b)
		switch CJB(a[0], b[0]) {
		case 1:
			A3++
			flag[0][strings.Index(str, a)]++
		case 0:
			A2++
		case -1:
			A1++
			flag[1][strings.Index(str, b)]++
		}
	}
	fmt.Printf("%d %d %d\n", A3, A2, A1)
	fmt.Printf("%d %d %d\n", A1, A2, A3)

	var A0, B0, indexA, indexB int
	for i := 0; i < 3; i++ {
		if A0 < flag[0][i] {
			A0 = flag[0][i]
			indexA = i
		}
		if B0 < flag[1][i] {
			B0 = flag[1][i]
			indexB = i
		}
	}
	fmt.Printf("%s %s", string(str[indexA]), string(str[indexB]))
}
