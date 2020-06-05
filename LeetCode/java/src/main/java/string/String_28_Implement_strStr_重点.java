package string;

/**
 * @author Charlie
 * @When
 * @Description Implement strStr().
 * Return the index of the first
 * occurrence of needle in haystack, or -1 if needle is not part of haystack.
 * 既判断 子串出现的第一次index下标,既实现String 的indexOf
 * @Detail 1. 双层for 循环 O(n^2) 匹配
 * 2. 截取子串进行判断,
 * @Attention: 遇到的坑:
 * 1. 试图用一个bool 变量判断是否是子串,但是却忽视了当 最后一个元素恰好在hayStack存在的时候,会为true
 * @Date 创建时间：2020-02-28 16:33
 */
public class String_28_Implement_strStr_重点
{
    // 错误的做法
    public static int strStr(String haystack, String needle)
    {
        if (haystack == null || needle == null) return -1;
        if (needle.length() == 0) return 0;
        if (needle.length() > haystack.length()) return -1;
        char[] hayStackArray = haystack.toCharArray();
        char[] needleArray = needle.toCharArray();
        int similarCount = 0;
        for (int i = 0; i < hayStackArray.length; i++)
        {
            for (int j = 0; j < needleArray.length; j++)
            {
                if (needleArray[j] == hayStackArray[i])
                {
                    similarCount++;
                } else
                {
                    continue;
                }
            }
            if (similarCount == needleArray.length)
            {
                return i;
            }
            similarCount = 0;
        }
        return -1;
    }

    public int strStrBySub(String h, String n)
    {
        if (n.length() == 0) return 0;
        for (int i = 0; i < h.length() - n.length(); i++)
        {
            if (h.substring(i, i + n.length()).equals(n))
            {
                return i;
            }
        }
        return -1;
    }

    public int strStr2(String haystack, String needle)
    {
        for (int i = 0; ; i++)
        { // The length of haystack
            for (int j = 0; ; j++)
            { // The length of needle
                if (j == needle.length())
                {
                    return i; // If at this point we have navigated through the entire length of needle, we have found a solution, haystack[i].
                }
                if (i + j == haystack.length())
                {
                    return -1; // This happens when we run out of elements in haystack, but there are still elements in needle.
                }
                if (needle.charAt(j) != haystack.charAt(i + j))
                {
                    break; // We stop comparing after needle[j], so i will be incremented and cycle repeats itself.
                }
            }
        }
    }

    public static void main(String[] args)
    {
        System.out.println(String_28_Implement_strStr_重点.strStr("mississippi", "mississippi"));
    }
}
