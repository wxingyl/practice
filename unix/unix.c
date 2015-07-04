#include "unix.h"

#define BSZ 48

void file_test()
{
	char buf[BSZ];
	if(getcwd(buf, BSZ) == NULL) {
		printf("getcwd failed!\n");
		return;
	}
	printf("cwd = %s\n", buf);
	FILE *fp;
	memset(buf, 'a', BSZ-2);
	buf[BSZ-2] = '\0';
	buf[BSZ-1] = 'X';
	printf("buf: %s\n", buf);
	printf("strlen: %lu\n", strlen(buf));
	/*
	if ((fp = fmemopen(buf, BSZ, "w+")) == NULL) {
		printf("fmemopen failed\n");
		return;
	}
	fprintf(fp, "hello world!");
	fflush(fp);
	printf("after fflush buf: %s\n", buf);
	printf("strlen: %lu\n", strlen(buf));
	*/
	
}
