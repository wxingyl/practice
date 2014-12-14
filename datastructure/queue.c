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
	Queue* p = (Queue*)malloc(sizeof(Queue));
	if (!p) {
		printf("memory error!\n");
		return NULL;
	}
	p = first;
	p->value = value;
	q->size = size++;
	return q;
}

Queue* queue_pop(Queue *q)
{
    Queue *p
    p = last;
    last->next = last;
    free(p);
    return q;
}



void queue_clear(Queue *q)
{
    while (first != NULL)
        first = queue_pop(q);
}

