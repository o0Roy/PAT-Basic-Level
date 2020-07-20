package main

import (
	"fmt"
	"math/rand"
	"os"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/20 9:29 上午
* @description：1018 文件读写构造
 */

func main() {
	str := "BCJ"
	f, err := os.Create("./1018.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// 写入
	var buf string
	for i := 0; i < 100000; i++ {
		buf = fmt.Sprintf("%s %s\n", string(str[rand.Intn(3)]), string(str[rand.Intn(3)]))
		_, err := f.WriteString(buf)
		if err != nil {
			fmt.Println(err)
		}
	}
	// 读取
	//ff, err := os.Open("./1018.txt")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer f.Close()
	//
	//r := bufio.NewReader(ff)
	//for {
	//	buf, err := r.ReadBytes('\n')
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//		fmt.Println(err)
	//	}
	//	fmt.Print(string(buf))
	//}
}
