package main

import (
	"fmt"
	"math"
	"math/big"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/17 9:44 上午
* @description：大整数运算
 */
func main() {
	var n int
	a := big.NewInt(math.MaxInt64)
	b := big.NewInt(math.MaxInt64)
	c := big.NewInt(math.MaxInt64)
	res := big.NewInt(math.MaxInt64)
	var str1, str2, str3 string
	fmt.Scan(&n)
	for i := 0;i < n;i++{
		res.SetString("0", 10)
		fmt.Scan(&str1, &str2, &str3)
		a.SetString(str1, 10)
		b.SetString(str2, 10)
		c.SetString(str3, 10)
		if res.Add(a, b).Cmp(c) > 0 {
			fmt.Printf("Case #%d: true\n", i+1)
		}else{
			fmt.Printf("Case #%d: false\n", i+1)
		}
	}

}
