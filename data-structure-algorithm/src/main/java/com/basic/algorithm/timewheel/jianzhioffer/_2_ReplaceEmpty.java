package com.basic.algorithm.timewheel.jianzhioffer;

import java.io.File;
import java.util.Stack;

/**
 * @author joker
 * @When
 * @Description 请实现一个函数，
 * 将一个字符串中的每个空格替换成“%20”。
 * 例如，当字符串为We Are Happy.则经过替换之后的字符串为We%20Are%20Happy。
 * @Detail 解题注意点
 * 1. 需要额外申请空间用来存储数据
 * 2. 当遍历匹配的时候需要设定两个下标来判断元素的位置,一个下标代表原先的数组,另外一个下标代表新建的数组
 * 并且要以原先的数组下标为基础,
 * @date 创建时间：2019-04-04 18:09
 */
public class _2_ReplaceEmpty

{
    public static String changeBlank2SpecialStr(StringBuffer str)
    {

        if (str.length() == 0)
        {
            return "";
        }

        Integer count = 0;
        for (int i = 0; i < str.length(); i++)
        {
            if (str.charAt(i) == ' ') count++;
        }
        char[] newArray = new char[str.length() + (count << 1)];
        for (int i = 0, j = 0; i < str.length(); j++, i++)
        {
            if (str.charAt(i) != ' ')
            {
                newArray[j] = str.charAt(i);
            } else
            {
                newArray[j++] = '%';
                newArray[j++] = '2';
                newArray[j] = '0';
            }
        }
        StringBuilder sb = new StringBuilder();
        for (char c : newArray)
        {
            sb.append(c);
        }
        return sb.toString();
    }

    public static void main(String[] args)
    {
        System.out.println(_2_ReplaceEmpty.changeBlank2SpecialStr(new StringBuffer("")));
    }

}
