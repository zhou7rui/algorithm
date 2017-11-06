# -*- coding: utf-8 -*
'''
    二分搜索树
'''
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
        print("insert",self.root.key)

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

    def preOrder(self,node):

        if node != None:
            print(node.key)
            self.preOrder(node.left)
            self.preOrder(node.right)

    def inOrder(self,node):

        if node != None:
            self.inOrder(node.left)
            print(node.key)
            self.inOrder(node.right)

    def postOrder(self,node):

        if node != None:
            self.postOrder(node.left)
            self.postOrder(node.right)
            print(node.key)




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
