/*
 * Copyright (C) Benesse China Company, 2017.All rights reserved.
 */

package xyz.zrui.algorithm.sorting;/*
 * Copyright (C) Benesse China Company, 2017.All rights reserved.
 */

import java.util.Arrays;

import java.util.Date;


public class MergeSort2 {


    public static void merge(Comparable[] a, int low, int mid, int high) {

        Comparable[] temp = Arrays.copyOfRange(a,low, high + 1);;
        int i = low;// 左指针
        int j = mid + 1;// 右指针
        int k = 0;
        // 把较小的数先移到新数组中
        while (i <= mid && j <= high) {
            if (a[i].compareTo(a[j]) < 0) {
                temp[k++] = a[i++];
            } else {
                temp[k++] = a[j++];
            }
        }
        // 把左边剩余的数移入数组
        while (i <= mid) {
            temp[k++] = a[i++];
        }
        // 把右边边剩余的数移入数组
        while (j <= high) {
            temp[k++] = a[j++];
        }
        // 把新数组中的数覆盖nums数组
        for (int k2 = 0; k2 < temp.length; k2++) {
            a[k2 + low] = temp[k2];
        }
    }

    public static void mergeSort(Comparable [] a, int low, int high) {
        int mid = (low + high) / 2;
        if (low < high) {
            // 左边
            mergeSort(a, low, mid);
            // 右边
            mergeSort(a, mid + 1, high);
            // 左右归并
            merge(a, low, mid, high);
//            System.out.println(Arrays.toString(a));
        }

    }

    public static void main(String[] args) {

        Comparable[] arr = SortTestHelper.generateRandomArray(1000000, 1, 999999);

        long stime = new Date().getTime();
        mergeSort(arr, 0, arr.length - 1);
        long etime = new Date().getTime();
        System.out.println("耗时:" + (etime - stime) + "ms");


    }
}



