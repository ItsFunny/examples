package string;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description Write a function
 * that takes a string as input and reverse only the vowels of a string.
 * 转换元音字母: a,e,i,o,u已经大写的也算 A,E,I,O,U
 * 但是注意,并不是反转字符串,而是后端也需要进行移动
 * @Detail
 * 双指针
 * 一左一右移动匹配即可
 * @Attention:
 * 犯了的错误:
 * 1. 在右端移动,匹配到时忘记右端也需要移动
 *
 * @Date 创建时间：2020-03-02 16:29
 */
public class String_345_Reverse_Vowels_of_a_String
{
    public static String reverseVowels(String s)
    {

        int start = 0;
        Map<Character, Byte> map = new HashMap<>();
        map.put('a', (byte) 1);
        map.put('e', (byte) 1);
        map.put('i', (byte) 1);
        map.put('o', (byte) 1);
        map.put('u', (byte) 1);

        map.put('A', (byte) 1);
        map.put('E', (byte) 1);
        map.put('I', (byte) 1);
        map.put('O', (byte) 1);
        map.put('U', (byte) 1);
        char[] chars = s.toCharArray();

        int end = s.length() - 1;

        while (start < end)
        {
            char st = chars[start];
            if (map.containsKey(st))
            {
                while (end > start)
                {
                    char lt = chars[end];
                    if (!map.containsKey(lt))
                    {
                        end--;
                        continue;
                    }
                    char temp = st;
                    chars[start] = lt;
                    chars[end] = temp;
                    end--;
                    break;
                }
            }
            start++;
        }
        StringBuilder stringBuilder = new StringBuilder();
        for (char aChar : chars)
        {
            stringBuilder.append(aChar);

        }

        return stringBuilder.toString();
    }

    public static void main(String[] args)
    {
        System.out.println(reverseVowels("leetcode"));
    }

}
