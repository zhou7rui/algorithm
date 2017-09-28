# -*- coding: utf-8 -*

'''
    三路快速排序
'''
import random
import sort_helper



def __sort(arr,l,r):

    if l > r:
        return
    k = random.randint(l, r)
    arr[l], arr[k] = arr[k], arr[l]

    v = arr[l]

    lt = l              #arr[l+1....lt] < v
    gt = r + 1          #arr[r...gt-1] > v
    i = l + 1           #arr[lt+1 ...i) == v

    while i > gt:
        if arr[i] < v:
            arr[lt + 1] , arr[i] = arr[i], arr[lt+1]
            i += 1
            lt += 1
        elif arr[i] > v:
            arr[gt - 1],arr[i] = arr[i], arr[gt - 1]
            gt -= 1
        else:
            i += 1
    arr[l], arr[lt] = arr[lt],arr[l]

    __sort(arr,l,lt - 1)
    __sort(arr,gt,r)

@sort_helper.test_sort
def sort(arr):
    __sort(arr,0,len(arr) - 1)


if __name__ == '__main__':
    a = sort_helper.generate_randoma_array(1000000,1,99999)
    sort(a)
