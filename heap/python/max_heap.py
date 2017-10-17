# -*- coding: utf-8 -*

'''
                        最大堆实现
                           98
                /                      \
               96                      84
          /          \            /          \
         92          82          78          47
       /    \      /    \      /    \      /    \
      33    26    51    85    50    15    44    60
     / \   / \   / \   / \   / \   / \   / \   / \
    40 51 98 51 7  17 94 82 32 21 64 60 7  44 63 63
'''
import random

class Maxheap(object):

    def __init__(self,cpacity,arr = None):

        self.data = [None] * (cpacity + 1)
        self.cpacity = cpacity
        if arr is None:
            self.count = 0
        else:
            for i in range(0,cpacity):
                self.data[i + 1]= arr[i]
            self.count = cpacity;
            for i in range(self.count / 2, 0, -1):
                self.__shifDown(i)

    def size(self):
        return self.count

    def isEmpty(self):
        return self.count == 0

    def __shiftUp(self,k):

        while k > 1 and self.data[k] > self.data[k / 2]:
            self.data[k],self.data[k / 2] = self.data[k / 2], self.data[k]
            k /= 2


    def insert(self,data):
        self.data[self.count + 1] = data
        self.count += 1
        self.__shiftUp(self.count)

    def __shifDown(self,k):
        while k * 2 <= self.count:

            j = k * 2

            if self.count >= j + 1 and self.data[j + 1] > self.data[j]:
                j += 1

            if self.data[k] > self.data[j]:
                break

            self.data[k], self.data[j] = self.data[j],self.data[k]

            k = j

    def extractMax(self):
        ret = self.data[1]
        self.data[1], self.data[self.count] = self.data[self.count], self.data[1]
        self.count -= 1
        self.__shifDown(1)
        return ret






if __name__ == '__main__':

    N = 31
    M = 100
    heap = Maxheap(N)
    for i in range(0,N):
        k = random.randint(1, M)
        heap.insert(k)

    # arr = [random.randint(1,M) for i in range(N)]
    # heap = Maxheap(len(arr),arr)
#[None, 95, 84, 85, 78, 80, 50, 74, 62, 70, 59, 51, 38, 47, 57, 71, 9, 16, 19, 53, 18, 25, 44, 7, 28, 2, 24, 16, 18, 13, 33, 70]
    print(heap.size())
    print(heap.data)
    print(heap.extractMax())
