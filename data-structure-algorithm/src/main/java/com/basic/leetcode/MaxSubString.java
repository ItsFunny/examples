package com.basic.leetcode;

import java.security.Key;
import java.util.HashSet;
import java.util.Set;

/**
 * @author joker
 * @When
 * @Description 最长重复子串
 * @Detail
 * @date 创建时间：2019-02-18 11:40
 */
public class MaxSubString
{

    // 第一种方法: 双重指针
    public static String maxSubString(String string)
    {
        // 最长长度
        int max = 0;
        // 最长子串的起始位置
        int startIndex = 0;
        // 临时变量,用于与max比较大小
        int tempLength = 0;
        for (int i = 1; i < string.length(); i++)
        {
            for (int j = 0; j < string.length() - i; j++)
            {
                if (string.charAt(i + j) == string.charAt(j))
                {
                    tempLength++;
                } else
                {
                    tempLength = 0;
                }
                if (tempLength > max)
                {
                    max = tempLength;
                    startIndex = j - tempLength + 1;
                }
            }
        }
        return string.substring(startIndex, startIndex + max + 1);
    }

    // 这是求最长无重复子串的
    public static int maxSubString2(String s)
    {
        Set<Character> set = new HashSet<>();
        int i = 0, j = 0, l = s.length(), max = 0, startIndex = 0;
        while (i < l & j < l)
        {
            if (!set.contains(s.charAt(i)))
            {
                set.add(s.charAt(i++));
                max = Math.max(max, i - j);
            } else
            {
                set.remove(s.charAt(j++));
            }
        }
        return max;
    }

    public static void main(String[] args)
    {
        System.out.println(maxSubString("asdjkaddddddddjqjqqqnasdm=d"));
        System.out.println(maxSubString2("asdjkaddddddddjqjqqqnasdm"));
    }
}
