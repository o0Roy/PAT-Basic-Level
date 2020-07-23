package main

import "fmt"

/**
* @Author: hsq_roy
* @Date: 2020/7/23 11:12 下午
* @description：
 */
func main() {
	var str string
	var num [10]int
	fmt.Scanf("%s", &str)
	for i := 0; i < len(str); i++ {
		num[str[i]-48]++  // 直接构造一个数组统计
	}
	for i := 0; i < len(num); i++ {
		if num[i] > 0 {
			fmt.Printf("%d:%d\n", i, num[i])
		}
	}
}
