package string;

/**
 * @author Charlie
 * @When
 * @Description Given a string, determine if it is a palindrome,
 * considering only alphanumeric characters and ignoring cases.
 * 既判断是否是回文字符串
 * @Detail 1. 陷入了与链表的思路,先获取到slow和fast,然后slow反转再进行判断,但是却忘记了String自带api,可以
 * 从后往前,而不是单链表
 * @Attention:
 * @Date 创建时间：2020-03-01 16:07
 */
public class String_125_Valid_Palindrome
{
    public static boolean isPalindrome(String s)
    {
        if (s.isEmpty())
        {
            return true;
        }

        int start = 0;
        int last = s.length() - 1;

        while (start < last)
        {
            char c = Character.toLowerCase(s.charAt(start));
            char l = Character.toLowerCase(s.charAt(last));
            if (Character.isLetterOrDigit(c) && Character.isLetterOrDigit(l))
            {
                if (c == l)
                {

                    start++;
                    last--;
                } else
                {
                    return false;
                }
            } else
            {
                if (!Character.isLetterOrDigit(l))
                {
                    last--;
                }

                if (!Character.isLetterOrDigit(c))
                {
                    start++;
                }
            }
        }

        return true;
    }

    public static void main(String[] args)
    {
        System.out.println(isPalindrome("A man, a plan, a canal: Panama"));
    }

}
