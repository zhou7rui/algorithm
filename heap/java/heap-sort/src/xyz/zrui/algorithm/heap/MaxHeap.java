package xyz.zrui.algorithm.heap;

public class MaxHeap<T extends Comparable> {

    private T[] data;

    private int count;

    private int capacity;

    //构造函数， 构造一个空的堆，可容纳capacity个元素
    public MaxHeap(int capacity){

        this.data = (T[]) new Object[capacity+1];

        this.count = 0;

        this.capacity = capacity;
    }

    public int size(){
        return count;
    }

    public boolean isEmpty(){
        return count == 0;
    }

    public void insert(T t){
        data[count + 1] = t;
        count ++;
        shiftUp(count);
    }

    private void shiftUp(int k){

        while (k > 1 && data[k/2].compareTo(data[k]) < 0){
            swap(data,k,k/2);
            k /= 2;
        }

    }

    /**
     * 交换数组间的位置
     *
     * @param arr 目标数组
     * @param i   交换位置索引
     * @param j   交换位置索引
     */
    private  void swap(Object[] arr, int i, int j) {
        Object t = arr[i];
        arr[i] = arr[j];
        arr[j] = t;
    }



}
