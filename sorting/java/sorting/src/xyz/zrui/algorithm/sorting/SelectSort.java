package xyz.zrui.algorithm.sorting;

/**
 * 选择排序
 */
public class SelectSort {

    // 我们的算法类不允许产生任何实例
    public static void sort(Comparable[] arr) {

        int n = arr.length;
        for (int i = 0; i < n; i++) {
            // 寻找[i, n)区间里的最小值的索引
            int minIndex = i;
            for (int j = i + 1; j < n; j++)
                // 使用compareTo方法比较两个Comparable对象的大小
                if (arr[j].compareTo(arr[minIndex]) < 0)
                    minIndex = j;

            swap(arr, i, minIndex);
        }
    }

    /**
     * 交换数组间的位置
     * @param arr 目标数组
     * @param i 交换位置索引
     * @param j 交换位置索引
     */
    private static void swap(Object[] arr, int i, int j) {
        Object t = arr[i];
        arr[i] = arr[j];
        arr[j] = t;
    }

    public static void main(String[] args) {

        // 测试Integer
        Integer[] a = {10, 9, 8, 7, 6, 5, 4, 3, 2, 1};
        SelectSort.sort(a);
        for (int i = 0; i < a.length; i++) {
            System.out.print(a[i]);
            System.out.print(' ');
        }
        System.out.println();


        // 测试自定义的类 Student
        Student[] d = new Student[4];
        d[0] = new Student("D", 67);
        d[1] = new Student("C", 100);
        d[2] = new Student("B", 34);
        d[3] = new Student("A", 95);
        SelectSort.sort(d);
        for (int i = 0; i < d.length; i++) {
            System.out.println(d[i]);
        }
    }


}
