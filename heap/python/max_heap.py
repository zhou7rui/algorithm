# -*- coding: utf-8 -*

class Maxheap(object):

    def __init__(self,cpacity):
        self.cpacity = cpacity
        self.count = 0
        self.data = [None] * cpacity

    def size(self):
        return self.count

    def isEmpty(self):
        return self.count == 0

    def __shiftUp(self,k):

        while k > 1 and self.data[k] < self.data[k / 2]:
            self.data[k],self.data[k / 2] = self.data[k / 2], self.data[k]
            k /= 2


    def insert(self,data):
        self.data[count + 1] = data
        count += 1
        __shiftUp(count)

    def __shifDown(self,k):
        while k * 2 <= count:

            j = k * 2

            if count >= j + 1 and self.data[j + 1] > self.data[j]:
                j += 1

            if self.data[k] > self.data[j]
                break

            self.data[k], self[j] = self.data[j],self.data[k]

            k = j

    def extractMax(self):
        ret = self.data[1]
        self.data[1], self.data[count] = self.data[count], self.data[1]
        count -= 1
        __shifDown(1)
        return ret






if __name__ == '__main__':
    a = Maxheap(1)
    print(a.data)
    print(a.size())
    print(a.isEmpty())
