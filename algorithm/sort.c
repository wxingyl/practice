#include<stdio.h>
#include<stdlib.h>
#include "sort.h"

//Ωªªª∫Ø ˝
static void swap(int* a, int* b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}

//√∞≈›≈≈–Ú
void simple_sort(int data[], int size)
{
    int i, j;
    for (i = 0; i < size-1; i++)
    {
        for (j = i+1; j < size; j++)
        {
            if (data[i] > data[j])
            {
                swap(&data[i], &data[j]);
            }
        }
    }
}


//—°‘Ò≈≈–Ú
void select_sort(int data[], int size)
{
    int i, j, min;
    for(i = 0; i < size - 1; i++)
    {
        min = i;
        for(j = i + 1; j < size; j++){
            if(data[min] > data[j]){
                min = j;
            }
        }
        if(min != i){
			swap(&data[min], &data[i]);
        }
    }
}

//≤Â»Î≈≈–Ú
void insert_sort(int data[], int size)
{
    int i, j;
    for(i = 1; i < size; i++)
    {
		j = i;
        while( j > 0 && data[j] < data[j-1])
        {
            swap(&data[j], &data[j-1]);
            j--;
        }
    }
}

