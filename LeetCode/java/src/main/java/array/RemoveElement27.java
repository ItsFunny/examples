package array;

/**
 * @author joker
 * @When
 * @Description 返回长度le, 同时对应数组的前l个不是目标元素
 * @Detail 空间复杂度必须要求O(1), 意味着不可以申请额外内存
 * 1. 第一想法就是2个下标,一前一后,
 * 前用来找目标值获得下标,后用来找非目标值也获得下标,然后交换位置,使得目标值相同的全挤在后面
 * 2. 大神的做法:
 * 直接将符合条件的数据扔到count下标所处的位置(不需要担心count是否会覆盖数据,count下标永远<=i)
 * @date 创建时间：2019-12-05 18:29
 * @Attention 1. 对于头尾指针或者是头尾下标,注意其临界点,如本题中的 i==j的时候,会少算i==j的那个下标
 * @Learn 1. 在Java中,数组是可以为空数组的,既内部无元素,但是非NULL,这时候只能够通过length+null复合判断
 */
public class RemoveElement27
{
    public int removeElement(int[] nums, int val)
    {
        if (nums == null || nums.length == 0)
        {
            return 0;
        }
        int count = 0;
        int i = 0, j = nums.length - 1;
        while (true)
        {
            if (i == j)
            {
                return nums[i] == val ? count : count + 1;
            }
            if (nums[i] == val)
            {
                while (j > i)
                {
                    if (nums[j] != val)
                    {
                        int temp = nums[i];
                        nums[i] = nums[j];
                        nums[j] = temp;
                        break;
                    } else
                    {
                        j--;
                    }
                }

            } else
            {
                count++;
                i++;

            }

        }
    }

    public int removeElementForSimple(int[] nums, int val)
    {
        int count = 0;
        for (int i = 0; i < nums.length; i++)
        {
            if (nums[i] != val)
            {
                nums[count++] = nums[i];
            }
        }
        return count;
    }

}
