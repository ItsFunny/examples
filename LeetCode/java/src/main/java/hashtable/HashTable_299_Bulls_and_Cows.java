package hashtable;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description You are playing the following Bulls and Cows game with your friend: You write down a number and ask your friend to guess what the number is. Each time your friend makes a guess, you provide a hint that indicates how many digits in said guess match your secret number exactly in both digit and position (called "bulls") and how many digits match the secret number but locate in the wrong position (called "cows"). Your friend will use successive guesses and hints to eventually derive the secret number.
 * <p>
 * Write a function to return a hint according to the secret number and friend's guess, use A to indicate the bulls and B to indicate the cows.
 * <p>
 * Please note that both secret number and friend's guess may contain duplicate digits.
 * 题目开始没看懂,直接搜的别人的题意:
 * 给出两个数字，输出（1）A有多少位是相同的（2）B有多少位不在正确的位置上。
 * 1A3B 代表 1个是相同的(元素相同,下标相同),3个不同的(B代表不同)
 * @Detail 1. 直接用map来存储下标对应的元素即可
 * @Attention: 3. 题目没看懂
 * 既然数字在0-9 ,所以可以定义一个int[10]的数组,index对应的值则为是否存在,secret ++ <0 代表 guess中存在这个值
 * guess-- >0 代表secret中存在这个值
 * @Date 创建时间：2020-03-20 16:09
 */
public class HashTable_299_Bulls_and_Cows
{
    public String getHint(String secret, String guess)
    {
        int same = 0;
        int dif = 0;
        int[] nums = new int[10];
        for (int i = 0; i < secret.length(); i++)
        {
            if (secret.charAt(i) == guess.charAt(i))
            {
                same++;
            } else
            {
                if (nums[secret.charAt(i) - '0']++ < 0) dif++;
                if (nums[guess.charAt(i) - '0']-- > 0) dif++;
            }

        }

        return same + "A" + dif + "B";
    }

    public static void main(String[] args)
    {

    }
}
