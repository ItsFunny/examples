package array;

/**
 * @author joker
 * @When
 * @Description 在一个整型数组中, 求连续的最大和
 * @Detail 1. O(n)
 * 2. 分治法
 * 思路:
 * 并不需要返回值所对应的具体数值,因而只需要通过2个变量进行for循环匹配即可
 * 一个变量保存最大值,另外一个变量则是根据遍历的元素进行数据加减,再与保存着的最大值进行比较
 * @date 创建时间：2019-12-12 18:41
 */
public class MaximumSubArrayNum53
{
    public int maxSubArray(int[] nums)
    {
        Integer max = nums[0];
        Integer indexResult = nums[0];
        // 1.
        for (int i = 1; i < nums.length; i++)
        {
            // 将这个下标对应的元素 加到result结果中,result的结果是包含了[0,i]下标的元素所有和
            // 与nums[i] 比较的原因在于,当此index的元素所累积的和小于当前下标的值的时候,根据题意所要获取的
            // 是最大的值,所以这里需要更改元素,从而再与最大值进行匹配
            indexResult = Math.max(indexResult + nums[i], nums[i]);
            max = Math.max(indexResult, max);
        }

        return max;

    }
}
