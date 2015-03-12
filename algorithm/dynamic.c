#include "dynamic.h" 

char* lcs(const char* str1, const int len1, const char* str2, const int len2)
{
	int arr[len1+1][len2+1];
	int i, j;
	for (i = 0; i < len1+1; i++) 
		arr[i][0] = 0;
	for (i = 0; i < len2+1; i++) 
		arr[0][j] = 0;
	for(i = 1; i < len1+1; i++)
		for(j = 1; j < len2+1; j++)
			if (str1[i] == str2[j]) 
				arr[i][j] = arr[i-1][j-1] + 1;
			else if (arr[i-1][j] > arr[i][j-1])
				arr[i][j] = arr[i-1][j];
			else
				arr[i][j] = arr[i][j-1];
	char* ret = (char *) malloc(sizeof(char) * 16);
	sprintf(ret, "%d", arr[len1][len2]);
	return ret;
}

