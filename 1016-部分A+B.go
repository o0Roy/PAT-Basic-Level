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
	var A, Da, B, Db string
	numA := new(big.Int)
	numB := new(big.Int)
	res := new(big.Int)
	fmt.Scan(&A, &Da, &B, &Db)
	// 以Da, Db分割
	arrA := strings.Split(A, Da)
	arrB := strings.Split(B, Db)
	// 根据长度判断出现次数
	lenA := len(arrA)
	lenB := len(arrB)
	strA := ""
	strB := ""
	// 根据分割数构造字符串，分割数为分割点 - 1
	for i := 0;i < lenA - 1;i++{
		strA += Da
	}
	for i := 0;i < lenB - 1;i++{
		strB += Db
	}
	// 大数，字符串转十进制
	numA.SetString(strA, 10)
	numB.SetString(strB, 10)
	res.Add(numA, numB)
	fmt.Println(res)
}
