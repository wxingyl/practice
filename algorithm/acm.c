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

int minSubArrayLen(int s, int* nums, int numsSize) {
	int i = 0, j = 0, ret = numsSize+1, sum = 0;
	while(j < numsSize) {
		while(j < numsSize && sum < s) {
			sum += nums[j];
			j++;
		}
		while(i < j && sum >= s) {
			sum -= nums[i];
			i++;
		}
		int len = j - i + 1;
		ret = ret < len ? ret : len;
	}
	return ret > numsSize ? 0 : ret;
}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 * 没意思, map解决, 或者排序后得到
 **/
int* twoSum(int* nums, int numsSize, int target) {
	int i = 0, j = numsSize-1;
	bool cont = true;
	for (; i < numsSize-1; i++) {
		int t = target - nums[i];
		for (j = i+1; j < numsSize; j++) {
			if (t == nums[j]) {
				cont = false;
				break;
			}
		}
		if (!cont)
			break;
	}
	int* ret = (int*) malloc(2 * sizeof(int));
	ret[0] = i+1;
	ret[1] = j+1;
	return ret;
}

void print_list_node(struct ListNode* l) {
	struct ListNode* p = l;
	while(p != NULL) {
		printf("%d->", p->val);
		p = p->next;
	}
	putchar('\n');
}

void free_list_node(struct ListNode* l) {
	struct ListNode* p = l;
	while(l != NULL) {
		p = l->next;
		free(l);
		l = p;
	}
}

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2) {
	struct ListNode *p1 = l1, *p2 = l2;
	int next_add = 0, sum, val1, val2;
	struct ListNode *ret = NULL, *p;
	while(p1 != NULL || p2 != NULL || next_add) {
		if (ret == NULL) {
			ret = (struct ListNode*) malloc(sizeof(struct ListNode));
			p = ret;
		} else {
			p->next = (struct ListNode*) malloc(sizeof(struct ListNode));
			p = p->next;
		}
		p->next = NULL;
		if (p1 == NULL) val1 = 0;
		else {
			val1 = p1->val;
			p1 = p1->next;
		}
		if (p2 == NULL) val2 = 0;
		else {
			val2 = p2->val;
			p2 = p2->next;
		}
		sum = val1 + val2 + next_add;
		next_add = 0;
		p->val = sum % 10;
		if (sum >= 10) next_add++;
	}
	return ret;
}

int lengthOfLongestSubstring(char* s) {
	int i = 0, j = 1, t, tmp, ret = 1;
	int len = strlen(s);
	if (!len) return 0;
	while(j < len) {
		char ch = s[j];
		for(t = i; t < j;t++) {
			if (s[t] == ch) break;
		}
		if (t != j) {
			tmp = j - i;
			ret = ret < tmp ? tmp : ret;
			i = t+1;
		}
		j++;
	}
	tmp = j - i;
	ret = ret < tmp ? tmp : ret;
	return ret;    
}
