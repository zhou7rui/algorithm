# -*- coding: utf-8 -*

'''
  堆排序
'''

import sort_helper

from max_heap import Maxheap
import quick_3_sort



@sort_helper.test_sort
def heap_sort(arr):
    heap = Maxheap(len(arr),arr)
    for i in range(len(arr) - 1,-1,-1):
        arr[i] = heap.extractMax()

def __shifDown(arr,n,k):

    while (2 * k + 1 < n ):
        j = 2 * k + 1
        if j + 1 < n and arr[j + 1] > arr[j] :
            j += 1
        if arr[k] >= arr[j]:
            break
        arr[j],arr[k] = arr[k],arr[j]
        k = j


def sort(arr):

    n = len(arr)
    for i in range((n + 1) / 2,-1,-1):
        __shifDown(arr,n,i)


    for i in range(n - 1,0, -1):
        arr[0], arr[i] = arr[i],arr[0]
        __shifDown(arr,i,0)

    print(arr)




if __name__ == '__main__':
    a = sort_helper.generate_randoma_array(100,1,99)
    print(a)
    # b = a[:]
    # heap_sort(a)
    # quick_3_sort.sort(b)
    sort(a)
