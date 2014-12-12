#ifndef _MAIN_H
#define _MAIN_H
typedef struct node {
	struct node *pre;
	int value;
	struct node *next;
} Node;
#endif
