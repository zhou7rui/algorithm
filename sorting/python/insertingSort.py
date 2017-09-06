# -*- coding: utf-8 -*

'''
    插入排序
'''
import sortHelper

@sortHelper.testSort
def insertingSort(arr):
    n = len(arr)
    for i in range(0,n,1):

        for j in range(i,0,-1):
            if arr[j] < arr[j-1]:
                # sawp 操作 交换位置
                temp = arr[j-1]
                arr[j] = arr[j-1]
                arr[j-1] = temp
            else:
                break


if __name__ == '__main__':
    arr = sortHelper.generateRandomArray(10000,1,10000)
    insertingSort(arr)
    # print(arr)
