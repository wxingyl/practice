#include <stdlib.h>
#include "datastructure/linked_list.h"
#include "datastructure/stack.h"

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
	
}

int main()
{
	printf("Hello World!\n");
	return 0;
}
