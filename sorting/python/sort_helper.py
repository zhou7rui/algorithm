# -*- coding: utf-8 -*
import random
import time

def generate_randoma_array(n,rangeL,rangeR):

    if rangeL >= rangeR:
        return;
    arr = [None] * n
    for i in range(0,n,1):
        arr[i] = random.randint(rangeL, rangeR)
    return arr

def is_sort(arr):
    n = len(arr)
    for i in range(0,n,1):
        if arr[i] < arr[i+1]:
            return False
    return True

def test_sort(func):
    def wrapper(*args,**kw):
        startTime = int(time.time()*1000)
        func(*args,**kw)
        endTime = int(time.time()*1000)
        # print(func.__name__,"runTime" + endTime - startTime + "ms")
        print("%s runTime %d ms"%(func.__name__,endTime - startTime))
    return wrapper
