package xyz.zrui.algorithm.search;


/**
 *  二叉搜索树
 * @param <Key>
 * @param <Value>
 */
public class BST<Key extends Comparable<Key>, Value> {


    private class Node {

        private Key key;

        private Value value;

        private Node left;

        private Node right;

        public Node(Key key, Value value) {
            this.key = key;
            this.value = value;
        }

        public Key getKey() {
            return key;
        }

        public void setKey(Key key) {
            this.key = key;
        }

        public Value getValue() {
            return value;
        }

        public void setValue(Value value) {
            this.value = value;
        }

        public Node getLeft() {
            return left;
        }

        public void setLeft(Node left) {
            this.left = left;
        }

        public Node getRight() {
            return right;
        }

        public void setRight(Node right) {
            this.right = right;
        }

    }

    private Node root;
    private int count;


    // 返回二分搜索树的节点个数
    public int size() {
        return count;
    }

    // 返回二分搜索树是否为空
    public boolean isEmpty() {
        return count == 0;
    }


    // 向以node为根的二分搜索树中, 插入节点(key, value), 使用递归算法
    // 返回插入新节点后的二分搜索树的根
    private Node insert(Node node, Key key, Value value) {

        if (node == null) {
            count++;
            return new Node(key, value);
        }

        if (node.getKey().compareTo(key) > 0) {

            node.setLeft(insert(node.getLeft(), key, value));

        } else if (node.getKey().compareTo(key) < 0) {

            node.setRight(insert(node.getRight(), key, value));

        } else {
            node.setValue(value);
        }
        return node;
    }

    // 查看以 node 为根二叉树 是否包含键值为 key的节点
    public boolean contain(Node node, Key key) {

        if (node == null) {
            return false;
        }

        if (key.equals(node.getKey())) {
            return true;
        } else if (key.compareTo(node.getKey()) < 0) {
            return contain(node.getLeft(), key);
        } else {
            return contain(node.getRight(), key);

        }

    }

    public void insert(Key key, Value value) {
        root = insert(root, key, value);
    }

    // 搜索以 node 为根二叉树 返回包含键值为 key的节点的value
    public Value search(Node node, Key key) {
        if (node == null) {
            return null;
        }

        if (key.equals(node.getKey())) {
            return node.getValue();
        } else if (key.compareTo(node.getKey()) < 0) {
            return search(node.getRight(), key);
        } else {
            return search(node.getRight(), key);
        }

    }



    public void preOrder(Node node){

        if(node != null){
            System.out.println(node.getKey());
            preOrder(node.getLeft());
            preOrder(node.getRight());
        }

    }

    public void inOrder(Node node){

        if(node != null){
            inOrder(node.getLeft());
            System.out.println(node.getKey());
            inOrder(node.getRight());
        }

    }

    public void postOrder(Node node){

        if(node != null){
            postOrder(node.getLeft());
            postOrder(node.getRight());
            System.out.println(node.getKey());
        }

    }



    public Node getRoot() {
        return root;
    }

    public void setRoot(Node root) {
        this.root = root;
    }

    public int getCount() {
        return count;
    }

    public void setCount(int count) {
        this.count = count;
    }


    public static void main(String[] args) {

        BST<Integer,String> bst = new BST<>();

        bst.insert(34,"AA");
        bst.insert(323,"BB");
        bst.insert(32,"CC");
        bst.insert(33,"DD");
        bst.insert(24,"EE");
        bst.insert(14,"FF");
        bst.insert(84,"GG");

        bst.inOrder(bst.root);




    }



}
