# -*- coding: utf-8 -*

'''
    并归排序
    通过循环 至底向上的的形式实现
'''
import sort_helper
def merge(arr,l,mid,r):
    aux = arr[l:r+1]
    i = l
    j = mid + 1
    for k in range(l,r+1,1):
        if i > mid:
            arr[k] = aux[j - l]
            j +=1
        elif j > r:
            arr[k] = aux[i - l]
            i +=1
        elif aux[i-l] < aux[j - l]:    #左半部分所指元素 < 右半部分所指元素
            arr[k] = aux[i-l]
            i +=1
        else:
            arr[k] = aux[j - l]
            j +=1

def sort(arr):
    n = len(arr)
    sz = 1
    while sz < n:
        i = 0
        while i < n - sz:
            merge(arr, i, i + sz - 1, min(n - 1, 2 * sz + i - 1))
            i += 2 * sz
        sz *= 2

if __name__ == '__main__':
    arr = sort_helper.generate_randoma_array(10,1,9)
    print(arr)
    sort(arr)
    print(arr)
