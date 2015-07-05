#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include "datastructure/linked_list.h"
#include "datastructure/stack.h"
#include "datastructure/queue.h"
#include "algorithm/acm.h"
#include "algorithm/sort.h"
#include "datastructure/tree.h"
#include "unix/unix.h"

void link_test()
{
	int i;
	Node* const head = link_create();
	for(i = 0; i < 10; i++)
	{
		link_append(head, rand() / (RAND_MAX + 1.0) * 100);
	}
    link_print(head);
    link_destory(head);
}

void stack_test()
{
	Node *top = stack_push(NULL, 100);
	printf("stack top->value: %d\n", top->value);
}

void rand_test()
{
	int a[100];
	init_rand(a, 100, 0, 100);
	int i = 0;
	for(; i < 100; i++)
	{
		printf("\t%d", a[i]);
		if((i+1) % 10 == 0)
			putchar('\n');
	}
	putchar('\n');
}

void tree_test()
{
	Tree* tr = NULL;
    Tr_Node* np = insert(tr, 18);
    printf("Orignal:\n");
    display(tr);
}

void test_sort()
{
	int a[100];
	init_rand(a, 100, 0, 100);
	quick_sort(a, 100);
	int i = 0;
	for(; i < 100; i++)
	{
		printf("\t%d", a[i]);
		if((i+1) % 10 == 0)
			putchar('\n');
	}
	putchar('\n');
}

void test_dynamic(const char* str1, const char* str2)
{
	printf("str1: %s\nstr2: %s\n", str1, str2);
	char* ret = lcs(str1, strlen(str1), str2, strlen(str2));
	printf("result: %s\n", ret);
	free(ret);
}

void test_unix()
{
	file_test();
}

void test_rotate(char* argv[])
{
	int n = atoi(argv[1]), k = atoi(argv[2]);
	int* p = (int* )malloc(sizeof(int) * n);
	init_rand(p, n, 0, 1000);
	printf("n = %d, k = %d :\n", n, k);
	int i;
	for (i = 0; i < n; i++) {
		if (i % 5 == 0) putchar('\n');
		printf("\t%d", p[i]);
	}
	putchar('\n');
	rotate(p, n, k);
	printf("after rotate:\n");
	for (i = 0; i < n; i++) {
		if (i % 5 == 0) putchar('\n');
		printf("\t%d", p[i]);
	}
	putchar('\n');
	free(p);
}

void testSingleNumber(int argc, char* argv[]) {
	argc--;
	int* p = (int*) malloc(sizeof(int) * argc);
	int i;
	for (i = 0; i < argc; i++)
		p[i] = atoi(argv[i+1]);
	for (i = 0; i < argc; i++) {
		printf("%d\t", p[i]);
	}
	putchar('\n');
	int ret = singleNumberll(p, argc);
	printf("singleNumber: %d\n", ret);
	free(p);
}

void testMinSubArrayLen(int argc, char* argv[]) {
	const int size = argc - 2;
	int* a = (int*) malloc(sizeof(int) * size);
	int i;
	for (i = 0; i < size; i++) {
		a[i] = atoi(argv[i+2]);
	}
	printf("%d\n", minSubArrayLen(atoi(argv[1]), a, size));
	free(a);
}

void testTwoSum(int argc, char* argv[]) {
	int index = atoi(argv[1])+1, i;
	struct ListNode *p1 = NULL, *p2 = NULL, *p = NULL;
	for (i = 2; i < argc; i++) {
		if (i <= index && p1 == NULL) {
			p1 = (struct ListNode*) malloc(sizeof(struct ListNode));
			p = p1;
		} else if (i > index && p2 == NULL) {
			p2 = (struct ListNode*) malloc(sizeof(struct ListNode));
			p = p2;
		} else {
			p->next = (struct ListNode*) malloc(sizeof(struct ListNode));
			p = p->next;
		}
		p->next = NULL;
		p->val = atoi(argv[i]);
	}
	printf("p1: ");
	print_list_node(p1);
	printf("p2: ");
	print_list_node(p2);
	struct ListNode* ret = addTwoNumbers(p1, p2);
	printf("ret: ");
	print_list_node(ret);
	free_list_node(p1);
	free_list_node(p2);
	free_list_node(ret);
}
void testLengthOfLongestSubstring(char* argv) {
	int len = lengthOfLongestSubstring(argv);
	printf("%s: %d\n", argv, len);
}

void testFindMedianSortedArrays(char* argv[]) {
	int len1 = atoi(argv[1]), len2 = atoi(argv[2]);
	int* nums1 = (int*) malloc(sizeof(int) * len1);
	int* nums2 = (int*) malloc(sizeof(int) * len2);
	int i;
	for (i = 0; i < len1; i++)
		nums1[i] = atoi(argv[i+3]);
	for (i = 0; i < len2; i++)
		nums2[i] = atoi(argv[i+3+len1]);
	printf("MedianSortedArrays: %f\n", findMedianSortedArrays(nums1, len1, nums2, len2));
	free(nums1);
	free(nums2);
}

void testLongestPalindrome(char* s) {
	char* pa = longestPalindrome(s);
	printf("LongestPalindrome: %s\n", pa);
	if (s != pa)
		free(pa);
}

void testConvert(char* s, int numRows) {
	char* ret = convert(s, numRows);
	printf("ret: %s\n", ret);
	free(ret);
}

void testLongestCommonPrefix(int argc, char* argv[]) {
	char* arg[argc-1];
	int i;
	for (i = 1; i < argc; i++)
		arg[i-1] = argv[i];
	puts(longestCommonPrefix(arg, argc-1));
}

void testRemoveNthFromEnd(int argc, char* argv[]) {
	struct ListNode* head = NULL;
	struct ListNode* p;
	int i;
	for (i = 1; i < argc - 1; i++) {
		struct ListNode* node = (struct ListNode*) malloc(sizeof(struct ListNode));
		node->val = atoi(argv[i]);
		node->next = NULL;
		if (head == NULL) {
			head = node;
			p = node;
		} else {
			p->next = node;
			p = p->next;
		}
	}
	head = removeNthFromEnd(head, atoi(argv[argc-1]));
	print_list_node(head);
	free_list_node(head);
}


void testGenerateParenthesis(int n) {
	int size = 0, i;
	char** ret = generateParenthesis(n, &size);
	for (i = 0; i < size; i++) {
		puts(ret[i]);
		putchar('\n');
		free(ret[i]);
	}
	free(ret);
    printf("size: %d\n", size);
}

int main(int argc, char* argv[])
{
//	rand_test();
//	stack_test();
//    tree_test();
//    test_sort();
//	test_dynamic(argv[1], argv[2]);
//	test_unix();
//	printf("%s: %d\n", argv[1], trailingZeroes(atoi(argv[1])));
//	printf("[%s, %s]: %d\n", argv[1], argv[2], rangeBitwiseAnd(atoi(argv[1]), atoi(argv[2])));
//	test_rotate(argv);
//	printf("%s, %d\n", argv[1], isHappy(atoi(argv[1])));
//	testSingleNumber(argc, argv);
//	testTwoSum(argc, argv);
//	testLengthOfLongestSubstring(argv[1]);
//	testFindMedianSortedArrays(argv);
//	testLongestPalindrome(argv[1]);
//	testConvert(argv[1], atoi(argv[2]));
//	int i = atoi(argv[1]);
//	printf("%s, %d: %d\n", argv[1], i, isPalindrome(atoi(argv[1])));
//	testLongestCommonPrefix(argc, argv);
//	testRemoveNthFromEnd(argc, argv);
//	testGenerateParenthesis(atoi(argv[1]));
    testGenerateParenthesis(10);
	return 0;
}
