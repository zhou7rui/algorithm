package xyz.zrui.algorithm.sorting;


/**
 *  快速排序
 *  递归实现
 *
 */
public class QuickSort implements Sorting {

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


    /**
     * 将目标 分割成 两块
     * 左边小于 l 位置的值
     * 右边大于 l 位置的值
     * @param arr  目标数组
     * @param l 分割的范围 左边的值
     * @param r 分割的范围 右边的值
     * @return  返回分割的位置
     */
    private int partition(Comparable[] arr,int l,int r){

        Comparable v = arr[l];
        int j = l;
        for (int i = l + 1; i < r + 1; i ++){
            if ( v.compareTo(arr[i]) > 0) {
                swap(arr,j + 1,i);
                j ++;
            }
        }
        swap(arr,l,j);

        return  j;
    }

    /**
     * 递归
     * @param arr 待排序数组
     * @param l 排序的范围 左边的值
     * @param r 排序的范围 右边的值
     */
    private void sort(Comparable[] arr,int l,int r){

        // 递归跳出条件
        if (l > r){
            return;
        }

        int p = partition(arr,l,r);

        sort(arr,l,p -1);

        sort(arr,p+1,r);


    }


    @Override
    public void sort(Comparable[] arr) {

        sort(arr,0,arr.length - 1);

    }


    public static void main(String[] args) {

        Comparable[] arr = SortTestHelper.generateRandomArray(100000,1,99999);


       SortTestHelper.testSort(QuickSort.class,arr);


    }

}
