package hashtable;

/**
 * @author Charlie
 * @When
 * @Description Given a non-empty array of integers, every element appears twice except for one. Find that single one.
 * <p>
 * Note:
 * <p>
 * Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 * @Detail 既然一组奇数长度的数组, 只有一个元素出现1次, 其他都是2次
 * Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?
 * @Attention: 1. 不可申请额外的数组空间,只能在这基础上进行
 * 1. 排序,前后不相等的则直接return
 * @Date 创建时间：2020-04-02 16:26
 */
public class HashTable_136_Single_Number
{
    public int singleNumber(int[] nums)
    {
        int result = 0;
        for (int i = 0; i < nums.length; i++)
        {
            result ^= nums[i];
        }
        return result;

    }

    public static void main(String[] args)
    {
        System.out.println(2 ^ 1);
    }
}
