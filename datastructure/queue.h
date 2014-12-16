#ifndef _QUEUE
#define _QUEUE

#include "node.h"

typedef struct {
	Node* first;
	Node* last;
	int size;
} Queue;

//创建队列
Queue* queue_create();
//队列的头添加value
Queue* queue_add(Queue *q, int value);
//从队列的尾删除元素
Queue* queue_pop(Queue *q);
//清空队列元素
void queue_clear(Queue *q);
#endif
