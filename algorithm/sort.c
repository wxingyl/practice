#include<stdio.h>
#include<stdlib.h>
#include "sort.h"

void simple_sort(int data[], int size)
{
    int i, j;
    for (i = 0; i < size-1; i++)
    {
        for (j = i+1; j < size; j++)
        {
            if (data[i] > data[j])
            {
                int tmp = data[i];
                data[i] = data[j];
                data[j] = tmp;
            }
        }
    }
}

static void swap(int*a,int*b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}

void select_sort(int data[],int n)
{
    register int i,j,min;
    for(i=0; i<n-1; i++)
    {
        min=i;//查找最小值
        for(j=i+1; j<n; j++)
        {
            if(data[min]>data[j])
            {
                min=j;
            }
        }
        if(min!=i)
        {
            swap(&data[min],&data[i]);
            printf("第%d趟排序结果为:\n",i+1);
        }
    }
}
