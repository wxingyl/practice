#include <stdlib.h>
#include <stdio.h>
#include "datastructure/linked_list.h"
#include "datastructure/stack.h"
#include "datastructure/queue.h"
#include "algorithm/random.h"
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

/*void rand_test()
{
	int a[100];
	init_rand(a, 100, 0, 10);
	int i = 0;
	for(; i < 100; i++)
	{
		printf("\t%d", a[i]);
		if((i+1) % 10 == 0)
			putchar('\n');
	}
	putchar('\n');
}
*/

int main()
{
    void tree_test();
//	rand_test();
//	stack_test();
    tree_test();
	return 0;
}


void tree_test()
{
	Tree* tr;
    Tr_Node* np;
    tr = NULL;
    tr = insert_value(tr, 18);
    tr = insert_value(tr, 56);
    tr = insert_value(tr, 23);
    tr = insert_value(tr, 8);
    tr = insert_value(tr, 2);
    tr = insert_value(tr, 5);
    printf("Orignal:\n");
    print_sorted_tree(tr);

    np = find_value(tr, 8);
    if(np != NULL)
    {
        delete_node(np);
        printf("After deletion:\n");
        print_sorted_tree(tr);
    }
}
