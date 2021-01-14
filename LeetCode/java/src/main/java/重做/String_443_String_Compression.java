package 重做;

/**
 * @author Charlie
 * @When
 * @Description Given an array of characters, compress it in-place.
 * The length after compression must always be
 * smaller than or equal to the original array.
 * <p>
 * Every element of the array
 * should be a character (not int) of length 1.
 * <p>
 * After you are done modifying the input array in-place, return the new length of the array.
 * <p>
 * <p>
 * Follow up:
 * Could you solve it using only O(1) extra space?
 * 类似于统计出现的次数,最终返回的是(字符串)长度
 * 条件
 * 1. 空间复杂度O(1)
 * 2. 最终结果要有序
 * 解题关键:
 * 并且输入数据input如果重复是一起重复的,不存在 abaca 这种情况,只有aaabc这种情况
 * 找到有哪些不同的数,为不同的数记录次数
 * @Detail 1. 通过sortedMap实现排序,最后遍历即可
 * 2.
 * @Attention:  
 * @Date 创建时间：2020-03-05 15:48
 */
public class String_443_String_Compression
{
    public int compress(char[] chars)
    {
        int len = 0; // also a pointer to modify array in-place
        for (int i = 0; i < chars.length; )
        {
            chars[len] = chars[i];
            int j = i + 1;

            // 计算该字符出现的次数
            while (j < chars.length && chars[j] == chars[i])
            {
                j++;
            }

            if (j - i > 1)
            // 说明需要压缩
            { // need compression
                String freq = j - i + "";
                for (char c : freq.toCharArray())
                    chars[++len] = c;
            }
            len++;
            i = j;
        }
        return len;
    }
}
