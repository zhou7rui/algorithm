package xyz.zrui.algorithm.sorting;


import java.util.Random;

/**
 * 双路快速排序
 */
public class Quick2Sort implements Sorting {


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


    private int partition2(Comparable[] arr, int l, int r) {

        Random random = new Random();
        int k = random.nextInt(r - l + 1) + l;
        swap(arr, l, k);

        Comparable v = arr[l];

        //arr[l+1....i] <= v    arr[r...j] >=v
        int i = l + 1;
        int j = r;
        while (true) {

            while (i <= r && arr[i].compareTo(v) < 0) {
                i++;
            }

            while (j >= l + 1 && arr[j].compareTo(v) > 0) {
                j--;
            }
            if (i > j) {
                break;
            }
            swap(arr, i, j);
            i++;
            j--;
        }
        // j 为从右往左的最后一个 小于等于v
        swap(arr, j, l);
        return j;

    }

    private void sort(Comparable[] arr, int l, int r) {
        if (l > r) {
            return;
        }

        int p = partition2(arr, l, r);
        sort(arr, l, p - 1);
        sort(arr, p + 1, r);

    }


    @Override
    public void sort(Comparable[] arr) {

        sort(arr, 0, arr.length - 1);

    }

    public static void main(String[] args) {

        Comparable[] arr = SortTestHelper.generateRandomArray(10, 1, 9);
        SortTestHelper.printArray(arr);
        SortTestHelper.testSort(Quick2Sort.class, arr);
    }


}
