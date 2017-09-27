package xyz.zrui.algorithm.sorting;

import java.util.Random;

/**
 * 三路快速排序
 */
public class Quick3Sort implements Sorting {


    /**
     * 交换数组间的位置
     *
     * @param arr 目标数组
     * @param i   交换位置索引
     * @param j   交换位置索引
     */
    private void swap(Object[] arr, int i, int j) {
        Object t = arr[i];
        arr[i] = arr[j];
        arr[j] = t;
    }

    private void sort(Comparable[] arr, int l, int r) {

        if (l > r){
            return;
        }

        Random random = new Random();
        int k = random.nextInt(r - l + 1) + l;
        swap(arr, l, k);

        Comparable v = arr[l];

        int lt = l;   // arr[l+1...lt] < v

        int gt = r + 1;  // arr[gt - 1...r] > v

        int i = l + 1;    // arr[lt+1....i) == v

        while (i < gt){
            if (arr[i].compareTo(v) < 0){                // arr[l+1...lt] < v
                swap(arr,i,lt + 1);
                lt ++;
                i ++;
            }else if (arr[i].compareTo(v) > 0){      // arr[gt - 1...r] > v
                swap(arr,i,gt - 1);
                gt --;
            }else {                               // arr[lt+1....i) == v
                i ++;
            }
        }
        swap(arr,l,lt);

        sort(arr,l,lt - 1);
        sort(arr, gt,r);
    }

    @Override
    public void sort(Comparable[] arr) {

        sort(arr, 0, arr.length - 1);

    }

    public static void main(String[] args) {

        Comparable[] arr2 = SortTestHelper.generateRandomArray(10000000, 1, 9);
        Comparable[] arr3 = arr2.clone();
        SortTestHelper.testSort(Quick2Sort.class, arr2);
        SortTestHelper.testSort(Quick3Sort.class, arr3);
    }
}
