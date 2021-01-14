package com.basic.algorithm.timewheel.jianzhioffer;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.zip.Inflater;

/**
 * @author joker
 * @When
 * @Description 输入一个整数，输出该数二进制表示中1的个数。其中负数用补码表示。
 * @Detail 整数转换为二进制的方式为: 一个数不停的除以2
 * 这题是我盲目了,其实压根就不需要我自己将其转换为二进制数,完全用&, <<,>>这种即可
 * 这题参考大神的做法, 这个数-1 之后 二进制中的最右边的1会变为0,1之后的0会变为1
 * 同时原数值与这个-1的数做与运算后,原先因为-1而导致的 0变为1的值重新变为0 ,所以 可以执行多少次
 * 就代表有多少个1  n=n&(n-1)
 * @date 创建时间：2019-05-20 22:37
 */
public class _11_number_of_1
{
    public static int NumberOf1(int n)
    {
        int count = 0;

        while (n != 0)
        {
            count++;
            n = n & (n - 1);
        }

        return count;
    }

    public static void main(String[] args)
    {
        System.out.println(_11_number_of_1.NumberOf1(3));
    }

}
