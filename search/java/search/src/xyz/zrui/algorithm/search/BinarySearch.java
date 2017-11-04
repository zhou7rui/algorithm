package xyz.zrui.algorithm.search;

public class BinarySearch {

    public static int rank(int key, int[] arr) {

        int hi = arr.length;
        int lo = 0;

        while (lo <= hi) {

            int mid = lo + (hi - lo) / 2;

            if (arr[mid] > key) {
                hi = mid - 1;
            } else if (arr[mid] < key) {
                lo = mid + 1;
            } else {
                return mid;
            }
        }
        return -1;

    }

    public static void main(String[] args) {
        int[] arr = {0,1,2,4,5,6,8,9};
        System.out.println(rank(6,arr));

    }




}
