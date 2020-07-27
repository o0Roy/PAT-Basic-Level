package main

import (
	"fmt"
	"sort"
)
/*
golang 此代码提交超时，C++版本与此代码思路一致。
 */

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

// toString
func (s Stu) String() string{
	return fmt.Sprintf("%d %d %d", s.Id, s.MoralScore, s.CapabilityScore)
}

func printRes(list []Stu) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d %d %d\n", list[i].Id, list[i].MoralScore, list[i].CapabilityScore)
	}
}

func main() {
	var lenA  int
	listA := make([]Stu,100000)
	var N, L, H int
	var id, moral, capability int
	fmt.Scan(&N, &L, &H)
	for i := 0; i < N; i++ {
		fmt.Scan(&id, &moral, &capability)
		if L <= moral && L <= capability {
			if H <= moral && H <= capability {
				//1. 才德全尽 - H <= 德分，才分，按德才总分排序。
				listA[lenA] = Stu{id, moral, capability, 1}
			} else if H <= moral && capability < H {
				//2. 德胜才 - H <= 德分，L <= 才分 < H，按德才总分排序且排在 case 1 之后。
				listA[lenA] = Stu{id, moral, capability, 2}
			} else if moral < H && capability < H && moral >= capability {
				//3. 才德兼亡 - L <= 德分，才分（德分>才分） <= H， 按德才总分排序且排在 case 2 之后。
				listA[lenA] = Stu{id, moral, capability, 3}
			} else {
				//4. 其他 - 总分排序，按德才总分排序且排在 case 3 之后。
				listA[lenA] = Stu{id, moral, capability,4}
			}
			lenA ++
		}
	}
	listB :=  listA[0:lenA]
	//fmt.Print(len(listC1), len(listC2), len(listC3), len(listC4))
	sort.Slice(listB, func(i, j int) bool {
		iGroup := listB[i].Group
		jGroup := listB[j].Group
		iScore := listB[i].MoralScore + listB[i].CapabilityScore
		jScore := listB[j].MoralScore + listB[j].CapabilityScore
		iMoralScore := listB[i].MoralScore
		jMoralScore := listB[j].MoralScore
		iId := listB[i].Id
		jId := listB[j].Id
		if iGroup > jGroup { // 类别
			return false
		} else if iGroup == jGroup {
			if iScore > jScore { // 同类别，按总分
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
	fmt.Println(len(listB))
	printRes(listB)
}
