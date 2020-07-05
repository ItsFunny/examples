package array;

/**
 * @author Charlie
 * @When
 * @Description 给定一个非负整数组成的非空数组，在该数的基础上加一，返回一个新的数组。
 * 最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。
 * 你可以假设除了整数 0 之外，这个整数不会以零开头。
 * @Detail 1. 常规做法是直接遍历然后加一,注意点就是当末尾为9的数字,1299 或者是1399,又或者是9999 全是9的数字
 * 当全是9的情况下,需要在原先的数组的长度下再扩容一位
 * <p>
 * 2. 最优解:
 * 他人的最优解其实本质相同,唯一不同在于代码行数能极大减少
 * @Attention:
 * 核心就是for循环遍历判断,唯一需要注意的只有 9999 全为9的特殊情况
 * @Date 创建时间：2020-02-14 12:26
 */
public class Array_66_Plus_One
{
    public int[] plusOne(int[] digits)
    {
        int length = digits.length;
        int count = 0;

        for (int i = length - 1; i >= 0; i--)
        {
            digits[i]++;
            if (digits[i] == 10)
            {
                digits[i] = 0;
                count++;
            } else
            {
                break;
            }
        }
        if (count == length)
        {
            int[] result = new int[length + 1];
            result[0] = 1;
            for (int i = 1; i < result.length; i++)
            {
                result[i] = digits[i - 1];
            }

            return result;
        }
        return digits;
    }


    public int[] plusOneWithMinCodes(int[] digits)
    {
        for (int i = digits.length - 1; i >= 0; i--)
        {
            digits[i]++;
            if (digits[i] < 9)
            {
                return digits;
            }
            digits[i] = 0;
        }

        int[] result = new int[digits.length + 1];
        result[0] = 1;
        return result;
    }
}
