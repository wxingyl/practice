#include "linked_list.h"
#include <stdlib.h>
#include <stdio.h>

Node* link_create()
{
	Node* head = (Node *) malloc(sizeof(Node));
	if (head == NULL) {
		printf("memory error!");
		return NULL;
	}
	head->next = NULL;
	return head;
}

Node* link_append(Node* head, int value)
{
	Node *tail = head;
	while(tail->next != NULL) {
		tail = tail->next;
	}
	Node *p = (Node *) malloc(sizeof(Node));
	if (p == NULL){
	    printf("memory error!");
		return NULL;
	}
	p->value = value;
	p->next = NULL;
	tail->next = p;
	return head;
}
//Á´±íÍ·²åÈëvalue
Node* link_push(Node *head, int value)
{
    Node *first = head->next;
    Node *p = (Node *)malloc(sizeof(Node));
    if (p == NULL){
        printf("memory error!");
        return NULL;
    }
    p->value = value;
    head->next = p;
    p->next = first;
    return head;
}

void link_clear(Node *head)
{
    Node *p, *pre;
    p = pre = head;
	while(head->next != NULL) {
		p = head->next;
		pre = p->next;
		free(p);
		head->next = pre;
	}
}

void link_destory(Node* head)
{

    link_clear(head);
    free(head);
    head = NULL;
}

void link_print(Node *head)
{
	int i = 0;
	while(head->next != NULL) {
		head = head->next;
		printf("index: %d, value: %d\n", i++, head->value);
	}
}

bool link_contain(Node* head, int value)
{
	Node *p = head->next;
	while(p != NULL) {
		if (p->value == value)
			return 1;
		p = p->next;
	}
	return 0;
}
