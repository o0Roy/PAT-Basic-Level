package main

import (
	"fmt"
	"math/big"
	"strings"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/19 6:33 下午
* @description：
 */
func main() {
	var a, da, b, db string
	resA := new(big.Int)
	resB := new(big.Int)
	res := new(big.Int)
	fmt.Scan(&a, &da, &b, &db)
	ta := strings.Split(a, da)
	tb := strings.Split(b, db)
	lenA := len(ta)
	lenB := len(tb)
	sa := ""
	sb := ""
	for i := 0;i < lenA - 1;i++{
		sa += da
	}
	for i := 0;i < lenB - 1;i++{
		sb += db
	}
	//fmt.Println(sa, sb)
	resA.SetString(sa, 10)
	resB.SetString(sb, 10)
	res.Add(resA, resB)
	fmt.Println(res)
}
