#include <stdlib.h>
#include <stdio.h>
#include "queue.h"

Queue* queue_create()
{
	Queue* q = (Queue*)malloc(sizeof(Queue));
	if (!q) {
		printf("memory error!\n");
		return NULL;
	}
	q->size = 0;
	q->first = NULL;
	q->last = NULL;
	return q;
}

Queue* queue_push(Queue *q, int value)
{
	//插入节点的是Node对象，Queue只是封装得一个队列对象
	//Queue对象中的first和last的Node指针对象是需要操作的
	Node* p = (Node*)malloc(sizeof(Node));
	if (!p) {
		printf("memory error!\n");
		return NULL;
	}
	p->value = value;
	p->next = q->first;
	q->first = p;
	if (!q->size) {
        q->last = q->first;
	}
	q->size += 1;
	return q;
}

int queue_pop(Queue *q)
{
    if (!q->size) {
        printf("size is 0, can not pop\n");
        return -1;
    }
    q->size--;
    Node *pre, *p;
    pre = p = q->first;
    while(p->next != NULL) {
        pre = p;
        p = p->next;
    }
    pre->next = NULL;
    if (pre == p) {
        q->last = NULL;
        q->first = NULL;
    } else {
        q->last = pre;
    }
    int ret = p->value;
    free(p);
    return ret;
}

void queue_clear(Queue *q)
{
    while(q->size) queue_pop(q);
}
