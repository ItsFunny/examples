package hashtable;

/**
 * @author Charlie
 * @When
 * @Description Given two strings s and t, determine if they are isomorphic.
 * <p>
 * Two strings are isomorphic if the characters in s can be replaced to get t.
 * <p>
 * All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.
 * 既判断s和t 的格式相同 aba,aaab 等格式相同即可
 * @Detail 没思路: 如何判断2个字符串的结构是否一致
 * ascii最长为256个字符,2个数组只需要判断下标是否一致即可
 * 数组中的值存储的是出现的下标,如果下标不一致说明结构就不一致
 * @Attention:
 * 解题的关键就是需要记录 字符与下标的关系,并且同一个字符下标要保持不变
 * 所以可以使用map或者数组(因为ascii有256个字符)
 * @Date 创建时间：2020-04-04 16:33
 */
public class HashTable_205_Isomorphic_Strings
{
    public boolean isIsomorphic(String s1, String s2)
    {
        int[] m = new int[512];
        for (int i = 0; i < s1.length(); i++)
        {
            if (m[s1.charAt(i)] != m[s2.charAt(i) + 256]) return false;
            m[s1.charAt(i)] = m[s2.charAt(i) + 256] = i + 1;
        }
        return true;
    }
}
