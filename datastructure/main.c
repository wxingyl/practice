#include <stdlib.h>
#include "linked_list.h"

int main()
{

	int i;
	Node *head = link_create();
	for(i = 0; i < 10; i++)
	{
		link_append(head, rand() / (RAND_MAX + 1.0) * 100);
	}
    link_print(head);
    link_clear(head);
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
