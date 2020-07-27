package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)
/**
* @Author: hsq_roy
* @Date: 2020/7/23 11:31 下午
* @description：
 */
func main() {
	stk := queue.New()
	stk.Enqueue(1)
	stk.Enqueue(2)
	fmt.Println(stk.Dequeue())
	fmt.Println(stk.Dequeue())
}
