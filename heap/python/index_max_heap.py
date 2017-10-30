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

class IndexMxaHeap(object):

    def __init__(self,cpacity):

        self.data = [None] * (cpacity + 1)
        self.cpacity = cpacity
        self.indexes = [None] * (cpacity + 1)
        self.count = 0

    def size(self):
        return self.count

    def isEmpty(self):
        return self.count == 0

    def __shiftUp(self,k):

        while k > 1 and self.data[self.indexes[k]] > self.data[self.indexes[k / 2]]:
            self.indexes[k],self.indexes[k / 2] = self.indexes[k / 2], self.indexes[k]
            k /= 2


    def insert(self,i,data):
        i += 1
        self.data[i] = data
        self.indexes[self.count+1] = i
        self.count += 1
        self.__shiftUp(self.count)

    def __shifDown(self,k):

        while k * 2 <= self.count:

            j = k * 2

            if self.count >= j + 1 and self.data[self.indexes[j + 1]] > self.data[self.indexes[j]]:
                j += 1

            if self.data[self.indexes[k]] > self.data[self.indexes[j]]:
                break

            self.indexes[k], self.indexes[j] = self.indexes[j],self.indexes[k]

            k = j

    def extractMax(self):
        ret = self.data[indexes[1]]
        self.indexes[1], self.indexes[self.count] = self.indexes[self.count], self.indexes[1]
        self.count -= 1
        self.__shifDown(1)
        self.data.pop()
        return ret


    def extractMaxIndex(self):
        ret = indexes[1] - 1
        self.indexes[1], self.indexes[self.count] = self.indexes[self.count], self.indexes[1]
        self.count -= 1
        self.__shifDown(1)
        del self.data[indexes[-1]]
        self.indexes.pop()
        return ret

    def change(self,k,newData):
        data[k + 1] = newData
        # 找到indexes[j] == i, i 表示data[i]在堆中的元素的位置
        # 之后shiftUp(j),再shiftDown(j)
        for i in range(len(self.indexes)):
            if self.indexes[i] == k:
                self.__shifDown(k)
                self.__shiftUp(k)

    def getItem(self,i):
        return self.data[i + 1]


if __name__ == '__main__':


    arr  =  [15,17,19,13,22,16,28,30,41,62]
    heap = IndexMxaHeap(len(arr))

    for i in range(0,len(arr)):
        heap.insert(i,arr[i])

    print(heap.indexes)
    print(heap.data)
