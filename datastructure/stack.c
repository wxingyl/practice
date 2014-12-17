#include <stdlib.h>
#include <stdio.h>
#include "stack.h"

Node *stack_create()
{
    return NULL;
}

Node* stack_push(Node* top, int value)
{
    Node *p = (Node *)malloc(sizeof(Node));
    if (p == NULL)
    {
        printf("memory error");
        return NULL;
    }
    p->value = value;
    p->next = top;
    top = p;
    return top;
}

Node* stack_pop(Node* top)
{
	if (!top) {
		printf("stack is empty!\n");
		return NULL;
	}
    Node *p = top;
    top = top->next;
    free(p);
    return top;
}

void stack_clear(Node *top)
{
    while((top = stack_pop(top)) != NULL);
}
