package xyz.zrui.algorithm.sorting;


import java.lang.reflect.Method;
import java.util.Random;

public class SortTestHelper {


    /**
     * 生成int型随机数组 [rangeL, rangeR]
     *
     * @param n      生成数组长度
     * @param rangeL 随机数最小值
     * @param rangeR 随机数最大值
     * @return 随机数组
     */
    public static Integer[] generateRandomArray(int n, int rangeL, int rangeR) {

        if (rangeL >= rangeR) {
            throw new RuntimeException();
        }

        Integer[] arr = new Integer[n];


        for (int i = 0; i < n; i++) {
            Random random = new Random();
            arr[i] = random.nextInt((rangeR - rangeL) + 1) + rangeL;
        }

        return arr;
    }


    /**
     * 验证排序是否正确
     *
     * @param arr 目标数组
     * @return 是否正确
     */
    public static boolean isSorted(Comparable[] arr) {

        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i].compareTo(arr[i + 1]) > 0) {
                return false;
            }
        }
        return true;

    }


    /**
     * 测试排序算法的性能
     * 通过反射的方式运行  参考链接: https://docs.oracle.com/javase/tutorial/reflect/
     * @param sortClass 排序函数class对象
     * @param arr 排序数组
     */
    public static void testSort(Class sortClass, Comparable[] arr) {

        try {

            //通过反射获取排序方法
            Method method = sortClass.getMethod("sort", new Class[]{Comparable[].class});
            // 排序参数只有一个，是可比较数组arr
            Object[] params = new Object[]{arr};

            long startTime = System.currentTimeMillis();
            //执行排序方法
            method.invoke(sortClass.newInstance(), params);
            long endTime = System.currentTimeMillis();

            if (!isSorted(arr)) {
                throw new RuntimeException("排序失败");
            }
            printArray(arr);
            System.out.println(sortClass.getSimpleName() + ":" + (endTime - startTime) + "ms");

        } catch (Exception e) {
            e.printStackTrace();
        }

    }


    /**
     * 打印数组
     *
     * @param arr 目标数组
     */
    public static void printArray(Object arr[]) {

        for (Object object : arr) {
            System.out.print(object);
            System.out.print(" ");
        }

        System.out.println();


    }


    public static void main(String[] args) {

        // 测试自定义的类 Student
        Student[] students = new Student[4];
        students[0] = new Student("D", 67);
        students[1] = new Student("C", 100);
        students[2] = new Student("B", 34);
        students[3] = new Student("A", 95);
        SortTestHelper.testSort(SelectSort.class,students);


        Integer[] integers = SortTestHelper.generateRandomArray(10000,7,10000);
        SortTestHelper.testSort(SelectSort.class,integers);
        SortTestHelper.printArray(integers);
    }



}
