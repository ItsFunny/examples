package string;

import sun.jvm.hotspot.debugger.RandomAccessFileDataSource;

import java.util.HashMap;
import java.util.Map;
import java.util.Random;
import java.util.stream.Stream;

/**
 * @author Charlie
 * @When
 * @Description Given an arbitrary ransom note string and another
 * string containing letters from all the magazines,
 * write a function that will return true
 * if the ransom note can be constructed from the magazines ;
 * otherwise, it will return false.
 * <p>
 * Each letter in the magazine string can only be used once in your ransom note.
 * <p>
 * Note:
 * You may assume that both strings contain only lowercase letters.
 * <p>
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 * <p>
 * 题目没看懂
 * 题目指的是,magazines中的字符串能否组成ransom中的值
 * @Detail 1. map搜集字符串,并且搜集次数
 * 2. 因为单词只有26个字母,因此可以定义长度为26的数组,存放这些字符型,value代表的是重复出现的次数
 * @Attention:
 * @Date 创建时间：2020-03-03 14:48
 */
public class String_383_Ransom_Note
{
    public boolean canConstruct(String ransomNote, String magazine)
    {
        Map<Character, Integer> map = new HashMap<>();
        char[] chars = magazine.toCharArray();
        for (char aChar : chars)
        {
            if (map.containsKey(aChar))
            {
                map.put(aChar, map.get(aChar) + 1);
            } else
            {
                map.put(aChar, 1);
            }
        }

        for (int i = 0; i < ransomNote.length(); i++)
        {
            char c = ransomNote.charAt(i);
            if (map.containsKey(c))
            {
                Integer integer = map.get(c);
                if (integer <= 0)
                {
                    return false;
                }
                map.put(c, integer - 1);
            } else
            {
                return false;
            }
        }
        return true;
    }


    public boolean canConstruct2(String ransomNote, String magazine)
    {
        int[] arr = new int[26];
        magazine.chars().forEach(c -> arr[c - 'a']++);

        // for 循环遍历判断是否为0;
        return true;
    }
}
