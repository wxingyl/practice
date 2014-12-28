#include <stdlib.h>
#include <stdio.h>
#include "datastructure/linked_list.h"
#include "datastructure/stack.h"
#include "datastructure/queue.h"
#include "algorithm/random.h"
#include "algorithm/sort.h"
#include "datastructure/tree.h"

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
	simple_sort(a, 100);
	int i = 0;
	for(; i < 100; i++)
	{
		printf("\t%d", a[i]);
		if((i+1) % 10 == 0)
			putchar('\n');
	}
	putchar('\n');
}
int main()
{
//	rand_test();
//	stack_test();
//    tree_test();
    test_sort();
	return 0;
}



