package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)
	strIn := fmt.Sprintf("%04d", num) // 输入即要拼接成四位数
	c := 0
	for {
		aStr := []byte(strIn)
		bStr := []byte(strIn)
		// 非递减排序
		sort.Slice(aStr, func(i, j int) bool {
			return aStr[i] >= aStr[j]
		})
		sort.Slice(bStr, func(i, j int) bool {
			return bStr[i] <= bStr[j]
		})
		// 转成数字
		a, _ := strconv.Atoi(string(aStr))
		b, _ := strconv.Atoi(string(bStr))
		c = a - b
		fmt.Printf("%04d - %04d = %04d\n", a, b, c)
		strIn = fmt.Sprintf("%04d", c)
		if c == 0 || c == 6174 {
			break
		}
	}
}
