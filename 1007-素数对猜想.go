package main

import (
	"fmt"
	"math"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/15 1:25 下午
* @description：
 */
var num [100000]int
var len int
var total int
func isPrinme(a int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(a)); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
func getPrimeArray(n int) {
	for i := 2; i <= n; i++ {
		if isPrinme(i) {
			num[len] = i
			if len >= 1 && num[len] - num[len-1] == 2{
				total ++
			}
			len++
		}
	}
}

func main() {
	len = 0
	total = 0
	var n int
	fmt.Scan(&n)
	getPrimeArray(n)
	fmt.Printf("%d\n",total)
}
