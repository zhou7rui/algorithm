package xyz.zrui.algorithm.sorting;

public class HeapSort implements Sorting {


    private void shiftDown(Comparable[] arr, int n, int k) {

        while (n >  k * 2 + 1) {

            int j = 2 * k  + 1;

            if (n > j + 1 &&  arr[ j + 1].compareTo(arr[j]) > 0 ) {
                j += 1;
            }

            if (arr[j].compareTo(arr[k]) <= 0){
                break;
            }
            swap(arr,k,j);

            k = j;
        }

    }

    private static void swap(Object[] arr, int i, int j) {
        Object t = arr[i];
        arr[i] = arr[j];
        arr[j] = t;
    }

    @Override
    public void sort(Comparable[] arr) {
        int n = arr.length;
        // 构建堆
        for (int i = ( n + 1) / 2; i >=0 ; i --) {
            shiftDown(arr,n,i);
        }

        //将根节点与尾节点交换后尾节点前的重新生成堆。
        for (int i = n - 1 ; i > 0; i --){
            swap(arr,i,0);
            shiftDown(arr,i,0);
        }

    }

    public static void main(String[] args) {
        Comparable[] arr2 = SortTestHelper.generateRandomArray(100000, 1, 99999);
        Comparable[] arr3 = arr2.clone();
        SortTestHelper.testSort(HeapSort.class, arr2);
        SortTestHelper.testSort(Quick3Sort.class, arr3);

    }


}
