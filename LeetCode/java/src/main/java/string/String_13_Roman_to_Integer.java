package string;

import sun.jvm.hotspot.debugger.bsd.amd64.BsdAMD64CFrame;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M
 * For example,
 * two is written as II in Roman numeral,
 * just two one's added together.
 * Twelve is written as, XII, which is simply X + II.
 * The number twenty seven is written as XXVII, which is XX + V + II.
 * <p>
 * Roman numerals are usually written largest to smallest from left to right.
 * However, the numeral for four is not IIII.
 * Instead, the number four is written as IV.
 * Because the one is before the five we subtract it making four.
 * The same principle applies to the number nine, which is written as IX.
 * There are six instances where subtraction is used:
 * <p>
 * 注意点:
 * 罗马数字,大的在左边,小的在右边
 * I can be placed before V (5) and X (10) to make 4 and 9.
 * X can be placed before L (50) and C (100) to make 40 and 90.
 * C can be placed before D (500) and M (1000) to make 400 and 900.
 * @Detail 1. 遍历判断
 * @Attention: 1. 只需要注意最后一个char 是单个还是多个,既判断之前的是单个还是多个
 * @Date 创建时间：2020-02-27 16:03
 */
public class String_13_Roman_to_Integer
{
    public static int romanToInt(String s)
    {
        Map<Character, Integer> map = new HashMap<>();
        map.put('I', 1);
        map.put('V', 5);
        map.put('X', 10);
        map.put('L', 50);
        map.put('C', 100);
        map.put('D', 500);
        map.put('M', 1000);
        char[] chars = s.toCharArray();
        int result = 0;
        int i = 0, j = 1;
        for(; j < chars.length; i++, j++) {
            if (map.get(chars[i]) >= map.get(chars[j])) {
                result += map.get(chars[i]);
            } else {
                result -= map.get(chars[i]);
            }
        }
        result += map.get(chars[i]);
        return result;
    }

    public static void main(String[] args)
    {
        System.out.println(romanToInt("III"));
    }

}
