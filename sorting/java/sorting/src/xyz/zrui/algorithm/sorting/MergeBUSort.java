package xyz.zrui.algorithm.sorting;

import java.util.Arrays;

/**
 * 并归排序  从底向上排序
 */
public class MergeBUSort implements Sorting {


    private void merge(Comparable[] arr, int l, int mid, int r) {

       Comparable[] aux = Arrays.copyOfRange(arr,l,r+1);

        int i = l, j = mid + 1;
        for (int k = l ; k < r + 1;k++){
            if (i > mid){                   //如果左半边的元素处理完了
                arr[k] = aux[j - l];
                j ++;
            } else if ( j > r) {           //如果右半边的元素处理完了
                arr[k] = aux[i - l];
                i ++;
            }else if (aux[i - l].compareTo(aux[j - l]) < 0){ // 左半部分所指元素 < 右半部分所指元素
                arr[k] = aux[i - l];
                i ++;
            }else{          // 左半部分所指元素 >= 右半部分所指元素
                arr[k] = aux[j - l];
                j ++;
            }

        }
    }


    @Override
    public void sort(Comparable[] arr) {

        int n = arr.length;
        for (int sz = 1; sz < n; sz *= 2) {
            for (int i = 0; i < n - sz; i += 2 * sz) {
                merge(arr, i, i + sz - 1, Math.min(n - 1, 2 * sz + i - 1));
            }

        }

    }

    public static void main(String[] args) {

        Comparable[] arr = SortTestHelper.generateRandomArray(10, 1, 9);

        SortTestHelper.testSort(MergeBUSort.class, arr);
    }


}
