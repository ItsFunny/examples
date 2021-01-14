package Top100;

/**
 * @author Charlie
 * @When
 * @Description You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
 * <p>
 * Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.
 * @Detail
 * 动态规划
 * @Attention:
 * @Date 创建时间：2020-04-10 16:28
 */
public class Too100_Easy_198_House_Robber
{
    public int rob(int[] nums)
    {
        return rob(nums, nums.length - 1);
    }

    private int rob(int[] nums, int i)
    {
        if (i < 0)
        {
            return 0;
        }
        return Math.max(rob(nums, i - 2) + nums[i], rob(nums, i - 1));
    }
}
