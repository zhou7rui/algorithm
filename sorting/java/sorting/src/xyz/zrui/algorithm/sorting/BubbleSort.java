package xyz.zrui.algorithm.sorting;

import java.util.Arrays;

/**
 * 冒泡排序
 */
public class BubbleSort implements Sorting {

    /**
     *
     * 冒泡排序
     * @param arr 排序数组
     */
    @Override
    public void sort(Comparable[] arr) {

        int n = arr.length;

        for (int i = 0; i < n - 1; i++) {

            for (int j = 0; j < n - i - 1; j++) {
                if (arr[j].compareTo(arr[j + 1]) > 0) {
                    swap(arr, j, j + 1);
                }

            }


        }

    }

    /**
     * 冒泡排序优化1：
     * 设置开关，当第二层循环不在进行交换时说明排序完成 提前结束外层循环
     * @param arr 排序数组
     */
    public void sort1(Comparable[] arr) {

        int n = arr.length;

        for (int i = 0; i < n - 1; i++) {
            boolean flag = true;
            for (int j = 0; j < n - i - 1; j++) {
                if (arr[j].compareTo(arr[j + 1]) > 0) {
                    swap(arr, j, j + 1);
                    flag = false;
                }

            }

            if (flag){
                break;
            }

        }

    }

    /**
     *
     * 冒泡排序优化2：
     * 在遍历后记录下最后发生交换的位置
     * 下一次只需要遍历循环索引位置以前的值
     * @param arr 排序数组
     */
    public void sort2(Comparable[] arr) {

        int n = arr.length;

        int flag = n;

        while (flag > 0){

            n = flag;
            flag = 0;
            for (int i = 1; i < n; i++){
                if(arr[i].compareTo(arr[i-1]) < 0){
                    swap(arr,i,i-1);
                    flag = i;
                }
            }

        }


    }


    private void swap(Object[] arr, int i, int j) {
        Object t = arr[i];
        arr[i] = arr[j];
        arr[j] = t;
    }

    public static void main(String[] args) {

        // 测试Integer
        Integer[] arr = SortTestHelper.generateRandomArray(10000, 1, 9999);
        Integer[] newArr = arr.clone();
        SortTestHelper.testSort(BubbleSort.class,arr);
    }

}
