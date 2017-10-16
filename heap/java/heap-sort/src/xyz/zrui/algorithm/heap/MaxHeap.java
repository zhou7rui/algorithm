package xyz.zrui.algorithm.heap;

public class MaxHeap<T extends Comparable> {

    protected T[] data;

    protected int count;

    protected int capacity;

    //构造函数， 构造一个空的堆，可容纳capacity个元素
    public MaxHeap(int capacity){

        this.data = (T[]) new Comparable[capacity+1];

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

    private void shiftDown(int k){

        while (2 * k <= count){

            int j = 2 * k;

            if (j + 1 <= count && data[j + 1].compareTo(data[j]) > 0){
                j += 1;
            }
            if (data[k].compareTo(data[j]) > 0){
                break;
            }
            swap(data,j,k);
            k = j;
        }



    }

    public T extractMax(){
        T ret = data[1];
        swap(data,count,1);
        count --;
        shiftDown(1);
        return ret;
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


    // 测试 MaxHeap
    public static void main(String[] args) {

        MaxHeap<Integer> maxHeap = new MaxHeap<Integer>(100);
        int N = 50; // 堆中元素个数
        int M = 100; // 堆中元素取值范围[0, M)
        for( int i = 0 ; i < N ; i ++ )
            maxHeap.insert( new Integer((int)(Math.random() * M)) );
        System.out.println(maxHeap.size());

    }


}
