package xyz.zrui.algorithm.sorting;


/**
 * 插入排序
 */
public class InsertingSort implements Sorting {


    @Override
    public void sort(Comparable[] arr) {

        for (int i = 1; i < arr.length; i++) {

            // 向前遍历 是否比前一位小 是)交换位置 否)跳出循环
            for (int j = i; j > 0; j--) {
                // 使用compareTo方法比较两个Comparable对象的大小
                if (arr[j].compareTo(arr[j - 1]) < 0) {
                    swap(arr, j, j - 1);
                } else {
                    break;
                }
            }

        }

    }

    public void sort1(Comparable[] arr) {

        for (int i = 1; i < arr.length; i++) {

            //插入值
            Comparable temp = arr[i];

            int j;
            // 向前遍历 插入值是否比前一位小
            for (j = i; j > 0 && (arr[j-1].compareTo(temp) > 0); j-- ) {
                //向后覆盖一位
                arr[j] = arr[j-1];
            }
            arr[j] = temp;

        }

    }




    /**
     * 交换数组间的位置
     *
     * @param arr 目标数组
     * @param i   交换位置索引
     * @param j   交换位置索引
     */
    private static void swap(Object[] arr, int i, int j) {
        Object t = arr[i];
        arr[i] = arr[j];
        arr[j] = t;
    }

    public static void main(String[] args) {

        // 测试Integer
        Integer[] a = SortTestHelper.generateRandomArray(10,1,10);
        InsertingSort sort = new InsertingSort();
        sort.sort1(a);
        SortTestHelper.printArray(a);
    }

}
