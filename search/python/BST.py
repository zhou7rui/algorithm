# -*- coding: utf-8 -*
'''
    二分搜索树
'''
import Queue
class BST(object):
    """docstring for BTR."""
    def __init__(self):
        self.root = None
        self.count = 0

    class Node(object):
        """docstring for Node."""
        def __init__(self, key,value):
            self.key = key
            self.value = value
            self.left = None
            self.right = None

    def size(self):
        return self.conut

    def isEmpty(self):
        return self.count == 0

    def __insert(self,node,key,value):

        if node == None:
            self.count += 1
            return self.Node(key,value)

        if node.key > key:
            node.left = self.__insert(node.left,key,value)

        elif node.key < key:
            node.right = self.__insert(node.right,key,value)
        else:
            node.value = value

        return node

    def insert(self,key,value):
        self.root = self.__insert(self.root,key,value)


    #查看以 node 为根二叉树 是否包含键值为 key的节点
    def contain(self,node,key):

        if node == None:
            return False

        if node.key > key:
            return self.contain(node.left,key)
        elif node.key < key:
            return self.contain(node.right,key)
        else:
            return True
    # 搜索以 node 为根二叉树 返回包含键值为 key的节点的value
    def search(self,node,key):

        if node == None:
            return None

        if node.key == key:
            return node.value
        elif node.key > key:
            return search(node.left,key)
        else:
            return search(node.right,key)

    #前序遍历
    def preOrder(self,node):

        if node != None:
            print(node.key)
            self.preOrder(node.left)
            self.preOrder(node.right)

    #中序遍历
    def inOrder(self,node):

        if node != None:
            self.inOrder(node.left)
            print(node.key)
            self.inOrder(node.right)

    #后序遍历
    def postOrder(self,node):

        if node != None:
            self.postOrder(node.left)
            self.postOrder(node.right)
            print(node.key)

    #层序遍历
    def levelOrder(self,node):

        q = Queue.Queue()
        q.put(node)
        while not q.empty():
            temp = q.get()
            print(temp.key)
            if not temp.left  == None:
                q.put(temp.left)
            if not temp.right == None:
                q.put(temp.right)

    # 返回node为根的最小节点
    def __minimum(self,node):
        if  node.left == None:
            return node

        return self.__minimum(node.left)

    #返回最小节点的key
    def minimum(self):
        if not self.root == None:
            node = self.__minimum(self.root)
            return node.key

    # 返回node为根的最大节点
    def __maximum(self,node):
        if node.right == None:
            return node
        return self.__maximum(node.right)

    #返回最大节点的key
    def maximum(self):
        if not self.root == None:
            node = self.__maximum(self.root)
            return node.key

    # 删除掉以node为根的二分搜索树中的最大节点
    # 返回删除节点后新的二分搜索树的根
    def __removeMax(self,node):
        if node.right == None:
            leftNode = node.left
            self.count -= 1
            return leftNode
        node.right = self.__removeMax(node.right)
        return node

    #删除最大节点
    def removeMax(self):
        if not self.root == None:
            self.root = self.__removeMax(self.root)



    # 删除掉以node为根的二分搜索树中的最小节点
    # 返回删除节点后新的二分搜索树的根
    def __removeMin(self,node):
        if node.left == None:
            rightNode = node.right
            self.count -= 1
            return rightNode

        node.left = self.__removeMax(node.left)
        return node

    #删除最小节点
    def removeMin(self):
        if not self.root == None:
            self.root = self.__removeMin(self.root)




if __name__ == '__main__':
    bst = BST()
    bst.insert(2,"AA")
    bst.insert(223,"BB")
    bst.insert(21,"CC")
    bst.insert(25,"DD")
    bst.insert(123,"EE")
    bst.insert(99,"FF")
    bst.insert(235,"GG")
    bst.inOrder(bst.root)
    print("max",bst.maximum())
    print("min",bst.minimum())
    print("************************")
    bst.removeMax()
    bst.inOrder(bst.root)
