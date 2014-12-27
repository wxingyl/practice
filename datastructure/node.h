#ifndef _NODE_H
#define _NODE_H

typedef struct node
{
    struct node *pre;
    int value;
    struct node *next;
} Node;

#endif

