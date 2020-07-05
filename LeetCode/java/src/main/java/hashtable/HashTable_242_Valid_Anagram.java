package hashtable;

/**
 * @author Charlie
 * @When
 * @Description Given two strings s and t , write a function to determine if t is an anagram of s.
 * 既判断是否是相同字母但是不同排列
 * @Detail 1. 排序 ,然后一个一个匹配即可
 * 2. 两个for循环遍历匹配
 * 3. 基于字符串的特性,判断是否包含,可以通过一个++ 一个-- 的形式 最终判断
 * @Attention:
 * @Date 创建时间：2020-04-07 16:36
 */
public class HashTable_242_Valid_Anagram
{
    public boolean isAnagram(String s, String t)
    {
        if (s.length() != t.length()) return false;

        int[] chars = new int[26];
        for (int i = 0; i < s.length(); i++) chars[s.charAt(i) - 'a']++;
        for (int i = 0; i < t.length(); i++) chars[t.charAt(i) - 'a']--;
        for (int i = 0; i < chars.length; i++)
        {
            if (chars[i] <0)
            {
                return false;
            }
        }

        return true;
    }
}
