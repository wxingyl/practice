#include <stdlib.h>
#include <time.h>

void init_rand(int* a, const int size, const int min, const int max)
{
	srand(time(0));
	int i;
	for(i = 0; i < size; i++)
		a[i] = rand() % (max - min) + min;
}
