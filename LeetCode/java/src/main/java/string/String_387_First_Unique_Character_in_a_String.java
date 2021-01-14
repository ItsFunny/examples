package string;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description Given a string,
 * find the first non-repeating character
 * in it and return it's index. If it doesn't exist, return -1.
 * 既找出第一个没有重复的元素的下标
 * @Detail 1. 第一反应,用map来做
 * 2. 第二反应用数组来做: 既然全是字符的话,26个字符长度数组,有疏忽,因为返回的
 * 是首个不重复的下标值
 * 2.1 所有可以考虑用 数组+map来实现:
 * 数组的长度根据字符串长度定义,map来辅助判断  : 既map存储的是下标,存在的话取得下标然后value++
 * 2.2 用数组是最好的,这里有疏忽,疏忽的就是第二次遍历的时候,我是通过下标遍历的
 * 却忽略了当初申请数组内存的时候就是以字符的顺序申请的,因此下标改为字符序即可,如firstUniqChar4
 * @Attention:
 * @Date 创建时间：2020-03-03 16:10
 */
public class String_387_First_Unique_Character_in_a_String
{
    public static int firstUniqChar(String s)
    {

        Map<Character, Byte> map = new HashMap<>();
        char[] chars = s.toCharArray();
        int minIndex = 0;
        for (int i = 0; i < chars.length; i++)
        {
            if (!map.containsKey(chars[i]))
            {
                map.put(chars[i], (byte) 1);
                minIndex = Math.min(minIndex, i);
            } else
            {

            }
        }

        // 说明全都重复
        if (map.size() == s.length() && map.size() % 2 == 0)
        {
            return -1;
        }

        return map.size() == s.length() ? 0 : minIndex;
    }


    public static int firstUniqChar2(String s)
    {
        int[] chars = new int[26];
        for (int i = 0; i < s.length(); i++)
        {
            char c = s.charAt(i);
            int i1 = c - 'a';
            chars[i1]++;
        }

        for (int i = 0; i < chars.length; i++)
        {
            if (chars[i] == 1)
            {
                return i;
            }

        }


        return -1;
    }


    public static int firstUniqChar3(String s)
    {
        Map<Character, Integer> map = new HashMap<>();
        int length = s.length();
        int[] ints = new int[length];
        char c;
        for (int i = 0; i < length; i++)
        {
            c = s.charAt(i);
            if (map.containsKey(c))
            {
                ints[map.get(c)]++;
            } else
            {
                ints[i] = 1;
                map.put(c, i);
            }

        }
        for (int i = 0; i < length; i++)
        {
            if (ints[i] == 1)
            {
                return i;
            }

        }
        return -1;
    }


    public static int firstUniqChar4(String s)
    {
        int[] chars = new int[26];

        s.chars().forEach(c -> chars[c - 'a']++);
        for (int i = 0; i < s.length(); i++)
        {
            if (chars[s.charAt(i) - 'a'] == 1)
            {
                return i;
            }
        }
        return -1;
    }

    public static void main(String[] args)
    {
        System.out.println(firstUniqChar3("loveleetcode"));

    }

}
