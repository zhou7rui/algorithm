package xyz.zrui.algorithm.heap;

public class IndexMxaHeap<T extends Comparable> {

    protected T[] data;

    protected int count;

    protected Integer[] indexes;

    protected int capacity;


    public IndexMxaHeap(int capacity){

        this.indexes = new Integer[capacity + 1];

        this.data = (T[]) new Comparable[capacity + 1];

        this.count = 0;

        this.capacity = capacity;

    }



    private void swap(Object[] arr,int i,int j){

        Object temp = arr[i];

        arr[i] = arr[j];

        arr[j] = temp;

    }



    private void shiftUp(int k){

        while ( k > 1 && data[indexes[k / 2]].compareTo(data[indexes[k]]) < 0){
            swap(indexes,k, k / 2);
            k /= 2;
        }

    }


    private void shiftDown(int k){
        while (k * 2 <= count ) {

            int j = k * 2;
            if (j + 1 <= count && data[indexes[j + 1]].compareTo(data[indexes[j]]) > 0) {
                j += 1;
            }

            if (data[indexes[j]].compareTo(data[indexes[k]]) <= 0){
                break;
            }

            swap(indexes,j,k);

            k = j;
        }
    }

    public T getItem(int i){
        return data[i + 1];
    }


    public int extractMaxIndex() {

        int ret = indexes[1] - 1;

        swap(indexes, count, 1);

        count--;

        shiftDown(1);
        return ret;
    }

    public T extractMax() {

        T ret = data[indexes[1]];

        swap(indexes, count, 1);

        count--;

        shiftDown(1);
        return ret;
    }


    public void insert(int i, T t){

        i += 1;
        data[i] = t;
        indexes[count + 1] = 1;

        count ++;
        shiftUp(count);


    }

    void change(int i, T t){

        data[i + 1] = t;
        // 找到indexes[j] == i, i 表示data[i]在堆中的元素的位置
        // 之后shiftUp(j),再shiftDown(j)
        for (int j = 1; j <= count; j ++){
            if(indexes[j] == i){
                shiftDown(j);
                shiftUp(j);
            }
        }

    }




}
