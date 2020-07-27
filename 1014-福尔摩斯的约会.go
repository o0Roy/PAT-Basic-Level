package main

import (
	"fmt"
	"strings"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/18 11:19 上午
* @description：
//1.字符相等，而且字符必须是大写字母，且必须在'A' ~ 'G' 之间。（星期几）
//2.字符相等，而且字符必须在'A' ~ 'N' 或 '0' ~ '9'之间 （小时）
//3.字符相等，而且必须是字母。（分钟）
 */
var DAY string
var HH, MM int
func judgeDayAndHH(a, b string) {
	strDay := []string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}
	strHH := "0123456789ABCDEFGHIJKLMN"
	var flag int = 0
	for i := 0; i < len(a) && i < len(b); i++ {
		if flag == 0 {
			// 限定大写字母A~G
			if a[i] == b[i] && 65 <= a[i] && a[i] < 72  {
					DAY = strDay[a[i]-65]
					flag = 1
			}
		} else if flag == 1 {
			// 限定数字，大写字母A~N
			if a[i] == b[i] && (65 <= a[i] && a[i] < 79  || 48 <= a[i] && a[i] < 58 ) {
				HH = strings.Index(strHH, a[i:i+1])
				break
			}
		}
	}
}
func judgeMM(c, d string) {
	for i := 0; i < len(c) && i < len(d); i++ {
		// 限定字母范围
		if c[i] == d[i] && (65 <= c[i] && c[i] < 91 || 97 <= c[i] && c[i] < 123) {
			MM = i
			break
		}
	}
}
func main() {
	var a, b, c, d string
	fmt.Scan(&a, &b, &c, &d)
	judgeDayAndHH(a, b)
	judgeMM(c, d)
	fmt.Printf("%s %02d:%02d\n", DAY, HH, MM)
}


