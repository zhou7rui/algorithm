# -*- coding: utf-8 -*

'''
    快速排序
    通过递归的形式实现
'''
import random
import sort_helper

def __partition(arr,l,r):
    # 随机取值
    k = random.randint(l, r)
    arr[l], arr[k] = arr[k], arr[l]

    v = arr[l]
    j = l
    for i in range(l + 1,r + 1):
        if v > arr[i]:
            arr[j + 1], arr[i] = arr[i], arr[j + 1]
            j += 1
    arr[l],arr[j] = arr[j],arr[l]
    return j


def  __quick_sort(arr,l,r):
    if l > r:
        return;
    p = __partition(arr,l,r)
    __quick_sort(arr,l, p - 1)
    __quick_sort(arr,p + 1, r)


@sort_helper.test_sort
def sort(arg):
    __quick_sort(arg,0,len(arg) - 1)


if __name__ == '__main__':
    a = sort_helper.generate_randoma_array(1000000,1,99999)
    sort(a)
