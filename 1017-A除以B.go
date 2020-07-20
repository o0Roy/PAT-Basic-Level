package main

import (
	"fmt"
	"math/big"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/20 12:09 上午
* @description：
 */
func main() {
	A := new(big.Int)
	B := new(big.Int)
	Q := new(big.Int)
	R := new(big.Int)
	var a, b string
	fmt.Scan(&a, &b)
	A.SetString(a, 10)
	B.SetString(b, 10)
	// 大整数求余数
	Q.DivMod(A, B, R)
	fmt.Println(Q, R)
}
