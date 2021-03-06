# -*- coding: utf-8 -*

'''
    并归排序
    通过递归的形式实现
'''
import sort_helper
import sys
sys.setrecursionlimit(10000000)
def __merage_sort(arr,l,r):
    if l >= r:
        return

    mid = l + int((r - l) / 2)
    __merage_sort(arr,l,mid)
    __merage_sort(arr,mid+1,r)
    __merage(arr,l,mid,r)

# 进行归并排序
def __merage(arr,l,mid,r):
    aux = arr[l:r+1]
    i = l
    j = mid + 1
    for k in range(l,r+1):
        if i > mid:
            arr[k] = aux[j - l]
            j +=1
        elif j > r:
            arr[k] = aux[i - l]
            i +=1
        elif aux[i-l] < aux[j - l]:
            arr[k] = aux[i-l]
            i +=1
        else:
            arr[k] = aux[j - l]
            j +=1

@sort_helper.test_sort
def sort(arr):
    __merage_sort(arr,0,len(arr)-1)

if __name__ == '__main__':
    a = sort_helper.generate_randoma_array(100000,1,99999)
    sort(a)
