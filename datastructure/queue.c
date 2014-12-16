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

Queue* queue_add(Queue *q, int value)
{
	//插入节点的是Node对象，Queue只是封装得一个队列对象
	//Queue对象中的first和last的Node指针对象是需要操作的
	Queue* p = (Queue*)malloc(sizeof(Queue));
	if (!p) {
		printf("memory error!\n");
		return NULL;
	}
	//同样地这儿给Node的指针对象赋值
	p->value = value;
	p->next = NULL;
	last->next = p;
	last = p;
	q->size++;
	return q;
}

Queue* queue_pop(Queue *q)
{
    Queue *p
    p = first;
    p->next = first;
    free(p);
    q->size--;
    return q;
}

Queue* queue_clear(Queue *q)
{
    while(first->next != NULL)
        first = queue_pop(q);
}
