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
    link_clear(head);
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


int main(int argc, char* argv[])
{
//	rand_test();
//	stack_test();
//    tree_test();
//    test_sort();
//	test_dynamic(argv[1], argv[2]);
//	test_unix();
//	printf("%s: %d\n", argv[1], trailingZeroes(atoi(argv[1])));
	printf("[%s, %s]: %d\n", argv[1], argv[2], rangeBitwiseAnd(atoi(argv[1]), atoi(argv[2])));
	return 0;
}


