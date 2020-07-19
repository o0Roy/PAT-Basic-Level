package main

import (
	"fmt"
	"sort"
)

/**
* @Author: hsq_roy
* @Date: 2020/7/19 10:01 上午
* @description：

1. 才德全尽 - H <= 德分，才分，按德才总分排序。
2. 德胜才 - H <= 德分，L <= 才分 < H，按德才总分排序且排在 case 1 之后。
3. 才德兼亡 - L <= 德分，才分（德分>=才分） <= H， 按德才总分排序且排在 case 2 之后。
4. 其他 - 总分排序，按德才总分排序且排在 case 3 之后。
思路：四种情况，分四个数列排序
*/
type Stu struct {
	Id              int
	MoralScore      int
	CapabilityScore int
	Group           int
}
func (s Stu) String() string{
	return fmt.Sprintf("%d %d %d", s.Id, s.MoralScore, s.CapabilityScore)
}
func printRes(listC1 []Stu) {
	for i := 0; i < len(listC1); i++ {
		fmt.Printf("%d %d %d\n", listC1[i].Id, listC1[i].MoralScore, listC1[i].CapabilityScore)
	}
}

func main() {
	var lenC1  int
	listC11 := make([]Stu,100000)
	var N, L, H int
	var id, moral, capability int
	fmt.Scan(&N, &L, &H)
	for i := 0; i < N; i++ {
		fmt.Scan(&id, &moral, &capability)
		if L <= moral && L <= capability {
			if H <= moral && H <= capability {
				//1. 才德全尽 - H <= 德分，才分，按德才总分排序。
				listC11[lenC1] = Stu{id, moral, capability, 1}
			} else if H <= moral && capability < H {
				//2. 德胜才 - H <= 德分，L <= 才分 < H，按德才总分排序且排在 case 1 之后。
				listC11[lenC1] = Stu{id, moral, capability, 2}
			} else if moral < H && capability < H && moral >= capability {
				//3. 才德兼亡 - L <= 德分，才分（德分>才分） <= H， 按德才总分排序且排在 case 2 之后。
				listC11[lenC1] = Stu{id, moral, capability, 3}
			} else {
				//4. 其他 - 总分排序，按德才总分排序且排在 case 3 之后。
				listC11[lenC1] = Stu{id, moral, capability,4}
			}
			lenC1 ++
		}
	}
	listC1 :=  listC11[0:lenC1]
	//fmt.Print(len(listC1), len(listC2), len(listC3), len(listC4))
	sort.Slice(listC1, func(i, j int) bool {
		iGroup := listC1[i].Group
		jGroup := listC1[j].Group
		iScore := listC1[i].MoralScore + listC1[i].CapabilityScore
		jScore := listC1[j].MoralScore + listC1[j].CapabilityScore
		iMoralScore := listC1[i].MoralScore
		jMoralScore := listC1[j].MoralScore
		iId := listC1[i].Id
		jId := listC1[j].Id
		if iGroup > jGroup { // 类别
			return false
		} else if iGroup == jGroup {
			if iScore > jScore { // 先按总分
				return true
			} else if iScore == jScore {
				if iMoralScore > jMoralScore { // 总分相同，再按德分
					return true
				} else if iMoralScore == jMoralScore { // 德分相同，再按学号，此时才分是一样的，故不用才分继续排序
					if iId > jId { // 学号唯一
						return false
					}
					return true
				}
				return false
			}
			return false
		}
		return true
	})
	fmt.Println(len(listC1))
	printRes(listC1)
}
