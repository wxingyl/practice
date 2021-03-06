#include<stdio.h>
#include<stdlib.h>
#include "sort.h"

//��������
static void swap(int* a, int* b)
{
    int temp;
    temp = *a;
    *a = *b;
    *b = temp;
}

//ð������
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


//ѡ������
void select_sort(int data[], int size)
{
    int i, j, min;
    for(i = 0; i < size - 1; i++)
    {
        min = i;
        for(j = i + 1; j < size; j++)
        {
            if(data[min] > data[j])
            {
                min = j;
            }
        }
        if(min != i)
        {
            swap(&data[min], &data[i]);
        }
    }
}

//��������
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

static void _quick_sort(int a[], int left, int right)
{
    if (left >= right) return;
    int i = left, j = right, key = a[left];
    while(i < j)
    {
        while(i < j && a[j] >= key)
        {
            j--;
        }
        a[i] = a[j];
        while(i < j && a[i] <= key)
        {
            i++;
        }
        a[j] = a[i];
    }
    a[i] = key;
    _quick_sort(a, left, i-1);
    _quick_sort(a, i+1, right);
}

//��������
void quick_sort(int data[], int size)
{
    _quick_sort(data, 0, size-1);
}
