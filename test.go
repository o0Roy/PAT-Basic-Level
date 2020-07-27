package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)
/**
* @Author: hsq_roy
* @Date: 2020/7/23 11:31 下午
* @description：
 */
func main() {
	stk := stack.New()
	stk.Push(1)
	fmt.Println(stk.Pop())
}
