#ifndef _LINKED_LIST_H_INCLUDED
#define _LINKED_LIST_H_INCLUDED

typedef struct node
{
    struct node *pre;
    int value;
    struct node *next;
} Node;

//�����������ش���������ͷ������ʧ�ܷ���NULL
Node* link_create();
//��չ������������β����һ��Ԫ��,����ʧ�ܷ���NULL
Node* link_append(Node* head, int value);

void link_clear(Node *head);

void link_print(Node *head);

#endif // LINKED_LIST_H_INCLUDED
