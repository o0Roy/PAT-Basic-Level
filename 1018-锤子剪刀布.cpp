#include <iostream>
#include "stdio.h"
using namespace std;
int main()
{
	int N, i;
	//甲的成绩，  赢 平 输，乙的成绩为甲的成绩相反
	int score[3] = {0, 0, 0};
	// 三个0分别代表    B  C  J   
	int flag[2][3] = { {0, 0, 0},{0, 0, 0} };
	int A0, B0, indexA, indexB;
	char a, b;
	char ch[3] = {'B', 'C', 'J'};
	scanf("%d", &N);
	for (i = 0; i < N; i++) {
		getchar();
		scanf("%c %c", &a, &b);
		if (a == b) {
			// 甲乙出相同，平手
			score[1]++;
		}
		else if (a == 'C' && b == 'J') {
			// 锤子 剪刀 甲赢
			score[0]++; 
			flag[0][1]++;  
		}
		else if (a == 'C' && b == 'B') {
			// 剪刀 布 乙赢
			score[2]++; 
			flag[1][0]++;
		}
		else if (a == 'J' && b == 'B') {
			// 剪刀 布 甲赢
			score[0]++; 
			flag[0][2]++;
		}
		else if (a == 'J' && b == 'C') {
			// 剪刀 锤子 乙赢
			score[2]++;  
			flag[1][1]++;
		}
		else if (a == 'B' && b == 'C') {
			// 布 锤子 甲赢
			score[0]++; 
			flag[0][0]++;
		}
		else if (a == 'B' && b == 'J') {
			// 布 剪刀 乙赢
			score[2]++; 
			flag[1][2]++;
		}
	}
	// 甲乙成绩相反
	printf("%d %d %d\n", score[0], score[1], score[2]);
	printf("%d %d %d\n", score[2], score[1], score[0]);
	A0 = B0 = indexA = indexB = 0;  // 没有初始化，导致测试点过不去
	// 以 BCJ 顺序判断，取第一个最大值
	for (i = 0; i < 3; i++) {
		if (A0 < flag[0][i]) {
			A0 = flag[0][i];
			indexA = i;
		}
		if (B0 < flag[1][i]) {
			B0 = flag[1][i];
			indexB = i;
		}
	}
	printf("%c %c", ch[indexA], ch[indexB]);
}
