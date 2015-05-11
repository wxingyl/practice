#include "acm.h"
#include "../datastructure/linked_list.h"
#include <time.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>

char* lcs(const char* str1, const int len1, const char* str2, const int len2)
{
	int arr[len1+1][len2+1];
	int i, j;
	for (i = 0; i < len1+1; i++)
		arr[i][0] = 0;
	for (i = 0; i < len2+1; i++)
		arr[0][j] = 0;
	for(i = 1; i < len1+1; i++)
		for(j = 1; j < len2+1; j++)
			if (str1[i] == str2[j])
				arr[i][j] = arr[i-1][j-1] + 1;
			else if (arr[i-1][j] > arr[i][j-1])
				arr[i][j] = arr[i-1][j];
			else
				arr[i][j] = arr[i][j-1];
	char* ret = (char *) malloc(sizeof(char) * 16);
	sprintf(ret, "%d", arr[len1][len2]);
	return ret;
}

void init_rand(int* a, const int size, const int min, const int max)
{
	srand(time(0));
	int i;
	for(i = 0; i < size; i++)
		a[i] = rand() % (max - min) + min;
}

int trailingZeroes(int n) {
	if (n < 5) return 0;
	int ret = n / 5;
	ret += trailingZeroes(n / 5);
	return ret;
}

int rangeBitwiseAnd(int m, int n) {
	int len1 = 0;
	long i = 1;
	for (; i <= m; i <<= 1, len1++);
	int len2 = len1;
	for (; i <= n; i <<= 1, len2++);
	if (len2 != len1) return 0;
	int ret = n;
	for (; m < n; m++) {
		ret &= m;
		if (!ret) break;
	}
	return ret;
}

void rotate(int nums[], int n, int k) {
	k = k % n;
	if (!(n && k)) return;
	int i = 0, next_val = nums[0];
	int shift = 0, c = 0, j;
	while(c++ < n) {
		j = (i+k) % n;
		int tmp = nums[j];
		nums[j] = next_val;
		next_val = tmp;
		i = j;
		if (i == shift) {
			i = ++shift;
			next_val = nums[i];
		}
	}
}

bool _isHappy(int n, Node* head) {
	printf("n: %d\n", n);
	char str[16];
	sprintf(str, "%d", n);
	int length = strlen(str);
	int i, newN = 0;
	for (i = 0; i < length; i++)
		newN += (str[i] - '0') * (str[i] - '0');
	if (newN == 1)
		return true;
	else if (link_contain(head, newN))
		return false;
	else {
		link_push(head, newN);
		return _isHappy(newN, head);
	}
}

bool isHappy(int n) {
	Node* head = link_create();
	bool ret = _isHappy(n, head);
	printf("ret: %d\n", ret);
	link_destory(head);
	return ret;
}

int singleNumber(int* nums, int numsSize) {
	int i, ret = 0;
	for (i = 0; i < numsSize; i++)
		ret ^= nums[i];
	return ret;
}

int singleNumberll(int* nums, int numsSize) {
	int ones = 0, twos = 0, threes = 0;
	int i;
	for (i = 0; i < numsSize; i++) {
		//升级到twos中的位
		int tmp = nums[i] & ones;
		twos |= tmp;
		//一次的进来，同时把出现2次的干掉
		ones ^= nums[i];
		threes = ones & twos;
		ones &= ~threes;
		twos &= ~threes;
	}
	return ones;
}
int longestConsecutive(int* nums, int numsSize) {
	//通过java实现了
	return 0;
}
