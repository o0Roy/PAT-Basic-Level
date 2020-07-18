package main

import "fmt"

/**
* @Author: hsq_roy
* @Date: 2020/7/17 2:07 下午
* @description：
 */
var a1, a2, a3, a4, a5 int
var aa4 float64
var opera int
var flag [5]int
func judge(num int) {
	d := num % 5
	switch(d){
	case 0:
		if num%2 == 0 {
			a1 += num
			flag[d] = 1
		}
	case 1:
		flag[d] = 1
		if opera == 0 {
			a2 += num
		}else{
			a2 -= num
		}
		opera = 1 - opera
	case 2:
		flag[d] = 1
		a3++
	case 3:
		flag[d] = 1
		aa4+=float64(num)
		a4++
	case 4:
		if num > a5 {
			flag[d] = 1
			a5 = num
		}
	}
}
func main() {
	var n, num int
	fmt.Scan(&n)
	for i := 0;i < n;i++ {
		fmt.Scan(&num)
		judge(num)
	}
	if flag[0] == 1 {
		fmt.Printf("%d", a1)
	}else{
		fmt.Print("N")
	}
	if flag[1] == 1 {
		fmt.Printf(" %d", a2)
	}else{
		fmt.Print(" N")
	}
	if flag[2] == 1 {
		fmt.Printf(" %d", a3)
	}else{
		fmt.Print(" N")
	}
	if flag[3] == 1 {
		fmt.Printf(" %.1f", aa4/float64(a4))
	}else{
		fmt.Print(" N")
	}
	if flag[4] == 1 {
		fmt.Printf(" %d", a5)
	}else{
		fmt.Print(" N")
	}
}
