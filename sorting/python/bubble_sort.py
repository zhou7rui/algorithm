# -*- coding: utf-8 -*

import sort_helper

'''
    冒泡排序
'''
@sort_helper.test_sort
def bubbleSort1(arr):
    n = len(arr)
    for i in range(0,n-1,1):
        for j in range(0,n-i-1,1):
            if arr[j] > arr[j+1]:
                arr[j],arr[j + 1] = arr[j + 1], arr[j]

'''
 冒泡排序优化思路一:
    设置开关，当第二层循环不在进行交换时说明排序完成 提前结束循环
'''
@sort_helper.test_sort
def bubble_sort2(arr):
    n = len(arr)
    for i in range(0,n - 1,1):
        b = True
        for j in range(0,n - i - 1,1):
            if arr[j] > arr[j+1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                b = False
        if b:
            break


'''
 冒泡排序优化思路二:
    在遍历后记录下最后发生交换的位置
    下一次只需要遍历循环索引位置以前的值
'''
@sort_helper.test_sort
def bubble_sort3(arr):
    n = len(arr)
    flag = n
    while flag > 0:
        k = flag
        flag = 0
        for j in range(1,k,1):
            if arr[j] < arr[j-1]:
                arr[j], arr[j - 1] = arr[j - 1], arr[j]
                flag = j



if __name__ == '__main__':
    a = sort_helper.generate_randoma_array(10,1,9)
    print (a)
    bubble_sort3(a)
    print(a)
