package array;

/**
 * @author Charlie
 * @When
 * @Description 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 * @Detail 1. 双指针,第一个指针当发现0的时候,启动第二个指针,与遇到的第一个非0的数交换
 * 2. 参考大神的思路: 利用一个变量(可以认为是count,记录0出现的次数遍历的时候,将非0的全部往左移动,最后再填充0)
 * 同时,当发生交换的时候,count下标所处的变量注定是无效的
 * @Attention:
 * @Date 创建时间：2020-02-21 16:06
 */
public class Array_283_move_zeroes
{
    public void moveZeroes(int[] nums)
    {
        int count = 0;
        for (int i = 0; i < nums.length; i++)
        {
            if (nums[i] == 0)
            {
                count++;
                for (int j = i + 1; j < nums.length; j++)
                {

                    if (nums[j] != 0)
                    {
                        int temp = nums[i];
                        nums[i] = nums[j];
                        nums[j] = temp;
                        break;
                    }
                }
            }

        }
    }


    public void moveZeroes2(int[] nums)
    {
        int count = 0;
        for (int i = 0; i < nums.length; i++)
        {
            if (nums[i] != 0)
            {
                int tempV = nums[count++];
                nums[count] = nums[i];
                nums[i] = tempV;
            }

        }
    }

}
