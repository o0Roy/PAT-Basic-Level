package main

import (
	"fmt"
	"math"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/18 11:00 上午
* @description：
 */
var prime [10000]int
var index int
func isPrime(num int) bool {
	for i := 2;float64(i) <= math.Sqrt(float64(num));i++ {
		if num % i == 0 {
			return false
		}
	}
	return true
}
func getPrime(n int) {
	for i := 2;index < n;i++ {
		if isPrime(i) {
			prime[index] = i
			index ++
		}
	}
}
func out(n int){
	var t int
	for i := n - 1;i < index;i++ {
		if t < 9 {
			if !(i == index -1) {
				fmt.Printf("%d ", prime[i])
				t++
			}else{
				fmt.Println(prime[i])
			}
		}else{
			fmt.Println(prime[i])
			t = 0
		}
	}
}
func main() {
	var m, n int
	fmt.Scan(&m, &n)
	getPrime(n)
	out(m)
}
