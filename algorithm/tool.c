#include <stdlib.h>
#include <time.h>

void init_rand(int* a, const int size, const int min, const int max)
{
	srand(time(0));
	int i;
	for(i = 0; i < size; i++)
		a[i] = rand() % (max - min) + min;
}

int trailingZeroes(int n) {
	if (n < 5) return 0;
	int ret = n / 5;
	ret += trailingZeroes(n / 5);
	return ret;
}
