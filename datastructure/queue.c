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
	return NULL;
}
