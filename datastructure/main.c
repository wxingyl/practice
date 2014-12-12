#include <stdio.h>
#include <stdlib.h>
#include "main.h"

//insert into head linked list
Node* insert(int value, Node* head)
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

int main()
{
	Node* head = (Node *) malloc(sizeof(Node));
	if (head == NULL) {
		printf("memory error!");
		return -1;
	}
	head->next = NULL;
	int i;
	Node *pre = head, *p;
	for(i = 0; i < 10; i++) 
	{
		p = (Node *) malloc(sizeof(Node));
		p->value = rand() / (RAND_MAX + 1.0) * 100; 
		if (p == NULL) {
			printf("memory error!");
			break;
		} else {
			pre->next = p;
			p->next = NULL;
			pre = p;
		}
	}

	p = head;
	i = 0;
	while(p->next != NULL) {
		p = p->next;
		printf("index: %d, value: %d\n", i++, p->value);
	}
	p = pre = head;
	i = 0;
	while(head->next != NULL) {
		p = head->next;
		pre = p->next;
		//printf("free index: %d, value: %d\n", i++, p->value);
		free(p);
		head->next = pre;
	}
	free(head);
	head = NULL;
	/*
	char ch = 'A';
	char ch1 = 65;
	char* chs = "AbCDEF";
	printf("ch %c, ch1 %c, chs: %s, chs[0]: %c\n", ch, ch1, chs, chs[0]);
	printf("ch %d, ch1 %d, chs: %s, chs[0]: %d\n", ch, ch1, chs, chs[1]);
	printf("a - A: %d\n", 'a'-'A');
	ch = getchar();
	if(ch >= 'A' && ch <= 'Z'){
		printf("you input: %c, is super char\n", ch);
	} else if (ch >= 'a' && ch <= 'z') {
		printf("you input: %c, is lower char\n", ch);
	} else {
		printf("you input char: %c is others\n", ch);
	}
	int i, j;
	for(i = 1; i < 10; i++) {
		for(j = 1; j <= i; j++) {
			printf("%d x %d = %d\t", j, i, j * i);
		}
		putchar('\n');
	}
	putchar('\n');
	*/
	return 0;
}
