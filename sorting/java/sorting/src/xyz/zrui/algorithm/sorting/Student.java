package xyz.zrui.algorithm.sorting;

/**
 * 测试学生类
 * 实现 Comparable 接口 实现支持排序
 */
public class Student implements Comparable<Student> {


    private String name;

    private int score;


    //定义Student 的 compareTo函数
    // 如果分数相等 则按照 名字的字母排序
    // 如果分数不等 则分数高的靠前
    @Override
    public int compareTo(Student that) {

        if (that.getScore() == this.getScore()) {
            return this.name.compareTo(that.getName());
        } else if (that.getScore() > this.getScore()) {
            return -1;
        } else {
            return 1;
        }

    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getScore() {
        return score;
    }

    public void setScore(int score) {
        this.score = score;
    }

    public Student(String name, int score) {
        this.name = name;
        this.score = score;
    }

    @Override
    public String toString() {
        return "Student{" +
                "name='" + name + '\'' +
                ", score=" + score +
                '}';
    }
}
