package xyz.zrui.algorithm.sorting;

import java.util.Arrays;

/**
 * 并归排序 递归的形式
 */
public class MergeSort implements Sorting{


    /**
     * 并归排序
     * @param arr 排序数组
     * @param l 起始位置
     * @param mid 中间位置
     * @param r 结束位置
     */
    private void sort(Comparable[] arr,int l,int mid,int r){

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


    /**
     * 使用递归排序
     * @param arr 排序数组
     * @param l 数组起始位置
     * @param r 数组结束位置
     */
    private void sort(Comparable [] arr,int l, int r){
        if (l >= r){
            return;
        }

        int mid = l + (r -l) / 2;

        sort(arr,l,mid);
        sort(arr,mid + 1,r);
        sort(arr,l,mid,r);


    }

    /**
     *
     * @param arr 目标数组
     */
    @Override
    public void sort(Comparable[] arr) {

        sort(arr,0, arr.length - 1);

    }


    public static void main(String[] args) {

        Comparable[] arr = SortTestHelper.generateRandomArray(10,1,9);

        MergeSort sort = new MergeSort();

        sort.sort(arr);


    }



}
