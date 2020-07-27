package main

import (
	"fmt"
	"io"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/16 1:20 下午
* @description：
 */
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if b == 0 {
		// 由于以指数递降给出，零多项式即为结束。
		fmt.Print("0 0")
		return
	} else {
		fmt.Printf("%d %d", a*b, b-1)
	}
	for {
		/// 循环读入致结束
		_, err := fmt.Scan(&a, &b)
		if err == io.EOF{
			break
		}
		if b != 0 {
			fmt.Printf(" %d %d", a*b, b-1)
		}
	}
}
