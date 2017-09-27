# -*- coding: utf-8 -*

'''
    两路快速排序
'''
import random
import sort_helper

def __partition2(arr,l,r):
    k = random.randint(l, r)
    arr[l], arr[k] = arr[k], arr[l]

    v = arr[l]


    i, j = l + 1, r

    while True:

        while i <= r and arr[i] < v :
            i += 1
        while j >= l + 1 and arr[j] > v:
            j -= 1

        if i > j:
            break

        arr[i], arr[j] = arr[j], arr[i]
        i += 1
        j -= 1
    arr[l] , arr[j] = arr[j], arr[l]
    return j

def __sort(arr,l,r):

    if l > r:
        return

    p = __partition2(arr,l,r)
    __sort(arr,l, p - 1)
    __sort(arr,p + 1, r)


@sort_helper.test_sort
def sort(arr):
    __sort(arr,0,len(arr) - 1)


if __name__ == '__main__':
    a = sort_helper.generate_randoma_array(1000000,1,99999)
    sort(a)
