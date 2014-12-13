#include <stdlib.h>
#include <stdio.h>
Node *stack_create()
{
//    Node *top = (Node *)malloc(sizeof(Node));
//    if (top == NULL){
//        printf("memory error");
//        return NULL;
//    }
//    top->next = NULL;
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
    Node *p = top;
    top = top->next;
    free(p);
    return top;
}

