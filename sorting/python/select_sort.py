# -*- coding: utf-8 -*
'''
    选择排序
'''
import sortHelper

@sortHelper.testSort
def selectSort(arr):
    n = len(arr)
    for i in range(0,n,1):
        # 寻找[i, n)区间里的最小值的索引
        minIndex = i;
        for j in range(i+1,n,1):
            if arr[j] < arr[minIndex]:
                minIndex = j

        # sawp 操作 交换位置
        temp = arr[i]
        arr[i] = arr[minIndex]
        arr[minIndex] = temp

if __name__ == '__main__':
    arr = sortHelper.generateRandomArray(10000,1,9999)
    selectSort(arr)
    print(arr)
