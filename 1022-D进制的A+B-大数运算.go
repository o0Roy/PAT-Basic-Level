package main

import (
	"fmt"
	"math/big"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/27 8:05 下午
* @description：直接大数整除，然后字符串拼接
 */
func main() {
	a := new(big.Int) // A
	b := new(big.Int) // B
	d := new(big.Int) // D
	zero := new(big.Int) // 大数 0
	sum := new(big.Int) // A + b
	mod := new(big.Int) // 余数 mod
	res := new(big.Int)
	var strA, strB, strC, resStr string
	fmt.Scanf("%s %s %s", &strA, &strB, &strC)
	zero.SetString("0", 10)
	// 大数初始化
	a.SetString(strA, 10) // A
	b.SetString(strB, 10) // B
	d.SetString(strC, 10) // D
	sum.Add(a, b)
	for {
		res.DivMod(sum, d, mod) // (A+B)/D .... mod
		resStr = mod.String() + resStr // 结果拼接
		if zero.Cmp(res) == 0 {
			break
		}
		sum.Div(sum, d) // sum = sum / D
	}
	fmt.Printf("%s", resStr)
}
