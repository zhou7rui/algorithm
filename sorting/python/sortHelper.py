# -*- coding: utf-8 -*
import random

def generateRandomArray(n,rangeL,rangeR):

    if rangeL >= rangeR:
        return;
    arr = [None] * n
    for i in range(0,n,1):
        arr[i] = random.randint(rangeL, rangeR)
    return arr




if __name__ == '__main__':
    print(generateRandomArray(10,1,9))
