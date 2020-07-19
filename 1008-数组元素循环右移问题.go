package main

import "fmt"

/**
* @Author: hsq_roy
* @Date: 2020/7/15 2:07 下午
* @description：
 */
func move(num *[100]int,n int,m int) {
	// 最后一位暂存
	t := num[n-1]
	for i := n-1;i > 0;i-- {
		num[i] = num[i-1]
	}
	num[0] = t
}

func main() {
	var n, m int
	var num [100]int
	fmt.Scan(&n,&m)
	for i := 0;i < n;i++ {
		fmt.Scan(&num[i])
	}
	// 每次移动一位，移动n位，循环n次
	for i := 0;i < m;i++ {
		move(&num ,n ,m)
	}
	fmt.Printf("%d", num[0])
	for i := 1;i < n;i++ {
		fmt.Printf(" %d", num[i])
	}
}
