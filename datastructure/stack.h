#ifndef _STACK
#define _STACK
#include "node.h"

Node *stack_create();

Node *stack_push(Node* top, int value);

Node* stack_pop(Node* top);

void stack_clear(Node *top);
#endif
