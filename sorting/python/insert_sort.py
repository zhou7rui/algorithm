# -*- coding: utf-8 -*

'''
    插入排序
'''
import sort_helper

@sort_helper.test_sort
def insert_sort(arr):
    n = len(arr)
    for i in range(1,n,1):
        temp = arr[i]
        j = i
        #向前查找插入位置
        while j > 0 and arr[j-1] > temp:
            arr[j] = arr[j-1]
            j -= 1

        arr[j] = temp

@sort_helper.test_sort
def insert_sort1(arr):
    for i in range(1,len(arr),1):
        #向前查找插入位置
        for j in range(i,0,-1):
            if arr[j] < arr[j-1]:
                #sawp 交换位置
                temp = arr[j-1]
                arr[j-1] = arr[j]
                arr[j] = temp
            else:
                break


if __name__ == '__main__':
    arr = sort_helper.generate_randoma_array(10,1,9)
    print(arr)
    insert_sort1(arr)
    print(arr)
