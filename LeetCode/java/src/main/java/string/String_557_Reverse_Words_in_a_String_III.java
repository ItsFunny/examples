package string;

import java.util.stream.Stream;

/**
 * @author Charlie
 * @When
 * @Description Given a string, you need to reverse the order of characters in each word
 * within a sentence while still preserving whitespace and initial word order.
 * @Detail
 * 1. 简单的拆分之后,反转再拼接
 * @Attention:
 * @Date 创建时间：2020-03-13 15:46
 */
public class String_557_Reverse_Words_in_a_String_III
{
    public String reverseWords(String s)
    {
        String[] chars = s.split(" ");
        String sb = "";

        for (int i = 0; i < chars.length - 1; i++)
        {
            sb += swap(chars[i]);
            sb += " ";
        }
        sb += swap(chars[chars.length - 1]);
        return sb;
    }

    private String swap(String s)
    {
        char[] chars = s.toCharArray();
        int i = 0;
        int j = s.length() - 1;
        while (i < j)
        {
            char temp = chars[i];
            chars[i++] = chars[j];
            chars[j--] = temp;
        }
        return new String(chars);
    }

}