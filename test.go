package main

import "fmt"

/**
* @Author: hsq_roy
* @Date: 2020/7/23 11:31 下午
* @description：
 */
func main() {
	num := 579
	for {
		num = num / 8
		fmt.Println(num)
		if num == 0 {
			break
		}
	}
	stack
}
