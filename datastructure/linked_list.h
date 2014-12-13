#ifndef _LINKED_LIST_H_INCLUDED
#define _LINKED_LIST_H_INCLUDED

typedef struct node
{
    struct node *pre;
    int value;
    struct node *next;
} Node;

//创建链表，返回创建的链表头，创建失败返回NULL
Node* link_create();
//扩展链表，即向链表尾插入一个元素,插入失败返回NULL
Node* link_append(Node* head, int value);

void link_clear(Node *head);

void link_print(Node *head);

#endif // LINKED_LIST_H_INCLUDED
