package com.basic.algorithm.timewheel.jianzhioffer;

/**
 * @author joker
 * @When
 * @Description 在一个二维数组中（每个一维数组的长度相同），
 * 每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
 * 请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
 * @Detail 题解: 想象成一个棋盘,起始位置可以从左下或者右上开始匹对,左下意味着对于行而言是最大的,列而言是最小的
 * 而当右上的时候代表着对于行而言是最小的,但是对于列而言却是最大的,
 * 所以若这个数大于当前值,则对左下而言意味着在右边(列++),而对于右上而言是下边(行++)
 * @date 创建时间：2019-04-04 18:03
 */
public class _1_FindInTwpoArray
{
    public static boolean Contains(int target, int[][] array)
    {
        // 从右上角开始
        int begin = array.length - 1;
        // key小于值则表示在左边,如果大于则表示在下边
        int len = array[begin].length;
        int j = 0;
        while (begin >= 0 && j < array[begin].length)
        {
            if (array[begin][j] > target)
            {
                begin--;
            } else if (array[begin][j] < target)
            {
                j++;
            } else
            {
                return true;
            }

        }

        return false;

    }

    public static void main(String[] args)
    {
//        [[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
        int[][] arr = {
                {1, 2, 8, 9},
                {2, 4, 9, 12},
                {4, 7, 10, 13},
                {6, 8, 11, 15}

        };
        boolean contains = _1_FindInTwpoArray.Contains(5, arr);
        System.out.println(contains);
    }

}
