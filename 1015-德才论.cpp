#include <iostream>
#include "stdio.h"
#include <algorithm>
using namespace std;
struct Stu {
	int id;
	int moralscore;
	int capabilityscore;
	int group;
}s[100050];

bool cmp(Stu a, Stu b) {
	int aScore = a.moralscore + a.capabilityscore;
	int bScore = b.moralscore + b.capabilityscore;
	if (a.group > b.group)  // 1. 按类别分
		return false;
	else if (a.group == b.group) {
		if (aScore > bScore)  // 2.类别相同，比较总成绩
			return true;
		else if (aScore == bScore) {
			if (a.moralscore > b.moralscore) // 3.总成绩相同，比较德分
				return true;
			else if (a.moralscore == b.moralscore) {
				if (a.id > b.id) //4.德分相同比较学号，学号唯一
					return false;
				return true;
			}
			return false;
		}
		return false;
	}
	return true;
	/* 简版写法
	if (a.group != b.group)
		return a.group < b.group;
	else if (aScore != bScore)
		return aScore > bScore;
	else if (a.moralscore != b.moralscore)
		return a.moralscore > b.moralscore;
	else
		return a.id < b.id;
	*/
}

int main()
{
	int len = 0;
	int N, L, H;
	int id, moral, capability, i;
	scanf("%d %d %d", &N, &L, &H);
	for (i = 0; i < N; i++) {
		scanf("%d %d %d", &id, &moral, &capability);
		if (L <= moral && L <= capability) {
			if (H <= moral && H <= capability) {
				//1. 才德全尽 - H <= 德分，才分，按德才总分排序。
				s[len++] = Stu{ id, moral, capability, 1 };
			}
			else if (H <= moral && capability < H) {
				//2. 德胜才 - H <= 德分，L <= 才分 < H，按德才总分排序且排在 case 1 之后。
				s[len++] = Stu{ id, moral, capability, 2 };
			}
			else if (moral < H && capability < H && moral >= capability) {
				//3. 才德兼亡 - L <= 德分，才分（德分>=才分） <= H， 按德才总分排序且排在 case 2 之后。
				s[len++] = Stu{ id, moral, capability, 3 };
			}
			else {
				//4. 其他 - 总分排序，按德才总分排序且排在 case 3 之后。
				s[len++] = Stu{ id, moral, capability, 4 };
			}
		}
	}
	sort(s, s + len, cmp);
	printf("%d\n", len);
	for (i = 0; i < len; i++) {
		printf("%d %d %d\n", s[i].id, s[i].moralscore, s[i].capabilityscore);
	}
}
