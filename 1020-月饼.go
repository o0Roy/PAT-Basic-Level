package main

import (
	"fmt"
	"io"
	"sort"
)

type Good struct {
	// 需要全部定义成float，否则测试点2过不去
	Num      float32
	SumPrice float32
	Price    float32
}

type goodList []Good

// 排序继承的三个接口
func (g goodList) Len() int           { return len(g) }           // 设置一个全局变量即可指定排序长度
func (g goodList) Swap(i, j int)      { g[i], g[j] = g[j], g[i] } // 设想：加上偏移量，排序指定区间
func (g goodList) Less(i, j int) bool { return g[i].Price >= g[j].Price }

func main() {
	var N, D int
	var t float32
	var good []Good
	var sum float32
	for {
		_, err := fmt.Scanf("%d %d", &N, &D)
		if err == io.EOF {
			break
		}
		for j := 0; j < N; j++ {
			fmt.Scanf("%f", &t)
			good = append(good, Good{t, 0, 0})
		}
		for j := 0; j < N; j++ {
			fmt.Scanf("%f", &t)
			good[j].SumPrice = t
		}
		// 求出单价
		for j := 0; j < N; j++ {
			good[j].Price = float32(good[j].SumPrice) / float32(good[j].Num)
		}
		// 按照单价排序
		sort.Sort(goodList(good))
		sum = 0
		for i := 0; i < N; i++ {
			if float32(D) >= good[i].Num { // 存货大于需求
				sum += float32(good[i].SumPrice)
			} else {
				sum += good[i].Price * float32(D)
				break
			}
			D -= int(good[i].Num)
		}
		fmt.Printf("%.2f\n", sum)
	}
	//for i := 0;i < N;i++ {i
	//	fmt.Println(good[i].Price)
	//}
}
