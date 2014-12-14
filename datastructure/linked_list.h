#ifndef _LINKED_LIST_H_INCLUDED
#define _LINKED_LIST_H_INCLUDED
#include "node.h"
//�����������ش���������ͷ������ʧ�ܷ���NULL
Node* link_create();
//��չ������������β����һ��Ԫ��,����ʧ�ܷ���NULL
Node* link_append(Node* head, int value);
//����ͷ����value
Node* link_push(Node *head, int value);

void link_clear(Node *head);

void link_print(Node *head);

void link_destory(Node* head);

#endif // LINKED_LIST_H_INCLUDED
