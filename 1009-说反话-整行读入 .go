package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/16 12:49 下午
* @description：
 */
func main() {
	cmdReader := bufio.NewReader(os.Stdin)
	cmdStr, err := cmdReader.ReadString('\n');
	if err != nil{
		fmt.Print(err)
	}
	res := strings.Split(cmdStr[:len(cmdStr)-1]," ")  //空格分割最后一个会带一个换行符
	fmt.Print(res[len(res) - 1])
	for i := len(res) - 2;i >= 0;i-- {
		fmt.Printf(" %s", res[i])
	}
}
