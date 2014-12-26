#ifndef _NODE
#define _NODE

typedef struct node
{
    struct node *pre;
    int value;
    struct node *next;
} Node;

#endif

