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

    def __init__(self,cpacity):
        self.cpacity = cpacity
        self.count = 0
        self.data = [None] * cpacity

    # def __init__(self,arr,n):
    #     self.data = [n for n in arr]
    #     self.count = len(arr)
    #     self.cpacity = len(arr)
    #     for i in range(self.count / 2,0, -1):
    #         self.__shifDown(i)

    def __init__(self):
        pass


    def size(self):
        return self.count

    def isEmpty(self):
        return self.count == 0

    def __shiftUp(self,k):

        while k > 1 and self.data[k] < self.data[k / 2]:
            self.data[k],self.data[k / 2] = self.data[k / 2], self.data[k]
            k /= 2


    def insert(self,data):
        self.data[self.count + 1] = data
        self.count += 1
        self.__shiftUp(self.count)

    def __shifDown(self,k):
        while k * 2 <= count:

            j = k * 2

            if count >= j + 1 and self.data[j + 1] > self.data[j]:
                j += 1

            if self.data[k] > self.data[j]:
                break

            self.data[k], self[j] = self.data[j],self.data[k]

            k = j

    def extractMax(self):
        ret = self.data[1]
        self.data[1], self.data[count] = self.data[count], self.data[1]
        self.count -= 1
        self.__shifDown(1)
        return ret






if __name__ == '__main__':

    N = 31
    M = 100
    heap = Maxheap(N)
    for i in range(1,N):
        k = random.randint(1, M)
        heap.insert(k)
    print(heap.size())
    print(heap.extractMax())
