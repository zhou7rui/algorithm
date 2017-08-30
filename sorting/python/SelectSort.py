# -*- coding: utf-8 -*
'''
    选择排序
'''
def selectSort(arr):
    n = len(arr)
    for i in range(0,n,1):
        minIndex = i;
        #找到最小的一个值
        for j in range(i+1,n,1):
            if arr[j] < arr[minIndex]:
                minIndex = j

        # sawp 操作 交换位置
        temp = arr[i]
        arr[i] = arr[minIndex]
        arr[minIndex] = temp

if __name__ == '__main__':
    arr = [3, 4, 2, 5, 8, 6, 9, 1]
    selectSort(arr)
    for i in arr:
        print(i)
