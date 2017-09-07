# -*- coding: utf-8 -*

import sortHelper

'''
    冒泡排序
'''
@sortHelper.testSort
def bubbleSort1(arr):
    n = len(arr)
    for i in range(0,n-1,1):
        for j in range(0,n-i-1,1):
            if arr[j] > arr[j+1]:
                temp = arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = temp


'''
 冒泡排序优化思路一:
    设置开关，当第二层循环不在进行交换时说明排序完成 提前结束循环
'''
@sortHelper.testSort
def bubbleSort2(arr):
    n = len(arr)
    for i in range(0,n - 1,1):
        b = True
        for j in range(0,n - i - 1,1):
            if arr[j] > arr[j+1]:
                temp = arr[j]
                arr[j] = arr[j+1]
                arr[j+1] = temp
                b = False
        if b:
            break


'''
 冒泡排序优化思路二:
    在遍历后记录下最后发生交换的位置
    下一次只需要遍历循环索引位置以前的值
'''
@sortHelper.testSort
def bubbleSort3(arr):
    n = len(arr)
    flag = n
    while flag > 0:
        k = flag
        flag = 0
        for j in range(1,k,1):
            if arr[j] < arr[j-1]:
                temp = arr[j]
                arr[j] = arr[j-1]
                arr[j-1] = temp
                flag = j



if __name__ == '__main__':
    a = sortHelper.generateRandomArray(10,1,9)
    print (a)
    bubbleSort3(a)
    print(a)