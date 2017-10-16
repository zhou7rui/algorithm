# -*- coding: utf-8 -*-
#栈的实现
class Stack(object):
    #初始化栈的大小
    def __init__(self,size):
        self.stack = []
        self.size = size
        self.top = -1

    #入栈操作，判断是否为空，栈顶的位置加一
    def push(self,content):
        if self.isFull():
            print("Stack is Full!")
        else:
            self.stack.append(content)
            self.top = self.top + 1

    #判断是否已满
    def isFull(self):
        if self.top == self.size:
            return True
        else:
            return False

    #判断是否为空
    def isEmpty(self):
        if self.top == -1:
            return True
        else:
            return False

    #出栈
    def out(self):
        if self.isEmpty():
            print("Stack is Empty!")
        else:
            self.top = self.top - 1
            self.stack.pop()

    def peek(self):
        if self.isEmpty():
            print("Stack is Empty!")
        else:
            return self.stack[self.top]
